package service

import (
	"context"
	"fmt"
	apiModel "github.com/anchamber/genetics-api/model"
	"github.com/anchamber/genetics-system/db"
	"github.com/anchamber/genetics-system/db/model"
	pb "github.com/anchamber/genetics-system/proto"
	"github.com/mennanov/fmutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
)

type TankService struct {
	pb.UnimplementedTankServiceServer
	db db.TankDB
}

var tankFilterKeys = []string{
	"number", "active", "size",
}

func (s *TankService) StreamTanks(in *pb.StreamTanksRequest, stream pb.TankService_StreamTanksServer) error {
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
			for _, fk := range tankFilterKeys {
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

	data, err := s.db.Select(db.Options{
		Pageination: paginationSettings,
		Filters:     filterSettings,
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, tank := range data {
		if err := stream.Send(mapToTankResponse(tank)); err != nil {
			fmt.Printf("%v\n", err)
			return status.Error(codes.Internal, "internal error")
		}
	}
	return nil
}

func (s *TankService) GetTank(_ context.Context, in *pb.GetTankRequest) (*pb.TankResponse, error) {
	log.Printf("GET: received for %d\n", in.Number)
	tank, err := s.db.SelectByNumber(in.Number)
	if err != nil {
		log.Panic(err)
	}
	if tank == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no tank with number %d found", in.Number))
	}
	return mapToTankResponse(tank), nil
}

func (s *TankService) CreateTank(_ context.Context, in *pb.CreateTankRequest) (*pb.CreateTankResponse, error) {
	log.Printf("CREATE: received for %v\n", in)
	if in.Number == 0 {
		return nil, status.Error(codes.InvalidArgument, "request needs to contain valid name")
	}
	tank := &model.Tank{
		Number:    in.Number,
		System:    in.System,
		Active:    in.Active,
		Size:      in.Size,
		FishCount: in.FishCount,
	}
	err := s.db.Insert(tank)
	if err != nil {
		re, ok := err.(*db.EntityAlreadyExists)
		if ok {
			return nil, status.Error(codes.AlreadyExists, re.Error())
		} else {
			return nil, status.Error(codes.Internal, re.Error())
		}
	}
	return &pb.CreateTankResponse{}, nil
}

func (s *TankService) UpdateTank(_ context.Context, in *pb.UpdateTankRequest) (*pb.UpdateTankResponse, error) {
	log.Printf("UPDATE: received for %v\n", in)
	entity, err := s.db.SelectByNumber(in.Number)
	if err != nil {

	}
	transformed := mapToTankProto(entity)
	in.Mask.Normalize()
	if !in.Mask.IsValid(transformed) {

	}
	fmutils.Filter(in.GetTank(), in.GetMask().GetPaths())
	proto.Merge(transformed, in.GetTank())
	err = s.db.Update(mapToTankModel(transformed))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTankResponse{}, nil
}

func (s *TankService) DeleteTank(_ context.Context, in *pb.DeleteTankRequest) (*pb.DeleteTankResponse, error) {
	log.Printf("DEL: received for %d\n", in.Number)
	err := s.db.Delete(in.Number)
	if err != nil {

	}
	return &pb.DeleteTankResponse{}, nil
}

func mapToTankResponse(tank *model.Tank) *pb.TankResponse {
	return &pb.TankResponse{
		Number:    tank.Number,
		System:    tank.System,
		Active:    tank.Active,
		Size:      tank.Size,
		FishCount: tank.FishCount,
	}
}
func mapToTankProto(tank *model.Tank) *pb.Tank {
	return &pb.Tank{
		Number:    tank.Number,
		System:    tank.System,
		Active:    tank.Active,
		Size:      tank.Size,
		FishCount: tank.FishCount,
	}
}

func mapToTankModel(tank *pb.Tank) *model.Tank {
	return &model.Tank{
		Number:    tank.Number,
		System:    tank.System,
		Active:    tank.Active,
		Size:      tank.Size,
		FishCount: tank.FishCount,
	}
}

func NewTankService(db db.TankDB) *TankService {
	return &TankService{
		db: db,
	}
}
