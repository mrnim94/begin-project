package main

import (
	"github.com/labstack/echo/v4"
	"begin-project/log"
)

func init() {
	os.Setenv("APP_NAME", "XXXX")
	logger := log.InitLogger(false)
	// Check if KUBERNETES_SERVICE_HOST is set
	if _, exists := os.LookupEnv("KUBERNETES_SERVICE_HOST"); !exists {
		// If not in Kubernetes, set LOG_LEVEL to DEBUG
		os.Setenv("LOG_LEVEL", "DEBUG")
	}
	logger.SetLevel(log.GetLogLevel("LOG_LEVEL"))
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
