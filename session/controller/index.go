package controller

import (
	"fmt"
	"golang-demo/session/session"
	"net/http"
)

type IndexController struct {
	session *session.Manager
}

func NewIndexController(sessionManager *session.Manager) *IndexController {
	return &IndexController{
		session: sessionManager,
	}
}

func (index *IndexController) Login(w http.ResponseWriter, r *http.Request) {
	sess := index.session.SessionStart(w, r)
	sess.Set("username", "tom")
}

func (index *IndexController) Home(w http.ResponseWriter, r *http.Request) {
	sess := index.session.SessionStart(w, r)
	username := sess.Get("username")
	fmt.Printf("username: %v\n", username)
}
