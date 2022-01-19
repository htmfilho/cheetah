package team

import "testing"

func TestManager_GetRandomGreeting(t *testing.T) {
	manager := &Manager{
		Name: "Justin Biber",
		Greetings: []string{"can you share your wisdom?",
			"have you contemplated the universe today?",
			"might or might not have something interesting to say.",
			"needs our attention.",
			"sometimes has nothing to say.",
			"is our boss, we should listen to him.",
			"might have missed most of this meeting, but that's ok.",
			"is here?"},
	}

	if greeting := manager.GetRandomGreeting(); greeting == "" {
		t.Errorf("Got empty greeting, but want something.")
	}
}
