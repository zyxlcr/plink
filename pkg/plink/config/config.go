package config

import "time"

type Config struct {
	Name string
	//服务绑定的IP地址
	IP string

	Heartbeat        time.Duration
	MaxMsgChanLen    int32
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32

	TcpConfig  *TcpConfig
	WsConfig   *WsConfig
	GrpcConfig *GrpcConfig
	MqttConfig *MqttConfig
}
type TcpConfig struct {
	//服务绑定的端口
	TcpPort int
	MaxConn uint32
}

type WsConfig struct {
	//服务绑定的端口
	WsPort  string
	MaxConn uint32
}
type GrpcConfig struct {
	//服务绑定的端口
	GrpcPort string
	MaxConn  uint32
}

type MqttConfig struct {
	//服务绑定的端口
	MqttPort string
	MaxConn  uint32
}

func NewConfig() *Config {
	t := &TcpConfig{
		TcpPort: 8989,
		MaxConn: 30,
	}
	w := &WsConfig{
		WsPort:  ":8988",
		MaxConn: 30,
	}
	g := &GrpcConfig{
		GrpcPort: ":8986",
		MaxConn:  5,
	}
	m := &MqttConfig{
		MqttPort: ":8984",
		MaxConn:  30,
	}
	return &Config{
		Name:      "plink",
		IP:        "0.0.0.0",
		Heartbeat: time.Second * 10,

		MaxMsgChanLen:    50,
		WorkerPoolSize:   5,
		MaxWorkerTaskLen: 11,

		TcpConfig:  t,
		WsConfig:   w,
		GrpcConfig: g,
		MqttConfig: m,
	}

}
