package books

import "github.com/gin-gonic/gin"

func StartApp() {
	r := gin.Default()
	r = AddBookRoutes(r)
	r.Run()
}
