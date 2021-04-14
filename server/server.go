package server

import (
	"context"
	"log"

	"github.com/anchamber/genetics-system/db"
	pb "github.com/anchamber/genetics-system/proto"
)

var systemDB db.SystemDB = &db.SytemDBMock{}

type SystemService struct {
	pb.UnimplementedSystemServiceServer
}

func (s *SystemService) GetSystems(ctx context.Context, in *pb.GetSystemsRequest) (*pb.GetSystemsResponse, error) {
	log.Printf("GET: received with %d filters\n", len(in.Filters))
	data := systemDB.SelctAll()
	log.Printf("%v", data)
	responseData := &pb.Systems{Systems: []*pb.System{}}
	for _, system := range data {
		responseData.Systems = append(responseData.Systems, &pb.System{
			Name:             system.Name,
			Location:         system.Location,
			Type:             pb.SystemType(system.Type),
			CleaningInterval: system.CleaningInterval,
			LastCleaned:      system.LastCleaned.Unix(),
		})
	}
	return &pb.GetSystemsResponse{Response: &pb.GetSystemsResponse_Systems{
		Systems: responseData,
	}}, nil
}

func (s *SystemService) GetSystem(ctx context.Context, in *pb.GetSystemRequest) (*pb.GetSystemResponse, error) {
	log.Printf("GET: received for %s\n", in.Name)
	return &pb.GetSystemResponse{Response: &pb.GetSystemResponse_System{}}, nil
}

func (s *SystemService) CreateSystem(ctx context.Context, in *pb.CreateSystemRequest) (*pb.CreateSystemResponse, error) {
	log.Printf("CREATE: received for %v\n", in.System)
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

func New() *SystemService {
	return &SystemService{}
}
