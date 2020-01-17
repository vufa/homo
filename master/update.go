package master

import (
	"fmt"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/countstarlight/homo/utils"
	"github.com/inconshreveable/go-update"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path"
	"strings"
)

var appDir = path.Join("var", "db", "homo")
var appConfigFile = path.Join(appDir, homo.AppConfFileName)
var appBackupFile = path.Join(appDir, homo.AppBackupFileName)

// UpdateSystem updates application or master
func (m *Master) UpdateSystem(trace, tp, target string) (err error) {
	switch tp {
	case homo.OTAMST:
		err = m.UpdateMST(trace, target, homo.DefaultBinBackupFile)
	default:
		err = m.UpdateAPP(trace, target)
	}
	if err != nil {
		err = fmt.Errorf("failed to update system: %s", err.Error())
		m.log.Errorf(err.Error())
	}
	m.infostats.setError(err)
	return err
}

// UpdateAPP updates application
func (m *Master) UpdateAPP(trace, target string) error {
	log := m.log
	isOTA := target != "" || utils.IsFile(m.cfg.OTALog.Path)
	if isOTA {
		log = logger.New(m.cfg.OTALog, homo.OTAKeyTrace, trace, homo.OTAKeyType, homo.OTAAPP)
		log.With(homo.OTAKeyStep, homo.OTAUpdating).Info("app is updating")
	}

	cur, old, err := m.loadAPPConfig(target)
	if err != nil {
		log.With(homo.OTAKeyStep, homo.OTARollingBack).Errorw("failed to reload config", zap.Error(err))
		rberr := m.rollBackAPP()
		if rberr != nil {
			log.With(homo.OTAKeyStep, homo.OTAFailure).Errorw("failed to roll back", zap.Error(rberr))
			return fmt.Errorf("failed to reload config: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		log.With(homo.OTAKeyStep, homo.OTARolledBack).Infof("app is rolled back")
		return fmt.Errorf("failed to reload config: %s", err.Error())
	}

	// prepare services
	keepServices := diffServices(cur, old)
	m.engine.Prepare(cur)

	// stop all removed or updated services
	m.stopServices(keepServices)
	// start all updated or added services
	err = m.startServices(cur)
	if err != nil {
		log.With(homo.OTAKeyStep, homo.OTARollingBack).Errorw("failed to start app", zap.Error(err))
		rberr := m.rollBackAPP()
		if rberr != nil {
			log.With(homo.OTAKeyStep, homo.OTAFailure).Errorw("failed to roll back", zap.Error(rberr))
			return fmt.Errorf("failed to start app: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		// stop all updated or added services
		m.stopServices(keepServices)
		// start all removed or updated services
		rberr = m.startServices(old)
		if rberr != nil {
			log.With(homo.OTAKeyStep, homo.OTAFailure).Errorw("failed to roll back", zap.Error(rberr))
			return fmt.Errorf("failed to restart old app: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		m.commitAPP(old.AppVersion)
		log.With(homo.OTAKeyStep, homo.OTARolledBack).Info("app is rolled back")
		return fmt.Errorf("failed to start app: %s", err.Error())
	}
	m.commitAPP(cur.AppVersion)
	if isOTA {
		log.With(homo.OTAKeyStep, homo.OTAUpdated).Info("app is updated")
	}
	return nil
}

func (m *Master) loadAPPConfig(target string) (cur, old homo.ComposeAppConfig, err error) {
	if target != "" {
		// backup
		if utils.IsFile(appConfigFile) {
			// application.yml --> application.yml.old
			err = os.Rename(appConfigFile, appBackupFile)
			if err != nil {
				return
			}
		} else {
			// none --> application.yml.old (empty)
			var f *os.File
			f, err = os.Create(appBackupFile)
			if err != nil {
				return
			}
			f.Close()
		}

		if utils.IsFile(target) {
			// copy {target} to application.yml
			err = utils.CopyFile(target, appConfigFile)
		} else {
			// copy {target}/application.yml to application.yml
			err = utils.CopyFile(path.Join(target, homo.AppConfFileName), appConfigFile)
		}
		if err != nil {
			return
		}
	}
	if utils.IsFile(appConfigFile) {
		cur, err = homo.LoadComposeAppConfigCompatible(appConfigFile)
		if err != nil {
			return
		}
	}
	if utils.IsFile(appBackupFile) {
		old, err = homo.LoadComposeAppConfigCompatible(appBackupFile)
		if err != nil {
			return
		}
	}
	return
}

func (m *Master) rollBackAPP() error {
	if !utils.IsFile(appBackupFile) {
		return nil
	}
	// application.yml.old --> application.yml
	return os.Rename(appBackupFile, appConfigFile)
}

func (m *Master) commitAPP(ver string) {
	defer m.log.Infof("app version (%s) committed", ver)

	// update config version
	m.infostats.setVersion(ver)
	// remove application.yml.old
	err := os.RemoveAll(appBackupFile)
	if err != nil {
		logger.S.Errorw(fmt.Sprintf("failed to remove backup file (%s)", appBackupFile), zap.Error(err))
	}
}

// UpdateMST updates master
func (m *Master) UpdateMST(trace, target, backup string) (err error) {
	log := logger.New(m.cfg.OTALog, homo.OTAKeyTrace, trace, homo.OTAKeyType, homo.OTAMST)

	if err = m.check(target); err != nil {
		log.With(homo.OTAKeyStep, homo.OTAFailure).Errorw("failed to check master", zap.Error(err))
		return fmt.Errorf("failed to check master: %s", err.Error())
	}

	log.With(homo.OTAKeyStep, homo.OTAUpdating).Info("master is updating")
	if err = apply(target, backup); err != nil {
		log.With(homo.OTAKeyStep, homo.OTARollingBack).Errorw("failed to apply master", zap.Error(err))
		rberr := RollBackMST()
		if rberr != nil {
			log.With(homo.OTAKeyStep, homo.OTAFailure).Errorw("failed to roll back", zap.Error(rberr))
			return fmt.Errorf("failed to apply master: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		log.With(homo.OTAKeyStep, homo.OTARolledBack).Info("master is rolled back")
		return fmt.Errorf("failed to apply master: %s", err.Error())
	}

	log.With(homo.OTAKeyStep, homo.OTARestarting).Info("master is restarting")
	return m.Close()
}

// RollBackMST rolls back master
func RollBackMST() error {
	// backward compatibility
	backup := homo.DefaultBinBackupFile
	if !utils.FileExists(backup) {
		if !utils.FileExists(homo.PreviousBinBackupFile) {
			return nil
		} else {
			backup = homo.PreviousBinBackupFile
		}
	}
	err := apply(backup, "")
	if err != nil {
		logger.S.Errorw("failed to apply backup master", zap.Error(err))
	}
	err = os.RemoveAll(backup)
	if err != nil {
		logger.S.Errorw(fmt.Sprintf("failed to remove backup file (%s)", backup), zap.Error(err))
	}
	return nil
}

func apply(target, backup string) error {
	f, err := os.Open(target)
	if err != nil {
		return fmt.Errorf("failed to open binary: %s", err.Error())
	}
	defer f.Close()
	err = update.Apply(f, update.Options{OldSavePath: backup})
	if err != nil {
		return fmt.Errorf("failed to apply binary: %s", err.Error())
	}
	return nil
}

func (m *Master) check(target string) error {
	m.log.Debugf("new binary: %s", target)
	os.Chmod(target, 0755)
	cmd := exec.Command(target, "check", "-w", m.pwd, "-c", m.cfg.File)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("check result: %s", err.Error())
	}
	if !strings.Contains(string(out), homo.CheckOK) {
		return fmt.Errorf("check result: OK expected, but get %s", string(out))
	}
	return nil
}
