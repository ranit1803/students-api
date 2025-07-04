package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ranit1803/students-api/internal/config"
	"github.com/ranit1803/students-api/internal/types"
)

type MySql struct {
	Db *sql.DB
}

//to make an instance of the sql this function is being made
func New(cfg *config.Config) (*MySql,error){
	db,err:= sql.Open("mysql", "root:ranit1803@/students_db")
	if err != nil{
		return nil, err
	}

	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id int primary key auto_increment,
    	name varchar(50),
    	email varchar(50),
    	age int
	)`)
	
	if err!=nil{
		return nil, err
	}
	
	return &MySql{
		Db: db,
	},nil
}

func (m *MySql) CreateStudent(name string, email string, age int) (int64,error){
	
	statement, err:= m.Db.Prepare(`insert into students(name, email, age) values(?, ?, ?)`)
	if err!= nil{
		return 0, nil
	}
	defer statement.Close()

	res, err:= statement.Exec(name, email, age)
	if err!=nil{
		return 0,err
	}

	lastid, err:= res.LastInsertId()
	if err!=nil{
		return 0, err
	}

	return lastid, nil
}

func (m *MySql) GetStudentByID(id int64) (types.Student, error){
	statement, err:= m.Db.Prepare(`select * from students where id = ? limit 1`)
	if err!= nil{
		return types.Student{}, nil
	}
	defer statement.Close()
	var student types.Student
	
	err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err!=nil{
		if err == sql.ErrNoRows{
			return types.Student{}, fmt.Errorf("no student found with id: %s",fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error %w", err)
	}
	return student,nil
}

func (m *MySql) GetStudents()([]types.Student, error){
	statement, err:= m.Db.Prepare(`select * from students`)
	if err!=nil{
		return nil, err
	}
	defer statement.Close()
	rows, err:= statement.Query()
	if err!=nil {
		return nil,err
	}
	
	defer rows.Close()
	var students []types.Student
	for rows.Next(){
		var student types.Student
		err:= rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err!=nil{
			return nil,err
		}

		students = append(students, student)
	}
	return students,nil
}