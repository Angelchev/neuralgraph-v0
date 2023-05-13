package core

type Node struct {
	ID      string
	Input   chan string
	Output  chan string
	Process func(string) string
}
