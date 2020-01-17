package test

import (
	"code.byted.org/baike/mykite/data"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/koding/kite"
	"log"
	"strconv"
	"testing"
)

// 注册服务到consul
func ConsulRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "134.175.80.121:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "222"
	registration.Name = "go-consul-mykite"
	registration.Port = 3636
	registration.Tags = []string{"go-consul-mykite"}
	registration.Address = "134.175.80.121"

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	//check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
}

// 从consul中发现服务
func ConsulFindServer() []data.Service {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "134.175.80.121:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 获取所有service
	services, _ := client.Agent().Services()
	for _, value := range services {
		fmt.Println(value.Address)
		fmt.Println(value.Port, value.ID, value.Service)
	}

	if _, found := services["111"]; found {
		log.Println("serverNode_1 found")
	}

	for _, service := range services {
		thisService := data.Service{Port: strconv.Itoa(service.Port), Ip: service.Address}
		data.SericeList = append(data.SericeList, thisService)
	}
	return data.SericeList
}

func TestRegister(t *testing.T) {
	//ConsulRegister()
	ConsulFindServer()
	thisService := data.SericeList[0]
	k := kite.New("exp2", "1.0.0")

	// Connect to our math kite
	mathWorker := k.NewClient("http://" + thisService.Ip + ":" + thisService.Port + "/kite")
	mathWorker.Dial()

	response, _ := mathWorker.Tell("square", 6) // call "square" method with argument 4
	fmt.Println("result new :", response.MustFloat64())
}
