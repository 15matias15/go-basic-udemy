package model

import "github.com/15matias15/go-basic-udemy/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	var todos []views.PostRequest
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	var todos []views.PostRequest
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}