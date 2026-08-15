package response

// AuthResult represents the complete authentication response
type AuthSuccessResult struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

// JWTClaims represents the decoded JWT payload
type JWTClaims struct {
	Iss          string                 `json:"iss"`
	Sub          string                 `json:"sub"`
	Aud          string                 `json:"aud"`
	Exp          int64                  `json:"exp"`
	Iat          int64                  `json:"iat"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	AppMetadata  AppMetadata            `json:"app_metadata"`
	UserMetadata UserMetadata           `json:"user_metadata"`
	Role         string                 `json:"role"`
	AAL          string                 `json:"aal"`
	AMR          []AuthenticationMethod `json:"amr"`
	SessionID    string                 `json:"session_id"`
	IsAnonymous  bool                   `json:"is_anonymous"`
}

// AppMetadata contains application-level metadata
type AppMetadata struct {
	Provider  string   `json:"provider"`
	Providers []string `json:"providers"`
}

// UserMetadata contains user-specific metadata
type UserMetadata struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PhoneVerified bool   `json:"phone_verified"`
	Sub           string `json:"sub"`
}

// AuthenticationMethod represents how the user authenticated
type AuthenticationMethod struct {
	Method    string `json:"method"`
	Timestamp int64  `json:"timestamp"`
}

// User represents the authenticated user
type User struct {
	ID    string `json:"id"`
	Aud   string `json:"aud"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

// TokenResponse is a helper for token refresh responses
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}

// AuthRequest represents login credentials
type SignInPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RefreshTokenRequest represents a token refresh request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type AuthErrorResponse struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}
