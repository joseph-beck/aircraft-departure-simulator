package service

import "github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"

type ManageService struct {
	db db.DB
}

func NewManageService() *ManageService {
	return &ManageService{}
}
