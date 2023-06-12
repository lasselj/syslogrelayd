package main

import (
	"syslogrelayd/http_server"
	"syslogrelayd/syslog_client"
)

type AppController struct {
	done                   chan struct{}
	httpServerController   *http_server.Controller
	syslogClientController *syslog_client.Controller
}

func NewAppController() *AppController {

	appController := &AppController{
		done:                   make(chan struct{}),
		syslogClientController: &syslog_client.Controller{},
		httpServerController:   &http_server.Controller{},
	}

	return appController
}

func (appController *AppController) Run() {
}
