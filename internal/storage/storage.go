package storage

import "github.com/ranit1803/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetStudents()([]types.Student, error)
}