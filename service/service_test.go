package service_test

import (
	"testing"

	apiProto "github.com/anchamber/genetics-api/proto"
	"github.com/anchamber/genetics-system/db"
	sm "github.com/anchamber/genetics-system/db/model"
	systemProto "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/service"
	grpc "google.golang.org/grpc"
)

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
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tc.responses)
			systemServer := service.New(db.NewMockDB())
			sericeMock := MockSystemService{
				t:         t,
				responses: tc.responses,
			}
			err := systemServer.GetSystems(tc.request, &sericeMock)
			if err != nil && !tc.expectedError {
				t.Errorf("response returned error and when it should be ok: %v", err)
			}
		})
	}
}

type MockSystemService struct {
	CallCount int
	t         *testing.T
	responses []*sm.System
	grpc.ServerStream
}

func (x *MockSystemService) Send(m *systemProto.SystemResponse) error {
	system := x.responses[x.CallCount]
	if system.Name != m.Name {
		x.t.Errorf("names do not match, expected: %s | got: %s", system.Name, m.Name)
	}
	if int(system.Type) != int(m.Type) {
		x.t.Errorf("types do not match, expected: %v | got: %v", system.Type, m.Type)
	}
	if system.Location != m.Location {
		x.t.Errorf("locations do not match, expected: %s | got: %s", system.Location, m.Location)
	}
	if system.CleaningInterval != m.CleaningInterval {
		x.t.Errorf("cleaning intervals do not match, expected: %d | got: %d", system.CleaningInterval, m.CleaningInterval)
	}
	if system.LastCleaned.Unix() != m.LastCleaned {
		x.t.Errorf("last cleaned do not match, expected: %d | got: %d", system.LastCleaned.Unix(), m.LastCleaned)
	}
	x.CallCount++
	return nil
}
