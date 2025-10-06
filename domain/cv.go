package domain

import (
	"github.com/kindalus/godx/pkg/event"
)

type CV struct {
	id       ID
	name     Name
	feedback string
	events   []event.Event
}

func NewCV(name Name) (CV, error) {
	id, err := ID{}.GenerateNew()
	if err != nil {
		return CV{}, err
	}
	cv := CV{
		id:       id,
		name:     name,
		feedback: "",
	}
	cv.AddEvent("CVCriado")
	return cv, nil
}

func (c *CV) ID() ID {
	return c.id
}

func (c *CV) Feeback() string {
	return c.feedback
}
func (c *CV) UpdateFeedback(feedback string) error {
	c.feedback = feedback
	c.AddEvent("FeedbackDado")

	return nil
}

func (c *CV) Name() string {
	return c.name.Value()
}
func (c *CV) AddEvent(name string) {
	c.events = append(c.events, event.New(name, event.WithPayload(c)))
}

func (c *CV) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}

	return events
}
