package main

import (
	bloglog "blog/log"
	"blog/webfrm"
	"context"
)

var fileLogger bloglog.Logger

func main() {
	log, err := bloglog.NewConsoleLogger()
	if err != nil {
		log.Fatal(err.Error())
	}
	webfrm, err := webfrm.NewWebfrm("configd", &log)
	if nil != err {
		log.ErrErr("webfrm.NewWebSrv failed", err)
		return
	}
	fileLogger = webfrm.Logger

	apiConfigRouter := webfrm.Router.Group("/api/config")
	setConfigRouter(apiConfigRouter)

	err = webfrm.Start(&fileLogger)
	if err != nil {
		fileLogger.ErrErr("webfrm.Start failed", err)
	}

	webfrm.WaitForExit(&fileLogger)
	webfrm.Stop(context.Background(), &fileLogger)
}
