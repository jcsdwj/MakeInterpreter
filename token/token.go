package token

type TokenType string

// 词法单元
type Token struct {
	Type    TokenType // 词法单元类型（标识是哪种类型）
	Literal string    // 保存字面量 比如是5还是10
}

const (
	ILLEGAL = "ILLEGAL" // 表示遇到未知的词法单元
	EOF     = "EOF"     // 表示文件结尾

	// 标识量+字面量
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	// 运算符
	PLUS     = "+"
	MINUS    = "-"
	ASSIGN   = "="
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	EQ     = "=="
	NOT_EQ = "!="

	LBRACKET = "["
	RBRACKET = "]"

	COLON = ":"
	MACRO = "MACRO"
)

// 关键字
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"marco":  MACRO,
}

/**
 * @Description: 区分语言关键字和用户定义标识符  判断给定的单词是否为关键字
 * @param ident
 * @return TokenType
 */
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		// 返回关键字
		return tok
	}
	return IDENT
}
