package response

import "github.com/sywung/gin-vue-admin/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
