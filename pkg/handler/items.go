package handler

import (
	"net/http"
	"strconv"
	"todo-app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user's id")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get list_id")
		return
	}
	var item todo.ToDoItem
	if err := c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't parse request data")
		return
	}
	itemId, err := h.services.ToDoItem.Create(userId, listId, item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't create item")
		return
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": itemId,
		})
	}

}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get list with such id")
	}

	items, err := h.services.ToDoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get items from list with this id")
	} else {
		c.JSON(http.StatusOK, items)
	}

}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid item_id")
	}

	item, err := h.services.ToDoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get item with such id")
	} else {
		c.JSON(http.StatusOK, item)
	}

}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid item_id")
	}

	var updateDataInput todo.UpdateItemInput
	if err := c.BindJSON(&updateDataInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't parse update item input")
	}
	if err := h.services.ToDoItem.Update(userId, itemId, updateDataInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't update item with such id")
	} else {
		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}

}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid item_id")
	}
	if err := h.services.ToDoItem.Delete(userId, itemId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't delete item with such id")
	} else {
		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}

}