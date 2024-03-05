package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    api := r.Group("/api")
    {
        v1 := api.Group("/v1")
        {
            v1.POST("/register", v1.RegisterUser)
        }
    }
    r.Run()
}
