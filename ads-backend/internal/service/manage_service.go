package service

import (
	"fmt"
	"sync"

	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/aircraftconstants"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/maxweight"
)

type ManageService struct {
	db   *db.DB
	dbMu sync.Mutex
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

func (s *ManageService) GetAircraftConstants(args string, reply *[]aircraftconstants.AircraftConstant) error {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	aircraftConstants, err := s.db.GetAircraftConstants()
	if err != nil {
		return err
	}

	*reply = aircraftConstants
	return nil
}

func (s *ManageService) GetMaximumWeights(args string, reply *[]maxweight.MaxWeight) error {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	maxWeights, err := s.db.GetMaximumWeights()
	if err != nil {
		return err
	}

	*reply = maxWeights
	return nil
}
