package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"ojire/message"
	"ojire/router"
	"ojire/service"
	"os"
	"os/signal"
	"sync"
)

var (
	Environment *string
	Http        *bool
)

func init() {
	Environment = flag.String("config_name", "local", "define environment")
	Http = flag.Bool("http", false, "define environment")
	flag.Parse()
}

func main() {
	mainHandler := service.MakeHandler(context.Background(), Environment, Http)
	var wg sync.WaitGroup

	if *Http {
		wg.Add(1)
		go func() {
			defer wg.Done()
			RunWithHTTP(mainHandler)
		}()
	}

	wg.Wait()
}

func RunWithHTTP(mainHandler *service.HandlerSetup) {
	handlerRouter := router.NewHandlerRouter(mainHandler)
	r := handlerRouter.ListRouter()
	messageRun := fmt.Sprintf("ojire run on %s environment", mainHandler.Env.Application.Env)
	message.Log(nil, logrus.InfoLevel, messageRun, "SETUP ENV HTTP")

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := r.Shutdown(); err != nil {
			message.Log(nil, logrus.ErrorLevel, fmt.Sprintf("HTTP server Shutdown: %v", err), "ERROR STOPPING HTTP ojire")
		}
		close(idleConsClosed)
		message.Log(nil, logrus.InfoLevel, "Bye. All", "SUCCESS STOPPING HTTP ojire")
	}()
	message.Log(nil, logrus.InfoLevel, fmt.Sprintf("Listening on port %s", mainHandler.Env.Application.Port), "SUCCESS RUNNING HTTP ojire")
	if err := r.Listen(fmt.Sprintf(":%s", mainHandler.Env.Application.Port)); err != nil {
		message.Log(nil, logrus.FatalLevel, err.Error(), "ERROR START HTTP ojire")
	}
	<-idleConsClosed
}
