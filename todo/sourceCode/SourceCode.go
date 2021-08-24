package source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type List struct {
	Task string
	Check    bool
}

type Lists struct {
	List []List
}

var file = "todo.csv"

var Task = Lists{}

//func init(){
//	Todo()
//}

func Todo() {
	file, err := ioutil.ReadFile("todo.csv")
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(file, &Task)
}

func (l Lists) add() {
	write, _ := json.Marshal(l)

	_ = ioutil.WriteFile(file, write, 9999)
}


func (l *Lists) Add(Tasks string) {
	list := List{Task: Tasks}
	l.List = append(l.List, list)

	l.add()
}



func (l *Lists) Done(i int) {
	if len(l.List) < i {
		fmt.Println("wrong number")
	} else {
		l.List[i-1].Check = true
		l.add()
	}
}

func (l *Lists) Undone(value int) {
	if len(l.List) < value {
		fmt.Println("wrong number")
	} else {
		l.List[value-1].Check = false
		l.add()
	}
}


func (l *Lists) Clear(){
	l.List = nil
	fmt.Println(l.List, "Cleaned List")
	l.add()
}

func (l *Lists) AllList() {
	for idx,val := range l.List {
		if val.Check {
			continue
		}
		fmt.Printf("%d \t %v\n", idx+1 , val.Task)
	}
}
