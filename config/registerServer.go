package config

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
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
	registration.ID = "111"
	registration.Name = "go-consul-test"
	registration.Port = 3636
	registration.Tags = []string{"go-consul-test"}
	registration.Address = "127.0.0.1"

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)

}
