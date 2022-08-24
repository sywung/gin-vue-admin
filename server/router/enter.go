package router

import (
	"github.com/sywung/gin-vue-admin/server/router/example"
	"github.com/sywung/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
