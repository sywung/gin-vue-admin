package global

{{- if .HasGlobal }}

import "github.com/sywung/gin-vue-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}