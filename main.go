package main

import (
	"fmt"
	"github.com/h-u-m-a-n/l4hw/tetris"
	"log"
)

type name struct {
	a int8
	c int8
	d int8
	e int8
	ee2 int8
	eee int8
	ee int8
	// 8
	f int32
	// 8
	b int64
	// 8
	f2 int32
	ff int32
	// 8
}


func main() {
	// Я написал 2 функции
	// 1) SortStruct находит лучший порядок параметров в структуре и возвращает его, а также переписывает файл со структурой
	// Она закомментирована и находится внизу
	// 2) Get3Top находит 3 самых лучших результата и просто возвращает их

	if results, err := tetris.Get3Top("tetris/teststruct.go"); err != nil {
		log.Printf("some error occured %v", err)
 	} else {
		for _, result := range results {
			for _, item := range result {
				fmt.Print(item[1:], " | ")
			}
			fmt.Println()
		}
	}

	//if ans, err := tetris.SortStruct("tetris/teststruct.go"); err != nil {
	//	log.Printf("some error occurred %v", err)
	//} else {
	//	fmt.Println(ans)
	//}
}
