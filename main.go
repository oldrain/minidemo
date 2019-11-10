package main

import (
	"flag"
	"github.com/oldrain/minigo"
	"minidemo/app/controllers"
	"minidemo/lib"
)

func main() {
	// Environment
	env := *flag.String("env", "dev", "go run main.go -env=[dev/sit/uat/prod]")
	flag.Parse()

	// Init config
	err := lib.InitCfg("config", env)
	if err != nil {
		panic(err)
	}

	router := minigo.Default()

	controllers.InitCtl(router)

	_ = router.Run(":9527")
}
