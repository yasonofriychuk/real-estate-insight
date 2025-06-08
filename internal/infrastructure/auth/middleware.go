package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"net/http"
	"strings"
)

type jwtService interface {
	ProfileIdFromToken(tokenString string) *uuid.UUID
}

type Auth struct {
	jwtService             jwtService
	next                   *api.Server
	excludeOperations      []string
	unauthorizedErrorBytes []byte
}

func MustNewMiddleware(next *api.Server, jwtService jwtService, excludeOperations ...string) *Auth {
	authErr := &api.Error{
		Status: api.ErrorStatusUnauthorized,
		Error: api.ErrorError{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		},
	}
	unauthorizedErrorBytes, err := authErr.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return &Auth{
		next:       next,
		jwtService: jwtService,
		excludeOperations: lo.Map(excludeOperations, func(op string, _ int) string {
			return strings.ToLower(op)
		}),
		unauthorizedErrorBytes: unauthorizedErrorBytes,
	}
}

func (m Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route, ok := m.next.FindRoute(r.Method, r.URL.Path)
	if !ok {
		m.next.ServeHTTP(w, r)
		return
	}

	if lo.Contains(m.excludeOperations, strings.ToLower(route.OperationID())) {
		m.next.ServeHTTP(w, r)
		return
	}

	if len(r.Header.Values("Authorization")) == 0 {
		_, _ = w.Write(m.unauthorizedErrorBytes)
		return
	}

	profileId := m.jwtService.ProfileIdFromToken(strings.ReplaceAll(r.Header.Values("Authorization")[0], "Bearer ", ""))
	if profileId == nil {
		_, _ = w.Write(m.unauthorizedErrorBytes)
		return
	}

	m.next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "profileId", *profileId)))
}

func ProfileIdFromCtx(ctx context.Context) *uuid.UUID {
	id, ok := ctx.Value("profileId").(uuid.UUID)
	if !ok {
		return nil
	}
	return &id
}
