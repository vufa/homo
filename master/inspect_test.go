package master

import (
	"path"
	"sync"
	"testing"

	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	info := infoStats{
		Inspect:  aiicy.Inspect{},
		services: nil,
		file:     path.Join("testdata", "application_docker.stats"),
		RWMutex:  sync.RWMutex{},
	}
	sss := map[string]map[string]interface{}{}
	b := info.LoadStats(&sss)
	assert.Equal(t, true, b)
}
