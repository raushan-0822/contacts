package contract

type (
	//ContactPostRequest is contract for post request validations
	ContactPostRequest struct {
		Number      string `form:"number" query:"number" json:"number" validate:"required"`
		FirstName   string `form:"first_name" query:"first_name" json:"first_name" validate:"required"`
		LastName    string `form:"last_name" query:"last_name" json:"last_name" validate:"required"`
		CompanyName string `form:"company_name" query:"company_name" json:"company_name" validate:"required"`
		Email       string `form:"email" query:"email" json:"email" validate:"required,email"`
	}

	//ContactUpdateRequest  is contract for put request validations
	ContactUpdateRequest struct {
		Number      string `form:"number" query:"number" json:"number"`
		FirstName   string `form:"first_name" query:"first_name" json:"first_name"`
		LastName    string `form:"last_name" query:"last_name" json:"last_name"`
		CompanyName string `form:"company_name" query:"company_name" json:"company_name"`
		Email       string `form:"email" query:"email" json:"email" validate:"required"`
	}
)
