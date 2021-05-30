package repo_mock

import (
	"errors"
	"github.com/ruspatrick/tasks/internal/models"
	"sync"
)

type Repo struct {
	lastIndex int64
	r         sync.Map
}

var ErrNoSuchKey = errors.New("repo: нет такой таски")

func New() *Repo {
	return &Repo{
		lastIndex: 0,
		r:         sync.Map{},
	}
}

func (r *Repo) CreateTask(task models.Task) (models.Task, error) {
	task.Id = r.lastIndex

	r.r.Store(task.Id, task)

	r.lastIndex++

	return task, nil
}

func (r Repo) GetTask(id int64) (models.Task, error) {
	task, ok := r.r.Load(id)
	if !ok {
		return models.Task{}, ErrNoSuchKey
	}

	return task.(models.Task), nil
}

func (r Repo) DeleteTask(id int64) error {
	panic("implement me")
}
