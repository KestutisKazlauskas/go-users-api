package users
// access to database for user code of database only herer

import (
	"github.com/KestutisKazlauskas/go-users-api/utils/errors"
	"github.com/KestutisKazlauskas/go-users-api/utils/date_utils"
	"github.com/KestutisKazlauskas/go-users-api/datasources/mysql/users_db"
	"github.com/KestutisKazlauskas/go-users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statment.Close()

	result := statment.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}



	return nil
}

func (user *User) Save() *errors.RestErr {

	statment, err := users_db.Clinet.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//important to close connection after code execution
	defer statment.Close()
	user.CreatedAt = date_utils.GetNowString()
	insertResult, insertErr := statment.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)

	if insertErr != nil {
		return mysql_utils.ParseError(insertErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statment.Close()

	_, err = statment.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil

}

func (user *User) Delete() *errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statment.Close()

	if _, err = statment.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

