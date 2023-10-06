package handler

type gameProcessor struct {
	Move        func(input [][]int)
	Merge       func(input [][]int) int
	AddRandCell func(input [][]int)
}
