package main

import (
	"github.com/cs-sysimpl/TasCalnder/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// 環境変数の読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// サーバーを立てる
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
	}))

	//ルーティング
	router.Routing(e)

	//ポート:9000で開始
	e.Start(":9000")
}
