package service_test

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"

	apiProto "github.com/anchamber/genetics-api/proto"
	"github.com/anchamber/genetics-system/db"
	sm "github.com/anchamber/genetics-system/db/model"
	systemProto "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/service"
	grpc "google.golang.org/grpc"
)

var testData = []*sm.System{
	{Name: "doctor", Location: "tardis", Type: sm.Techniplast, Responsible: "", CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "rick", Location: "c-137", Type: sm.Techniplast, Responsible: "", CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "morty", Location: "herry-herpson", Type: sm.Techniplast, Responsible: "", CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "obi", Location: "high_ground", Type: sm.Techniplast, Responsible: "", CleaningInterval: 90, LastCleaned: time.Now()},
}

func TestGetSystems(t *testing.T) {
	testCases := []struct {
		name          string
		request       *systemProto.GetSystemsRequest
		responses     []*sm.System
		expectedError bool
	}{
		{
			name:          "request all entries",
			request:       &systemProto.GetSystemsRequest{},
			responses:     db.MockDataSystems,
			expectedError: false,
		},
		{
			name: "request with limit",
			request: &systemProto.GetSystemsRequest{
				Pageination: &apiProto.Pagination{
					Limit: 2,
				},
			},
			responses:     db.MockDataSystems[0:2],
			expectedError: false,
		},
		{
			name: "request with offset",
			request: &systemProto.GetSystemsRequest{
				Pageination: &apiProto.Pagination{
					Offset: 2,
				},
			},
			responses:     db.MockDataSystems[2:],
			expectedError: false,
		},
		{
			name: "request with offset and limit",
			request: &systemProto.GetSystemsRequest{
				Pageination: &apiProto.Pagination{
					Offset: 2,
					Limit:  1,
				},
			},
			responses:     db.MockDataSystems[2:3],
			expectedError: false,
		},
		{
			name: "request with name filter EQ",
			request: &systemProto.GetSystemsRequest{
				Filters: []*apiProto.Filter{
					{
						Key:      "name",
						Operator: apiProto.Operator_EQ,
						Value:    &apiProto.Filter_S{S: db.MockDataSystems[1].Name},
					},
				},
			},
			responses: []*sm.System{
				db.MockDataSystems[1],
			},
			expectedError: false,
		},
		{
			name: "request with name filter CONTAINS",
			request: &systemProto.GetSystemsRequest{
				Filters: []*apiProto.Filter{
					{
						Key:      "name",
						Operator: apiProto.Operator_CONTAINS,
						Value:    &apiProto.Filter_S{S: db.MockDataSystems[2].Name[1:4]},
					},
				},
			},
			responses: []*sm.System{
				db.MockDataSystems[2],
			},
			expectedError: false,
		},
		{
			name: "request with invalid filter key",
			request: &systemProto.GetSystemsRequest{
				Filters: []*apiProto.Filter{
					{
						Key:      "nme",
						Operator: apiProto.Operator_CONTAINS,
						Value:    &apiProto.Filter_S{S: db.MockDataSystems[2].Name},
					},
				},
			},
			responses:     db.MockDataSystems,
			expectedError: false,
		},
	}

	systemServer := service.New(db.NewMockDB(testData))
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tc.responses)
			serviceMock := MockSystemService{
				t:         t,
				responses: tc.responses,
			}
			err := systemServer.GetSystems(tc.request, &serviceMock)
			if err != nil && !tc.expectedError {
				t.Errorf("response returned error and when it should be ok: %v", err)
			}
			if serviceMock.CallCount != len(tc.responses) {
				t.Errorf("Call count of mock does not match, expected: %d | actual: %d", len(tc.responses), serviceMock.CallCount)
			}
		})
	}
}

func TestGetSystem(t *testing.T) {
	testCases := []struct {
		name          string
		request       *systemProto.GetSystemRequest
		response      *sm.System
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:          "request existing system",
			response:      testData[0],
			expectedError: false,
			request: &systemProto.GetSystemRequest{
				Name: testData[0].Name,
			},
		},
		{
			name:          "request none existing system",
			response:      nil,
			expectedError: true,
			request: &systemProto.GetSystemRequest{
				Name: "does not exists",
			},
			errorCode: codes.NotFound,
		},
	}

	systemServer := service.New(db.NewMockDB(testData))
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resp, err := systemServer.GetSystem(context.Background(), tc.request)
			if err != nil {
				if !tc.expectedError {
					t.Errorf("response returned error and when it should be ok: %v", err)
				}
				st, ok := status.FromError(err)
				if !ok {
					t.Errorf("was not a status: %v", err)
				}
				if st.Code() != tc.errorCode {
					t.Errorf("wrong status code: expected %v | actual: %v", tc.errorCode, st.Code())
				}
				return
			}
			compareResponseToSystem(t, resp, tc.response)
		})
	}
}

type MockSystemService struct {
	CallCount int
	t         *testing.T
	responses []*sm.System
	grpc.ServerStream
}

func (x *MockSystemService) Send(resp *systemProto.SystemResponse) error {
	compareResponseToSystem(x.t, resp, x.responses[x.CallCount])
	x.CallCount++
	return nil
}

func compareResponseToSystem(t *testing.T, resp *systemProto.SystemResponse, system *sm.System) {

	if system.Name != resp.Name {
		t.Errorf("names do not match, expected: %s | actual: %s", system.Name, resp.Name)
	}
	if int(system.Type) != int(resp.Type) {
		t.Errorf("types do not match, expected: %v | actual: %v", system.Type, resp.Type)
	}
	if system.Location != resp.Location {
		t.Errorf("locations do not match, expected: %s | actual: %s", system.Location, resp.Location)
	}
	if system.CleaningInterval != resp.CleaningInterval {
		t.Errorf("cleaning intervals do not match, expected: %d | actual: %d", system.CleaningInterval, resp.CleaningInterval)
	}
	if system.LastCleaned.Unix() != resp.LastCleaned {
		t.Errorf("last cleaned do not match, expected: %d | actual: %d", system.LastCleaned.Unix(), resp.LastCleaned)
	}
}
