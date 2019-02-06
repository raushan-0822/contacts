package middlewares

import (
	contract "contacts/contracts"
	"contacts/error"
	"contacts/model"
	"contacts/service/handlers"
	"contacts/utils"

	"github.com/labstack/echo"
)

// BasicAuth provides BasicAuth
func BasicAuth(accountID, password string, c echo.Context) (bool, error) {
	account, err := model.GetAccountByID(accountID)
	if err != nil {
		response := new(contract.ContactsResponse)
		response.SetErrorData(errorpkg.ErrInternalServerError("Error while Basic Auth"))
		response.SetHTTPCode(response.BaseResponse.HTTPCode)
		response.SetRequestID(c.Get("RequestID").(string))
		response.SetMethod(c.Get("Method").(string))
		return false, handler.RawResponse(c, response, response.BaseResponse.HTTPCode)
	} else if utils.Hash(password) == account.Password {
		return true, nil
	}
	return false, nil
}
