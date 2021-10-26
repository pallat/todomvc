package todo

import (
	"strconv"

	dom "honnef.co/go/js/dom/v2"
)

func RemoveTodoEvent(event dom.Event) {
	input := event.Target().(*dom.HTMLButtonElement)
	println("remove at", input.Value())

	id, err := strconv.Atoi(input.Value())
	if err != nil {
		println("error", err.Error())
		return
	}

	for i := range todos {
		if todos[i].ID == uint(id) {
			// todos = append(todos[:i], todos[i+1:]...)
			todos[i].Completed = true
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
