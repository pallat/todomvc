package todo

import (
	"strconv"
	"strings"

	dom "honnef.co/go/js/dom/v2"
)

func RemoveTodoEvent(event dom.Event) {
	input := event.Target().(*dom.HTMLButtonElement)

	id, err := strconv.Atoi(input.Value())
	if err != nil {
		println("error", err.Error())
		return
	}

	for i := range todos {
		if todos[i].ID == uint(id) {
			todos[i].Completed = !todos[i].Completed
			refreshTodoList()
			return
		}
	}

}

func AddTodoEvent(event dom.Event) {
	ke := event.(*dom.KeyboardEvent)
	if ke.KeyCode() == 13 {
		input := event.Target().(*dom.HTMLInputElement)
		if input.Value() == "" {
			return
		}
		AddTodo(input.Value())
		input.SetValue("")
	}
}

func ClickFilterEvent(event dom.Event) {
	input := event.Target().(*dom.HTMLAnchorElement)
	switch {
	case strings.HasSuffix(input.Href(), "all"):
		filter = "all"
	case strings.HasSuffix(input.Href(), "active"):
		filter = "active"
	case strings.HasSuffix(input.Href(), "completed"):
		filter = "completed"
	}

	refreshTodoList()
}

func ClickClearCompletedEvent(event dom.Event) {
	bound := len(todos) - 1
	for i := bound; i >= 0; i-- {
		if todos[i].Completed {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
	refreshTodoList()
}
