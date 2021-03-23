package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Cfg  *ini.File
	Conf Config
)

type DefaultConfig struct {
	RunMode  string `ini:"run_mode"`
	LogLevel string `ini:"log_level"`
}

type ServerConfig struct {
	HttpIp       string `ini:"http_ip"`
	HttpPort     string `ini:"http_port"`
	ReadTimeout  int    `ini:"read_timeout"`
	WriteTimeout int    `ini:"write_timeout"`
}

type Config struct {
	DefaultConf DefaultConfig
	ServerConf  ServerConfig
}

func InitCfg(config string) (err error) {
	Cfg, err = ini.Load(config)
	if err != nil {
		fmt.Printf("failed to load %s with err: %s\n", config, err.Error())
		return
	}

	dc, err := LoadDefaultConfig()
	if err != nil {
		fmt.Printf("failed to load default config with err: %s\n", err.Error())
		return
	}

	sc, err := LoadServerConfig()
	if err != nil {
		fmt.Printf("failed to load default config with err: %s\n", err.Error())
		return
	}

	Conf = Config{
		DefaultConf: dc,
		ServerConf:  sc,
	}

	return
}

func LoadDefaultConfig() (dc DefaultConfig, err error) {
	sec, err := Cfg.GetSection("")
	if err != nil {
		return
	}

	err = sec.MapTo(&dc)
	if err != nil {
		return
	}

	return
}

func LoadServerConfig() (sc ServerConfig, err error) {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		return
	}

	err = sec.MapTo(&sc)
	if err != nil {
		return
	}

	return
}
