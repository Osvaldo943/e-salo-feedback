package ports

import "github.com/Osvaldo943/domain"

type FakeFeedbackSystem interface {
	Analyze(id domain.CV) (string, error)
}
