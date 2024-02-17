package room

type Request struct {
	Name     string `json:"name"`
	CreateBy int64  `json:"create_by"`
	Type     int    `json:"type"`
}

type JoinRequest struct {
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type Room struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CreateBy int64  `json:"create_by"`
}
