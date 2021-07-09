package brackets

type stack struct {
	items   []rune
	current int
}

func (s *stack) empty() bool {
	return s.current == -1
}

func (s *stack) push(item rune) {
	s.current += 1
	s.items[s.current] = item
}

func (s *stack) pop() {
	s.current -= 1
}

func (s *stack) top() rune {
	return s.items[s.current]
}

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

// Bracket returns if a string parenthesis are balanced
func Bracket(input string) bool {
	s := stack{items: make([]rune, len(input)), current: -1}

	for _, c := range []rune(input) {
		switch c {
		case '(', '[', '{':
			s.push(c)
		case ')', ']', '}':
			if s.empty() {
				return false
			}

			if t := s.top(); t != pairs[c] {
				return false
			}
			s.pop()
		default:
			continue
		}
	}
	return s.empty()
}
