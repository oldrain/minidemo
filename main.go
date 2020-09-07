package main

import (
	"github.com/oldrain/minigo"
	"minidemo/app/entry"
)

func main() {
	// go run main.go -env=[dev/sit/uat/prod]
	// fyi: Initialization configuration in lib/config.go

	router := minigo.Default()

	entry.InitEntry(router)

	_ = router.Run(":9527")
}
