package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"time"

	apiModel "github.com/anchamber/genetics-api/model"
	"github.com/anchamber/genetics-system/db"
	"github.com/anchamber/genetics-system/db/model"
	pb "github.com/anchamber/genetics-system/proto"
	"github.com/mennanov/fmutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SystemService struct {
	pb.UnimplementedSystemServiceServer
	db db.SystemDB
}

var systemFilterKeys = []string{
	"id", "name", "location", "type", "responsible", "cleaning_interval", "last_cleaned",
}

func (s *SystemService) StreamSystems(in *pb.StreamSystemsRequest, stream pb.SystemService_StreamSystemsServer) error {
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
			found := false
			for _, fk := range systemFilterKeys {
				if fk == filter.Key {
					filterSettings = append(filterSettings, apiModel.NewFilterFromProto(filter))
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("invalid filter key '%s' will be ignored\n", filter.Key)
			}
		}
	}

	data, _ := s.db.Select(db.Options{
		Pageination: paginationSettings,
		Filters:     filterSettings,
	})
	for _, system := range data {
		if err := stream.Send(mapToSystemResponse(system)); err != nil {
			fmt.Printf("%v\n", err)
			return status.Error(codes.Internal, "internal error")
		}
	}
	return nil
}

func (s *SystemService) GetSystem(_ context.Context, in *pb.GetSystemRequest) (*pb.SystemResponse, error) {
	log.Printf("GET: received for %s\n", in.Name)
	system, err := s.db.SelectByName(in.Name)
	if err != nil {
		log.Panic(err)
	}
	if system == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no system with name '%s' found", in.Name))
	}
	return mapToSystemResponse(system), nil
}

func (s *SystemService) CreateSystem(_ context.Context, in *pb.CreateSystemRequest) (*pb.CreateSystemResponse, error) {
	log.Printf("CREATE: received for %v\n", in)
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "request needs to contain valid name")
	}
	if in.CleaningInterval <= 0 {
		return nil, status.Error(codes.InvalidArgument, "cleaning interval to short")
	}
	system := &model.System{
		Name:             in.Name,
		Location:         in.Location,
		Type:             model.SystemType(in.Type),
		CleaningInterval: in.CleaningInterval,
		LastCleaned:      time.Unix(in.LastCleaned, 0),
	}
	err := s.db.Insert(system)
	if err != nil {
		re, ok := err.(*db.EntityAlreadyExists)
		if ok {
			return nil, status.Error(codes.AlreadyExists, re.Error())
		} else {
			return nil, status.Error(codes.Internal, re.Error())
		}
	}
	return &pb.CreateSystemResponse{}, nil
}

func (s *SystemService) UpdateSystem(_ context.Context, in *pb.UpdateSystemRequest) (*pb.UpdateSystemResponse, error) {
	log.Printf("UPDATE: received for %v\n", in)
	entity, err := s.db.SelectByName(in.Name)
	if err != nil {

	}
	transformed := mapToSystemProto(entity)
	in.Mask.Normalize()
	if !in.Mask.IsValid(transformed) {

	}
	fmutils.Filter(in.GetSystem(), in.GetMask().GetPaths())
	proto.Merge(transformed, in.GetSystem())
	err = s.db.Update(mapToSystemModel(transformed))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSystemResponse{}, nil
}

func (s *SystemService) DeleteSystem(_ context.Context, in *pb.DeleteSystemRequest) (*pb.DeleteSystemResponse, error) {
	log.Printf("DEL: received for %s\n", in.Name)
	err := s.db.Delete(in.Name)
	if err != nil {

	}
	return &pb.DeleteSystemResponse{}, nil
}

func mapToSystemResponse(system *model.System) *pb.SystemResponse {
	return &pb.SystemResponse{
		Name:             system.Name,
		Location:         system.Location,
		Type:             pb.SystemType(system.Type),
		CleaningInterval: system.CleaningInterval,
		LastCleaned:      system.LastCleaned.Unix(),
	}
}
func mapToSystemProto(system *model.System) *pb.System {
	return &pb.System{
		Name:             system.Name,
		Location:         system.Location,
		Type:             pb.SystemType(system.Type),
		CleaningInterval: system.CleaningInterval,
		LastCleaned:      system.LastCleaned.Unix(),
	}
}

func mapToSystemModel(system *pb.System) *model.System {
	return &model.System{
		Name:             system.Name,
		Location:         system.Location,
		Type:             model.SystemType(system.Type),
		CleaningInterval: system.CleaningInterval,
		LastCleaned:      time.Unix(system.LastCleaned, 0),
	}
}

func NewSystemService(db db.SystemDB) *SystemService {
	return &SystemService{
		db: db,
	}
}
