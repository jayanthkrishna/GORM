package main

import (
	"GORM-Practice/controllers"
	"GORM-Practice/database"
	"GORM-Practice/populate"
	"encoding/json"
	"fmt"
)

func main() {

	database.LoadDB()

	populate.Populate()

	d := controllers.FacultyModel{}

	fac, stu, _ := d.FindAll()

	r, _ := json.Marshal(fac[1])

	fmt.Println("Result after seeding  faculty :", string(r))

	r, _ = json.Marshal(stu[1])

	fmt.Println("Result after seeding  student :", string(r))

}
