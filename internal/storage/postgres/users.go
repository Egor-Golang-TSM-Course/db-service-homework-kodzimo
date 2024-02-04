package postgres

import (
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID   int    `json:"user_id"`
	Name string `json:"user_name"`
	Role Role   `json:"user_role"`
}

type Role struct {
	Read   bool
	Write  bool
	Delete bool
}

func (p *PostgresDB) GetAllUsers() (map[int]User, error) {

	var users map[int]User

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
