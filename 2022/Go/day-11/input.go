package main

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int
	divisor   int
}

func Input() []Monkey {
	return []Monkey{
		{ // 0
			items: []int{98, 70, 75, 80, 84, 89, 55, 98},
			operation: func(old int) int {
				return old * 2
			},
			test: func(item int) int {
				if item%11 == 0 {
					return 1
				}
				return 4
			},
			divisor: 11,
		},
		{ // 1
			items: []int{59},
			operation: func(old int) int {
				return old * old
			},
			test: func(item int) int {
				if item%19 == 0 {
					return 7
				}
				return 3
			},
			divisor: 19,
		},
		{ // 2
			items: []int{77, 95, 54, 65, 89},
			operation: func(old int) int {
				return old + 6
			},
			test: func(item int) int {
				if item%7 == 0 {
					return 0
				}
				return 5
			},
			divisor: 7,
		},
		{ //3
			items: []int{71, 64, 75},
			operation: func(old int) int {
				return old + 2
			},
			test: func(item int) int {
				if item%17 == 0 {
					return 6
				}
				return 2
			},
			divisor: 17,
		},
		{ // 4
			items: []int{74, 55, 87, 98},
			operation: func(old int) int {
				return old * 11
			},
			test: func(item int) int {
				if item%3 == 0 {
					return 1
				}
				return 7
			},
			divisor: 3,
		},
		{ // 5
			items: []int{90, 98, 85, 52, 91, 60},
			operation: func(old int) int {
				return old + 7
			},
			test: func(item int) int {
				if item%5 == 0 {
					return 0
				}
				return 4
			},
			divisor: 5,
		},
		{ // 6
			items: []int{99, 51},
			operation: func(old int) int {
				return old + 1
			},
			test: func(item int) int {
				if item%13 == 0 {
					return 5
				}
				return 2
			},
			divisor: 13,
		},
		{ // 7
			items: []int{98, 94, 59, 76, 51, 65, 75},
			operation: func(old int) int {
				return old + 5
			},
			test: func(item int) int {
				if item%2 == 0 {
					return 3
				}
				return 6
			},
			divisor: 2,
		},
	}

}
func InputExample() []Monkey {
	return []Monkey{
		{ // 0
			items: []int{79, 98},
			operation: func(old int) int {
				return old * 19
			},
			test: func(item int) int {
				if item%23 == 0 {
					return 2
				}
				return 3
			},
			divisor: 23,
		},
		{ // 1
			items: []int{54, 65, 75, 74},
			operation: func(old int) int {
				return old + 6
			},
			test: func(item int) int {
				if item%19 == 0 {
					return 2
				}
				return 0
			},
			divisor: 19,
		},
		{ // 3
			items: []int{79, 60, 97},
			operation: func(old int) int {
				return old * old
			},
			test: func(item int) int {
				if item%13 == 0 {
					return 1
				}
				return 3
			},
			divisor: 13,
		},
		{ // 4
			items: []int{74},
			operation: func(old int) int {
				return old + 3
			},
			test: func(item int) int {
				if item%17 == 0 {
					return 0
				}
				return 1
			},
			divisor: 17,
		},
	}
}
