package service

 {{- if .NeedModel }}
import (
   "github.com/tongxz/xs-admin-vue/plugin/{{ .Snake}}/model"
)
{{ end }}

type {{ .PlugName}}Service struct{}

func (e *{{ .PlugName}}Service) PlugService({{- if .HasRequest }}req model.Request {{ end -}}) ({{- if .HasResponse }}res model.Response,{{ end -}} err error) {
    // 写你的业务逻辑
	return nil{{- if .HasResponse }},res {{ end }}
}
