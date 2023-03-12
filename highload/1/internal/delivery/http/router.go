package http

import (
	"github.com/gin-gonic/gin"
	"github.com/keepcalmist/otus/highload/1/internal/database"
	"github.com/keepcalmist/otus/highload/1/internal/manager/users"
	"log"
)

// NewRouter создает роутер для сервера
// после создания роутера необходимо вызвать http.ListenAndServe()
func NewRouter(repo *database.Repo) *gin.Engine {
	r := gin.Default()

	usersManager := users.New(repo)
	usersHandler := NewUsersHandler(usersManager)

	app := r.Group("/app")

	app.POST("/user/register", usersHandler.Register())
	app.GET("/user/get/:id", usersHandler.GetByID())

	r.POST("/login", usersHandler.Login())

	for _, info := range r.Routes() {
		log.Printf("%s %s", info.Method, info.Path)
	}
	//r.Group()
	return r
}
