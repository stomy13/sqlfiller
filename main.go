package main

import (
	"fmt"
	"log"
	"sqlfiller/pkg/sqlform"
)

func main() {
	sql, err := sqlform.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sql)
}
