package handler

import (
	"net/http"
	"strconv"
	todo "todo-app/app-models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input todo.ToDoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "could not parse body of request")
		return
	}
	listId, err := h.services.ToDoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

type getAllListResponse struct {
	Data []todo.ToDoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.ToDoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "couldn't get lists for this id")
		return
	}
	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can not get list_id from request")
		return
	}
	list, err := h.services.ToDoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "couldn't get lists for this id")
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can'not get list_id from request")
		return
	}

	if err := h.services.ToDoList.Delete(userId, listId); err != nil {
		newErrorResponse(c, http.StatusNoContent, "can not delete list")
		return
	} else {
		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}

}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list_id")
	}

	var updateData todo.UpdateListInput

	if err := c.BindJSON(&updateData); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid request body")
	}

	if err := h.services.ToDoList.Update(userId, listId, updateData); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't update list")
	} else {
		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}

}
