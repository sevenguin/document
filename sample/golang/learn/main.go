package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
	math "sample.com/learn/calculation"
	"sample.com/learn/calculation/advanced"
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println("new learn.")
	fmt.Println(math.Add(2, 1))
	fmt.Println(math.Subtract(2, 1))
	fmt.Println(math.Mul(2, 1))
	fmt.Println(advanced.Square(2))
	fmt.Println(uuid)
}
