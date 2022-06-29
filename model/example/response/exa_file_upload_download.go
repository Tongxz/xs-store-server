package response

import "github.com/tongxz/xs-admin-vue/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
