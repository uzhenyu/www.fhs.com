package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"zg5/z311/framework/vipers"
)

func NewClient(name, fileName string) error {
	err := vipers.GetYaml(fileName)
	if err != nil {
		return err
	}
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	err = c.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    name,
		Tags:    []string{"GRPC"},
		Port:    viper.GetInt("Grpc.Port"),
		Address: viper.GetString("Nacos.Ip"),
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%v:%v", viper.GetString("Nacos.Ip"), viper.GetInt("Grpc.Port")),
			DeregisterCriticalServiceAfter: "30s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

//func C(name,fileName string) error {
//	c,err := api.NewClient(api.DefaultConfig())
//	if err != nil{
//		return err
//	}
//	err = c.Agent().ServiceRegister(&api.AgentServiceRegistration{
//		ID:                uuid.New().String(),
//		Name:              name,
//		Tags:              []string{"GRPC"},
//		Port:              8081,
//		Address:           "127.0.0.1",
//		Check:             &api.AgentServiceCheck{
//			Interval:                       "5s",
//			Timeout:                        "5s",
//			GRPC:                           "127.0.0.1:8081",
//			DeregisterCriticalServiceAfter: "30s",
//		},
//	})
//	return nil
//}

func c(name, fileName string) error {
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	err = c.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    name,
		Tags:    []string{"GRPC"},
		Port:    8081,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           "1332422:8081",
			DeregisterCriticalServiceAfter: "30s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func NewClients(name, ip string, port int64) (string, int64, error) {
	c, err := api.NewClient(&api.Config{Address: fmt.Sprintf("%v:%v", ip, "8500")})
	if err != nil {
		return "", 0, err
	}
	_, data, err := c.Agent().AgentHealthServiceByName(name)
	if err != nil {
		return "", 0, err
	}
	return data[0].Service.Address, int64(data[0].Service.Port), nil
}
