package main

import (
	"github.com/labstack/echo/v4"
	"begin-project/log"
)

func init() {
	os.Setenv("APP_NAME", "XXXX")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
