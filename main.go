package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type task struct {
	id   int
	name string
	desc string
}

func newTask(id int, name, desc string) task {
	return task{id: id, name: name, desc: desc}
}

func main() {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Println("failed to open connection: ", err)
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Println("failed to perform ping(): ", err)
	}

	fmt.Println("Connected successfully :)")

	// Connected successfully , so let's insery insert some data

	prepStatementInsertion, err := db.Prepare("INSERT INTO task VALUES( ?, ?, ? )")
	if err != nil {
		fmt.Println("failed to prepare statements: ", err)
	}
	defer prepStatementInsertion.Close()

	fmt.Println("Prepared statment successfully !!")

	insertionResult, err := prepStatementInsertion.Exec("45452", "1222", "45454")
	if err != nil {
		fmt.Println("failed to insert the data: ", err)
	}

	id, err := insertionResult.LastInsertId()
	if err != nil {
		fmt.Printf("failed to insert the to table: %s", err.Error())
	}

	fmt.Printf("%d\t", id)
	fmt.Println("Successfully inserted the data !!")

	// Let's read the data from database !! :)
	prepStatementSelection, err := db.Prepare("SELECT * FROM task")
	if err != nil {
		fmt.Println("failed to preapre SQL statement")
	}
	defer prepStatementSelection.Close()

	rows, err := prepStatementSelection.Query()
	if err != nil {
		fmt.Println("failed to read the rows: ", err)
	}
	defer rows.Close()

	var list []task
	for rows.Next() {
		var id int
		var name, desc string
		err = rows.Scan(&id, &name, &desc)

		list = append(list, newTask(id, name, desc))
	}

	fmt.Println("Reading the data done !:)", list)

	// Let's update the data
	prepStatementUpdation, err := db.Prepare("UPDATE task SET name=?,  description=? WHERE id=?")
	if err != nil {
		fmt.Println("failed to update the sql prepared statement: ", err)
	}
	defer prepStatementUpdation.Close()

	updationResult, err := prepStatementUpdation.Exec("learn", "learn_more", "34")
	if err != nil {
		fmt.Println("failed to execute the update statment: ", err)
	}

	hasUpdated, err := updationResult.RowsAffected()
	if err != nil {
		fmt.Println("failed to update the record it seems")
	}

	if hasUpdated > 0 {
		fmt.Println("It seems we update the resource easily !!")
	} else {
		fmt.Println("failed to update the resource !!")
	}

	fmt.Println("Somehow update operation completed !!")

	// Let's delete some data
	prepStatementDeletion, err := db.Prepare("DELETE FROM task WHERE id=?")
	if err != nil {
		fmt.Println("failed to prepare statment to delete: ", err)
	}
	defer prepStatementDeletion.Close()

	deletionResult, err := prepStatementDeletion.Exec("34")
	if err != nil {
		fmt.Println("error occurred while deleting the data: ", err)
	}

	hasDeleted, err := deletionResult.RowsAffected()
	if err != nil {
		fmt.Println("while checking effected rows , I got the data: ", err)
	}

	if hasDeleted > 0 {
		fmt.Println("deleted successfully :)")
	} else {
		fmt.Println("failed to delete the rows")
	}

	fmt.Println("Somehow delete operation completed successfully")
}
