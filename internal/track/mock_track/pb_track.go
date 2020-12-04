// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_CodeExpress/internal/track/proto_track (interfaces: TrackServiceClient)

// Package mock_proto_track is a generated GoMock package.
package mock_track

import (
	context "context"
	proto_track "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/track/proto_track"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTrackServiceClient is a mock of TrackServiceClient interface
type MockTrackServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTrackServiceClientMockRecorder
}

// MockTrackServiceClientMockRecorder is the mock recorder for MockTrackServiceClient
type MockTrackServiceClientMockRecorder struct {
	mock *MockTrackServiceClient
}

// NewMockTrackServiceClient creates a new mock instance
func NewMockTrackServiceClient(ctrl *gomock.Controller) *MockTrackServiceClient {
	mock := &MockTrackServiceClient{ctrl: ctrl}
	mock.recorder = &MockTrackServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrackServiceClient) EXPECT() *MockTrackServiceClientMockRecorder {
	return m.recorder
}

// AddToFavourites mocks base method
func (m *MockTrackServiceClient) AddToFavourites(arg0 context.Context, arg1 *proto_track.Favorites, arg2 ...grpc.CallOption) (*proto_track.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddToFavourites", varargs...)
	ret0, _ := ret[0].(*proto_track.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToFavourites indicates an expected call of AddToFavourites
func (mr *MockTrackServiceClientMockRecorder) AddToFavourites(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFavourites", reflect.TypeOf((*MockTrackServiceClient)(nil).AddToFavourites), varargs...)
}

// CreateTrack mocks base method
func (m *MockTrackServiceClient) CreateTrack(arg0 context.Context, arg1 *proto_track.Track, arg2 ...grpc.CallOption) (*proto_track.Track, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrack", varargs...)
	ret0, _ := ret[0].(*proto_track.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrack indicates an expected call of CreateTrack
func (mr *MockTrackServiceClientMockRecorder) CreateTrack(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrack", reflect.TypeOf((*MockTrackServiceClient)(nil).CreateTrack), varargs...)
}

// DeleteFromFavourites mocks base method
func (m *MockTrackServiceClient) DeleteFromFavourites(arg0 context.Context, arg1 *proto_track.Favorites, arg2 ...grpc.CallOption) (*proto_track.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFromFavourites", varargs...)
	ret0, _ := ret[0].(*proto_track.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFromFavourites indicates an expected call of DeleteFromFavourites
func (mr *MockTrackServiceClientMockRecorder) DeleteFromFavourites(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromFavourites", reflect.TypeOf((*MockTrackServiceClient)(nil).DeleteFromFavourites), varargs...)
}

// DeleteTrack mocks base method
func (m *MockTrackServiceClient) DeleteTrack(arg0 context.Context, arg1 *proto_track.TrackID, arg2 ...grpc.CallOption) (*proto_track.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTrack", varargs...)
	ret0, _ := ret[0].(*proto_track.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTrack indicates an expected call of DeleteTrack
func (mr *MockTrackServiceClientMockRecorder) DeleteTrack(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrack", reflect.TypeOf((*MockTrackServiceClient)(nil).DeleteTrack), varargs...)
}

// GetByAlbumID mocks base method
func (m *MockTrackServiceClient) GetByAlbumID(arg0 context.Context, arg1 *proto_track.AlbumID, arg2 ...grpc.CallOption) (*proto_track.Tracks, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByAlbumID", varargs...)
	ret0, _ := ret[0].(*proto_track.Tracks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAlbumID indicates an expected call of GetByAlbumID
func (mr *MockTrackServiceClientMockRecorder) GetByAlbumID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAlbumID", reflect.TypeOf((*MockTrackServiceClient)(nil).GetByAlbumID), varargs...)
}

// GetByArtistId mocks base method
func (m *MockTrackServiceClient) GetByArtistId(arg0 context.Context, arg1 *proto_track.GetByArtistIdMessage, arg2 ...grpc.CallOption) (*proto_track.Tracks, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByArtistId", varargs...)
	ret0, _ := ret[0].(*proto_track.Tracks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByArtistId indicates an expected call of GetByArtistId
func (mr *MockTrackServiceClientMockRecorder) GetByArtistId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByArtistId", reflect.TypeOf((*MockTrackServiceClient)(nil).GetByArtistId), varargs...)
}

// GetByID mocks base method
func (m *MockTrackServiceClient) GetByID(arg0 context.Context, arg1 *proto_track.TrackID, arg2 ...grpc.CallOption) (*proto_track.Track, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByID", varargs...)
	ret0, _ := ret[0].(*proto_track.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTrackServiceClientMockRecorder) GetByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTrackServiceClient)(nil).GetByID), varargs...)
}

// GetByParams mocks base method
func (m *MockTrackServiceClient) GetByParams(arg0 context.Context, arg1 *proto_track.GetByParamsMessage, arg2 ...grpc.CallOption) (*proto_track.Tracks, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByParams", varargs...)
	ret0, _ := ret[0].(*proto_track.Tracks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParams indicates an expected call of GetByParams
func (mr *MockTrackServiceClientMockRecorder) GetByParams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParams", reflect.TypeOf((*MockTrackServiceClient)(nil).GetByParams), varargs...)
}

// GetByPlaylistID mocks base method
func (m *MockTrackServiceClient) GetByPlaylistID(arg0 context.Context, arg1 *proto_track.PlaylistID, arg2 ...grpc.CallOption) (*proto_track.Tracks, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByPlaylistID", varargs...)
	ret0, _ := ret[0].(*proto_track.Tracks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPlaylistID indicates an expected call of GetByPlaylistID
func (mr *MockTrackServiceClientMockRecorder) GetByPlaylistID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPlaylistID", reflect.TypeOf((*MockTrackServiceClient)(nil).GetByPlaylistID), varargs...)
}

// GetFavoritesByUserID mocks base method
func (m *MockTrackServiceClient) GetFavoritesByUserID(arg0 context.Context, arg1 *proto_track.UserID, arg2 ...grpc.CallOption) (*proto_track.Tracks, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFavoritesByUserID", varargs...)
	ret0, _ := ret[0].(*proto_track.Tracks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoritesByUserID indicates an expected call of GetFavoritesByUserID
func (mr *MockTrackServiceClientMockRecorder) GetFavoritesByUserID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritesByUserID", reflect.TypeOf((*MockTrackServiceClient)(nil).GetFavoritesByUserID), varargs...)
}

// UpdateTrack mocks base method
func (m *MockTrackServiceClient) UpdateTrack(arg0 context.Context, arg1 *proto_track.Track, arg2 ...grpc.CallOption) (*proto_track.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrack", varargs...)
	ret0, _ := ret[0].(*proto_track.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrack indicates an expected call of UpdateTrack
func (mr *MockTrackServiceClientMockRecorder) UpdateTrack(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrack", reflect.TypeOf((*MockTrackServiceClient)(nil).UpdateTrack), varargs...)
}

// UpdateTrackAudio mocks base method
func (m *MockTrackServiceClient) UpdateTrackAudio(arg0 context.Context, arg1 *proto_track.Track, arg2 ...grpc.CallOption) (*proto_track.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrackAudio", varargs...)
	ret0, _ := ret[0].(*proto_track.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrackAudio indicates an expected call of UpdateTrackAudio
func (mr *MockTrackServiceClientMockRecorder) UpdateTrackAudio(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrackAudio", reflect.TypeOf((*MockTrackServiceClient)(nil).UpdateTrackAudio), varargs...)
}
