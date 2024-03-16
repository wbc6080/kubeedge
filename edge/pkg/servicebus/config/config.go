package config

import (
	"sync"

	"github.com/kubeedge/kubeedge/staging/src/github.com/kubeedge/api/componentconfig/edgecore/v1alpha2"
)

var Config Configure
var once sync.Once

type Configure struct {
	v1alpha2.ServiceBus
}

func InitConfigure(s *v1alpha2.ServiceBus) {
	once.Do(func() {
		Config = Configure{
			ServiceBus: *s,
		}
	})
}
