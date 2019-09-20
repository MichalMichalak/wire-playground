package store

import "strconv"

type Options struct {
	Host string
	Port int
	DB   string
}

func NewPostgres(opt Options) Postgres {
	return Postgres{opt}
}

type Postgres struct {
	opt Options
}

func (pg Postgres) GetByID(id int) string {
	// Implementation doesn't really matter
	return strconv.Itoa(id)
}
