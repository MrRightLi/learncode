package container

import "fmt"

func main() {
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "immoc",
		"quality": "not bad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	fmt.Println(m["course"])
	fmt.Println(m["cause"])
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)



}