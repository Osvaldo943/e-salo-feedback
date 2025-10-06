package services

import (
	"github.com/Osvaldo943/domain"
	"github.com/Osvaldo943/ports"
	"github.com/kindalus/godx/pkg/event"
)

type FeedbackService struct {
	repo           ports.CVRepository
	feedbackSystem ports.FakeFeedbackSystem
	bus            event.Bus
}

func NewFeedbackService(repo ports.CVRepository, feedbackSystem ports.FakeFeedbackSystem, bus event.Bus) *FeedbackService {
	return &FeedbackService{repo: repo, feedbackSystem: feedbackSystem, bus: bus}
}

func (f *FeedbackService) GiveFeedback(id string) (domain.CV, error) {
	cvID, err := domain.NewIDFromString(id)
	if err != nil {
		return domain.CV{}, err
	}

	cv, err := f.repo.FindById(cvID)
	if err != nil {
		return domain.CV{}, err
	}

	feedback, err := f.feedbackSystem.Analyze(cv)
	if err != nil {
		return domain.CV{}, err
	}

	if err = cv.UpdateFeedback(feedback); err != nil {
		return domain.CV{}, err
	}

	f.repo.Save(cv)

	events := cv.PullEvents()
	for _, event := range events {
		f.bus.Publish(event)
	}

	return cv, nil
}
