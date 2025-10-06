package services

import (
	"github.com/Osvaldo943/application/dto"
	"github.com/Osvaldo943/domain"
	"github.com/Osvaldo943/ports"
	"github.com/kindalus/godx/pkg/event"
)

type CVService struct {
	repository ports.CVRepository
	bus        event.Bus
}

func NewCVService(repo ports.CVRepository, bus event.Bus) *CVService {
	return &CVService{repository: repo, bus: bus}
}

func (c *CVService) CreateCV(cvDTO dto.CVDTO) (domain.CV, error) {
	name, err := domain.NewName(cvDTO.Name)
	if err != nil {
		return domain.CV{}, err
	}

	cv, err := domain.NewCV(name)
	if err != nil {
		return domain.CV{}, err
	}

	if err = c.repository.Save(cv); err != nil {
		return domain.CV{}, err
	}

	events := cv.PullEvents()
	for _, event := range events {
		c.bus.Publish(event)
	}

	return cv, nil
}
