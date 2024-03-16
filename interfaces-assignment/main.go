package main

func main() {
	rc := triangle{
		height: 5.6,
		side:   6.5,
	}

	sq := square{
		sideLength: 6,
	}

	printArea(rc)
	printArea(sq)
}
