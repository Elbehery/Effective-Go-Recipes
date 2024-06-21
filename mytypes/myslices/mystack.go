package main

import "fmt"

type Token struct {
	Loc  int
	Char rune
}

type Stack []Token

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(tok Token) {
	*s = append(*s, tok)
}

func (s *Stack) Pop() (Token, error) {
	size := s.Len()
	if size == 0 {
		return Token{}, fmt.Errorf("empty stack")
	}

	sl := *s
	val := sl[size-1]
	sl = sl[:size-1]

	if len(sl) > 1024 && 2*len(sl) < cap(sl) {
		sl2 := make([]Token, len(sl))
		copy(sl2, sl)
		sl = sl2
	}

	*s = sl
	return val, nil
}

func main() {
	var s Stack
	fmt.Println(s)
	s.Push(Token{19, '('})
	s.Push(Token{49, '['})
	fmt.Println(s)
	if v, err := s.Pop(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("pop", v)
	}
}
