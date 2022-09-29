package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gunSlaveUnit/TaskTracker/pkg/entities"
)

func (h *Handler) SignUp(cxt *gin.Context) {
	var user entities.User

	if err := cxt.BindJSON(&user); err != nil {
		return
	}

	if _, err := h.service.Authorization.CreateUser(user); err != nil {
		return
	}
}

func (h *Handler) SignIn(cxt *gin.Context) {

}

func (h *Handler) SignOut(cxt *gin.Context) {

}
