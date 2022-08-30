package port

import "github.com/ce-final-project/backend_rest_api/authentication_service/internal/core/domain"

type AuthRepository interface {
	GetAllSession() ([]domain.AuthSession, error)
	GetOneSession(sessionID string) (*domain.AuthSession, error)
	StoreSession(session *domain.AuthSession) (*domain.AuthSession, error)
	UpdateSession(session *domain.AuthSession) (*domain.AuthSession, error)
	DeletedSession(sessionID string) (*domain.AuthSession, error)
}

type AuthService interface {
	GetAllSession() ([]domain.AuthResponse, error)
	GetOneSession(sessionID string) (*domain.AuthResponse, error)
	CreateSession(session *domain.AuthRequest) (*domain.AuthRequest, error)
	Register(resister *domain.ResisterPayload) (*domain.AuthResponse, error)
	Login(login *domain.LoginPayload) (*domain.AuthResponse, error)
	Logout() error
}
