package grpc

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"zg5/z311/framework/nacos"
)

type T struct {
	App struct {
		Ip   string `json:"Ip"`
		Port string `json:"Port"`
	} `json:"App"`
}

func Client(fileName string) (*grpc.ClientConn, error) {
	config, err := nacos.GetConfig(fileName)
	if err != nil {
		return nil, err
	}
	cnf := new(T)
	err = json.Unmarshal([]byte(config), &cnf)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(fmt.Sprintf("%v:%v", cnf.App.Ip, cnf.App.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
