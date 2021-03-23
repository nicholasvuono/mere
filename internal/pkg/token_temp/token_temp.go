package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	NUM   = "NUM"
	STR   = "STR"
	BOOL  = "BOOL"

	LBRCK = "LBRCK"
	RBRCK = "RBRCK"

	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	FUNC   = "FUNC"
	END    = "END"
	RETURN = "RETURN"
)

type Token struct {
	Type    string
	Literal string
}

var keywords = map[string]string{
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"func":   FUNC,
	"end":    END,
	"return": RETURN,
}

func lookup(ident string) string {
	token, bool := keywords[ident]
	if bool {
		return token
	}
	return IDENT
}
