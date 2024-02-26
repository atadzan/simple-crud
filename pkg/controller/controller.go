package controller

import "github.com/atadzan/simple-crud/pkg/repository"

type Controller struct {
	repo repository.Repo
}

func New(repo repository.Repo) *Controller {
	return &Controller{repo: repo}
}

func Init() {

}
