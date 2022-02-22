package main

import "fmt"

type Coordinates struct {
	lat, long float32
}

func main() {
	my_home := Coordinates{72.67, -89.54}
	my_office := Coordinates{69.34, -93.12}

	m := make(map[string]Coordinates)
	m["Home"] = my_home
	m["Office"] = my_office
	fmt.Println(m)

	// map literal approach to initialization
	var m_literal = map[string]Coordinates{
		"Google office": Coordinates{
			19.76, -23.4,
		},
		"Bell labs": Coordinates{
			37.42, -122.4,
		},
	}

	coord, ok := m_literal["Google office"]
	if ok {
		fmt.Println("Google office coordinates: ", coord)
	}

	// Now delete the key
	delete(m_literal, "Google office")
	coord, ok = m_literal["Google office"]
	if ok {
		fmt.Println("Google office coordinates: ", coord)
	} else {
		fmt.Println("Google office coordinates have now been deleted from map!")
	}
}
