package lexer

import "lexer/token"

type Lexer struct {
	input        string // 输入的字符串
	position     int    // 所输入字符串中的当前位置（指向当前字符）
	readPosition int    // 所输入字符串中的当前读取位置（指向当前字符之后的一个字符，用来查看下一个字符）
	ch           byte   // 当前正在查看的字符
}

/**
 * @Description: 创建
 * @param input
 * @return *Lexer
 */
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/**
 * @Description: 读取下一个字符且调整position位置
 * @receiver l
 */
func (l *Lexer) readChar() {
	// 判断是否读到末尾 读到设置当前字符为0(NULL的ASCII编码)
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

/**
 * @Description:遍历源代码生成词法单元(获取下一个词法单元)
 * @receiver l
 * @return token.Token
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	//case byte(0):
	//	tok.Type = token.STRING
	//	tok.Literal = l.readString()
	case '=':
		// 查看下一个字符是不是'='
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch) // == 情况
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch) // !=情况
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// 如果是字母 读取剩余部分
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// 数字词法单元
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

/**
 * @Description: 构建词法单元
 * @param t
 * @param ch
 * @return token.Token
 */
func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal: string(ch)}
}

/**
 * @Description: 如果当前字符为字母就往后走  直到遇见一个非字母
 * @receiver l
 * @return string
 */
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/**
 * @Description: 判断是否为字母
 * @param ch
 * @return bool
 */
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/**
 * @Description: 跳过空白字符
 * @receiver l
 */
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/**
 * @Description: 返回下一个字符但不移动指针
 * @receiver l
 * @return byte
 */
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// Go中怎么表示byte("")
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == byte(0) || l.ch == 0 {
			break
		}
	}

	return l.input[position : position+1]
}
