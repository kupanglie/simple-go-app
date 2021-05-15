package user

import (
	"errors"
	"log"
	"time"

	"github.com/kupanglie/simple-go-app/gql/graph/model"
	user_proto "github.com/kupanglie/simple-go-app/grpc/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewUserClient() *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:8901", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Fail to connect GRPC")
	}

	return conn
}

type userHandler struct {
	userClient *grpc.ClientConn
}

func NewUserHandler() UserHandler {
	userClient := NewUserClient()

	return &userHandler{
		userClient: userClient,
	}
}

func (u *userHandler) UserAdd(ctx context.Context, name string) (*model.UserData, error) {
	var result *model.UserData

	req := &user_proto.AddRequest{
		Name: name,
	}

	resp, err := user_proto.NewUserHandlerClient(u.userClient).Add(ctx, req)
	if err != nil {
		return result, err
	}

	if resp.Error != nil {
		return result, errors.New(resp.Error.Message)
	}

	result.ID = resp.GetData().GetId()
	result.Name = resp.GetData().GetName()
	result.CreatedAt = resp.GetData().GetCreatedAt()
	result.UpdatedAt = resp.GetData().GetUpdatedAt()
	result.DeletedAt = resp.GetData().GetDeletedAt()

	return result, nil
}

func (u *userHandler) UserEdit(ctx context.Context, userID string, name string) (*model.UserData, error) {
	var result *model.UserData

	req := &user_proto.EditRequest{
		Id:   userID,
		Name: name,
	}

	resp, err := user_proto.NewUserHandlerClient(u.userClient).Edit(ctx, req)
	if err != nil {
		return result, err
	}

	if resp.Error != nil {
		return result, errors.New(resp.Error.Message)
	}

	result.ID = resp.GetData().GetId()
	result.Name = resp.GetData().GetName()
	result.CreatedAt = resp.GetData().GetCreatedAt()
	result.UpdatedAt = resp.GetData().GetUpdatedAt()
	result.DeletedAt = resp.GetData().GetDeletedAt()

	return result, nil
}

func (u *userHandler) UserDelete(ctx context.Context, userID string) (*model.UserData, error) {
	var result *model.UserData

	req := &user_proto.DeleteRequest{
		Id: userID,
	}

	resp, err := user_proto.NewUserHandlerClient(u.userClient).Delete(ctx, req)
	if err != nil {
		return result, err
	}

	if resp.Error != nil {
		return result, errors.New(resp.Error.Message)
	}

	result.ID = resp.GetData().GetId()
	result.Name = resp.GetData().GetName()
	result.CreatedAt = resp.GetData().GetCreatedAt()
	result.UpdatedAt = resp.GetData().GetUpdatedAt()
	result.DeletedAt = resp.GetData().GetDeletedAt()

	return result, nil
}

func (u *userHandler) UserFind(ctx context.Context, userID string) (*model.UserData, error) {
	var result *model.UserData

	req := &user_proto.FindRequest{
		Id: userID,
	}

	resp, err := user_proto.NewUserHandlerClient(u.userClient).Find(ctx, req)
	if err != nil {
		return result, err
	}

	if resp.Error != nil {
		return result, errors.New(resp.Error.Message)
	}

	result.ID = resp.GetData().GetId()
	result.Name = resp.GetData().GetName()
	result.CreatedAt = resp.GetData().GetCreatedAt()
	result.UpdatedAt = resp.GetData().GetUpdatedAt()
	result.DeletedAt = resp.GetData().GetDeletedAt()

	return result, nil
}
