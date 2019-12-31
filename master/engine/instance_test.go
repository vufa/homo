package engine

import (
	"testing"

	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	static := []string{"name=homo", "org=linux", homo.EnvKeyServiceToken + "=key", homo.EnvServiceTokenKey + "=key"}
	dyn := map[string]string{
		"repo":    "github",
		"project": "homo",
	}
	envs := GenerateInstanceEnv(t.Name(), static, dyn)
	assert.Contains(t, envs, "name=homo")
	assert.Contains(t, envs, "org=linux")
	assert.Contains(t, envs, "repo=github")
	assert.Contains(t, envs, "project=homo")
	assert.Contains(t, envs, homo.EnvKeyServiceInstanceName+"="+t.Name())
	assert.NotContains(t, envs, homo.EnvKeyServiceToken+"=key")
	assert.NotContains(t, envs, homo.EnvServiceTokenKey+"=key")
}
