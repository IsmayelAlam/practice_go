package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "delete", -1, "Delete todo from the list by id")
	flag.StringVar(&cf.Edit, "edit", "", "Edit todo from the list by id:title")
	flag.BoolVar(&cf.List, "list", false, "Show all added todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Mark todo as done")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("please use id:tile format to edit the todo.")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	default:
		fmt.Println("Invalid Command")
	}
}
