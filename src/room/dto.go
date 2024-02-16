package room

type Request struct {
	Name     string `json:"name"`
	CreateBy string `json:"create_by"`
}

type JoinRequest struct {
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type Room struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CreateBy string `json:"create_by"`
}
