package service_test

import (
	"testing"

	"github.com/anchamber/genetics-system/db"
	pb "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/service"
	grpc "google.golang.org/grpc"
)

func TestGetSystems(t *testing.T) {
	testCases := []struct {
		name          string
		request       *pb.GetSystemsRequest
		message       string
		expectedError bool
	}{
		{
			name:          "request all entries",
			request:       &pb.GetSystemsRequest{},
			message:       "",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			systemServer := service.New(db.NewMockDB())
			sericeMock := MockService{
				t: t,
			}
			err := systemServer.GetSystems(tc.request, &sericeMock)
			if err != nil && !testCase.expectedError {
				t.Errorf("respons returned error and when it should be ok: %v", err)
			}
		})
	}
}

type MockService struct {
	CallCount int
	t         *testing.T
	grpc.ServerStream
}

func (x *MockService) Send(m *pb.SystemResponse) error {
	system := db.MockDataSystems[x.CallCount]
	if system.Name != m.Name {
		x.t.Errorf("names do not match, expected: %s | got: %s", system.Name, m.Name)
	}
	if int(system.Type) != int(m.Type) {
		x.t.Errorf("types do not match, expected: %s | got: %s", system.Name, m.Name)
	}
	if system.Location != m.Location {
		x.t.Errorf("locations do not match, expected: %s | got: %s", system.Name, m.Name)
	}
	if system.CleaningInterval != m.CleaningInterval {
		x.t.Errorf("cleaning intervals do not match, expected: %s | got: %s", system.Name, m.Name)
	}
	if system.LastCleaned.Unix() != m.LastCleaned {
		x.t.Errorf("last cleaned do not match, expected: %s | got: %s", system.Name, m.Name)
	}
	x.CallCount++
	return nil
}
