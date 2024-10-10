package problem6

import (
	"fmt"
)

type Animal struct {
	name    string
	legsNum int
}

type Insect struct {
	name    string
	legsNum int
}

func sumOfAllLegsNum(animalsAndInsects ...interface{}) int {
	totalLegs := 0
	for _, item := range animalsAndInsects {
		switch v := item.(type) {
		case *Animal:
			totalLegs += v.legsNum
		case *Insect:
			totalLegs += v.legsNum
		}
	}
	return totalLegs
}

func main() {
	horse := &Animal{
		name:    "horse",
		legsNum: 4,
	}
	kangaroo := &Animal{
		name:    "kangaroo",
		legsNum: 2,
	}
	ant := &Insect{
		name:    "ant",
		legsNum: 6,
	}
	spider := &Insect{
		name:    "spider",
		legsNum: 8,
	}

	totalLegs := sumOfAllLegsNum(horse, kangaroo, ant, spider)
	fmt.Println(totalLegs) // Output: 20
}
