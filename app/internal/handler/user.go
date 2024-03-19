package handler

import (
	"app/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) userCreate(c *gin.Context) {
	var input db.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,
			"Bad request")
		return
	}

	Create := db.User{
		Name: input.Name,
	}

	id, err := h.services.UserCreate(Create.Name)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) userInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var userList db.UserInfo

	userList, err = h.services.UserInfo(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance":   userList.Balance,
		"questList": userList.Quests,
	})
}

func (h *Handler) userQuest(c *gin.Context) {
	var input db.History
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,
			"Bad request")
		return
	}

	Quest := db.History{
		QuestID: input.QuestID,
		UserID:  input.UserID,
		Status:  input.Status,
	}

	err := h.services.UserQuest(&Quest)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
