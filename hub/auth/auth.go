package auth

import (
	"github.com/256dpi/gomqtt/topic"
	"github.com/countstarlight/homo/hub/config"
	"github.com/countstarlight/homo/utils"
)

// all permit actions
const (
	Publish   = "pub"
	Subscribe = "sub"
)

// Auth auth
type Auth struct {
	// for client certs
	certs map[string]cert
	// for client account
	accounts map[string]account
}

// NewAuth creates auth
func NewAuth(principals []config.Principal) *Auth {
	_certs := make(map[string]cert)
	_accounts := make(map[string]account)
	for _, principal := range principals {
		authorizer := NewAuthorizer()
		for _, p := range duplicatePubSubPermitRemove(principal.Permissions) {
			for _, topic := range p.Permits {
				authorizer.Add(topic, p.Action)
			}
		}
		if principal.Password == "" {
			_certs[principal.Username] = cert{
				Authorizer: authorizer,
			}
		} else {
			_accounts[principal.Username] = account{
				Password:   principal.Password,
				Authorizer: authorizer,
			}
		}
	}
	return &Auth{certs: _certs, accounts: _accounts}
}

func duplicatePubSubPermitRemove(permission []config.Permission) []config.Permission {
	PubPermitList := make(map[string]struct{})
	SubPermitList := make(map[string]struct{})
	for _, _permission := range permission {
		switch _permission.Action {
		case Publish:
			for _, v := range _permission.Permits {
				PubPermitList[v] = struct{}{}
			}
		case Subscribe:
			for _, v := range _permission.Permits {
				SubPermitList[v] = struct{}{}
			}
		}
	}
	return []config.Permission{
		{Action: Publish, Permits: utils.GetKeys(PubPermitList)},
		{Action: Subscribe, Permits: utils.GetKeys(SubPermitList)},
	}
}

type account struct {
	Password   string
	Authorizer *Authorizer
}

type cert struct {
	Authorizer *Authorizer
}

// Authorizer checks topic permission
type Authorizer struct {
	*topic.Tree
}

// NewAuthorizer create a new authorizer
func NewAuthorizer() *Authorizer {
	return &Authorizer{Tree: topic.NewStandardTree()}
}
