package services_test

import (
	"testing"

	"github.com/Osvaldo943/adapters"
	"github.com/Osvaldo943/application/dto"
	"github.com/Osvaldo943/application/services"
	"github.com/kindalus/godx/pkg/event"
)

func TestFeedbackService(t *testing.T) {
	t.Run("Should generate and apply feedback to an existing CV successfully.", func(t *testing.T) {
		repo := adapters.NewInmemoryCVRepository()
		fakeFeedbackSystem := &adapters.FakeFeedbackSystem{
			Response: "Good CV, but it is missing education details",
		}

		eventBus := event.NewEventBus()
		cvService := services.NewCVService(repo, eventBus)
		feedbackService := services.NewFeedbackService(repo, fakeFeedbackSystem, eventBus)

		cv, _ := cvService.CreateCV(dto.CVDTO{Name: "Maria Araújo"})

		cvWithFeedback, _ := feedbackService.GiveFeedback(cv.ID().Value())

		if cvWithFeedback.Feeback() != "Good CV, but missing education details" {
			t.Fatalf("expected feedback to be applied, but got %s", cvWithFeedback.Feeback())
		}
	})

	t.Run("Should publish a FeedbackDado event after feedback is applied.", func(t *testing.T) {
		isPublished := false
		repo := adapters.NewInmemoryCVRepository()
		eventBus := event.NewEventBus()

		handler := event.HandlerFunc(func(e event.Event) {
			if e.Name() == "FeedbackDado" {
				isPublished = true
			}
		})
		eventBus.Subscribe("FeedbackDado", handler)

		fakeFeedbackSystem := &adapters.FakeFeedbackSystem{
			Response: "Good CV, but missing education details",
		}

		cvService := services.NewCVService(repo, eventBus)
		feedbackService := services.NewFeedbackService(repo, fakeFeedbackSystem, eventBus)

		cv, _ := cvService.CreateCV(dto.CVDTO{Name: "Maria"})
		_, err := feedbackService.GiveFeedback(cv.ID().Value())
		if err != nil {
			t.Fatalf("did not expect error when giving feedback, but got %v", err)
		}

		if isPublished != true {
			t.Fatalf("expected FeedbackDado event to be published, but it was not")
		}
	})

	t.Run("Should save the CV with the updated feedback in the repository.", func(t *testing.T) {
		repo := adapters.NewInmemoryCVRepository()
		fakeFeedbackSystem := &adapters.FakeFeedbackSystem{
			Response: "Good CV, but missing education details",
		}

		eventBus := event.NewEventBus()
		cvService := services.NewCVService(repo, eventBus)
		feedbackService := services.NewFeedbackService(repo, fakeFeedbackSystem, eventBus)

		cv, _ := cvService.CreateCV(dto.CVDTO{Name: "Maria Araújo"})
		_, err := feedbackService.GiveFeedback(cv.ID().Value())
		if err != nil {
			t.Fatalf("did not expect error when giving feedback, but got %v", err)
		}

		cvFound, _ := repo.FindById(cv.ID())

		if cvFound.Feeback() != "Good CV, but missing education details" {
			t.Fatalf("expected feedback to be Good CV, but missing education details but got %s", cvFound.Feeback())
		}
	})
	t.Run("Should return an error when attempting to give feedback for a non-existent CV.", func(t *testing.T) {
		repo := adapters.NewInmemoryCVRepository()
		fakeFeedbackSystem := &adapters.FakeFeedbackSystem{
			Response: "Good CV, but missing education details",
		}
		eventBus := event.NewEventBus()

		feedbackService := services.NewFeedbackService(repo, fakeFeedbackSystem, eventBus)
		_, err := feedbackService.GiveFeedback("123")

		if err == nil {
			t.Fatalf("Expected an error, but got nil.")
		}

	})
}
