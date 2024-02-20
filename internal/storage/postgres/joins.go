package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Подключение к базе данных
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
	//
	//	// INNER JOIN
	//	fmt.Println("INNER JOIN")
	//	innerJoin(db)
	//
	//	// LEFT JOIN
	//	fmt.Println("LEFT JOIN")
	//	leftJoin(db)
	//
	//	// RIGHT JOIN
	//	fmt.Println("RIGHT JOIN")
	//	rightJoin(db)
	//
	//	// FULL JOIN
	//	fmt.Println("FULL JOIN")
	//	fullJoin(db)
	//}
	//
	//func innerJoin(db *sql.DB) {
	//	rows, err := db.Query(`
	//        SELECT s.StudentName, c.CourseName
	//        FROM Students s
	//        INNER JOIN Enrollments e ON s.StudentID = e.StudentID
	//        INNER JOIN Courses c ON e.CourseID = c.CourseID
	//    `)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer rows.Close()
	//
	//	for rows.Next() {
	//		var studentName, courseName string
	//		if err := rows.Scan(&studentName, &courseName); err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Printf("%s is enrolled in %s\n", studentName, courseName)
	//	}
	//}
	//
	//func leftJoin(db *sql.DB) {
	//	// Аналогичная логика для LEFT JOIN
	//}
	//
	//func rightJoin(db *sql.DB) {
	//	// Аналогичная логика для RIGHT JOIN
	//}
	//
	//func fullJoin(db *sql.DB) {
	//	// Аналогичная логика для FULL JOIN
}
