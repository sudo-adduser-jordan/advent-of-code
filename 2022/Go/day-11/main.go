package main

import (
	"fmt"
	"sort"
)

func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {
	monkeys := Input()
	// monkeys := InputExample()
	inspected := make([]int, len(monkeys))

	for round := 1; round <= 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspectedItem := monkey.operation(item) / 3
				recievingMonkey := monkey.test(inspectedItem)
				monkeys[recievingMonkey].items =
					append(monkeys[recievingMonkey].items, inspectedItem)
			}
			inspected[i] = inspected[i] + len(monkey.items)
			monkeys[i].items = []int{}
		}

	}

	sort.Ints(inspected)
	level := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	fmt.Printf("The level of monkey business after 20 rounds of stuff-slinging simian shenanigans is: %v", level)
}

func PartTwo() {
	monkeys := Input()
	// monkeys := InputExample()
	inspected := make([]int, len(monkeys))

	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= m.divisor
	}

	for round := 1; round <= 10000; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspectedItem := monkey.operation(item)
				inspectedItem %= commonDivisor

				recievingMonkey := monkey.test(inspectedItem)
				monkeys[recievingMonkey].items =
					append(monkeys[recievingMonkey].items, inspectedItem)
			}
			inspected[i] = inspected[i] + len(monkey.items)
			monkeys[i].items = []int{}
		}
	}

	fmt.Println(inspected)
	sort.Ints(inspected)
	level := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	fmt.Printf("The level of monkey business after 20 rounds of stuff-slinging simian shenanigans is: %v", level)
}
