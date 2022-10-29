package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/enylvia/boilerplate-go/domain/model"
	"github.com/enylvia/boilerplate-go/domain/service"
	"github.com/enylvia/boilerplate-go/utils"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{service}
}

type PostController interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
}

func (c *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	// Add code here
	post := &model.Post{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.BadRequestResponse(err, "Failed to read body")
		return
	}
	err = json.Unmarshal(body, post)
	if err != nil {
		utils.BadRequestResponse(err, "Failed to unmarshal body")
		return
	}
	err = c.service.CreatePost(post)
	if err != nil {
		utils.BadRequestResponse(err, "Failed to create post")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.SuccessResponse("Successfully created post"))
}

func (c *Controller) GetPost(w http.ResponseWriter, r *http.Request) {
	// Add code here
	posts, err := c.service.GetPost()
	if err != nil {
		utils.BadRequestResponse(err, "Failed to get post")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.SuccessResponse(posts))
}
