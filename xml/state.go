package xml

type parserState int

func (s parserState) Is(state parserState) bool {
	return s == state
}

const (
	parserStateTagOpenStarted parserState = iota + 1
	parserStateTagOpenNamingStarted
	parserStateTagOpenNamingEnded
	parserStateTagOpenEnded
	parserStateTagAttrNamingStarted
	parserStateTagAttrNamingEnded
	parserStateTagAttrValueStarted
	parserStateTagAttrValueEnded
	parserStateTagCloseStarted
	parserStateTagCloseNamingStarted
	parserStateTagCloseNamingEnded
	parserStateTagCloseEnded
)

type parserStateStack struct {
	stack []parserState
}

func (s *parserStateStack) Length() int {
	return len(s.stack)
}

func (s *parserStateStack) IsEmpty() bool {
	return s.Length() < 1
}

func (s *parserStateStack) Last() parserState {
	if s.IsEmpty() {
		return 0
	}

	return s.stack[s.Length() - 1]
}

func (s *parserStateStack) Push(state parserState) {
	s.stack = append(s.stack, state)
}

func (s *parserStateStack) Pop() parserState {
	if s.IsEmpty() {
		return 0
	}

	e := s.Last()
	s.stack = s.stack[:s.Length() - 1]

	return e
}
