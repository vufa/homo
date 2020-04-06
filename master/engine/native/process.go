package native

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/aiicy/aiicy-go/logger"

	"github.com/aiicy/aiicy/master/engine"
	"github.com/aiicy/aiicy/utils"
	"github.com/shirou/gopsutil/process"
)

type processConfigs struct {
	exec string
	pwd  string
	argv []string
	env  []string
}

func (e *nativeEngine) startProcess(cfg processConfigs) (*os.Process, error) {
	err := os.Chmod(cfg.exec, os.ModePerm)
	if err != nil {
		e.log.Warn(fmt.Sprintf("chmod exec %s to %s failed", cfg.exec, os.ModePerm), logger.Error(err))
		return nil, err
	}
	p, err := os.StartProcess(
		cfg.exec,
		cfg.argv,
		&os.ProcAttr{
			Dir: cfg.pwd,
			Env: cfg.env,
			Files: []*os.File{
				os.Stdin,
				os.Stdout,
				os.Stderr,
			},
		},
	)
	if err != nil {
		e.log.Warn("failed to start process", logger.Error(err))
		return nil, err
	}
	e.log.Debugf("process (%d) started", p.Pid)
	return p, nil
}

func (e *nativeEngine) waitProcess(p *os.Process) error {
	ps, err := p.Wait()
	if err != nil {
		e.log.Warn(fmt.Sprintf("failed to wait process (%d)", p.Pid), logger.Error(err))
		return err
	}
	e.log.Debugf("process (%d) %s", p.Pid, ps.String())
	if !ps.Success() {
		return fmt.Errorf("process exit code: %s", ps.String())
	}
	return nil
}

func (e *nativeEngine) stopProcess(p *os.Process) error {
	e.log.Debugf("process (%d) is stopping", p.Pid)

	err := p.Signal(syscall.SIGTERM)
	if err != nil {
		e.log.Debugf("failed to stop process (%d): %s", p.Pid, err.Error())
		return nil
	}

	done := make(chan error, 1)
	go func() {
		_, err := p.Wait()
		done <- err
	}()
	select {
	case <-time.After(e.grace):
		e.log.Warnf("timed out to wait process (%d)", p.Pid)
		err = p.Kill()
		if err != nil {
			e.log.Warn(fmt.Sprintf("failed to kill process (%d)", p.Pid), logger.Error(err))
		}
		return fmt.Errorf("timed out to wait process (%d)", p.Pid)
	case err := <-done:
		if err != nil {
			e.log.Debugf("failed to wait process (%d): %s", p.Pid, err.Error())
		}
		return nil
	}
}

func (e *nativeEngine) statsProcess(p *os.Process) engine.PartialStats {
	proc, err := process.NewProcess(int32(p.Pid))
	if err != nil {
		return engine.PartialStats{"error": err.Error()}
	}
	cpu := utils.CPUInfo{Time: time.Now().UTC()}
	cpu.UsedPercent, err = proc.CPUPercent()
	if err != nil {
		cpu.Error = err.Error()
	}
	mem := utils.MemInfo{Time: time.Now().UTC()}
	meminfo, err := proc.MemoryInfo()
	if err != nil {
		mem.Error = err.Error()
	} else {
		mem.Used = meminfo.RSS
		mem.SwapUsed = meminfo.Swap
		mup, err := proc.MemoryPercent()
		if err != nil {
			mem.Error = err.Error()
		} else {
			mem.UsedPercent = float64(mup)
		}
	}
	return engine.PartialStats{
		"cpu_stats": cpu,
		"mem_stats": mem,
	}
}
