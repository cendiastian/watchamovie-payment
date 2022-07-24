package main

import (
	"os"
	"watchamovie-payment/driver"
	_middleware "watchamovie-payment/middleware"
	"watchamovie-payment/routes"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	params := os.Args
	paramsLength := len(params)
	driver.InitDB()
	e := routes.New()

	// Log Middleware
	_middleware.LogMiddlewareInit(e)
	if paramsLength < 2 {
		log.Println("Please add SERVER or CRONJOB along with go run main.go command")
		log.Println("SERVER or CRONJOB not found")
		e.Start(":" + port)
	}
	if paramsLength > 1 {
		inputMethod := os.Args[1]
		valid := IsValidInputMethod(inputMethod)
		if valid {
			if inputMethod == "SERVER" {
				// Starting The Server
				e.Start(":" + port)
			}
		}

		// if inputMethod == "CRONJOB" {
		// presenter := factory.Init()
		// // Cron Job
		// log.Info("Create new cron")
		// c := cron.New()
		// c.AddFunc("@every 3s", func() {
		// 	fmt.Println("Every minute")
		// })

		// log.Info("Start cron")
		// c.Start()
		// time.Sleep(10080 * time.Hour)
		// }
	}
}

func IsValidInputMethod(method string) bool {
	switch method {
	case
		"SERVER":
		// "CRONJOB":
		return true
	}
	return false
}
