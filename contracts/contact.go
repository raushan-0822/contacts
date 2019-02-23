package contract

import (
	"github.com/raushan-0822/contacts/model"
)

// ContactsResponse is contract for crud response
type ContactsResponse struct {
	BaseResponse
	ResourceData *model.Contact `json:"data,omitempty"`
}

// ContactsSearchResponse is ContactsSearch response
type ContactsSearchResponse struct {
	BaseResponse
	ResourceData *[]model.Contact `json:"data,omitempty"`
}
