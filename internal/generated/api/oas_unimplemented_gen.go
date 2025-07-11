// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddToFavoriteSelection implements addToFavoriteSelection operation.
//
// Add or remove a development to/from the selected user's favorite selection.
//
// POST /selection/favorite
func (UnimplementedHandler) AddToFavoriteSelection(ctx context.Context, req *AddToFavoriteSelectionReq) (r AddToFavoriteSelectionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// BuildRoutesByPoints implements buildRoutesByPoints operation.
//
// Build a route between points.
//
// GET /routes/build/points
func (UnimplementedHandler) BuildRoutesByPoints(ctx context.Context, params BuildRoutesByPointsParams) (r BuildRoutesByPointsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSelection implements createSelection operation.
//
// Create a new selection for the user with name, comment, and form.
//
// POST /selection/create
func (UnimplementedHandler) CreateSelection(ctx context.Context, req *CreateSelectionReq) (r CreateSelectionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteSelection implements deleteSelection operation.
//
// Delete a selection for the user by selection ID.
//
// POST /selection/delete
func (UnimplementedHandler) DeleteSelection(ctx context.Context, params DeleteSelectionParams) (r DeleteSelectionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DevelopmentSearch implements developmentSearch operation.
//
// POST /developments/search
func (UnimplementedHandler) DevelopmentSearch(ctx context.Context, req *DevelopmentSearchReq) (r DevelopmentSearchRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EditSelection implements editSelection operation.
//
// Edit new selection.
//
// POST /selection/edit
func (UnimplementedHandler) EditSelection(ctx context.Context, req *EditSelectionReq) (r EditSelectionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GenerateInfrastructureHeatmap implements generateInfrastructureHeatmap operation.
//
// Returns a grid-based heatmap for infrastructure objects based on type weights within a selected
// bounding box.
//
// POST /infrastructure/heatmap
func (UnimplementedHandler) GenerateInfrastructureHeatmap(ctx context.Context, req *GenerateInfrastructureHeatmapReq) (r GenerateInfrastructureHeatmapRes, _ error) {
	return r, ht.ErrNotImplemented
}

// InfrastructureRadiusBoard implements infrastructureRadiusBoard operation.
//
// Search for infrastructure around the selected residential complex.
//
// GET /infrastructure/radius
func (UnimplementedHandler) InfrastructureRadiusBoard(ctx context.Context, params InfrastructureRadiusBoardParams) (r InfrastructureRadiusBoardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// LocationList implements locationList operation.
//
// Get location list.
//
// GET /location/list
func (UnimplementedHandler) LocationList(ctx context.Context) (r LocationListRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SelectionById implements selectionById operation.
//
// Get selection.
//
// GET /selection/{selectionId}
func (UnimplementedHandler) SelectionById(ctx context.Context, params SelectionByIdParams) (r SelectionByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SelectionList implements selectionList operation.
//
// Get selection list.
//
// GET /selection/list
func (UnimplementedHandler) SelectionList(ctx context.Context) (r SelectionListRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UserLogin implements userLogin operation.
//
// Authenticate the user using email and password.
//
// POST /profile/login
func (UnimplementedHandler) UserLogin(ctx context.Context, req *UserLoginReq) (r UserLoginRes, _ error) {
	return r, ht.ErrNotImplemented
}
