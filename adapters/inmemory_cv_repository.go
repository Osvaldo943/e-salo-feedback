package adapters

import (
	"errors"

	"github.com/Osvaldo943/domain"
)

type InmemoryCVRepository struct {
	cvs map[string]domain.CV
}

func NewInmemoryCVRepository() *InmemoryCVRepository {
	return &InmemoryCVRepository{cvs: make(map[string]domain.CV)}
}

func (r *InmemoryCVRepository) Save(cv domain.CV) error {
	r.cvs[cv.ID().Value()] = cv
	return nil
}

func (r *InmemoryCVRepository) FindById(id domain.ID) (domain.CV, error) {
	cv, exists := r.cvs[id.Value()]

	if !exists {
		return domain.CV{}, errors.New("CV não existe")
	}

	return cv, nil
}

func (r *InmemoryCVRepository) Get() []domain.CV {
	cvs := make([]domain.CV, 0, len(r.cvs))
	for _, cv := range r.cvs {
		cvs = append(cvs, cv)
	}
	return cvs
}
