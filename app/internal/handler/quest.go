package handler

import (
	"app/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) questCreate(c *gin.Context) {
	var input db.Quest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,
			"Bad request")
		return
	}

	Create := db.Quest{
		Name: input.Name,
		Cost: input.Cost,
		Time: input.Time,
	}

	err := h.services.QuestCreate(&Create)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
