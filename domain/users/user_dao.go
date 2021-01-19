package users

import (
	"github.com/uuthman/bookstore_users-api/logger"
	"fmt"
	"github.com/uuthman/bookstore_users-api/datasources/mysql/users_db"
	"github.com/uuthman/bookstore_users-api/utils/errors"
)

const(

	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created,password,status) values (?,?,?,?,?,?);"
	queryGetUser = "SELECT id,first_name,last_name,email,date_created,status from users where id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users where status = ?;"
)

func (user *User) Get() *errors.RestErr{
	stmt,err := users_db.Client.Prepare(queryGetUser)
	if err != nil{
		logger.Error("error when trying to prepare get user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	
	if getErr := result.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); getErr != nil{

		logger.Error("error when trying to get user by id",getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr{

	stmt,err := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		logger.Error("error when trying to prepare save user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()


	insertResult,err := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated,user.Password,user.Status)
 
	if err != nil{
		logger.Error("error when trying to  save user ",err)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil{
		logger.Error("error when trying to get last insert id after creating a new user",err)
		return errors.NewInternalServerError("database error")
		
	}

	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestErr{
	stmt,err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		logger.Error("error when trying to prepare update user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_,err = stmt.Exec(user.FirstName,user.LastName,user.Email,user.ID)
	if err != nil{
		logger.Error("error when trying to update user",err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr{
	stmt,err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		logger.Error("error when trying to prepare delete user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_,err = stmt.Exec(user.ID)
	if err != nil{
		logger.Error("error when trying to delete user",err)
		return errors.NewInternalServerError("database error")
	}

	return nil


}

func (user *User) FindByStatus(status string) ([]User,*errors.RestErr){
	stmt,err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil{
		logger.Error("error when trying to prepare find user by status statement",err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows,err := stmt.Query(status)
	if err != nil{
		logger.Error("error when trying to prepare find user by status",err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User,0)
	for rows.Next(){
		var u User
		if err := rows.Scan(&u.ID,&u.FirstName,&u.LastName,&u.Email,&u.DateCreated,&u.Status); err != nil{
			logger.Error("error when trying to scan user row into user struct",err)
		return nil, errors.NewInternalServerError("database error")
		}
		results = append(results,u)
	}

	if len(results) == 0{
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s",status))
	}

	return results,nil
}
