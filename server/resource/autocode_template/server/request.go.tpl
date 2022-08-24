package request

import (
	"github.com/sywung/gin-vue-admin/server/model/{{.Package}}"
	"github.com/sywung/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
