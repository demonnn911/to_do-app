package repository

import (
	"fmt"
	"strings"
	todo "todo-app/app-models"

	"github.com/jmoiron/sqlx"
)

type ToDoItemPostgres struct {
	db *sqlx.DB
}

func NewToDoItemPostgres(db *sqlx.DB) *ToDoItemPostgres {
	return &ToDoItemPostgres{
		db: db,
	}
}

func (r *ToDoItemPostgres) Create(listId int, item todo.ToDoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var id int
	queryToDoItem := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemTable)
	row := tx.QueryRow(queryToDoItem, item.Title, item.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	queryListItem := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES($1, $2)", listItemsTable)
	if _, err := tx.Exec(queryListItem, listId, id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ToDoItemPostgres) GetAll(listId int) ([]todo.ToDoItem, error) {
	var itemData []todo.ToDoItem
	//	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
	//		todoItemTable, listItemsTable, usersListsTable)
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id WHERE li.list_id = $1", todoItemTable, listItemsTable)
	if err := r.db.Select(&itemData, query, listId); err != nil {
		return nil, err
	}
	return itemData, nil
}

func (r *ToDoItemPostgres) GetById(userId, itemId int) (todo.ToDoItem, error) {
	var item todo.ToDoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2", todoItemTable, listItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *ToDoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ti.id = $1 AND ul.user_id = $2", todoItemTable, listItemsTable, usersListsTable)
	if _, err := r.db.Exec(query, itemId, userId); err != nil {
		return err
	}
	return nil
}

func (r *ToDoItemPostgres) Update(userId, itemId int, updateData todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if updateData.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, updateData.Title)
		argId++
	}
	if updateData.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, updateData.Description)
		argId++
	}
	if updateData.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, updateData.Done)
		argId++
	}
	querySetValues := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ti.id = $%d AND ul.user_id = $%d",
		todoItemTable, querySetValues, listItemsTable, usersListsTable, argId, argId+1)
	args = append(args, itemId, userId)
	_, err := r.db.Exec(query, args...)
	return err
}
