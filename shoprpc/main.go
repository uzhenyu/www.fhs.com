package main

import (
	"flag"
	grpc2 "google.golang.org/grpc"
	"zg5/z311/framework/app"
	"zg5/z311/framework/grpc"
	"zg5/z311/shoprpc/api"
	"zg5/z311/shoprpc/conts"
	"zg5/z311/shoprpc/model"
)

var (
	port = flag.Int("port", 8081, "wzy")
)

func main() {
	flag.Parse()
	err := app.Init(conts.FileName, "mysql")
	if err != nil {
		return
	}

	err = model.AutoTable()
	if err != nil {
		return
	}

	err = grpc.GetGrpc(int64(*port), func(c *grpc2.Server) {
		api.Reg(c)
	}, conts.FileName)
	if err != nil {
		return
	}
}
