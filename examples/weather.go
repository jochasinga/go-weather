package main

import (
	"fmt"
	temp "github.com/jochasinga/gotemp"
)

func main() {
	mytemp := temp.Now("New York", "US")
	fmt.Println(mytemp)
}

