package web

import (
	"github.com/gin-gonic/gin"
)

// func InitDB() (db *gorm.DB, err error) {
// 	dsn := "host=192.168.3.150 user=misen password=root dbname=theatre port=5432 sslmode=disable TimeZone=Europe/Tirane"
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, err
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "SMPP server running on port 9999")
	})

	return r
}