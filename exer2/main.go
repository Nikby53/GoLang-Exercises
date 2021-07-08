package main

import "fmt"

func search(arg []string) []string {
	key := make(map[string]struct{})
	arr := []string{}
	for i, v := range arg {
		if _, ok := key[v]; !ok {
			key[arg[i]] = struct{}{}
			arr = append(arr, v)
		}

	}

	return arr
}

func main() {
	arg := []string{"blue", "blue", "blue", "red", "red", "yellow"}
	fmt.Println(arg)
	newArg := search(arg)
	fmt.Println(newArg)
}
