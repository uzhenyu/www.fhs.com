package grpc

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"zg5/z311/framework/consul"
	"zg5/z311/framework/vipers"
)

func GetGrpc(port int64, reg func(c *grpc.Server), fileName string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = vipers.GetYaml(fileName)
	if err != nil {
		return err
	}
	err = consul.NewClient(viper.GetString("Wzy.DataId"), fileName)
	s := grpc.NewServer()
	reflection.Register(s)
	reg(s)

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
