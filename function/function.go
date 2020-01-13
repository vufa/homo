package main

import (
	"github.com/countstarlight/homo/utils"
	pool "github.com/jolestar/go-commons-pool"
	"go.uber.org/zap"
)

// Function function
type Function struct {
	p    Producer
	cfg  FunctionInfo
	ids  chan uint32
	pool *pool.ObjectPool
	log  *zap.SugaredLogger
	tomb utils.Tomb
}

// NewFunction creates a new function
func NewFunction(cfg FunctionInfo, p Producer) *Function {
	f := &Function{
		p:   p,
		cfg: cfg,
		ids: make(chan uint32, cfg.Instance.Max),
		log: logger.WithField("function", cfg.Name),
	}
	for index := 1; index <= cfg.Instance.Max; index++ {
		f.ids <- uint32(index)
	}
	pc := pool.NewDefaultPoolConfig()
	pc.MinIdle = cfg.Instance.Min
	pc.MaxIdle = cfg.Instance.Max
	pc.MaxTotal = cfg.Instance.Max
	pc.MinEvictableIdleTime = cfg.Instance.IdleTime
	pc.TimeBetweenEvictionRuns = cfg.Instance.EvictTime
	f.pool = pool.NewObjectPool(context.Background(), f, pc)
	return f
}
