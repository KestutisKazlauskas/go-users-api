package users
// access to database for user code of database only herer

import (
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-utils/logger"
	"github.com/KestutisKazlauskas/go-users-api/datasources/mysql/users_db"
	"github.com/KestutisKazlauskas/go-users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, created_at, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status=?;"
	queryFindEmailAndPassword = "SELECT id, first_name, last_name, email, created_at, status, password FROM users WHERE email=?"
)

func (user *User) Get() *rest_errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryGetUser)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing mysql statment", err, logger.Log)
	}
	defer statment.Close()

	result := statment.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *rest_errors.RestErr {

	statment, err := users_db.Clinet.Prepare(queryInsertUser)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing mysql statment", err, logger.Log)
	}
	//important to close connection after code execution
	defer statment.Close()
	insertResult, insertErr := statment.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Status, user.Password)

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

func (user *User) Update() *rest_errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryUpdateUser)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing stamtent", err, logger.Log)
	}
	defer statment.Close()

	_, err = statment.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil

}

func (user *User) Delete() *rest_errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryDeleteUser)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing mysql statment", err, logger.Log)
	}
	defer statment.Close()

	if _, err = statment.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr) {
	statment, err := users_db.Clinet.Prepare(queryFindByStatus)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error preparing mysql statment", err, logger.Log)
	}
	defer statment.Close()

	rows, err := statment.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) ==  0 {
		return nil, rest_errors.NewNotFoundError("No users found.")
	}

	return results, nil
}

func (user *User) FindByEmail() *rest_errors.RestErr {
	statment, err := users_db.Clinet.Prepare(queryFindEmailAndPassword)
	if err != nil {
		return rest_errors.NewInternalServerError("Error preparing mysql statment", err, logger.Log)
	}
	defer statment.Close()

	result := statment.QueryRow(user.Email)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status, &user.Password); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}