package contract

import (
	"time"

	"github.com/raushan-0822/contacts/model"
)

//Contact defines the structure to store details of contact resource
type Contact struct {
	ID          uint      `form:"id" query:"id" json:"id,omitempty"`
	DateCreated time.Time `form:"date_created" query:"date_created" json:"date_created,omitempty"`
	DateUpdated time.Time `form:"date_updated" query:"date_updated" json:"date_updated,omitempty"`
	AccountID   string    `form:"account_id" query:"account_id" json:"account_id,omitempty"`
	Number      string    `form:"number" query:"number" json:"number,omitempty"`
	FirstName   string    `form:"first_name" query:"first_name" json:"first_name,omitempty"`
	LastName    string    `form:"last_name" query:"last_name" json:"last_name,omitempty"`
	CompanyName string    `form:"company_name" query:"company_name" json:"company_name,omitempty"`
	Email       string    `form:"email" query:"email" json:"email,omitempty"`
	URI         string    `form:"uri" query:"uri" json:"uri,omitempty"`
}

// ContactsResponse is contract for crud response
type ContactsResponse struct {
	BaseResponse
	ResourceData *Contact `json:"data,omitempty"`
}

// ContactsSearchResponse is ContactsSearch response
type ContactsSearchResponse struct {
	BaseResponse
	ResourceData *[]model.Contact `json:"data,omitempty"`
}
