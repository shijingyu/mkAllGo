package main

import (
	"github.com/gin-gonic/gin"
	_ "gopkg.in/mgo.v2/bson"
	_ "log"
)






func main() {
	r := gin.Default()
	r = CollectRoute(r)

	r.Run(":8080")
}


