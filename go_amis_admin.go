package main

import (
	"flag"
	"fmt"
	"github.com/fs714/go-amis-admin/api"
	"github.com/fs714/go-amis-admin/utils/config"
	"github.com/fs714/go-amis-admin/utils/log"
	"github.com/fs714/go-amis-admin/utils/version"
	"net/http"
	"os"
	"time"
)

var isShowVersion bool
var confPath string

func init() {
	flag.BoolVar(&isShowVersion, "v", false, "Show Version")
	flag.StringVar(&confPath, "c", "", "Config Path")
	flag.Parse()

	if isShowVersion {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	if confPath == "" {
		if _, err := os.Stat("conf/go_amis_admin.conf"); err == nil {
			confPath = "conf/go_amis_admin.conf"
		} else if _, err := os.Stat("/etc/go_amis_admin/go_amis_admin.conf"); err == nil {
			confPath = "/etc/go_amis_admin/go_amis_admin.conf"
		} else {
			fmt.Println("invalid config path " + confPath)
			os.Exit(1)
		}
	}

	err := config.InitCfg(confPath)
	if err != nil {
		fmt.Println("failed to initialize configuration")
		os.Exit(1)
	}

	err = log.SetLevel(config.Conf.DefaultConf.LogLevel)
	if err != nil {
		fmt.Println("failed to set log level")
		os.Exit(1)
	}

	err = log.SetFormat("text")
	if err != nil {
		fmt.Println("failed to set log format")
		os.Exit(1)
	}

	log.SetOutput(os.Stdout)
}

func main() {
	log.Infoln("------ Start go_amis_admin ------")
	log.Infof("%s build on %s", version.BaseVersion, version.BuildTime)
	log.Infof("Git Commit on %s", version.GitVersion)
	log.Infoln(version.GoVersion)

	router := api.InitRouter()

	log.Infof("Listening on %s:%s", config.Conf.ServerConf.HttpIp, config.Conf.ServerConf.HttpPort)
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", config.Conf.ServerConf.HttpIp, config.Conf.ServerConf.HttpPort),
		Handler:        router,
		ReadTimeout:    time.Duration(config.Conf.ServerConf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Conf.ServerConf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
