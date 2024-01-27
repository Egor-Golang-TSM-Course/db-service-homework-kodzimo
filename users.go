package main

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

var Admin Role = Role{Read: true, Write: true, Delete: true}
var Member Role = Role{Read: true, Write: true, Delete: false}
var Guest Role = Role{Read: true, Write: false, Delete: false}
