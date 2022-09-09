package controllers

import (
	"GORM-Practice/database"
	"GORM-Practice/models"
)

type FacultyModel struct {
}

func (facultyModel FacultyModel) FindAll() ([]models.Faculty, []models.Student, error) {

	var faculties []models.Faculty
	database.DB.Preload("Students").Find(&faculties)
	var students []models.Student
	database.DB.Preload("Faculty").Find(&students)
	return faculties, students, nil

}
