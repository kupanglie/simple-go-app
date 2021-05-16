package helper

import (
	"strconv"

	grpc_user "github.com/kupanglie/simple-go-app/grpc/user"
	"github.com/kupanglie/simple-go-app/service/user/internal/constant"
)

func GRPCErrorGenerator(code int) *grpc_user.Error {
	return &grpc_user.Error{
		Code:    strconv.Itoa(code),
		Message: constant.ERROR_MAPPING[code],
	}
}
