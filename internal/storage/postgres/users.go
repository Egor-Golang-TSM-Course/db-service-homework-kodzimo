package postgres

import (
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID   *int    `json:"user_id"`
	Name string  `json:"user_name"`
	Role *string `json:"user_role"`
}

type Role struct {
	Read   bool
	Write  bool
	Delete bool
}

var admin *Role = &Role{Read: true, Write: true, Delete: true}
var member *Role = &Role{Read: true, Write: true, Delete: false}
var guest *Role = &Role{Read: true, Write: false, Delete: false}

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

	println(newUserName)

	sqlStatement := `
        INSERT INTO blog.users (name, role)
        VALUES ($1, $2);
        `
	_, err := p.DB.Exec(sqlStatement, newUserName, newUserRole)
	if err != nil {
		log.Error("Failed to insert row %d: %v", err)
		return err
	}
	return nil
}
