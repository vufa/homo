package master

import (
	"reflect"
	"testing"

	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/aiicy/aiicy/utils"
	"github.com/stretchr/testify/assert"
)

var cfgV1 = `
version: V1
services:
  - name: a
    image: 'aiicy-a:latest'
    replica: 1
    mounts:
      - name: a-conf-V1
        path: etc/aiicy
        readonly: true
  - name: b
    image: 'aiicy-b:latest'
    replica: 1
    mounts:
      - name: b-conf-V1
        path: etc/aiicy
        readonly: true
  - name: c
    image: 'aiicy-c:latest'
    replica: 1
    mounts:
      - name: c-conf-V1
        path: etc/aiicy
        readonly: true
volumes:
  - name: a-conf-V1
    path: a-conf/V1
  - name: b-conf-V1
    path: b-conf/V1
  - name: c-conf-V1
    path: c-conf/V1
`

var cfgV2 = `
version: V2
services:
  - name: a
    image: 'aiicy-a:latest'
    replica: 1
    mounts:
      - name: a-conf-V1
        path: etc/aiicy
        readonly: true
  - name: b
    image: 'aiicy-b:0.1.4'
    replica: 1
    mounts:
      - name: b-conf-V1
        path: etc/aiicy
        readonly: true
  - name: d
    image: 'aiicy-d:latest'
    replica: 1
    mounts:
      - name: d-conf-V1
        path: etc/aiicy
        readonly: true
volumes:
  - name: a-conf-V1
    path: a-conf/V1
  - name: b-conf-V1
    path: b-conf/V1
  - name: d-conf-V1
    path: d-conf/V1
`

var cfgV3 = `
version: V3
services:
  - name: a
    image: 'aiicy-a:latest'
    replica: 0
    mounts:
      - name: a-conf-V1
        path: etc/aiicy
        readonly: true
  - name: b
    image: 'aiicy-b:0.1.4'
    replica: 1
    mounts:
      - name: b-conf-V1
        path: etc/aiicy
        readonly: true
      - name: b-data-V1
        path: var/db/aiicy/data
  - name: d
    image: 'aiicy-d:latest'
    replica: 1
    mounts:
      - name: d-conf-V1
        path: etc/aiicy
        readonly: true
volumes:
  - name: a-conf-V1
    path: a-conf/V1
  - name: b-conf-V1
    path: b-conf/V1
  - name: d-conf-V1
    path: d-conf/V22
`

var cfgV4 = `
version: V4
`

func Test_diffServices(t *testing.T) {
	var V1 aiicy.AppConfig
	err := utils.UnmarshalYAML([]byte(cfgV1), &V1)
	assert.NoError(t, err)

	var V2 aiicy.AppConfig
	err = utils.UnmarshalYAML([]byte(cfgV2), &V2)
	assert.NoError(t, err)

	var V3 aiicy.AppConfig
	err = utils.UnmarshalYAML([]byte(cfgV3), &V3)
	assert.NoError(t, err)

	var V4 aiicy.AppConfig
	err = utils.UnmarshalYAML([]byte(cfgV4), &V4)
	assert.NoError(t, err)

	type args struct {
		cur aiicy.AppConfig
		old aiicy.AppConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "a,b,c --> a,b',d",
			args: args{
				cur: V2,
				old: V1,
			},
			want: map[string]struct{}{
				"a": {},
			},
		},
		{
			name: "a,b,d --> a',b',d'",
			args: args{
				cur: V3,
				old: V2,
			},
			want: map[string]struct{}{},
		},
		{
			name: "a,b,d --> nil",
			args: args{
				cur: V4,
				old: V3,
			},
			want: map[string]struct{}{},
		},
		{
			name: "nil --> a,b,d",
			args: args{
				cur: V3,
				old: V4,
			},
			want: map[string]struct{}{},
		},
		{
			name: "a,b,d --> a,b,d",
			args: args{
				cur: V3,
				old: V3,
			},
			want: map[string]struct{}{
				"a": {},
				"b": {},
				"d": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccur := tt.args.cur.ToComposeAppConfig()
			cold := tt.args.old.ToComposeAppConfig()
			if got := diffServices(ccur, cold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceSort(t *testing.T) {
	services := map[string]aiicy.ComposeService{}
	services["a"] = aiicy.ComposeService{
		DependsOn: []string{},
	}
	services["b"] = aiicy.ComposeService{
		DependsOn: []string{"a"},
	}
	services["c"] = aiicy.ComposeService{
		DependsOn: []string{"a", "b"},
	}
	services["d"] = aiicy.ComposeService{
		DependsOn: []string{"b", "c"},
	}
	services["e"] = aiicy.ComposeService{
		DependsOn: []string{"c", "a", "b"},
	}
	services["f"] = aiicy.ComposeService{
		DependsOn: []string{"b", "c"},
	}
	services["h"] = aiicy.ComposeService{
		DependsOn: []string{"d", "f"},
	}
	order := ServiceSort(services)
	om := map[string]int{}
	for i, o := range order {
		om[o] = i
	}
	// order of depended services are less than the service
	assert.True(t, om["a"] < om["b"])
	assert.True(t, om["a"] < om["c"])
	assert.True(t, om["b"] < om["c"])
	assert.True(t, om["b"] < om["d"])
	assert.True(t, om["c"] < om["d"])
	assert.True(t, om["a"] < om["e"])
	assert.True(t, om["b"] < om["e"])
	assert.True(t, om["c"] < om["e"])
	assert.True(t, om["b"] < om["f"])
	assert.True(t, om["c"] < om["f"])
	assert.True(t, om["d"] < om["h"])
	assert.True(t, om["f"] < om["h"])
}
