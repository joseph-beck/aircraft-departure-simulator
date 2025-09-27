package service

import (
	"fmt"

	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"
)

type ManageService struct {
	db *db.DB
}

func NewManageService(db *db.DB) *ManageService {
	return &ManageService{
		db: db,
	}
}

func (s *ManageService) Ping(args string, reply *string) error {
	*reply = fmt.Sprintf("ping: %s", args)
	return nil
}
