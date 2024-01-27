package main

import "time"

type Comment struct {
	ID           int       `json:"id"`
	Author       string    `json:"author"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
}

func (c *Comment) createComment(Comment) {}

func (c *Comment) deleteComment(id int) {}
