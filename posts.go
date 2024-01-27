package main

import "time"

type Post struct {
	ID           int       `json:"post_id"`
	Header       string    `json:"post_header"`
	Body         string    `json:"post_body"`
	Comments     []Comment `json:"post_comments"`
	CreationDate time.Time `json:"creation_date"`
}

func createPost() {}

func (p *Post) displayPost(id int) {}

func (p *Post) deletePost(id int) {}

func displayAllPosts() {}
