package todo

import (
	"fmt"
	"strconv"

	"github.com/augustoroman/promise"
	dom "honnef.co/go/js/dom/v2"
)

var (
	running  uint
	token    string
	filter   string = "all"
	todos    []Todo
	doc      dom.Document
	body     *dom.HTMLBodyElement
	todoapp  *dom.BasicHTMLElement
	newtodo  *dom.HTMLInputElement
	todolist *dom.HTMLUListElement
	count    *dom.HTMLSpanElement
	footer   *dom.BasicHTMLElement
	footerUL *dom.HTMLUListElement
	clearBtn *dom.HTMLButtonElement

	todoLI []*dom.HTMLLIElement
)

func Start() {
	doc = dom.GetWindow().Document()

	body = doc.GetElementByID("bodypage").(*dom.HTMLBodyElement)
	todoapp = doc.GetElementByID("todoapp").(*dom.BasicHTMLElement)

	newtodo = doc.GetElementByID("new-todo").(*dom.HTMLInputElement)
	todolist = doc.GetElementByID("todo-list").(*dom.HTMLUListElement)
	count = doc.GetElementByID("todo-count").(*dom.HTMLSpanElement)

	footer = doc.GetElementByID("footer").(*dom.BasicHTMLElement)
	footerUL = doc.CreateElement("ul").(*dom.HTMLUListElement)
	clearBtn = doc.CreateElement("button").(*dom.HTMLButtonElement)

	footer.AppendChild(footerUL)
	footer.AppendChild(clearBtn)

	newtodo.AddEventListener("keyup", false, AddTodoEvent)

	refreshTodoList()
}

func FetchTodoList() {
	var p = &promise.Promise{}

	p.Then(
		func(value interface{}) interface{} {
			todos = value.([]Todo)
			refreshTodoList()
			return p
		}, func(value interface{}) interface{} {
			println("error", value)
			return p
		},
	)

	PromiseTodoList(token, p)
}

func AddTodo(val string) {
	newToDo := Todo{ID: running, Text: val}
	todos = append(todos, newToDo)
	refreshTodoList()
	running++
}

func refreshTodoList() {
	for _, li := range todoLI {
		todolist.RemoveChild(li)
	}

	todoLI = []*dom.HTMLLIElement{}

	for _, todo := range todos {
		if filter == "active" && todo.Completed {
			continue
		}
		if filter == "completed" && !todo.Completed {
			continue
		}

		li := doc.CreateElement("li").(*dom.HTMLLIElement)

		div := doc.CreateElement("div").(*dom.HTMLDivElement)
		div.SetClass("view")
		cb := doc.CreateElement("input").(*dom.HTMLInputElement)
		cb.SetType("checkbox")
		cb.SetClass("toggle")
		cb.SetID("toggle" + strconv.Itoa(int(todo.ID)))
		cb.SetDefaultChecked(todo.Completed)

		lb := doc.CreateElement("label").(*dom.HTMLLabelElement)
		lb.SetInnerHTML(todo.Text)
		btn := doc.CreateElement("button").(*dom.HTMLButtonElement)
		btn.SetClass("destroy")
		btn.SetValue(strconv.Itoa(int(todo.ID)))
		btn.AddEventListener("click", false, RemoveTodoEvent)

		div.AppendChild(cb)
		div.AppendChild(lb)
		div.AppendChild(btn)

		inp := doc.CreateElement("input").(*dom.HTMLInputElement)
		inp.SetClass("edit")

		li.AppendChild(div)
		li.AppendChild(inp)
		if todo.Completed {
			li.SetClass("completed")
		} else {
			li.SetClass("todo-list")
		}
		li.SetAttribute("data-id", strconv.Itoa(int(todo.ID)))

		todolist.AppendChild(li)
		todoLI = append(todoLI, li)

		refreshFooter()
	}
}

func refreshFooter() {
	count.SetInnerHTML(fmt.Sprintf("<strong>%d</strong> %s left", left(), unitWord()))
	footer.RemoveChild(footerUL)
	footer.RemoveChild(clearBtn)

	footerUL = doc.CreateElement("ul").(*dom.HTMLUListElement)
	footerUL.SetID("filters")

	liall := doc.CreateElement("li").(*dom.HTMLLIElement)
	aall := doc.CreateElement("a").(*dom.HTMLAnchorElement)
	aall.SetHref("#/all")
	aall.SetText("All")
	aall.AddEventListener("click", false, ClickFilterEvent)

	if filter == "all" {
		aall.SetClass("selected")
	}

	liall.AppendChild(aall)

	liact := doc.CreateElement("li").(*dom.HTMLLIElement)
	aact := doc.CreateElement("a").(*dom.HTMLAnchorElement)
	aact.SetHref("#/active")
	aact.SetText("Active")
	aact.AddEventListener("click", false, ClickFilterEvent)

	if filter == "active" {
		aact.SetClass("selected")
	}

	liact.AppendChild(aact)

	licomp := doc.CreateElement("li").(*dom.HTMLLIElement)
	acomp := doc.CreateElement("a").(*dom.HTMLAnchorElement)
	acomp.SetHref("#/completed")
	acomp.SetText("Completed")
	acomp.AddEventListener("click", false, ClickFilterEvent)

	if filter == "completed" {
		acomp.SetClass("selected")
	}

	licomp.AppendChild(acomp)

	footerUL.AppendChild(liall)
	footerUL.AppendChild(liact)
	footerUL.AppendChild(licomp)

	footer.AppendChild(footerUL)

	clearBtn = doc.CreateElement("button").(*dom.HTMLButtonElement)
	clearBtn.SetID("clear-completed")
	clearBtn.SetInnerHTML("Clear completed")
	clearBtn.AddEventListener("click", false, ClickClearCompletedEvent)

	footer.AppendChild(clearBtn)
}

func left() int {
	i := 0
	for _, todo := range todos {
		if !todo.Completed {
			i++
		}
	}
	return i
}

func unitWord() string {
	if left() == 1 {
		return "item"
	}
	return "items"
}
