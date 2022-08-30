package domain

type AuthSession struct {
	SessionID    string `json:"session_id" bson:"session_id"`
	AccountID    uint64 `json:"account_id" bson:"account_id"`
	IP           string `json:"ip" bson:"ip"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	Service      string `json:"service" bson:"service"`
	CreatedAt    int64  `json:"created_at" bson:"created_at" db:"created_at"`
}

type AuthRequest struct {
	JWTToken     string `json:"jwt_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	AccountID int64    `json:"account_id"`
	PlayerID  string   `json:"player_id"`
	Username  string   `json:"username"`
	Roles     []string `json:"roles"`
}

type ResisterPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
