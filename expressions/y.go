// Code generated by goyacc expressions.y. DO NOT EDIT.

//line expressions.y:2
package expressions

import __yyfmt__ "fmt"

//line expressions.y:2
import (
	"fmt"
	"github.com/Funnelish/liquid/values"
)

func init() {
	// This allows adding and removing references to fmt in the rules below,
	// without having to comment and un-comment the import statement above.
	_ = ""
}

//line expressions.y:15
type yySymType struct {
	yys           int
	name          string
	val           any
	f             func(Context) values.Value
	s             string
	ss            []string
	exprs         []Expression
	cycle         Cycle
	cyclefn       func(string) Cycle
	loop          Loop
	loopmods      loopModifiers
	filter_params []valueFn
}

const LITERAL = 57346
const IDENTIFIER = 57347
const KEYWORD = 57348
const PROPERTY = 57349
const ASSIGN = 57350
const CYCLE = 57351
const LOOP = 57352
const WHEN = 57353
const EQ = 57354
const NEQ = 57355
const GE = 57356
const LE = 57357
const IN = 57358
const AND = 57359
const OR = 57360
const CONTAINS = 57361
const DOTDOT = 57362

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LITERAL",
	"IDENTIFIER",
	"KEYWORD",
	"PROPERTY",
	"ASSIGN",
	"CYCLE",
	"LOOP",
	"WHEN",
	"EQ",
	"NEQ",
	"GE",
	"LE",
	"IN",
	"AND",
	"OR",
	"CONTAINS",
	"DOTDOT",
	"'.'",
	"'|'",
	"'<'",
	"'>'",
	"';'",
	"'='",
	"':'",
	"','",
	"'['",
	"']'",
	"'('",
	"')'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 117

var yyAct = [...]int8{
	9, 47, 42, 18, 8, 2, 76, 23, 14, 15,
	10, 11, 43, 34, 3, 4, 5, 6, 35, 25,
	38, 10, 11, 60, 41, 43, 46, 51, 52, 53,
	54, 55, 56, 57, 58, 25, 25, 12, 44, 39,
	25, 26, 69, 24, 61, 62, 65, 63, 12, 66,
	64, 68, 45, 21, 7, 16, 48, 26, 26, 19,
	70, 77, 26, 14, 15, 72, 73, 1, 75, 36,
	37, 71, 78, 79, 74, 20, 25, 80, 14, 15,
	81, 27, 28, 31, 32, 40, 13, 17, 33, 59,
	49, 50, 30, 29, 25, 22, 67, 0, 26, 27,
	28, 31, 32, 0, 0, 0, 33, 0, 0, 0,
	30, 29, 0, 0, 0, 0, 26,
}

var yyPact = [...]int16{
	6, -1000, 61, 50, 55, 48, 17, -1000, 21, 87,
	-1000, -1000, 17, -1000, 17, 17, -6, 14, -3, -1000,
	13, 36, 1, 28, 85, -1000, 17, 17, 17, 17,
	17, 17, 17, 17, 69, -9, -1000, -1000, 17, -1000,
	-1000, 55, -1000, 55, -1000, 17, -1000, -1000, 17, -1000,
	17, 12, 33, 33, 33, 33, 33, 33, 33, 17,
	-1000, 46, -16, -16, 21, 33, 28, -22, 33, -1000,
	29, -1000, -1000, -1000, 67, -1000, 17, -1000, -1000, 17,
	33, 33,
}

var yyPgo = [...]int8{
	0, 0, 54, 4, 5, 96, 95, 1, 87, 85,
	2, 75, 74, 3, 67,
}

var yyR1 = [...]int8{
	0, 14, 14, 14, 14, 14, 8, 9, 9, 10,
	10, 6, 7, 7, 13, 11, 12, 12, 12, 1,
	1, 1, 1, 1, 1, 3, 3, 3, 5, 5,
	2, 2, 2, 2, 2, 2, 2, 2, 4, 4,
	4,
}

var yyR2 = [...]int8{
	0, 2, 5, 3, 3, 3, 2, 3, 1, 0,
	3, 2, 0, 3, 1, 4, 0, 2, 3, 1,
	1, 2, 4, 5, 3, 1, 3, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 1, 3,
	3,
}

var yyChk = [...]int16{
	-1000, -14, -4, 8, 9, 10, 11, -2, -3, -1,
	4, 5, 31, 25, 17, 18, 5, -8, -13, 4,
	-11, 5, -6, -1, 22, 7, 29, 12, 13, 24,
	23, 14, 15, 19, -1, -4, -2, -2, 26, 25,
	-9, 27, -10, 28, 25, 16, 25, -7, 28, 5,
	6, -1, -1, -1, -1, -1, -1, -1, -1, 20,
	32, -4, -13, -13, -3, -1, -1, -5, -1, 30,
	-1, 25, -10, -10, -12, -7, 28, 32, 5, 6,
	-1, -1,
}

var yyDef = [...]int8{
	0, -2, 0, 0, 0, 0, 0, 38, 30, 25,
	19, 20, 0, 1, 0, 0, 0, 0, 9, 14,
	0, 0, 0, 12, 0, 21, 0, 0, 0, 0,
	0, 0, 0, 0, 25, 0, 39, 40, 0, 3,
	6, 0, 8, 0, 4, 0, 5, 11, 0, 26,
	0, 0, 31, 32, 33, 34, 35, 36, 37, 0,
	24, 0, 9, 9, 16, 25, 12, 27, 28, 22,
	0, 2, 7, 10, 15, 13, 0, 23, 17, 0,
	29, 18,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	31, 32, 3, 3, 28, 3, 21, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 27, 25,
	23, 26, 24, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 30, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 22,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:45
		{
			yylex.(*lexer).val = yyDollar[1].f
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
//line expressions.y:46
		{
			yylex.(*lexer).Assignment = Assignment{yyDollar[2].name, &expression{yyDollar[4].f}}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:49
		{
			yylex.(*lexer).Cycle = yyDollar[2].cycle
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:50
		{
			yylex.(*lexer).Loop = yyDollar[2].loop
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:51
		{
			yylex.(*lexer).When = When{yyDollar[2].exprs}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:54
		{
			yyVAL.cycle = yyDollar[2].cyclefn(yyDollar[1].s)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:57
		{
			h, t := yyDollar[2].s, yyDollar[3].ss
			yyVAL.cyclefn = func(g string) Cycle { return Cycle{g, append([]string{h}, t...)} }
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:61
		{
			vals := yyDollar[1].ss
			yyVAL.cyclefn = func(h string) Cycle { return Cycle{Values: append([]string{h}, vals...)} }
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions.y:68
		{
			yyVAL.ss = []string{}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:69
		{
			yyVAL.ss = append([]string{yyDollar[2].s}, yyDollar[3].ss...)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:72
		{
			yyVAL.exprs = append([]Expression{&expression{yyDollar[1].f}}, yyDollar[2].exprs...)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions.y:74
		{
			yyVAL.exprs = []Expression{}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:75
		{
			yyVAL.exprs = append([]Expression{&expression{yyDollar[2].f}}, yyDollar[3].exprs...)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:78
		{
			s, ok := yyDollar[1].val.(string)
			if !ok {
				panic(SyntaxError(fmt.Sprintf("expected a string for %q", yyDollar[1].val)))
			}
			yyVAL.s = s
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions.y:86
		{
			name, expr, mods := yyDollar[1].name, yyDollar[3].f, yyDollar[4].loopmods
			yyVAL.loop = Loop{name, &expression{expr}, mods}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions.y:92
		{
			yyVAL.loopmods = loopModifiers{}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:93
		{
			switch yyDollar[2].name {
			case "reversed":
				yyDollar[1].loopmods.Reversed = true
			default:
				panic(SyntaxError(fmt.Sprintf("undefined loop modifier %q", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:102
		{
			switch yyDollar[2].name {
			case "cols":
				yyDollar[1].loopmods.Cols = &expression{yyDollar[3].f}
			case "limit":
				yyDollar[1].loopmods.Limit = &expression{yyDollar[3].f}
			case "offset":
				yyDollar[1].loopmods.Offset = &expression{yyDollar[3].f}
			default:
				panic(SyntaxError(fmt.Sprintf("undefined loop modifier %q", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:118
		{
			val := yyDollar[1].val
			yyVAL.f = func(Context) values.Value { return values.ValueOf(val) }
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:119
		{
			name := yyDollar[1].name
			yyVAL.f = func(ctx Context) values.Value { return values.ValueOf(ctx.Get(name)) }
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:120
		{
			yyVAL.f = makeObjectPropertyExpr(yyDollar[1].f, yyDollar[2].name)
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions.y:121
		{
			yyVAL.f = makeIndexExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
//line expressions.y:122
		{
			yyVAL.f = makeRangeExpr(yyDollar[2].f, yyDollar[4].f)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:123
		{
			yyVAL.f = yyDollar[2].f
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:128
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, nil)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions.y:129
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, yyDollar[4].filter_params)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:133
		{
			yyVAL.filter_params = []valueFn{yyDollar[1].f}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:135
		{
			yyVAL.filter_params = append(yyDollar[1].filter_params, yyDollar[3].f)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:139
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(a.Equal(b))
			}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:146
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(!a.Equal(b))
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:153
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(b.Less(a))
			}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:160
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(a.Less(b))
			}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:167
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(b.Less(a) || a.Equal(b))
			}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:174
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(a.Less(b) || a.Equal(b))
			}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:181
		{
			yyVAL.f = makeContainsExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:186
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				return values.ValueOf(fa(ctx).Test() && fb(ctx).Test())
			}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:192
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				return values.ValueOf(fa(ctx).Test() || fb(ctx).Test())
			}
		}
	}
	goto yystack /* stack new state and value */
}
