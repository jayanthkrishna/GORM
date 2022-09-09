package models

import (
	"fmt"
)

type Faculty struct {
	Id       int `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string
	Students []Student `gorm:"ForeignKey:FacultyID"`
}

func (faculty *Faculty) TableName() string {
	return "faculty"
}

func (faculty Faculty) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", faculty.Id, faculty.Name)
}

type Student struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Name      string
	Address   string
	FacultyID int
	Faculty   Faculty
}

func (student *Student) TableName() string {
	return "student"
}

func (student Student) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\naddress: %s", student.Id, student.Name, student.Address)
}
