package service_test

import (
	"context"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"testing"
	"time"

	apiProto "github.com/anchamber/genetics-api/proto"
	"github.com/anchamber/genetics-system/db"
	sm "github.com/anchamber/genetics-system/db/model"
	systemProto "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/service"
	grpc "google.golang.org/grpc"
)

var systemTestData = db.MockDataSystems

var testSystemsToCreate = []*sm.System{
	{Name: "clara", Location: "london", Type: sm.Glass, Responsible: "doctor", CleaningInterval: 50, LastCleaned: time.Now().Add(time.Hour * 24 * 49)},
	{Name: "emilia", Location: "pond", Type: sm.Glass, Responsible: "doctor", CleaningInterval: 50, LastCleaned: time.Now().Add(time.Hour * 24 * 50)},
	{Name: "song", Location: "river", Type: sm.Glass, Responsible: "doctor", CleaningInterval: 50, LastCleaned: time.Now().Add(time.Hour * 24 * 51)},
}

func TestStreamSystems(t *testing.T) {
	testCases := []struct {
		name          string
		request       *systemProto.StreamSystemsRequest
		responses     []*sm.System
		expectedError bool
	}{
		{
			name:          "request all entries",
			request:       &systemProto.StreamSystemsRequest{},
			responses:     systemTestData,
			expectedError: false,
		},
		{
			name: "request with limit",
			request: &systemProto.StreamSystemsRequest{
				Pageination: &apiProto.Pagination{
					Limit: 2,
				},
			},
			responses:     systemTestData[0:2],
			expectedError: false,
		},
		{
			name: "request with offset",
			request: &systemProto.StreamSystemsRequest{
				Pageination: &apiProto.Pagination{
					Offset: 2,
				},
			},
			responses:     systemTestData[2:],
			expectedError: false,
		},
		{
			name: "request with offset and limit",
			request: &systemProto.StreamSystemsRequest{
				Pageination: &apiProto.Pagination{
					Offset: 2,
					Limit:  1,
				},
			},
			responses:     systemTestData[2:3],
			expectedError: false,
		},
		{
			name: "request with name filter EQ",
			request: &systemProto.StreamSystemsRequest{
				Filters: []*apiProto.Filter{
					{
						Key:      "name",
						Operator: apiProto.Operator_EQ,
						Value:    &apiProto.Filter_S{S: systemTestData[1].Name},
					},
				},
			},
			responses: []*sm.System{
				systemTestData[1],
			},
			expectedError: false,
		},
		{
			name: "request with name filter CONTAINS",
			request: &systemProto.StreamSystemsRequest{
				Filters: []*apiProto.Filter{
					{
						Key:      "name",
						Operator: apiProto.Operator_CONTAINS,
						Value:    &apiProto.Filter_S{S: systemTestData[2].Name[1:4]},
					},
				},
			},
			responses: []*sm.System{
				systemTestData[2],
			},
			expectedError: false,
		},
		{
			name: "request with invalid filter key",
			request: &systemProto.StreamSystemsRequest{
				Filters: []*apiProto.Filter{
					{
						Key:      "nme",
						Operator: apiProto.Operator_CONTAINS,
						Value:    &apiProto.Filter_S{S: systemTestData[2].Name},
					},
				},
			},
			responses:     systemTestData,
			expectedError: false,
		},
	}

	systemServer := service.NewSystemService(db.NewSystemDBMock(systemTestData))
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tc.responses)
			serviceMock := MockSystemService{
				t:         t,
				responses: tc.responses,
			}
			err := systemServer.StreamSystems(tc.request, &serviceMock)
			if validateError(t, err, codes.Unknown, tc.expectedError) {
				return
			}
			if serviceMock.CallCount != len(tc.responses) {
				t.Errorf("Call count of mock does not match, expected: %d | actual: %d", len(tc.responses), serviceMock.CallCount)
			}
		})
	}
}

func TestGetSystem(t *testing.T) {
	index := rand.Intn(len(systemTestData))
	testCases := []struct {
		name          string
		request       *systemProto.GetSystemRequest
		response      *sm.System
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:          "request existing system",
			response:      systemTestData[index],
			expectedError: false,
			request: &systemProto.GetSystemRequest{
				Name: systemTestData[index].Name,
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

	systemServer := service.NewSystemService(db.NewSystemDBMock(systemTestData))
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resp, err := systemServer.GetSystem(context.Background(), tc.request)
			if validateError(t, err, tc.errorCode, tc.expectedError) {
				return
			}
			compareResponseToSystem(t, resp, tc.response)
		})
	}
}

func TestCreateSystem(t *testing.T) {
	system := testSystemsToCreate[rand.Intn(len(testSystemsToCreate))]
	testCases := []struct {
		name          string
		request       *systemProto.CreateSystemRequest
		response      *sm.System
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name:          "create valid system",
			response:      system,
			expectedError: false,
			request: &systemProto.CreateSystemRequest{
				Name:             system.Name,
				Location:         system.Location,
				Type:             systemProto.SystemType(system.Type),
				Responsible:      system.Responsible,
				CleaningInterval: system.CleaningInterval,
				LastCleaned:      system.LastCleaned.Unix(),
			},
		},
		{
			name:          "create system with invalid name",
			response:      nil,
			expectedError: true,
			request: &systemProto.CreateSystemRequest{
				Name:             "",
				Location:         system.Location,
				Type:             systemProto.SystemType(system.Type),
				Responsible:      system.Responsible,
				CleaningInterval: system.CleaningInterval,
				LastCleaned:      system.LastCleaned.Unix(),
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name:          "create system with invalid cleaning interval",
			response:      nil,
			expectedError: true,
			request: &systemProto.CreateSystemRequest{
				Name:             system.Name,
				Location:         "",
				Type:             systemProto.SystemType(system.Type),
				Responsible:      system.Responsible,
				CleaningInterval: 0,
				LastCleaned:      system.LastCleaned.Unix(),
			},
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tc := range testCases {
		systemServer := service.NewSystemService(db.NewSystemDBMock(systemTestData))
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := systemServer.CreateSystem(context.Background(), tc.request)
			if validateError(t, err, tc.errorCode, tc.expectedError) {
				return
			}
			if res == nil {
				t.Error("response should not be nil")
			}
		})
	}
}

func TestUpdateSystem(t *testing.T) {
	system := systemTestData[rand.Intn(len(systemTestData))]
	testCases := []struct {
		name          string
		request       *systemProto.UpdateSystemRequest
		expected      sm.System
		expectedError bool
		errorCode     codes.Code
	}{
		{
			name: "update location",
			request: &systemProto.UpdateSystemRequest{
				Name: system.Name,
				System: &systemProto.System{
					Location: "test location",
				},
				Mask: &field_mask.FieldMask{
					Paths: []string{"location"},
				},
			},
			expected:      sm.System{Name: system.Name, Location: "test location", Type: system.Type, Responsible: system.Responsible, CleaningInterval: system.CleaningInterval, LastCleaned: system.LastCleaned},
			expectedError: false,
		},
		{
			name: "update type",
			request: &systemProto.UpdateSystemRequest{
				Name: system.Name,
				System: &systemProto.System{
					Type: systemProto.SystemType((system.Type + 1) % 2),
				},
				Mask: &field_mask.FieldMask{
					Paths: []string{"type"},
				},
			},
			expected:      sm.System{Name: system.Name, Location: system.Location, Type: (system.Type + 1) % 2, Responsible: system.Responsible, CleaningInterval: system.CleaningInterval, LastCleaned: system.LastCleaned},
			expectedError: false,
		},
		{
			name: "update responsible",
			request: &systemProto.UpdateSystemRequest{
				Name: system.Name,
				System: &systemProto.System{
					Responsible: "darth",
				},
				Mask: &field_mask.FieldMask{
					Paths: []string{"responsible"},
				},
			},
			expected:      sm.System{Name: system.Name, Location: system.Location, Type: system.Type, Responsible: "darth", CleaningInterval: system.CleaningInterval, LastCleaned: system.LastCleaned},
			expectedError: false,
		},
		{
			name: "update cleaning interval",
			request: &systemProto.UpdateSystemRequest{
				Name: system.Name,
				System: &systemProto.System{
					CleaningInterval: system.CleaningInterval + 10,
				},
				Mask: &field_mask.FieldMask{
					Paths: []string{"cleaning_interval"},
				},
			},
			expected:      sm.System{Name: system.Name, Location: system.Location, Type: system.Type, Responsible: system.Responsible, CleaningInterval: system.CleaningInterval + 10, LastCleaned: system.LastCleaned},
			expectedError: false,
		},
		{
			name: "update last cleaned",
			request: &systemProto.UpdateSystemRequest{
				Name: system.Name,

				System: &systemProto.System{
					LastCleaned: system.LastCleaned.Add(-10 * time.Hour * 24).Unix(),
				},
				Mask: &field_mask.FieldMask{
					Paths: []string{"last_cleaned"},
				},
			},
			expected:      sm.System{Name: system.Name, Location: system.Location, Type: system.Type, Responsible: system.Responsible, CleaningInterval: system.CleaningInterval, LastCleaned: system.LastCleaned.Add(-10 * time.Hour * 24)},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		systemServer := service.NewSystemService(db.NewSystemDBMock(systemTestData))
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := systemServer.UpdateSystem(context.Background(), tc.request)
			validateError(t, err, tc.errorCode, tc.expectedError)
			resp, err := systemServer.GetSystem(context.Background(), &systemProto.GetSystemRequest{Name: tc.expected.Name})
			if validateError(t, err, tc.errorCode, tc.expectedError) {
				return
			}
			compareResponseToSystem(t, resp, &tc.expected)
		})
	}
}

func TestDeleteSystem(t *testing.T) {
	system := systemTestData[rand.Intn(len(systemTestData))]
	testCases := []struct {
		name             string
		request          *systemProto.DeleteSystemRequest
		expectedErrorGet bool
		expectedErrorDel bool
		errorCode        codes.Code
	}{
		{
			name: "delete existing system",
			request: &systemProto.DeleteSystemRequest{
				Name: system.Name,
			},
			expectedErrorDel: false,
			expectedErrorGet: true,
			errorCode:        codes.NotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			systemServer := service.NewSystemService(db.NewSystemDBMock(systemTestData))
			_, err := systemServer.DeleteSystem(context.Background(), tc.request)
			if validateError(t, err, tc.errorCode, tc.expectedErrorDel) {
				return
			}
			resp, err := systemServer.GetSystem(context.Background(), &systemProto.GetSystemRequest{Name: system.Name})
			if validateError(t, err, tc.errorCode, tc.expectedErrorGet) {
				return
			}
			compareResponseToSystem(t, resp, system)
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

func validateError(t *testing.T, err error, code codes.Code, expected bool) bool {
	done := false
	if err != nil {
		if !expected {
			t.Fatalf("response returned error and when it should be ok: %v", err)
		}
		st, ok := status.FromError(err)
		if !ok {
			t.Fatalf("was not a status: %v", err)
		}
		if st.Code() != code {
			t.Fatalf("wrong status code: expected %v | actual: %v", code, st.Code())
		}
		done = true
	} else {
		if expected {
			t.Fatalf("response should have had an error")
		}
	}
	return done
}
