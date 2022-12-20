package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/danh996/go-destiny/api"
	db "github.com/danh996/go-destiny/db/sqlc"
	"github.com/danh996/go-destiny/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router:= gin.New()
	router.Use(gin.Logger())
	
}
