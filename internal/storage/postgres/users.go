package postgres

import (
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID       *int    `json:"user_id"`
	Name     string  `json:"user_name"`
	Role     *string `json:"user_role"`
	Password string  `json:"user_password"`
}

type Role struct {
	Read   bool
	Write  bool
	Delete bool
}

//var admin *Role = &Role{Read: true, Write: true, Delete: true}
//var member *Role = &Role{Read: true, Write: true, Delete: false}
//var guest *Role = &Role{Read: true, Write: false, Delete: false}

func (p *PostgresDB) GetAllUsers() ([]User, error) {

	var users []User

	query := `SELECT * FROM blog.users;`
	err := p.DB.Select(&users, query)
	if err != nil {
		log.Error("Failed to select users: ", err)
		return nil, err
	}
	//rows, err := p.DB.Query(query)
	//if err != nil {
	//	log.Error("Failed to select students: ", err)
	//	return nil, err
	//}
	//for rows.Next() {
	//	var s Student
	//	err := rows.Scan(&s.ID, &s.Name, &s.BirthDate)
	//	if err != nil {
	//		log.Error("Failed to Scan() students: ", err)
	//		return nil, err
	//	}
	//	res = append(res, s)
	//}

	return users, nil
}

func (p *PostgresDB) RegisterUser(newUser User) error {
	//role := "admin"
	//var newUser = User{Name: "John", Role: &role}
	newUserName := newUser.Name
	newUserRole := newUser.Role
	newUserPassword := newUser.Password

	println(newUserName)

	sqlStatement := `
        INSERT INTO blog.users (name, role, password)
        VALUES ($1, $2, $3);
        `
	_, err := p.DB.Exec(sqlStatement, newUserName, newUserRole, newUserPassword)
	if err != nil {
		log.Error("Failed to insert row %d: %v", err)
		return err
	}
	return nil
}

func (p *PostgresDB) CheckCredentials(username, password string) bool {
	// Запрос к базе данных для получения пароля пользователя
	var dbPassword string
	err := p.DB.Get(&dbPassword, "SELECT password FROM blog.users WHERE name = $1", username)
	if err != nil {
		p.Logger.Println(err)
		return false
	}

	// Сравнение пароля из базы данных с введенным паролем
	return dbPassword == password
}
