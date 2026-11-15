package main

import (
	"fmt"
	"football-club/internal/database"
	"football-club/internal/routes"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//ini the db
	database.ConnectDB()
	// set up routes
	r := routes.SetUpRouter()

	err := r.Run(":8080")
	if err != nil {
		fmt.Print("Error starting server: ", err.Error())
	}
}
