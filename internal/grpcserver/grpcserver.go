package grpcserver

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maximprokopchuk/address_service/internal/sqlc"
	"github.com/maximprokopchuk/address_service/internal/store"
	"github.com/maximprokopchuk/address_service/pkg/api"
)

type GRPCServer struct {
	Store *store.Store
}

func New(st *store.Store) *GRPCServer {
	return &GRPCServer{Store: st}
}

func (server *GRPCServer) CreateAddress(ctx context.Context, req *api.CreateAddressRequest) (*api.AddressResponse, error) {
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
	return &api.AddressResponse{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, ParentId: rec.ParentID.Int32}, nil
}

func (server *GRPCServer) ListGetAddressesByParent(ctx context.Context, req *api.GetAddressesByParentIdRequest) (*api.GetAddressesByParentIdResponse, error) {
	var (
		rec      []sqlc.Address
		err      error
		parentId = req.GetParentId()
	)
	if parentId == 0 {
		rec, err = server.Store.Queries.GetAddressesByParent(ctx)
	} else {
		rec, err = server.Store.Queries.ListAddressesForParent(ctx, pgtype.Int4{Int32: parentId, Valid: true})
	}

	if err != nil {
		return nil, err
	}

	var result []*api.AddressResponse
	for _, element := range rec {
		newRec := api.AddressResponse{Id: int32(element.ID), Name: element.Name, Type: element.Type, ParentId: element.ParentID.Int32}
		result = append(result, &newRec)
	}
	return &api.GetAddressesByParentIdResponse{Items: result}, nil

}

func (server *GRPCServer) GetAddressById(ctx context.Context, req *api.GetAddressByIdRequest) (*api.AddressResponse, error) {
	rec, err := server.Store.Queries.GetAddress(ctx, int64(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &api.AddressResponse{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, ParentId: rec.ParentID.Int32}, nil
}

func (server *GRPCServer) DeleteAddress(ctx context.Context, req *api.DeleteAddressRequest) (*api.Empty, error) {
	err := server.Store.Queries.DeleteAddress(ctx, int64(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}
