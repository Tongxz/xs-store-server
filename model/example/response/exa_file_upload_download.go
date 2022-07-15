package response

import "github.com/tongxz/xs-admin-vue/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
