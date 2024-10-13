package main

type RoomData struct {
	Id       string
	Name     string
	UserName string
}

type AvailavleRoomData struct {
	Rooms []RoomData
}

type UserData struct {
	Name  string
	Score int
}

type GameOverData struct {
	WinnerName string
	Score      int
}
