package repository_test

import (
	"testing"

	"github.com/iamseki/dev-to/domain"
	"github.com/iamseki/dev-to/infra/repository"
)

func TestInMemoryRepository(t *testing.T) {
	m := repository.NewInMemoryRepository()

	err := m.Add(domain.Event{Title: "fake title"})
	if err != nil {
		t.Error("expect err to be nil but got:", err)
	}

	events, err := m.Get("fake title")
	if len(events) == 0 {
		t.Error("expect len(events) > 0 but got:", len(events))
	}

	if err != nil {
		t.Error("expect err to be nil but got:", err)
	}

	events, err = m.Get("")
	// 5 because there are already 4 events populated
	if len(events) != 5 {
		t.Error("expect len(events) equals to 1 but got:", len(events))
	}

	if err != nil {
		t.Error("expect err to be nil but got:", err)
	}
}

func TestInMemorySingleton(t *testing.T) {
	// as a singleton, must already have a "fake event" inserted on previous test
	m := repository.NewInMemoryRepository()

	events, _ := m.Get("fake title")
	if len(events) == 0 {
		t.Error("expect len(events) > 0 due to singleton but got:", len(events))
	}
}
