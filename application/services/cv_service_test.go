package services_test

import (
	"testing"

	"github.com/Osvaldo943/adapters"
	"github.com/Osvaldo943/application/dto"
	"github.com/Osvaldo943/application/services"
	"github.com/kindalus/godx/pkg/event"
)

func TestCVService(t *testing.T) {

	t.Run("Should create a CV successfully", func(t *testing.T) {
		cv := dto.CVDTO{Name: "Maria Araújo"}

		eventBus := event.NewEventBus()
		repo := adapters.NewInmemoryCVRepository()
		service := services.NewCVService(repo, eventBus)

		_, err := service.CreateCV(cv)
		if err != nil {
			t.Fatalf("expected CV to be created successfully, but got error: %v", err)
		}
	})

	t.Run("Should publish CVCreated event when a CV is created", func(t *testing.T) {
		isPublished := false
		cv := dto.CVDTO{Name: "Maria Araújo"}

		eventBus := event.NewEventBus()
		handler := event.HandlerFunc(func(e event.Event) {
			if e.Name() == "CVCriado" {
				isPublished = true
			}
		})
		eventBus.Subscribe("CVCriado", handler)

		repo := adapters.NewInmemoryCVRepository()
		service := services.NewCVService(repo, eventBus)
		service.CreateCV(cv)

		if !isPublished {
			t.Fatalf("expected CVCriado event to be published, but it was not")
		}
	})

	t.Run("Should fail to create CV when name is empty", func(t *testing.T) {
		cv := dto.CVDTO{Name: ""}

		eventBus := event.NewEventBus()
		repo := adapters.NewInmemoryCVRepository()
		service := services.NewCVService(repo, eventBus)

		_, err := service.CreateCV(cv)
		if err == nil {
			t.Fatalf("expected error when creating CV with empty name, but got nil")
		}
	})

	t.Run("Should store CV in repository after successful creation", func(t *testing.T) {
		cv := dto.CVDTO{Name: "Maria Araújo"}

		eventBus := event.NewEventBus()
		repo := adapters.NewInmemoryCVRepository()
		service := services.NewCVService(repo, eventBus)

		cvCreated, err := service.CreateCV(cv)
		if err != nil {
			t.Fatalf("did not expect error when creating CV, but got: %v", err)
		}

		cvStored, err := repo.FindById(cvCreated.ID())
		if err != nil {
			t.Fatalf("expected to find stored CV in repository, but got error: %v", err)
		}

		if cvStored.ID() != cvCreated.ID() {
			t.Fatalf("expected stored CV ID to be %s, but got %s", cvCreated.ID(), cvStored.ID())
		}
	})
}
