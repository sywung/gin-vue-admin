package service

import (
	"github.com/sywung/gin-vue-admin/server/service/example"
	"github.com/sywung/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
