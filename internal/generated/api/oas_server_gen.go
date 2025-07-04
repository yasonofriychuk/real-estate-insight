// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddToFavoriteSelection implements addToFavoriteSelection operation.
	//
	// Add or remove a development to/from the selected user's favorite selection.
	//
	// POST /selection/favorite
	AddToFavoriteSelection(ctx context.Context, req *AddToFavoriteSelectionReq) (AddToFavoriteSelectionRes, error)
	// BuildRoutesByPoints implements buildRoutesByPoints operation.
	//
	// Build a route between points.
	//
	// GET /routes/build/points
	BuildRoutesByPoints(ctx context.Context, params BuildRoutesByPointsParams) (BuildRoutesByPointsRes, error)
	// CreateSelection implements createSelection operation.
	//
	// Create a new selection for the user with name, comment, and form.
	//
	// POST /selection/create
	CreateSelection(ctx context.Context, req *CreateSelectionReq) (CreateSelectionRes, error)
	// DeleteSelection implements deleteSelection operation.
	//
	// Delete a selection for the user by selection ID.
	//
	// POST /selection/delete
	DeleteSelection(ctx context.Context, params DeleteSelectionParams) (DeleteSelectionRes, error)
	// DevelopmentSearch implements developmentSearch operation.
	//
	// POST /developments/search
	DevelopmentSearch(ctx context.Context, req *DevelopmentSearchReq) (DevelopmentSearchRes, error)
	// EditSelection implements editSelection operation.
	//
	// Edit new selection.
	//
	// POST /selection/edit
	EditSelection(ctx context.Context, req *EditSelectionReq) (EditSelectionRes, error)
	// GenerateInfrastructureHeatmap implements generateInfrastructureHeatmap operation.
	//
	// Returns a grid-based heatmap for infrastructure objects based on type weights within a selected
	// bounding box.
	//
	// POST /infrastructure/heatmap
	GenerateInfrastructureHeatmap(ctx context.Context, req *GenerateInfrastructureHeatmapReq) (GenerateInfrastructureHeatmapRes, error)
	// InfrastructureRadiusBoard implements infrastructureRadiusBoard operation.
	//
	// Search for infrastructure around the selected residential complex.
	//
	// GET /infrastructure/radius
	InfrastructureRadiusBoard(ctx context.Context, params InfrastructureRadiusBoardParams) (InfrastructureRadiusBoardRes, error)
	// LocationList implements locationList operation.
	//
	// Get location list.
	//
	// GET /location/list
	LocationList(ctx context.Context) (LocationListRes, error)
	// SelectionById implements selectionById operation.
	//
	// Get selection.
	//
	// GET /selection/{selectionId}
	SelectionById(ctx context.Context, params SelectionByIdParams) (SelectionByIdRes, error)
	// SelectionList implements selectionList operation.
	//
	// Get selection list.
	//
	// GET /selection/list
	SelectionList(ctx context.Context) (SelectionListRes, error)
	// UserLogin implements userLogin operation.
	//
	// Authenticate the user using email and password.
	//
	// POST /profile/login
	UserLogin(ctx context.Context, req *UserLoginReq) (UserLoginRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
