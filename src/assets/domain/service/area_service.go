package service

import (
	"github.com/Tanibox/tania-server/src/assets/domain"
	"github.com/Tanibox/tania-server/src/assets/query"
	"github.com/Tanibox/tania-server/src/assets/storage"
	uuid "github.com/satori/go.uuid"
)

type AreaServiceInMemory struct {
	FarmReadQuery      query.FarmReadQuery
	ReservoirReadQuery query.ReservoirReadQuery
}

func (s AreaServiceInMemory) FindFarmByID(uid uuid.UUID) (domain.AreaFarmServiceResult, error) {
	result := <-s.FarmReadQuery.FindByID(uid)

	if result.Error != nil {
		return domain.AreaFarmServiceResult{}, result.Error
	}

	farm, ok := result.Result.(storage.FarmRead)

	if !ok {
		return domain.AreaFarmServiceResult{}, domain.AreaError{Code: domain.AreaErrorFarmNotFound}
	}

	if farm == (storage.FarmRead{}) {
		return domain.AreaFarmServiceResult{}, domain.AreaError{Code: domain.AreaErrorFarmNotFound}
	}

	return domain.AreaFarmServiceResult{
		UID:  farm.UID,
		Name: farm.Name,
	}, nil
}

func (s AreaServiceInMemory) FindReservoirByID(reservoirUID uuid.UUID) (domain.AreaReservoirServiceResult, error) {
	result := <-s.ReservoirReadQuery.FindByID(reservoirUID)

	if result.Error != nil {
		return domain.AreaReservoirServiceResult{}, result.Error
	}

	res, ok := result.Result.(storage.ReservoirRead)

	if !ok {
		return domain.AreaReservoirServiceResult{}, domain.AreaError{Code: domain.AreaErrorReservoirNotFound}
	}

	if res.UID == (uuid.UUID{}) {
		return domain.AreaReservoirServiceResult{}, domain.AreaError{Code: domain.AreaErrorReservoirNotFound}
	}

	return domain.AreaReservoirServiceResult{
		UID:  res.UID,
		Name: res.Name,
	}, nil
}
