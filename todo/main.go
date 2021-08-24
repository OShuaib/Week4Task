/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"TodoList/cmd"
	_ "TodoList/sourceCode"
)

//
//type List struct {
//		Task string
//		Check    bool
//}
//
//type Lists struct {
//	List []List
//}
//
//var file = "todo.csv"
//
//var Task = Lists{}
//
//
//func Todo() {
//	file, err := ioutil.ReadFile("todo.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	_ = json.Unmarshal(file, &Task)
//}
//
//
//
//func (l *Lists) Add(Tasks string) {
//	list := List{Task: Tasks}
//	l.List = append(l.List, list)
//
//	l.add()
//}
//
//func (l Lists) add() {
//write, _ := json.Marshal(l)
//
//_ = ioutil.WriteFile(file, write, 0666)
//}
//
//func (l *Lists) Done(i int) {
//	if len(l.List) < i {
//		fmt.Println("wrong number")
//	} else {
//		l.List[i-1].Check = true
//		l.add()
//	}
//}
//
//func (l *Lists) Undone(value int) {
//	if len(l.List) < value {
//		fmt.Println("wrong number")
//	} else {
//		l.List[value-1].Check = false
//		l.add()
//	}
//}
//
//
//func (l *Lists) Clear(){
//	l.List = nil
//	fmt.Println(l.List, "Cleaned List")
//	 l.add()
//}

//func (l *Lists) DeleteTask(value string) {
//	for _ , v := range l.List {
//		if v == value {
//			l.List = append(l.List, l.List[i])
//			l.List = append(l.List[:key], l.List[key+1:]...)
//			//fmt.Println(l.List,"saved List")
//		}
//	}
//	 l.add()
//}





func main() {
	cmd.Execute()

}



