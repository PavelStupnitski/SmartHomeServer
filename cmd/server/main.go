package main

import (
	"smartHome/pkg/mqtt"
	"smartHome/pkg/opensearch"
	"smartHome/pkg/server"
)

func main() {
	opensearch.Init()
	mqtt.Init()
	server.Init()
}
