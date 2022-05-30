package user

type User struct {
	Id      int64  `db:"Id,omitempty"`
	Name    string `db:"Name"`
	Age     int64  `db:"Age"`
	City    string `db:"City"`
	Country string `db:"Country"`
}
