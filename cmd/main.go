package main

import (
	"fmt"
	"go-gin-api-medium/pkg/books"
	"go-gin-api-medium/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	fmt.Printf("%s", dbUrl)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
