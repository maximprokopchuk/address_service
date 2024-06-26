package grpcserver

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maximprokopchuk/address_service/internal/sqlc"
	"github.com/maximprokopchuk/address_service/internal/store"
	"github.com/maximprokopchuk/address_service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	Store *store.Store
}

func New(st *store.Store) *GRPCServer {
	return &GRPCServer{Store: st}
}

func (server *GRPCServer) CreateAddress(ctx context.Context, req *api.CreateAddressRequest) (*api.CreateAddressResponse, error) {
	params := sqlc.CreateAddressParams{
		Name: req.GetName(),
		Type: req.GetType(),
	}
	parentId := req.GetParentId()
	if parentId != 0 {
		params.ParentID = pgtype.Int4{Int32: parentId, Valid: true}
	}
	rec, err := server.Store.Queries.CreateAddress(ctx, params)

	if err != nil {
		return nil, err
	}
	return &api.CreateAddressResponse{Result: &api.Address{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, ParentId: rec.ParentID.Int32}}, nil
}

func (server *GRPCServer) ListAddressesByParentAndType(ctx context.Context, req *api.ListAddressesByParentIdAndTypeRequest) (*api.ListAddressesByParentIdAndTypeResponse, error) {
	var (
		rec         []sqlc.Address
		err         error
		parentId    = req.GetParentId()
		addressType = req.GetType()
	)
	if parentId == 0 {
		rec, err = server.Store.Queries.ListTopLevelAddresses(ctx)
	} else {
		params := sqlc.ListAddressesByParentIdAndTypeParams{
			ParentID: pgtype.Int4{Int32: parentId, Valid: true},
			Type:     addressType,
		}
		rec, err = server.Store.Queries.ListAddressesByParentIdAndType(ctx, params)
	}

	if err != nil {
		return nil, err
	}

	var result []*api.Address
	for _, element := range rec {
		newRec := api.Address{Id: int32(element.ID), Name: element.Name, Type: element.Type, ParentId: element.ParentID.Int32}
		result = append(result, &newRec)
	}
	return &api.ListAddressesByParentIdAndTypeResponse{Result: result}, nil

}

func (server *GRPCServer) GetAddressById(ctx context.Context, req *api.GetAddressByIdRequest) (*api.GetAddressResponse, error) {
	rec, err := server.Store.Queries.GetAddress(ctx, int64(req.GetId()))
	if err != nil {
		if err == pgx.ErrNoRows {
			err = status.Error(codes.NotFound, "address not found")
		}
		return nil, err
	}
	return &api.GetAddressResponse{Result: &api.Address{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, ParentId: rec.ParentID.Int32}}, nil
}

func (server *GRPCServer) DeleteAddress(ctx context.Context, req *api.DeleteAddressRequest) (*api.DeleteAddressResponse, error) {
	params := sqlc.DeleteAddressParams{
		ID:   int64(req.GetId()),
		Type: req.GetType(),
	}
	err := server.Store.Queries.DeleteAddress(ctx, params)
	if err != nil {
		return nil, err
	}
	return &api.DeleteAddressResponse{}, nil
}
