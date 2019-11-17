package config

import (
	"code.byted.org/baike/mykite/consts"
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
	registration.ID = consts.CONSUL_REGISTER_ID
	registration.Name = consts.CONSUL_REGISTER_NAME
	registration.Port = 3636
	registration.Tags = []string{consts.CONSUL_REGISTER_NAME}
	registration.Address = "134.175.80.121"

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	//健康检查的url
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/kite")
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
