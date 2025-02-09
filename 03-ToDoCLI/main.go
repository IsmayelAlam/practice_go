package main

func main() {
	todo := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todo)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todo)
	storage.save(todo)
}
