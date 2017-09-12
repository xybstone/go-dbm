package main

import (
	"fmt"
	"graceful"
	"os"
	"os/signal"
	"syscall"
	"time"

	"apigateway.gaodun.com/api"
)

const (
	//NetworkAddr 端口
	NetworkAddr = "0.0.0.0:6064"
	maxWaitTime = 2 * time.Second
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt,
		syscall.SIGTERM, syscall.SIGINT,
		syscall.SIGHUP, syscall.SIGUSR1,
		syscall.SIGUSR2,
	)
	networkAddr := graceful.NewGracefulListener(NetworkAddr, maxWaitTime)
	go api.Run(networkAddr)
	s := <-exit
	networkAddr.Close()
	fmt.Println("shutting down server .", s)
}
