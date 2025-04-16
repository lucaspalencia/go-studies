package main

import "fmt"

// map maps keys to values

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Mogi das Cruzes"] = Vertex{
		40.68333,
		-74.39967,
	}

	fmt.Println(m["Mogi das Cruzes"])

	// map literals
	var mm = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(mm)

	// can omit type name, most used
	var mmm = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(mmm)

	//mutating maps
	mmmm := make(map[string]int)

	mmmm["Answer"] = 42
	fmt.Println("The value:", mmmm["Answer"])

	mmmm["Answer"] = 48
	fmt.Println("The value:", mmmm["Answer"])

	//delete element
	delete(mmmm, "Answer")
	fmt.Println("The value:", mmmm["Answer"])

	//check if kaey is present, if it's ok = true
	v, ok := mmmm["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
