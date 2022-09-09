package populate

import (
	"GORM-Practice/database"
	"GORM-Practice/models"
	"log"
)

func Populate() {
	faculty := []models.Faculty{

		{
			Name: "Faculty 1",
		},
		{
			Name: "Faculty 2",
		},
		{
			Name: "faculty 3",
		},
	}

	students := []models.Student{
		{
			Name:      "Name 1",
			Address:   "Address 1",
			FacultyID: 1,
		},
		{
			Name:      "Name 2",
			Address:   "Address 2",
			FacultyID: 1,
		},
		{
			Name:      "Name 3",
			Address:   "Address 3",
			FacultyID: 2,
		},
		{
			Name:      "Name 4",
			Address:   "Address 4",
			FacultyID: 2,
		},
		{
			Name:      "Name 5",
			Address:   "Address 5",
			FacultyID: 2,
		},
	}

	for i, _ := range faculty {
		err := database.DB.Create(&faculty[i]).Error
		if err != nil {
			log.Fatal("Cannot seed faculty table :", err)
		}
	}

	for i, _ := range students {
		err := database.DB.Create(&students[i]).Error
		if err != nil {
			log.Fatal("Cannot seed students table :", err)
		}
	}
}
