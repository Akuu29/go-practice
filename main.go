package main

import (
	"fmt"

	"github.com/Akuu29/go-practice/app/controllers"
	"github.com/Akuu29/go-practice/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
