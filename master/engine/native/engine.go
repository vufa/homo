package native

import (
	"fmt"
	"github.com/countstarlight/homo/sdk/homo-go"
	"go.uber.org/zap"
	"os"
	"path"
	"strings"
	"time"

	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/master/engine"
	"github.com/countstarlight/homo/utils"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/shirou/gopsutil/process"
)

// TODO: need unit test for engine
// NAME of this engine
const NAME = "native"

func init() {
	engine.Factories()[NAME] = New
}

// New native engine
func New(stats engine.InfoStats, opts engine.Options) (engine.Engine, error) {
	e := &nativeEngine{
		InfoStats: stats,
		pwd:       opts.Pwd,
		grace:     opts.Grace,
		log:       logger.New(logger.LogInfo{Level: "debug"}, "engine", NAME),
	}
	return e, nil
}

type nativeEngine struct {
	engine.InfoStats
	pwd   string // work directory
	grace time.Duration
	log   *zap.SugaredLogger
}

// Name of engine
func (e *nativeEngine) Name() string {
	return NAME
}

// Recover recover old services when master restart
func (e *nativeEngine) Recover() {
	// clean old services in native mode
	e.clean()
}

// Prepare prepares all images
func (e *nativeEngine) Prepare(homo.ComposeAppConfig) {
	// do nothing in native mode
}

// Clean clean all old instances
func (e *nativeEngine) clean() {
	sss := map[string]map[string]attribute{}
	if e.LoadStats(&sss) {
		for sn, instances := range sss {
			for in, instance := range instances {
				id := int32(instance.Process.ID)
				if id == 0 {
					e.log.Warnf("[%s][%s] process id not found, maybe running mode changed", sn, in)
					continue
				}
				name := instance.Process.Name
				p, err := process.NewProcess(id)
				if err != nil {
					e.log.Warn(fmt.Sprintf("[%s][%s] failed to get old process (%d)", sn, in, id), zap.Error(err))
					continue
				}
				pn, err := p.Name()
				if err != nil {
					e.log.Warn(fmt.Sprintf("[%s][%s] failed to get name of old process (%d)", sn, in, id), zap.Error(err))
					continue
				}
				if pn != name {
					e.log.Debugf("[%s][%s] name of old process (%d) not matched, %s -> %s", sn, in, id, name, pn)
					continue
				}
				err = p.Kill()
				if err != nil {
					e.log.Warnf("[%s][%s] failed to stop the old process (%d)", sn, in, id)
				} else {
					e.log.Infof("[%s][%s] old process (%d) stopped", sn, in, id)
				}
			}
		}
	}
}

// Run new service
func (e *nativeEngine) Run(name string, cfg homo.ComposeService, _ map[string]homo.ComposeVolume) (engine.Service, error) {
	spwd := path.Join(e.pwd, "var", "run", "homo", "services", name)
	err := os.RemoveAll(spwd)
	if err != nil {
		return nil, err
	}
	err = mountAll(e.pwd, spwd, cfg.Volumes)
	if err != nil {
		os.RemoveAll(spwd)
		return nil, err
	}
	var pkg packageConfig
	image := strings.Replace(strings.TrimSpace(cfg.Image), ":", "/", -1)
	pkgDir := path.Join(spwd, "lib", "homo", image)
	err = utils.LoadYAML(path.Join(pkgDir, packageConfigPath), &pkg)
	if err != nil {
		os.RemoveAll(spwd)
		return nil, err
	}
	params := processConfigs{
		exec: path.Join(pkgDir, pkg.Entry),
		env:  utils.AppendEnv(cfg.Environment.Envs, true),
		argv: cfg.Command.Cmd,
		pwd:  spwd,
	}
	s := &nativeService{
		name:      name,
		cfg:       cfg,
		engine:    e,
		params:    params,
		instances: cmap.New(),
		log:       e.log.With(zap.String("service", name)),
	}
	err = s.Start()
	if err != nil {
		s.Stop()
		return nil, err
	}
	return s, nil
}

// Close engine
func (e *nativeEngine) Close() error {
	return nil
}

func mountAll(epwd, spwd string, ms []*homo.ServiceVolume) error {
	for _, m := range ms {
		if len(m.Source) == 0 {
			return fmt.Errorf("host path is empty")
		}
		// for preventing path escape
		m.Source = path.Join(epwd, path.Join("/", m.Source))
		err := mount(m.Source, path.Join(spwd, strings.TrimSpace(m.Target)))
		if err != nil {
			return err
		}
	}
	sock := utils.GetEnv(homo.EnvKeyMasterAPISocket)
	if sock != "" {
		err := mount(sock, path.Join(spwd, homo.DefaultSockFile))
		if err != nil {
			return err
		}
	}
	grpcSock := utils.GetEnv(homo.EnvKeyMasterGRPCAPISocket)
	if grpcSock != "" {
		return mount(grpcSock, path.Join(spwd, homo.DefaultGRPCSockFile))
	}
	return nil
}

func mount(src, dst string) error {
	// if it is a file mapping, the file must exist, otherwise it
	// will be used as a dir mapping and make the dir.
	if !utils.PathExists(src) {
		err := os.MkdirAll(src, 0755)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(path.Dir(dst), 0755)
	if err != nil {
		return err
	}
	return os.Symlink(src, dst)
}
