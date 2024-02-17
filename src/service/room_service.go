package service

import (
	"context"

	"github.com/andres-mfv/chat-backend/src/room"
)

type RoomService interface {
	CreateRoom(ctx context.Context, req *room.Request) error
}

type roomService struct {
	roomMng room.Manager
}

func (r *roomService) CreateRoom(ctx context.Context, req *room.Request) error {
	_, e := r.roomMng.CreateRoom(ctx, req)
	return e
}

func NewRoomService() RoomService {
	roomMng := room.NewRoomManager()
	return &roomService{
		roomMng: roomMng,
	}
}
