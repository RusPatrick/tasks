package tasks

import (
	"github.com/ruspatrick/tasks/internal/controllers"
	repo_mock "github.com/ruspatrick/tasks/internal/repositories/repo-mock"
	"github.com/ruspatrick/tasks/internal/router"
	"log"
	"net/http"
)

func main() {
	r := repo_mock.New()
	c := controllers.New(r)
	router := router.New(c)

	log.Println("started")
	http.ListenAndServe(":8080", router)
}
