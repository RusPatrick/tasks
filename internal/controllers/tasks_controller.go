package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ruspatrick/tasks/internal/models"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	r Repo
}

type Repo interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTask(id int64) (models.Task, error)
	DeleteTask(id int64) error
}

func New(r Repo) *Controller {
	return &Controller{r: r}
}

func (c Controller) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	task, err := c.r.CreateTask(task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	respBody, _ := json.Marshal(task)

	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func (c Controller) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)

	task, err := c.r.GetTask(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	respBody, _ := json.Marshal(task)

	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}
