package room

import (
	"context"

	"github.com/andres-mfv/chat-backend/src/db"
)

type Manager interface {
	CreateRoom(ctx context.Context, req *Request) (*Room, error)
	DeleteRoom() error
	JoinRoom() error
}

type roomManager struct {
	dbInstance db.PostgresDB
}

func (r *roomManager) CreateRoom(ctx context.Context, req *Request) (*Room, error) {
	exec, err := r.dbInstance.DB.Exec("INSERT INTO rooms (name, create_by) VALUES ($1, $2)", req.Name, req.CreateBy)
	if err != nil {
		return nil, err
	}

	lastInsertID, _ := exec.LastInsertId()
	return &Room{
		ID:       lastInsertID,
		Name:     req.Name,
		CreateBy: req.CreateBy,
	}, nil
}

func (r *roomManager) DeleteRoom() error {
	//TODO implement me
	panic("implement me")
}

func (r *roomManager) JoinRoom() error {
	//TODO implement me
	panic("implement me")
}

func NewRoomManager() Manager {
	return &roomManager{}
}
