package api

import (
	"context"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zg5/z311/message/user"
	"zg5/z311/shoprpc/service"
)

type ServerUser struct {
	user.UnimplementedUserServer
}

func (s ServerUser) Add(ctx context.Context, request *user.AddRequest) (*user.AddResponse, error) {
	create, err := service.Create(request.Info)
	if err != nil {
		return nil, err
	}
	if create.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "添加失败")
	}
	shopInfos, err := service.Select()
	if err != nil {
		return nil, err
	}
	for _, v := range shopInfos {
		err = service.Add(v.ID, v.Num, float64(v.Price), v.Name)
		if err != nil {
			return nil, err
		}
	}
	return &user.AddResponse{Info: create}, nil
}

func (s ServerUser) Search(ctx context.Context, request *user.SearchRequest) (*user.SearchResponse, error) {
	search, err := service.Search(request.Name)
	if err != nil {
		return nil, err
	}
	j, _ := json.Marshal(search)
	logs.Info(j)
	return &user.SearchResponse{Info: j}, nil
}

func (s ServerUser) UpdateNum(ctx context.Context, request *user.UpdateNumRequest) (*user.UpdateNumResponse, error) {
	err := service.Delete(request.ID)
	if err != nil {
		return nil, err
	}
	return &user.UpdateNumResponse{}, nil
}
