// Code generated by goyacc -o filter_parser.go filter_parser.y. DO NOT EDIT.

//line filter_parser.y:2
package main

import __yyfmt__ "fmt"

//line filter_parser.y:2

import (
	"strings"
	"time"
)

//line filter_parser.y:9
type yySymType struct {
	yys   int
	token Token
	expr  Expression
}

const BY = 57346
const TO = 57347
const ADDED = 57348
const ASSIGNED = 57349
const SUBTASK = 57350
const SHARED = 57351
const STRING = 57352
const NUMBER = 57353
const MONTH_IDENT = 57354
const TWELVE_CLOCK_IDENT = 57355
const HOURS = 57356
const PRIORITY = 57357
const RECURRING = 57358
const TODAY_IDENT = 57359
const TOMORROW_IDENT = 57360
const YESTERDAY_IDENT = 57361
const DAYS = 57362
const VIEW = 57363
const ALL = 57364
const DUE = 57365
const CREATED = 57366
const BEFORE = 57367
const AFTER = 57368
const OVER = 57369
const OVERDUE = 57370
const NO = 57371
const DATE = 57372
const TIME = 57373
const LABELS = 57374

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BY",
	"TO",
	"ADDED",
	"ASSIGNED",
	"SUBTASK",
	"SHARED",
	"STRING",
	"NUMBER",
	"MONTH_IDENT",
	"TWELVE_CLOCK_IDENT",
	"HOURS",
	"PRIORITY",
	"RECURRING",
	"TODAY_IDENT",
	"TOMORROW_IDENT",
	"YESTERDAY_IDENT",
	"DAYS",
	"VIEW",
	"ALL",
	"DUE",
	"CREATED",
	"BEFORE",
	"AFTER",
	"OVER",
	"OVERDUE",
	"NO",
	"DATE",
	"TIME",
	"LABELS",
	"'#'",
	"'@'",
	"'&'",
	"'|'",
	"'!'",
	"'('",
	"')'",
	"':'",
	"'-'",
	"'/'",
	"'+'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line filter_parser.y:294

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 105

var yyAct = [...]int8{
	21, 26, 19, 18, 17, 5, 30, 31, 90, 63,
	65, 14, 33, 34, 35, 59, 13, 61, 16, 15,
	91, 2, 24, 25, 11, 39, 40, 65, 22, 23,
	77, 76, 4, 3, 30, 31, 29, 64, 36, 62,
	33, 34, 35, 75, 38, 37, 74, 36, 70, 55,
	73, 72, 38, 37, 64, 45, 71, 28, 50, 68,
	69, 51, 52, 48, 29, 56, 36, 49, 78, 83,
	47, 44, 46, 84, 85, 86, 93, 92, 82, 81,
	80, 79, 67, 66, 60, 58, 89, 88, 87, 43,
	42, 41, 54, 53, 57, 9, 8, 20, 7, 6,
	10, 12, 27, 32, 1,
}

var yyPact = [...]int16{
	-5, -1000, 17, -5, -5, -1000, 81, 80, 79, -1000,
	-1000, 40, -1000, 45, -1000, 33, 36, -1000, -1000, 88,
	-1000, -1000, 16, -1000, 42, -1000, 90, 4, -1000, 73,
	-3, 72, -1000, -1000, -1000, -1000, 71, -5, -5, 9,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 26, -1000,
	11, 10, 6, 3, -9, -1000, -1000, -10, -1000, 14,
	48, -1000, 70, 69, 68, -1000, 67, 55, -1000, -1000,
	-1000, -1000, 23, 23, 23, 78, 77, 76, -1000, -34,
	-1000, -20, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	66, 65, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 104, 21, 0, 103, 102, 101, 100, 99, 98,
	57, 97, 96, 95,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 11, 11, 11, 9,
	8, 12, 13, 7, 7, 6, 6, 3, 3, 3,
	3, 3, 5, 5, 5, 5, 5, 5, 5, 4,
	4, 4, 10, 10, 10, 10,
}

var yyR2 = [...]int8{
	0, 0, 1, 3, 3, 3, 2, 1, 2, 2,
	2, 1, 1, 2, 2, 1, 2, 1, 4, 4,
	4, 1, 1, 1, 1, 1, 4, 4, 4, 2,
	1, 1, 2, 2, 3, 2, 1, 2, 1, 1,
	3, 2, 5, 3, 3, 1, 1, 1, 1, 2,
	2, 3, 3, 5, 3, 2,
}

var yyChk = [...]int16{
	-1000, -1, -2, 38, 37, 10, -8, -9, -12, -13,
	-7, 29, -6, 21, 16, 24, 23, 9, 8, 7,
	-11, -3, 33, 34, 27, 28, 6, -5, -10, 41,
	11, 12, -4, 17, 18, 19, 43, 36, 35, -2,
	-2, 10, 10, 10, 31, 15, 32, 30, 23, 22,
	25, 25, 26, 5, 4, 33, 23, 4, -10, 11,
	11, 20, 42, 12, 40, 13, 11, 11, -2, -2,
	39, 30, 40, 40, 40, 40, 40, 40, 20, 11,
	11, 11, 11, 14, -3, -3, -3, 10, 10, 10,
	42, 40, 11, 11,
}

var yyDef = [...]int8{
	1, -2, 2, 0, 0, 7, 0, 0, 0, 11,
	12, 0, 15, 0, 17, 0, 0, 21, 22, 23,
	24, 25, 30, 31, 0, 36, 0, 38, 39, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 0, 0,
	6, 8, 9, 10, 13, 14, 32, 33, 0, 16,
	0, 0, 0, 0, 0, 29, 35, 0, 37, 0,
	0, 41, 0, 50, 0, 55, 49, 0, 3, 4,
	5, 34, 0, 0, 0, 0, 0, 0, 40, 51,
	44, 52, 43, 54, 18, 19, 20, 26, 27, 28,
	0, 0, 42, 53,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 3, 33, 3, 3, 35, 3,
	38, 39, 3, 43, 3, 41, 3, 42, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 40, 3,
	3, 3, 3, 3, 34, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 36,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32,
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
		yyDollar = yyS[yypt-0 : yypt+1]
//line filter_parser.y:30
		{
			yyVAL.expr = VoidExpr{}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:34
		{
			yyVAL.expr = yyDollar[1].expr
			yylex.(*Lexer).result = yyVAL.expr
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:41
		{
			yyVAL.expr = BoolInfixOpExpr{left: yyDollar[1].expr, operator: '|', right: yyDollar[3].expr}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:45
		{
			yyVAL.expr = BoolInfixOpExpr{left: yyDollar[1].expr, operator: '&', right: yyDollar[3].expr}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:49
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:53
		{
			yyVAL.expr = NotOpExpr{expr: yyDollar[2].expr}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:57
		{
			yyVAL.expr = StringExpr{literal: yyDollar[1].token.literal}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:61
		{
			yyVAL.expr = ProjectExpr{isAll: false, name: yyDollar[2].token.literal}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:65
		{
			yyVAL.expr = ProjectExpr{isAll: true, name: yyDollar[2].token.literal}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:69
		{
			yyVAL.expr = LabelExpr{name: yyDollar[2].token.literal}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:73
		{
			yyVAL.expr = LabelExpr{name: ""}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:77
		{
			yyVAL.expr = DateExpr{operation: NO_DUE_DATE}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:81
		{
			yyVAL.expr = DateExpr{operation: NO_TIME}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:85
		{
			yyVAL.expr = NoPriorityExpr{}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:89
		{
			yyVAL.expr = DateExpr{allDay: false, datetime: now(), operation: DUE_BEFORE}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:93
		{
			yyVAL.expr = ViewAllExpr{}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:97
		{
			yyVAL.expr = RecurringExpr{}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:101
		{
			e := yyDollar[4].expr.(DateExpr)
			e.operation = CREATED_BEFORE
			yyVAL.expr = e
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:107
		{
			e := yyDollar[4].expr.(DateExpr)
			e.operation = DUE_BEFORE
			yyVAL.expr = e
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:113
		{
			e := yyDollar[4].expr.(DateExpr)
			e.operation = DUE_AFTER
			yyVAL.expr = e
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:119
		{
			yyVAL.expr = SharedExpr{}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:123
		{
			yyVAL.expr = SubtaskExpr{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:127
		{
			yyVAL.expr = AssignedExpr{}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:135
		{
			yyVAL.expr = PersonExpr{operation: ASSIGNED_TO, person: yyDollar[4].token.literal}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:139
		{
			yyVAL.expr = PersonExpr{operation: ASSIGNED_BY, person: yyDollar[4].token.literal}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:143
		{
			yyVAL.expr = PersonExpr{operation: ADDED_BY, person: yyDollar[4].token.literal}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:149
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:155
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:161
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:167
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:173
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:177
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:183
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:187
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:194
		{
			date := yyDollar[1].expr.(time.Time)
			time := yyDollar[2].expr.(time.Duration)
			yyVAL.expr = DateExpr{allDay: false, datetime: date.Add(time)}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:200
		{
			yyVAL.expr = DateExpr{allDay: true, datetime: yyDollar[1].expr.(time.Time)}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:204
		{
			nd := now().Sub(today())
			d := yyDollar[1].expr.(time.Duration)
			if d <= nd {
				d = d + time.Duration(int64(time.Hour)*24)
			}
			yyVAL.expr = DateExpr{allDay: false, datetime: today().Add(d)}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:213
		{
			date := today().AddDate(0, 0, -atoi(yyDollar[2].token.literal))
			yyVAL.expr = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:218
		{
			date := today().AddDate(0, 0, atoi(yyDollar[1].token.literal))
			yyVAL.expr = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
//line filter_parser.y:225
		{
			yyVAL.expr = time.Date(atoi(yyDollar[5].token.literal), time.Month(atoi(yyDollar[1].token.literal)), atoi(yyDollar[3].token.literal), 0, 0, 0, 0, timezone())
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:229
		{
			yyVAL.expr = time.Date(atoi(yyDollar[3].token.literal), MonthIdentHash[strings.ToLower(yyDollar[1].token.literal)], atoi(yyDollar[2].token.literal), 0, 0, 0, 0, timezone())
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:233
		{
			yyVAL.expr = time.Date(atoi(yyDollar[3].token.literal), MonthIdentHash[strings.ToLower(yyDollar[2].token.literal)], atoi(yyDollar[1].token.literal), 0, 0, 0, 0, timezone())
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:237
		{
			tod := today()
			date := yyDollar[1].expr.(time.Time)
			if date.Before(tod) {
				date = date.AddDate(1, 0, 0)
			}
			yyVAL.expr = date
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:246
		{
			yyVAL.expr = today()
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:250
		{
			yyVAL.expr = today().AddDate(0, 0, 1)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:254
		{
			yyVAL.expr = today().AddDate(0, 0, -1)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:260
		{
			yyVAL.expr = time.Date(today().Year(), MonthIdentHash[strings.ToLower(yyDollar[1].token.literal)], atoi(yyDollar[2].token.literal), 0, 0, 0, 0, timezone())
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:264
		{
			yyVAL.expr = time.Date(today().Year(), MonthIdentHash[strings.ToLower(yyDollar[2].token.literal)], atoi(yyDollar[1].token.literal), 0, 0, 0, 0, timezone())
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:268
		{
			yyVAL.expr = time.Date(now().Year(), time.Month(atoi(yyDollar[3].token.literal)), atoi(yyDollar[1].token.literal), 0, 0, 0, 0, timezone())
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:274
		{
			yyVAL.expr = time.Duration(int64(time.Hour)*int64(atoi(yyDollar[1].token.literal)) + int64(time.Minute)*int64(atoi(yyDollar[3].token.literal)))
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
//line filter_parser.y:278
		{
			yyVAL.expr = time.Duration(int64(time.Hour)*int64(atoi(yyDollar[1].token.literal)) + int64(time.Minute)*int64(atoi(yyDollar[3].token.literal)) + int64(time.Second)*int64(atoi(yyDollar[5].token.literal)))
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:282
		{
			yyVAL.expr = time.Duration(int64(time.Hour) * int64(atoi(yyDollar[2].token.literal)))
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:286
		{
			hour := atoi(yyDollar[1].token.literal)
			if TwelveClockIdentHash[yyDollar[2].token.literal] {
				hour = hour + 12
			}
			yyVAL.expr = time.Duration(int64(time.Hour) * int64(hour))
		}
	}
	goto yystack /* stack new state and value */
}
