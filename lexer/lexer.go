package lexer

type Lexer struct {
	input 			string
	position		int		// current position in input (points to current character)
	readPosition	int		// current reading position in input (after current character)
	ch 				byte 	// current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	// File end condition
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Null character in the ASCII code table
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
