package main

func main() {
	todos := Todos{}
	store := NewStorage[Todos]("storage.json")
	store.Load(&todos)
	todos.add("Buy bread")
	todos.add("Buy milk")
	todos.toggle(0)
	todos.print()
	store.Save(todos)
}
