package entity

type DisplayCart struct {
	Id int
	Name        string
	Description string
	Published   string
}

type AddCart struct {
	Id     int
	UserId int
	GameId int
}

type DeleteCart struct {
	Id     int
	UserId int
	GameId int
}
