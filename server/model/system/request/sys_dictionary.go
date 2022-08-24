package request

import (
	"github.com/sywung/gin-vue-admin/server/model/common/request"
	"github.com/sywung/gin-vue-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
