package master

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/aiicy/aiicy/utils"
	"github.com/inconshreveable/go-update"
)

// TODO: need test
var appDir = path.Join("var", "db", "aiicy")
var appConfigFile = path.Join(appDir, aiicy.AppConfFileName)
var appBackupFile = path.Join(appDir, aiicy.AppBackupFileName)

// UpdateSystem updates application or master
func (m *Master) UpdateSystem(trace, tp, target string) (err error) {
	switch tp {
	case aiicy.OTAMST:
		err = m.UpdateMST(trace, target, aiicy.DefaultBinBackupFile)
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
		log = logger.New(m.cfg.OTALog, aiicy.OTAKeyTrace, trace, aiicy.OTAKeyType, aiicy.OTAAPP)
		log.With(aiicy.OTAKeyStep, aiicy.OTAUpdating).Info("app is updating")
	}

	cur, old, err := m.loadAPPConfig(target)
	if err != nil {
		log.With(aiicy.OTAKeyStep, aiicy.OTARollingBack).Errorw("failed to reload config", logger.Error(err))
		rberr := m.rollBackAPP()
		if rberr != nil {
			log.With(aiicy.OTAKeyStep, aiicy.OTAFailure).Errorw("failed to roll back", logger.Error(rberr))
			return fmt.Errorf("failed to reload config: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		log.With(aiicy.OTAKeyStep, aiicy.OTARolledBack).Infof("app is rolled back")
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
		log.With(aiicy.OTAKeyStep, aiicy.OTARollingBack).Errorw("failed to start app", logger.Error(err))
		rberr := m.rollBackAPP()
		if rberr != nil {
			log.With(aiicy.OTAKeyStep, aiicy.OTAFailure).Errorw("failed to roll back", logger.Error(rberr))
			return fmt.Errorf("failed to start app: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		// stop all updated or added services
		m.stopServices(keepServices)
		// start all removed or updated services
		rberr = m.startServices(old)
		if rberr != nil {
			log.With(aiicy.OTAKeyStep, aiicy.OTAFailure).Errorw("failed to roll back", logger.Error(rberr))
			return fmt.Errorf("failed to restart old app: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		m.commitAPP(old.AppVersion)
		log.With(aiicy.OTAKeyStep, aiicy.OTARolledBack).Info("app is rolled back")
		return fmt.Errorf("failed to start app: %s", err.Error())
	}
	m.commitAPP(cur.AppVersion)
	if isOTA {
		log.With(aiicy.OTAKeyStep, aiicy.OTAUpdated).Info("app is updated")
	}
	return nil
}

func (m *Master) loadAPPConfig(target string) (cur, old aiicy.ComposeAppConfig, err error) {
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
			err = utils.CopyFile(path.Join(target, aiicy.AppConfFileName), appConfigFile)
		}
		if err != nil {
			return
		}
	}
	if utils.IsFile(appConfigFile) {
		cur, err = aiicy.LoadComposeAppConfigCompatible(appConfigFile)
		if err != nil {
			return
		}
	}
	if utils.IsFile(appBackupFile) {
		old, err = aiicy.LoadComposeAppConfigCompatible(appBackupFile)
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
		logger.S.Errorw(fmt.Sprintf("failed to remove backup file (%s)", appBackupFile), logger.Error(err))
	}
}

// UpdateMST updates master
func (m *Master) UpdateMST(trace, target, backup string) (err error) {
	log := logger.New(m.cfg.OTALog, aiicy.OTAKeyTrace, trace, aiicy.OTAKeyType, aiicy.OTAMST)

	if err = m.check(target); err != nil {
		log.With(aiicy.OTAKeyStep, aiicy.OTAFailure).Errorw("failed to check master", logger.Error(err))
		return fmt.Errorf("failed to check master: %s", err.Error())
	}

	log.With(aiicy.OTAKeyStep, aiicy.OTAUpdating).Info("master is updating")
	if err = apply(target, backup); err != nil {
		log.With(aiicy.OTAKeyStep, aiicy.OTARollingBack).Errorw("failed to apply master", logger.Error(err))
		rberr := RollBackMST()
		if rberr != nil {
			log.With(aiicy.OTAKeyStep, aiicy.OTAFailure).Errorw("failed to roll back", logger.Error(rberr))
			return fmt.Errorf("failed to apply master: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		log.With(aiicy.OTAKeyStep, aiicy.OTARolledBack).Info("master is rolled back")
		return fmt.Errorf("failed to apply master: %s", err.Error())
	}

	log.With(aiicy.OTAKeyStep, aiicy.OTARestarting).Info("master is restarting")
	return m.Close()
}

// RollBackMST rolls back master
func RollBackMST() error {
	// backward compatibility
	backup := aiicy.DefaultBinBackupFile
	if !utils.FileExists(backup) {
		return nil
	}
	err := apply(backup, "")
	if err != nil {
		logger.S.Errorw("failed to apply backup master", logger.Error(err))
	}
	err = os.RemoveAll(backup)
	if err != nil {
		logger.S.Errorw(fmt.Sprintf("failed to remove backup file (%s)", backup), logger.Error(err))
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
	if !strings.Contains(string(out), aiicy.CheckOK) {
		return fmt.Errorf("check result: OK expected, but get %s", string(out))
	}
	return nil
}
