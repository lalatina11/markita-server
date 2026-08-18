package response

import "github.com/lalatina11/markita.git/src/model"

type AuthSuccessResult struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         UserData `json:"user"`
}

func (this *AuthSuccessResult) IsSuccess() bool {
	return this.AccessToken != ""
}

type UserData struct {
	ID           string       `json:"id"`
	UserMetadata UserMetadata `json:"user_metadata"`
}

type UserMetadata struct {
	DisplayName string         `json:"display_name"`
	Email       string         `json:"email"`
	Role        model.UserRole `json:"role"`
}

type User struct {
	ID          string         `json:"id"`
	DisplayName string         `json:"display_name"`
	Email       string         `json:"email"`
	Role        model.UserRole `json:"role"`
}

type AuthSuccessPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type AuthUserPayload struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         *model.User `json:"user"`
}

func (this *AuthSuccessPayload) ToAuthUserPayload(User *model.User) *AuthUserPayload {
	return &AuthUserPayload{
		AccessToken:  this.AccessToken,
		RefreshToken: this.RefreshToken,
		User:         User,
	}
}

func (this *AuthSuccessResult) ToPayload() *AuthSuccessPayload {
	return &AuthSuccessPayload{
		AccessToken:  this.AccessToken,
		RefreshToken: this.RefreshToken,
		User: User{
			ID:          this.User.ID,
			DisplayName: this.User.UserMetadata.DisplayName,
			Email:       this.User.UserMetadata.Email,
			Role:        this.User.UserMetadata.Role,
		},
	}
}

type AuthErrorResult struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}

type AuthGetUserSuccessResponse struct {
	ID string `json:"id"`
}

func (this *AuthGetUserSuccessResponse) IsGetUserSuccess() bool {
	return this.ID != ""
}
