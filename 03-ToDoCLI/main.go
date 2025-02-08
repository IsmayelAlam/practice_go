package main

func main() {
	todo := Todos{}
	todo.add("hello")
	todo.add("world")
	todo.toggle(1)
	todo.add("delete1")
	todo.add("delete2")
	todo.delete(2)
	todo.edit(2, "edit")
	todo.print()
}
