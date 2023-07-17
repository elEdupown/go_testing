package modelos

import (
	"time"
)

type User struct {
	Id        int
	Nombre    string
	CreatedAt time.Time
	status    bool
}

func (this *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	this.Id = id
	this.Nombre = name
	this.CreatedAt = createdAt
	this.status = status
}
