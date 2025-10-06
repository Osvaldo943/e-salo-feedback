package ports

import "github.com/Osvaldo943/domain"

type CVRepository interface {
	Save(cv domain.CV) error
	Get() []domain.CV
	FindById(id domain.ID) (domain.CV, error)
}
