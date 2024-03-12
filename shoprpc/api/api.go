package api

import (
	"google.golang.org/grpc"
	"zg5/z311/message/user"
)

func Reg(s grpc.ServiceRegistrar) {
	user.RegisterUserServer(s, ServerUser{})
}
