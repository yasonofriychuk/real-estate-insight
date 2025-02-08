// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

type BuildRoutesByPointsBadRequest Error

func (*BuildRoutesByPointsBadRequest) buildRoutesByPointsRes() {}

type BuildRoutesByPointsInternalServerError Error

func (*BuildRoutesByPointsInternalServerError) buildRoutesByPointsRes() {}

type BuildRoutesByPointsNotFound Error

func (*BuildRoutesByPointsNotFound) buildRoutesByPointsRes() {}

type BuildRoutesByPointsOK map[string]jx.Raw

func (s *BuildRoutesByPointsOK) init() BuildRoutesByPointsOK {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

func (*BuildRoutesByPointsOK) buildRoutesByPointsRes() {}

// Ref: #/components/schemas/development
type Development struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	Coords      DevelopmentCoords `json:"coords"`
	ImageUrl    string            `json:"imageUrl"`
	Description string            `json:"description"`
}

// GetID returns the value of ID.
func (s *Development) GetID() int64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Development) GetName() string {
	return s.Name
}

// GetCoords returns the value of Coords.
func (s *Development) GetCoords() DevelopmentCoords {
	return s.Coords
}

// GetImageUrl returns the value of ImageUrl.
func (s *Development) GetImageUrl() string {
	return s.ImageUrl
}

// GetDescription returns the value of Description.
func (s *Development) GetDescription() string {
	return s.Description
}

// SetID sets the value of ID.
func (s *Development) SetID(val int64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Development) SetName(val string) {
	s.Name = val
}

// SetCoords sets the value of Coords.
func (s *Development) SetCoords(val DevelopmentCoords) {
	s.Coords = val
}

// SetImageUrl sets the value of ImageUrl.
func (s *Development) SetImageUrl(val string) {
	s.ImageUrl = val
}

// SetDescription sets the value of Description.
func (s *Development) SetDescription(val string) {
	s.Description = val
}

type DevelopmentCoords struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// GetLat returns the value of Lat.
func (s *DevelopmentCoords) GetLat() float64 {
	return s.Lat
}

// GetLon returns the value of Lon.
func (s *DevelopmentCoords) GetLon() float64 {
	return s.Lon
}

// SetLat sets the value of Lat.
func (s *DevelopmentCoords) SetLat(val float64) {
	s.Lat = val
}

// SetLon sets the value of Lon.
func (s *DevelopmentCoords) SetLon(val float64) {
	s.Lon = val
}

type DevelopmentSearchBadRequest Error

func (*DevelopmentSearchBadRequest) developmentSearchRes() {}

type DevelopmentSearchInternalServerError Error

func (*DevelopmentSearchInternalServerError) developmentSearchRes() {}

type DevelopmentSearchOK struct {
	Developments []Development           `json:"developments"`
	Meta         DevelopmentSearchOKMeta `json:"meta"`
}

// GetDevelopments returns the value of Developments.
func (s *DevelopmentSearchOK) GetDevelopments() []Development {
	return s.Developments
}

// GetMeta returns the value of Meta.
func (s *DevelopmentSearchOK) GetMeta() DevelopmentSearchOKMeta {
	return s.Meta
}

// SetDevelopments sets the value of Developments.
func (s *DevelopmentSearchOK) SetDevelopments(val []Development) {
	s.Developments = val
}

// SetMeta sets the value of Meta.
func (s *DevelopmentSearchOK) SetMeta(val DevelopmentSearchOKMeta) {
	s.Meta = val
}

func (*DevelopmentSearchOK) developmentSearchRes() {}

type DevelopmentSearchOKMeta struct {
	Total int64 `json:"total"`
}

// GetTotal returns the value of Total.
func (s *DevelopmentSearchOKMeta) GetTotal() int64 {
	return s.Total
}

// SetTotal sets the value of Total.
func (s *DevelopmentSearchOKMeta) SetTotal(val int64) {
	s.Total = val
}

type DevelopmentSearchReq struct {
	SearchQuery OptString                         `json:"searchQuery"`
	Pagination  OptDevelopmentSearchReqPagination `json:"pagination"`
	Board       OptDevelopmentSearchReqBoard      `json:"board"`
}

// GetSearchQuery returns the value of SearchQuery.
func (s *DevelopmentSearchReq) GetSearchQuery() OptString {
	return s.SearchQuery
}

// GetPagination returns the value of Pagination.
func (s *DevelopmentSearchReq) GetPagination() OptDevelopmentSearchReqPagination {
	return s.Pagination
}

// GetBoard returns the value of Board.
func (s *DevelopmentSearchReq) GetBoard() OptDevelopmentSearchReqBoard {
	return s.Board
}

// SetSearchQuery sets the value of SearchQuery.
func (s *DevelopmentSearchReq) SetSearchQuery(val OptString) {
	s.SearchQuery = val
}

// SetPagination sets the value of Pagination.
func (s *DevelopmentSearchReq) SetPagination(val OptDevelopmentSearchReqPagination) {
	s.Pagination = val
}

// SetBoard sets the value of Board.
func (s *DevelopmentSearchReq) SetBoard(val OptDevelopmentSearchReqBoard) {
	s.Board = val
}

type DevelopmentSearchReqBoard struct {
	TopLeftLon     float64 `json:"topLeftLon"`
	TopLeftLat     float64 `json:"topLeftLat"`
	BottomRightLon float64 `json:"bottomRightLon"`
	BottomRightLat float64 `json:"bottomRightLat"`
}

// GetTopLeftLon returns the value of TopLeftLon.
func (s *DevelopmentSearchReqBoard) GetTopLeftLon() float64 {
	return s.TopLeftLon
}

// GetTopLeftLat returns the value of TopLeftLat.
func (s *DevelopmentSearchReqBoard) GetTopLeftLat() float64 {
	return s.TopLeftLat
}

// GetBottomRightLon returns the value of BottomRightLon.
func (s *DevelopmentSearchReqBoard) GetBottomRightLon() float64 {
	return s.BottomRightLon
}

// GetBottomRightLat returns the value of BottomRightLat.
func (s *DevelopmentSearchReqBoard) GetBottomRightLat() float64 {
	return s.BottomRightLat
}

// SetTopLeftLon sets the value of TopLeftLon.
func (s *DevelopmentSearchReqBoard) SetTopLeftLon(val float64) {
	s.TopLeftLon = val
}

// SetTopLeftLat sets the value of TopLeftLat.
func (s *DevelopmentSearchReqBoard) SetTopLeftLat(val float64) {
	s.TopLeftLat = val
}

// SetBottomRightLon sets the value of BottomRightLon.
func (s *DevelopmentSearchReqBoard) SetBottomRightLon(val float64) {
	s.BottomRightLon = val
}

// SetBottomRightLat sets the value of BottomRightLat.
func (s *DevelopmentSearchReqBoard) SetBottomRightLat(val float64) {
	s.BottomRightLat = val
}

type DevelopmentSearchReqPagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// GetPage returns the value of Page.
func (s *DevelopmentSearchReqPagination) GetPage() int {
	return s.Page
}

// GetPerPage returns the value of PerPage.
func (s *DevelopmentSearchReqPagination) GetPerPage() int {
	return s.PerPage
}

// SetPage sets the value of Page.
func (s *DevelopmentSearchReqPagination) SetPage(val int) {
	s.Page = val
}

// SetPerPage sets the value of PerPage.
func (s *DevelopmentSearchReqPagination) SetPerPage(val int) {
	s.PerPage = val
}

// Ref: #/components/schemas/error
type Error struct {
	Status ErrorStatus `json:"status"`
	// Ошибка.
	Error ErrorError `json:"error"`
}

// GetStatus returns the value of Status.
func (s *Error) GetStatus() ErrorStatus {
	return s.Status
}

// GetError returns the value of Error.
func (s *Error) GetError() ErrorError {
	return s.Error
}

// SetStatus sets the value of Status.
func (s *Error) SetStatus(val ErrorStatus) {
	s.Status = val
}

// SetError sets the value of Error.
func (s *Error) SetError(val ErrorError) {
	s.Error = val
}

// Ошибка.
type ErrorError struct {
	// Код ошибки.
	Code int `json:"code"`
	// Сообщение об ошибке.
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *ErrorError) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ErrorError) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorError) SetMessage(val string) {
	s.Message = val
}

type ErrorStatus string

const (
	ErrorStatusNotFound      ErrorStatus = "not-found"
	ErrorStatusBadRequest    ErrorStatus = "bad-request"
	ErrorStatusInternalError ErrorStatus = "internal-error"
	ErrorStatusUnauthorized  ErrorStatus = "unauthorized"
)

// AllValues returns all ErrorStatus values.
func (ErrorStatus) AllValues() []ErrorStatus {
	return []ErrorStatus{
		ErrorStatusNotFound,
		ErrorStatusBadRequest,
		ErrorStatusInternalError,
		ErrorStatusUnauthorized,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ErrorStatus) MarshalText() ([]byte, error) {
	switch s {
	case ErrorStatusNotFound:
		return []byte(s), nil
	case ErrorStatusBadRequest:
		return []byte(s), nil
	case ErrorStatusInternalError:
		return []byte(s), nil
	case ErrorStatusUnauthorized:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ErrorStatus) UnmarshalText(data []byte) error {
	switch ErrorStatus(data) {
	case ErrorStatusNotFound:
		*s = ErrorStatusNotFound
		return nil
	case ErrorStatusBadRequest:
		*s = ErrorStatusBadRequest
		return nil
	case ErrorStatusInternalError:
		*s = ErrorStatusInternalError
		return nil
	case ErrorStatusUnauthorized:
		*s = ErrorStatusUnauthorized
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type InfrastructureRadiusBoardBadRequest Error

func (*InfrastructureRadiusBoardBadRequest) infrastructureRadiusBoardRes() {}

type InfrastructureRadiusBoardInternalServerError Error

func (*InfrastructureRadiusBoardInternalServerError) infrastructureRadiusBoardRes() {}

type InfrastructureRadiusBoardNotFound Error

func (*InfrastructureRadiusBoardNotFound) infrastructureRadiusBoardRes() {}

type InfrastructureRadiusBoardOKApplicationJSON []InfrastructureRadiusBoardOKItem

func (*InfrastructureRadiusBoardOKApplicationJSON) infrastructureRadiusBoardRes() {}

type InfrastructureRadiusBoardOKItem struct {
	ID      int                                   `json:"id"`
	Name    OptString                             `json:"name"`
	ObjType string                                `json:"objType"`
	Coords  InfrastructureRadiusBoardOKItemCoords `json:"coords"`
}

// GetID returns the value of ID.
func (s *InfrastructureRadiusBoardOKItem) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *InfrastructureRadiusBoardOKItem) GetName() OptString {
	return s.Name
}

// GetObjType returns the value of ObjType.
func (s *InfrastructureRadiusBoardOKItem) GetObjType() string {
	return s.ObjType
}

// GetCoords returns the value of Coords.
func (s *InfrastructureRadiusBoardOKItem) GetCoords() InfrastructureRadiusBoardOKItemCoords {
	return s.Coords
}

// SetID sets the value of ID.
func (s *InfrastructureRadiusBoardOKItem) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *InfrastructureRadiusBoardOKItem) SetName(val OptString) {
	s.Name = val
}

// SetObjType sets the value of ObjType.
func (s *InfrastructureRadiusBoardOKItem) SetObjType(val string) {
	s.ObjType = val
}

// SetCoords sets the value of Coords.
func (s *InfrastructureRadiusBoardOKItem) SetCoords(val InfrastructureRadiusBoardOKItemCoords) {
	s.Coords = val
}

type InfrastructureRadiusBoardOKItemCoords struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// GetLon returns the value of Lon.
func (s *InfrastructureRadiusBoardOKItemCoords) GetLon() float64 {
	return s.Lon
}

// GetLat returns the value of Lat.
func (s *InfrastructureRadiusBoardOKItemCoords) GetLat() float64 {
	return s.Lat
}

// SetLon sets the value of Lon.
func (s *InfrastructureRadiusBoardOKItemCoords) SetLon(val float64) {
	s.Lon = val
}

// SetLat sets the value of Lat.
func (s *InfrastructureRadiusBoardOKItemCoords) SetLat(val float64) {
	s.Lat = val
}

// NewOptDevelopmentSearchReqBoard returns new OptDevelopmentSearchReqBoard with value set to v.
func NewOptDevelopmentSearchReqBoard(v DevelopmentSearchReqBoard) OptDevelopmentSearchReqBoard {
	return OptDevelopmentSearchReqBoard{
		Value: v,
		Set:   true,
	}
}

// OptDevelopmentSearchReqBoard is optional DevelopmentSearchReqBoard.
type OptDevelopmentSearchReqBoard struct {
	Value DevelopmentSearchReqBoard
	Set   bool
}

// IsSet returns true if OptDevelopmentSearchReqBoard was set.
func (o OptDevelopmentSearchReqBoard) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDevelopmentSearchReqBoard) Reset() {
	var v DevelopmentSearchReqBoard
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDevelopmentSearchReqBoard) SetTo(v DevelopmentSearchReqBoard) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDevelopmentSearchReqBoard) Get() (v DevelopmentSearchReqBoard, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDevelopmentSearchReqBoard) Or(d DevelopmentSearchReqBoard) DevelopmentSearchReqBoard {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDevelopmentSearchReqPagination returns new OptDevelopmentSearchReqPagination with value set to v.
func NewOptDevelopmentSearchReqPagination(v DevelopmentSearchReqPagination) OptDevelopmentSearchReqPagination {
	return OptDevelopmentSearchReqPagination{
		Value: v,
		Set:   true,
	}
}

// OptDevelopmentSearchReqPagination is optional DevelopmentSearchReqPagination.
type OptDevelopmentSearchReqPagination struct {
	Value DevelopmentSearchReqPagination
	Set   bool
}

// IsSet returns true if OptDevelopmentSearchReqPagination was set.
func (o OptDevelopmentSearchReqPagination) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDevelopmentSearchReqPagination) Reset() {
	var v DevelopmentSearchReqPagination
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDevelopmentSearchReqPagination) SetTo(v DevelopmentSearchReqPagination) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDevelopmentSearchReqPagination) Get() (v DevelopmentSearchReqPagination, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDevelopmentSearchReqPagination) Or(d DevelopmentSearchReqPagination) DevelopmentSearchReqPagination {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
