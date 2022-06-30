package model

import "fmt"

func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer deleteQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
