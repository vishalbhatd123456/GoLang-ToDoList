package main
var addFlag string;
var listFlag bool;

import "fmt"

//global var
var tasks[] Task;
var nextID = 1;


type Task struct{
	ID int;
	Name string;
	Done bool;
}

func addTask(name string){
	task:=Task{
		ID: nextID,
		Name: name,
		Done:false,
	}
	tasks= append(tasks, task)
	nextID++;
}

func listTasks(){
	for _, task := range tasks{
		status := "Pending"
		if task.Done{
			status = "Done"
		}
		fmt.Printf("[%d] %s -%s\n",task.ID,task.Name,status);
	}
}

func main() {
	flag.StringVar(&addFlag,"add","","Add a new task");
	flag.BoolVar(&listFlag,"list", false, "List all tasks");
	flag.Parse();

	if addFlag!=""{
		addTask(addFlag);
		fmt.Println("Task added");
	}
	if listFlag{
		listTasks()
	}
}
