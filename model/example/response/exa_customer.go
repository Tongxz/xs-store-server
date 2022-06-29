package response

import "github.com/tongxz/xs-admin-vue/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
