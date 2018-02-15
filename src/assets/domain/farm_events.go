package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type FarmCreated struct {
	UID         uuid.UUID
	Name        string
	Type        string
	Latitude    string
	Longitude   string
	CountryCode string
	CityCode    string
	IsActive    bool
	CreatedDate time.Time
}

type FarmGeolocationChanged struct {
	FarmUID   uuid.UUID
	Latitude  string
	Longitude string
}

type FarmRegionChanged struct {
	FarmUID     uuid.UUID
	CountryCode string
	CityCode    string
}
