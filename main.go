package main

func main() {
	todos := Todos{}

	store := NewStorage[Todos]("data.json")
	store.Load(&todos)

	cmdFlag := NewCommandFlag()
	cmdFlag.Execute(&todos)

	store.Save(todos)
}
