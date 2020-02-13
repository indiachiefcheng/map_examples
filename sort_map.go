package main

import (
	"fmt"
	"sort"
)

func main(){
	ages := make(map[string]int)
	ages["Tom"] = 32
	ages["Jerry"] = 28
	ages["Alice"] = 26

	var names []string
	for name := range ages{
		names = append(names,name)
	}
	sort.Strings(names)
	for _,name := range names{
		fmt.Printf("%s\t%d\n",name,ages[name])	
	}
}
