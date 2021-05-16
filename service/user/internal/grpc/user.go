package grpc

import (
	"context"
	"log"

	grpc_user "github.com/kupanglie/simple-go-app/grpc/user"
	"github.com/kupanglie/simple-go-app/service/user/helper"
	"github.com/kupanglie/simple-go-app/service/user/internal/constant"
	"github.com/kupanglie/simple-go-app/service/user/internal/usecase"
)

type usersServer struct {
	userUC usecase.User
}

func NewUserGRPC(userUC usecase.User) *usersServer {
	return &usersServer{
		userUC: userUC,
	}
}

func (u *usersServer) Add(ctx context.Context, in *grpc_user.AddRequest) (*grpc_user.AddResponse, error) {
	if in == nil || in.GetName() == "" {
		return &grpc_user.AddResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.EMPTY_REQUEST),
		}, nil
	}

	resp, err := u.userUC.Add(ctx, in.GetName())
	if err != nil {
		log.Println("[GRPC] Fail to Add User", err)
		return &grpc_user.AddResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.INTERNAL_SERVER_ERROR),
		}, nil
	}

	res := &grpc_user.AddResponse{
		Data: &grpc_user.Data{
			Id:        resp.ID,
			Name:      resp.Name,
			CreatedAt: resp.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: resp.UpdatedAt.Format("02-01-2006 15:04:05"),
			DeletedAt: resp.DeletedAt.Format("02-01-2006 15:04:05"),
		},
		Error: &grpc_user.Error{},
	}

	return res, nil
}

func (u *usersServer) Find(ctx context.Context, in *grpc_user.FindRequest) (*grpc_user.FindResponse, error) {
	if in == nil || in.GetId() == "" {
		return &grpc_user.FindResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.EMPTY_REQUEST),
		}, nil
	}

	resp, err := u.userUC.Find(ctx, in.GetId())
	if err != nil {
		log.Println("[GRPC] Fail to Find User", err)
		return &grpc_user.FindResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.INTERNAL_SERVER_ERROR),
		}, nil
	}

	res := &grpc_user.FindResponse{
		Data: &grpc_user.Data{
			Id:        resp.ID,
			Name:      resp.Name,
			CreatedAt: resp.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: resp.UpdatedAt.Format("02-01-2006 15:04:05"),
			DeletedAt: resp.DeletedAt.Format("02-01-2006 15:04:05"),
		},
		Error: &grpc_user.Error{},
	}

	return res, nil
}

func (u *usersServer) Edit(ctx context.Context, in *grpc_user.EditRequest) (*grpc_user.EditResponse, error) {
	if in == nil || in.GetId() == "" || in.GetName() == "" {
		return &grpc_user.EditResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.EMPTY_REQUEST),
		}, nil
	}

	resp, err := u.userUC.Edit(ctx, in.GetId(), in.GetName())
	if err != nil {
		log.Println("[GRPC] Fail to Find User", err)
		return &grpc_user.EditResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.INTERNAL_SERVER_ERROR),
		}, nil
	}

	res := &grpc_user.EditResponse{
		Data: &grpc_user.Data{
			Id:        resp.ID,
			Name:      resp.Name,
			CreatedAt: resp.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: resp.UpdatedAt.Format("02-01-2006 15:04:05"),
			DeletedAt: resp.DeletedAt.Format("02-01-2006 15:04:05"),
		},
		Error: &grpc_user.Error{},
	}

	return res, nil
}

func (u *usersServer) Delete(ctx context.Context, in *grpc_user.DeleteRequest) (*grpc_user.DeleteResponse, error) {
	if in == nil || in.GetId() == "" {
		return &grpc_user.DeleteResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.EMPTY_REQUEST),
		}, nil
	}

	resp, err := u.userUC.Delete(ctx, in.GetId())
	if err != nil {
		log.Println("[GRPC] Fail to Find User", err)
		return &grpc_user.DeleteResponse{
			Data:  &grpc_user.Data{},
			Error: helper.GRPCErrorGenerator(constant.INTERNAL_SERVER_ERROR),
		}, nil
	}

	res := &grpc_user.DeleteResponse{
		Data: &grpc_user.Data{
			Id:        resp.ID,
			Name:      resp.Name,
			CreatedAt: resp.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: resp.UpdatedAt.Format("02-01-2006 15:04:05"),
			DeletedAt: resp.DeletedAt.Format("02-01-2006 15:04:05"),
		},
		Error: &grpc_user.Error{},
	}

	return res, nil
}
