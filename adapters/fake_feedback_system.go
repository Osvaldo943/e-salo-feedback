package adapters

import "github.com/Osvaldo943/domain"

type FakeFeedbackSystem struct {
	Response string
	Err      error
}

func (f *FakeFeedbackSystem) Analyze(cv domain.CV) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}

	return f.Response, nil
}
