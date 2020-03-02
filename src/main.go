package main

import (
	"configparser/src/config"
	"fmt"
)

func main() {
	d := config.Dictionary{}
	d.Set("name", "barathwaj")
	lt := config.List{}
	lt.Append(123)
	ls := config.List{}
	ls.Append(1)
	ls.Append("barathwaj")
	ls.Append(d)
	ls.Append(lt)
	fmt.Print(ls)

}
