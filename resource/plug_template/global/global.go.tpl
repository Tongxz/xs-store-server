package global

{{- if .HasGlobal }}

import "github.com/tongxz/xs-admin-vue/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}