package response

import "github.com/sywung/gin-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
