package mysql

import (
	"database/sql"

	"github.com/ranit1803/students-api/internal/config"
	_ "github.com/go-sql-driver/mysql"
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