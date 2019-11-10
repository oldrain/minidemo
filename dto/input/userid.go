package input

type UserIdIn struct {
	Id int `json:"id" regexp:"^[0-9]{0,9}$" tips:"Invalid id"`
}
