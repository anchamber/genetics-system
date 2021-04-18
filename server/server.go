package server

import (
	"context"
	"fmt"
	"log"
	"time"

	apiModel "github.com/anchamber/genetics-api/model"
	"github.com/anchamber/genetics-system/db"
	"github.com/anchamber/genetics-system/db/model"
	pb "github.com/anchamber/genetics-system/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var systemDB db.SystemDB = db.NewMockDB()

type SystemService struct {
	pb.UnimplementedSystemServiceServer
}

func (s *SystemService) GetSystems(in *pb.GetSystemsRequest, stream pb.SystemService_GetSystemsServer) error {
	log.Printf("GET: received with %d filters\n", len(in.Filters))
	var paginationSettings *apiModel.Pageination
	if in.Pageination != nil {
		paginationSettings = &apiModel.Pageination{
			Limit:  in.Pageination.Limit,
			Offset: in.Pageination.Offset,
		}
	}

	var filterSettings []*apiModel.Filter
	if in.Filters != nil {
		for _, filter := range in.Filters {
			filterSettings = append(filterSettings, apiModel.NewFilterFromProto(filter))
		}
	}

	data, _ := systemDB.Select(db.Options{
		Pageination: paginationSettings,
		Filters:     filterSettings,
	})
	for _, system := range data {
		if err := stream.Send(mapToProto(system)); err != nil {
			fmt.Printf("%v\n", err)
			return status.Error(codes.Internal, "internal error")
		}
	}
	return nil
}

func (s *SystemService) GetSystem(ctx context.Context, in *pb.GetSystemRequest) (*pb.SystemResponse, error) {
	log.Printf("GET: received for %s\n", in.Name)
	system, error := systemDB.SelectByName(in.Name)
	if error != nil {
		log.Panic(error)
	}
	if system == nil {
		return nil, status.Error(codes.NotFound, "no system with name found")
	}
	return mapToProto(system), nil
}

func (s *SystemService) CreateSystem(ctx context.Context, in *pb.CreateSystemRequest) (*pb.CreateSystemResponse, error) {
	log.Printf("CREATE: received for %v\n", in)
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "request needs to contain valid name")
	}
	system := &model.System{
		Name:             in.Name,
		Location:         in.Location,
		Type:             model.SystemType(in.Type),
		CleaningInterval: in.CleaningInterval,
		LastCleaned:      time.Unix(in.LastCleaned, 0),
	}
	err := systemDB.Insert(system)
	if err != nil {
		switch err.Error() {
		case string(db.SystemAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, "system already exists")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}
	return &pb.CreateSystemResponse{}, nil
}

func (s *SystemService) UpdateSystem(ctx context.Context, in *pb.UpdateSystemRequest) (*pb.UpdateSystemResponse, error) {
	log.Printf("UPDATE: received for %v\n", in)
	system := &model.System{
		Name:             in.Name,
		Location:         in.Location,
		Type:             model.SystemType(in.Type),
		CleaningInterval: in.CleaningInterval,
		LastCleaned:      time.Unix(in.LastCleaned, 0),
	}
	systemDB.Update(system)
	return &pb.UpdateSystemResponse{}, nil
}

func (s *SystemService) DeleteSystem(ctx context.Context, in *pb.DeleteSystemRequest) (*pb.DeleteSystemResponse, error) {
	log.Printf("DEL: received for %s\n", in.Name)
	return &pb.DeleteSystemResponse{}, nil
}

func mapToProto(system *model.System) *pb.SystemResponse {
	return &pb.SystemResponse{
		Name:             system.Name,
		Location:         system.Location,
		Type:             pb.SystemType(system.Type),
		CleaningInterval: system.CleaningInterval,
		LastCleaned:      system.LastCleaned.Unix(),
	}
}

func New() *SystemService {
	return &SystemService{}
}
