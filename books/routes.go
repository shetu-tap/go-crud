package books

import "github.com/gin-gonic/gin"

func AddBookRoutes(router *gin.Engine) *gin.Engine {
	bookRepository := NewRepository()
	bookService := NewBookService(bookRepository)
	bookController := NewBookController(bookService)
	bookRouter := router.Group("/books")
	bookRouter.GET("", bookController.GetBooks)
	bookRouter.GET("/:id", bookController.GetDetail)
	bookRouter.PUT("/:id", bookController.UpdateBook)
	bookRouter.DELETE("/:id", bookController.DeleteBook)
	bookRouter.POST("", bookController.CreateBook)
	return router;
}
