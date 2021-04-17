package server

import (
	"context"
	"log"
	"time"

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

func (s *SystemService) GetSystems(ctx context.Context, in *pb.GetSystemsRequest) (*pb.GetSystemsResponse, error) {
	log.Printf("GET: received with %d filters\n", len(in.Filters))
	data, _ := systemDB.SelectAll()
	log.Printf("%v", data)
	responseData := &pb.Systems{Systems: []*pb.System{}}
	for _, system := range data {
		responseData.Systems = append(responseData.Systems, mapToProto(system))
	}
	return &pb.GetSystemsResponse{Response: &pb.GetSystemsResponse_Systems{
		Systems: responseData,
	}}, nil
}

func (s *SystemService) GetSystem(ctx context.Context, in *pb.GetSystemRequest) (*pb.GetSystemResponse, error) {
	log.Printf("GET: received for %s\n", in.Name)
	system, error := systemDB.SelectByName(in.Name)
	if error != nil {
		log.Panic(error)
	}
	if system == nil {
		return nil, status.Error(codes.NotFound, "no system with name found")
	}
	return &pb.GetSystemResponse{Response: &pb.GetSystemResponse_System{
		System: mapToProto(system),
	}}, nil
}

func (s *SystemService) CreateSystem(ctx context.Context, in *pb.CreateSystemRequest) (*pb.CreateSystemResponse, error) {
	log.Printf("CREATE: received for %v\n", in.System)
	system := mapFromProto(in.System)
	systemDB.Insert(system)
	return &pb.CreateSystemResponse{Response: &pb.CreateSystemResponse_Systems{}}, nil
}

func (s *SystemService) UpdateSystem(ctx context.Context, in *pb.UpdateSystemRequest) (*pb.UpdateSystemResponse, error) {
	log.Printf("UPDATE: received for %v\n", in.System)
	return &pb.UpdateSystemResponse{Response: &pb.UpdateSystemResponse_Systems{}}, nil
}

func (s *SystemService) DeleteSystem(ctx context.Context, in *pb.DeleteSystemRequest) (*pb.DeleteSystemResponse, error) {
	log.Printf("DEL: received for %s\n", in.Name)
	return &pb.DeleteSystemResponse{Response: &pb.DeleteSystemResponse_Systems{}}, nil
}

func mapToProto(system *model.System) *pb.System {
	return &pb.System{
		Name:             system.Name,
		Location:         system.Location,
		Type:             pb.SystemType(system.Type),
		CleaningInterval: system.CleaningInterval,
		LastCleaned:      system.LastCleaned.Unix(),
	}
}

func mapFromProto(system *pb.System) *model.System {
	return &model.System{
		Name:             system.Name,
		Location:         system.Location,
		Type:             model.SystemType(system.Type),
		CleaningInterval: system.CleaningInterval,
		LastCleaned:      time.Unix(system.LastCleaned, 0),
	}
}

func New() *SystemService {
	return &SystemService{}
}
