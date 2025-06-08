package profile_login

import (
	"context"
	"github.com/google/uuid"
)

type jwtAuth interface {
	CreatePermanentToken(profileId uuid.UUID) (string, error)
}

type profileStorage interface {
	Login(ctx context.Context, username string, password string) (*uuid.UUID, error)
}
