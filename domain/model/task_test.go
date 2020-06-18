package model

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	cases := []struct {
		title   string
		content string
	}{
		{title: "test1", content: "success1"},
		{title: "test2", content: "success2"},
	}

	for i, c := range cases {
		task, err := NewTask(c.title, c.content)
		if err != nil {
			t.Error(err)
		}

		if task.Title != c.title {
			t.Errorf("case[%d] title expected: %s, actual: %s", i, c.title, task.Title)
		}
		if task.Content != c.content {
			t.Errorf("case[%d] content expected: %s, actual: %s", i, c.content, task.Content)
		}
	}
}
