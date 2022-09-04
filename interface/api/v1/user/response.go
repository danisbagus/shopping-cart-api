package user

import (
	"github.com/danisbagus/shopping-cart-api/core/domain"
	"github.com/danisbagus/shopping-cart-api/interface/api/common"
)

type (
	ResponseRegister struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		RoleID      uint64 `json:"role_id"`
		AccessToken string `json:"access_token"`
	}
)

func NewResponseUserAuth(message string, userAuth *domain.UserAuth) *common.DefaultResponse {

	var data ResponseRegister
	data.ID = userAuth.ID
	data.Name = userAuth.Name
	data.Email = userAuth.Email
	data.RoleID = userAuth.RoleID
	data.AccessToken = userAuth.AccessToken

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, data)
	return responseData
}
