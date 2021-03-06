package main

import (
	"github.com/GrappigPanda/Olivia/cache"
	"github.com/GrappigPanda/Olivia/config"
	"github.com/GrappigPanda/Olivia/network"
	"github.com/GrappigPanda/Olivia/network/message_handler"
)

func Init() {
	config := config.ReadConfig()

	internalCache := cache.NewCache()

	messageHandler := message_handler.NewMessageHandler()
	networkHandler.StartIncomingNetwork(
		messageHandler,
		internalCache,
		config,
		nil,
	)
}

func main() {
	Init()
}
