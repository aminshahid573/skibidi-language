package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Token types
type TokenType int

const (
	SKIBIDI TokenType = iota // variable declaration
	RIZZ                     // assignment
	CAP                      // if statement
	NOCAP                    // else
	BUSSIN                   // while loop
	GYATT                    // print
	OHIO                     // end statement
	SIGMA                    // function declaration
	ALPHA                    // return
	BETA                     // function call
	BRUH                     // comment
	NUMBER
	STRING
	IDENTIFIER
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	EQUALS
	LESS_THAN
	GREATER_THAN
	LESS_EQUAL
	GREATER_EQUAL
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	SEMICOLON
	EOF
	MODULO
	AND
	OR
	FOR
	INPUT
	TRUE
	FALSE
	COMMA
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

type Lexer struct {
	input    string
	position int
	line     int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
		line:     1,
	}
}

func (l *Lexer) peek() byte {
	if l.position >= len(l.input) {
		return 0
	}
	return l.input[l.position]
}

func (l *Lexer) advance() byte {
	if l.position >= len(l.input) {
		return 0
	}
	ch := l.input[l.position]
	l.position++
	if ch == '\n' {
		l.line++
	}
	return ch
}

func (l *Lexer) skipWhitespace() {
	for l.peek() == ' ' || l.peek() == '\t' || l.peek() == '\r' {
		l.advance()
	}
}

func (l *Lexer) readString() string {
	var result strings.Builder
	l.advance() // skip opening quote

	for l.peek() != '"' && l.peek() != 0 {
		if l.peek() == '\\' {
			l.advance()
			switch l.peek() {
			case 'n':
				result.WriteByte('\n')
			case 't':
				result.WriteByte('\t')
			case 'r':
				result.WriteByte('\r')
			case '\\':
				result.WriteByte('\\')
			case '"':
				result.WriteByte('"')
			default:
				result.WriteByte(l.peek())
			}
		} else {
			result.WriteByte(l.peek())
		}
		l.advance()
	}

	if l.peek() == '"' {
		l.advance() // skip closing quote
	}

	return result.String()
}

func (l *Lexer) readNumber() string {
	var result strings.Builder

	for (l.peek() >= '0' && l.peek() <= '9') || l.peek() == '.' {
		result.WriteByte(l.peek())
		l.advance()
	}

	return result.String()
}

func (l *Lexer) readIdentifier() string {
	var result strings.Builder

	for (l.peek() >= 'a' && l.peek() <= 'z') ||
		(l.peek() >= 'A' && l.peek() <= 'Z') ||
		(l.peek() >= '0' && l.peek() <= '9') ||
		l.peek() == '_' {
		result.WriteByte(l.peek())
		l.advance()
	}

	return result.String()
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.peek() == '\n' {
		l.advance()
		return l.NextToken()
	}

	if l.peek() == 0 {
		return Token{EOF, "", l.line}
	}

	// Single-line comments starting with "bruh"
	if l.position+4 <= len(l.input) && l.input[l.position:l.position+4] == "bruh" {
		for l.peek() != '\n' && l.peek() != 0 {
			l.advance()
		}
		return l.NextToken()
	}

	// Keywords
	keywords := map[string]TokenType{
		"skibidi": SKIBIDI,
		"rizz":    RIZZ,
		"cap":     CAP,
		"nocap":   NOCAP,
		"bussin":  BUSSIN,
		"gyatt":   GYATT,
		"ohio":    OHIO,
		"sigma":   SIGMA,
		"alpha":   ALPHA,
		"beta":    BETA,
		"gyatfor": FOR,
		"input":   INPUT,
		"true":    TRUE,
		"false":   FALSE,
	}

	if (l.peek() >= 'a' && l.peek() <= 'z') || (l.peek() >= 'A' && l.peek() <= 'Z') {
		identifier := l.readIdentifier()
		if tokenType, exists := keywords[identifier]; exists {
			return Token{tokenType, identifier, l.line}
		}
		// Built-in functions: len, abs, str
		if identifier == "len" || identifier == "abs" || identifier == "str" {
			return Token{IDENTIFIER, identifier, l.line}
		}
		return Token{IDENTIFIER, identifier, l.line}
	}

	if l.peek() >= '0' && l.peek() <= '9' {
		number := l.readNumber()
		return Token{NUMBER, number, l.line}
	}

	if l.peek() == '"' {
		str := l.readString()
		return Token{STRING, str, l.line}
	}

	switch l.peek() {
	case '+':
		l.advance()
		return Token{PLUS, "+", l.line}
	case '-':
		l.advance()
		return Token{MINUS, "-", l.line}
	case '*':
		l.advance()
		return Token{MULTIPLY, "*", l.line}
	case '/':
		l.advance()
		return Token{DIVIDE, "/", l.line}
	case '=':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{EQUALS, "==", l.line}
		}
		return Token{RIZZ, "=", l.line}
	case '<':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{LESS_EQUAL, "<=", l.line}
		}
		return Token{LESS_THAN, "<", l.line}
	case '>':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{GREATER_EQUAL, ">=", l.line}
		}
		return Token{GREATER_THAN, ">", l.line}
	case '(':
		l.advance()
		return Token{LPAREN, "(", l.line}
	case ')':
		l.advance()
		return Token{RPAREN, ")", l.line}
	case '{':
		l.advance()
		return Token{LBRACE, "{", l.line}
	case '}':
		l.advance()
		return Token{RBRACE, "}", l.line}
	case ';':
		l.advance()
		return Token{SEMICOLON, ";", l.line}
	case '%':
		l.advance()
		return Token{MODULO, "%", l.line}
	case '&':
		l.advance()
		if l.peek() == '&' {
			l.advance()
			return Token{AND, "&&", l.line}
		}
		return Token{EOF, "&", l.line}
	case '|':
		l.advance()
		if l.peek() == '|' {
			l.advance()
			return Token{OR, "||", l.line}
		}
		return Token{EOF, "|", l.line}
	case ',':
		l.advance()
		return Token{COMMA, ",", l.line}
	default:
		ch := l.advance()
		return Token{EOF, string(ch), l.line}
	}
}

// AST Nodes
type ASTNode interface {
	String() string
}

type Program struct {
	Statements []ASTNode
}

func (p *Program) String() string {
	return "Program"
}

type VarDecl struct {
	Name  string
	Value ASTNode
}

func (v *VarDecl) String() string {
	return fmt.Sprintf("VarDecl(%s)", v.Name)
}

type Assignment struct {
	Name  string
	Value ASTNode
}

func (a *Assignment) String() string {
	return fmt.Sprintf("Assignment(%s)", a.Name)
}

type PrintStmt struct {
	Value ASTNode
}

func (p *PrintStmt) String() string {
	return "PrintStmt"
}

type IfStmt struct {
	Condition ASTNode
	ThenBlock []ASTNode
	ElseBlock []ASTNode
}

func (i *IfStmt) String() string {
	return "IfStmt"
}

type WhileStmt struct {
	Condition ASTNode
	Body      []ASTNode
}

func (w *WhileStmt) String() string {
	return "WhileStmt"
}

type BinaryOp struct {
	Left     ASTNode
	Operator string
	Right    ASTNode
}

func (b *BinaryOp) String() string {
	return fmt.Sprintf("BinaryOp(%s)", b.Operator)
}

type NumberLiteral struct {
	Value float64
}

func (n *NumberLiteral) String() string {
	return fmt.Sprintf("Number(%f)", n.Value)
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf("String(%s)", s.Value)
}

type Identifier struct {
	Name string
}

func (i *Identifier) String() string {
	return fmt.Sprintf("Identifier(%s)", i.Name)
}

type SigmaFunc struct {
	Name   string
	Params []string
	Body   []ASTNode
}

func (f *SigmaFunc) String() string {
	return fmt.Sprintf("SigmaFunc(%s)", f.Name)
}

type BetaCall struct {
	Name string
	Args []ASTNode
}

func (b *BetaCall) String() string {
	return fmt.Sprintf("BetaCall(%s)", b.Name)
}

type AlphaReturn struct {
	Value ASTNode
}

func (a *AlphaReturn) String() string {
	return "AlphaReturn"
}

type ForStmt struct {
	Init      ASTNode
	Condition ASTNode
	Post      ASTNode
	Body      []ASTNode
}

func (f *ForStmt) String() string {
	return "ForStmt"
}

type InputExpr struct{}

func (i *InputExpr) String() string {
	return "InputExpr"
}

type BoolLiteral struct {
	Value bool
}

func (b *BoolLiteral) String() string {
	return fmt.Sprintf("Bool(%v)", b.Value)
}

// Parser
type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.currentToken = parser.lexer.NextToken()
	return parser
}

func (p *Parser) eat(expectedType TokenType) {
	if p.currentToken.Type == expectedType {
		p.currentToken = p.lexer.NextToken()
	} else {
		panic(fmt.Sprintf("Expected token %d, got %d at line %d", expectedType, p.currentToken.Type, p.currentToken.Line))
	}
}

func (p *Parser) parseExpression() ASTNode {
	return p.parseLogicalOr()
}

func (p *Parser) parseLogicalOr() ASTNode {
	node := p.parseLogicalAnd()

	for p.currentToken.Type == OR {
		op := p.currentToken.Value
		p.eat(OR)
		right := p.parseLogicalAnd()
		node = &BinaryOp{Left: node, Operator: op, Right: right}
	}

	return node
}

func (p *Parser) parseLogicalAnd() ASTNode {
	node := p.parseComparison()

	for p.currentToken.Type == AND {
		op := p.currentToken.Value
		p.eat(AND)
		right := p.parseComparison()
		node = &BinaryOp{Left: node, Operator: op, Right: right}
	}

	return node
}

func (p *Parser) parseComparison() ASTNode {
	node := p.parseArithmetic()

	for p.currentToken.Type == EQUALS || p.currentToken.Type == LESS_THAN || p.currentToken.Type == GREATER_THAN || p.currentToken.Type == LESS_EQUAL || p.currentToken.Type == GREATER_EQUAL {
		op := p.currentToken.Value
		p.eat(p.currentToken.Type)
		right := p.parseArithmetic()
		node = &BinaryOp{Left: node, Operator: op, Right: right}
	}

	return node
}

func (p *Parser) parseArithmetic() ASTNode {
	node := p.parseTerm()

	for p.currentToken.Type == PLUS || p.currentToken.Type == MINUS {
		op := p.currentToken.Value
		p.eat(p.currentToken.Type)
		right := p.parseTerm()
		node = &BinaryOp{Left: node, Operator: op, Right: right}
	}

	return node
}

func (p *Parser) parseTerm() ASTNode {
	node := p.parseFactor()

	for p.currentToken.Type == MULTIPLY || p.currentToken.Type == DIVIDE || p.currentToken.Type == MODULO {
		op := p.currentToken.Value
		p.eat(p.currentToken.Type)
		right := p.parseFactor()
		node = &BinaryOp{Left: node, Operator: op, Right: right}
	}

	return node
}

func (p *Parser) parseFactor() ASTNode {
	token := p.currentToken

	if token.Type == NUMBER {
		p.eat(NUMBER)
		value, _ := strconv.ParseFloat(token.Value, 64)
		return &NumberLiteral{Value: value}
	} else if token.Type == STRING {
		p.eat(STRING)
		return &StringLiteral{Value: token.Value}
	} else if token.Type == TRUE {
		p.eat(TRUE)
		return &BoolLiteral{Value: true}
	} else if token.Type == FALSE {
		p.eat(FALSE)
		return &BoolLiteral{Value: false}
	} else if token.Type == IDENTIFIER {
		name := token.Value
		p.eat(IDENTIFIER)
		if p.currentToken.Type == LPAREN {
			// Built-in or user function call
			p.eat(LPAREN)
			args := []ASTNode{}
			if p.currentToken.Type != RPAREN {
				args = append(args, p.parseExpression())
				for p.currentToken.Type == COMMA {
					p.eat(COMMA)
					args = append(args, p.parseExpression())
				}
			}
			p.eat(RPAREN)
			return &BetaCall{Name: name, Args: args}
		}
		return &Identifier{Name: name}
	} else if token.Type == INPUT {
		p.eat(INPUT)
		return &InputExpr{}
	} else if token.Type == BETA {
		p.eat(BETA)
		name := p.currentToken.Value
		p.eat(IDENTIFIER)
		p.eat(LPAREN)
		args := []ASTNode{}
		if p.currentToken.Type != RPAREN {
			args = append(args, p.parseExpression())
			for p.currentToken.Type == COMMA {
				p.eat(COMMA)
				args = append(args, p.parseExpression())
			}
		}
		p.eat(RPAREN)
		return &BetaCall{Name: name, Args: args}
	} else if token.Type == MINUS {
		p.eat(MINUS)
		factor := p.parseFactor()
		return &BinaryOp{Left: &NumberLiteral{Value: 0}, Operator: "-", Right: factor}
	} else if token.Type == LPAREN {
		p.eat(LPAREN)
		node := p.parseExpression()
		p.eat(RPAREN)
		return node
	}

	panic(fmt.Sprintf("Unexpected token %s at line %d", token.Value, token.Line))
}

func (p *Parser) parseBlock() []ASTNode {
	statements := []ASTNode{}
	p.eat(LBRACE)

	for p.currentToken.Type != RBRACE && p.currentToken.Type != EOF {
		stmt := p.parseStatement()
		statements = append(statements, stmt)
	}

	p.eat(RBRACE)
	return statements
}

func (p *Parser) parseStatement() ASTNode {
	switch p.currentToken.Type {
	case SKIBIDI:
		return p.parseVarDecl()
	case IDENTIFIER:
		return p.parseAssignment()
	case GYATT:
		return p.parsePrintStmt()
	case CAP:
		return p.parseIfStmt()
	case BUSSIN:
		return p.parseWhileStmt()
	case FOR:
		return p.parseForStmt()
	case SIGMA:
		return p.parseSigmaFunc()
	case BETA:
		return p.parseBetaCallStmt()
	case ALPHA:
		return p.parseAlphaReturn()
	default:
		panic(fmt.Sprintf("Unexpected token %s at line %d", p.currentToken.Value, p.currentToken.Line))
	}
}

func (p *Parser) parseVarDecl() ASTNode {
	p.eat(SKIBIDI)
	name := p.currentToken.Value
	p.eat(IDENTIFIER)

	// Handle both 'rizz' keyword and '=' symbol for assignment
	if p.currentToken.Type == RIZZ {
		p.eat(RIZZ)
	} else {
		panic(fmt.Sprintf("Expected 'rizz' after variable name, got %s at line %d", p.currentToken.Value, p.currentToken.Line))
	}

	value := p.parseExpression()
	p.eat(OHIO)
	return &VarDecl{Name: name, Value: value}
}

func (p *Parser) parseAssignment() ASTNode {
	name := p.currentToken.Value
	p.eat(IDENTIFIER)

	// Handle both 'rizz' keyword and '=' symbol for assignment
	if p.currentToken.Type == RIZZ {
		p.eat(RIZZ)
	} else {
		panic(fmt.Sprintf("Expected 'rizz' after variable name, got %s at line %d", p.currentToken.Value, p.currentToken.Line))
	}

	value := p.parseExpression()
	p.eat(OHIO)
	return &Assignment{Name: name, Value: value}
}

func (p *Parser) parsePrintStmt() ASTNode {
	p.eat(GYATT)
	value := p.parseExpression()
	p.eat(OHIO)
	return &PrintStmt{Value: value}
}

func (p *Parser) parseIfStmt() ASTNode {
	p.eat(CAP)
	p.eat(LPAREN)
	condition := p.parseExpression()
	p.eat(RPAREN)
	thenBlock := p.parseBlock()

	var elseBlock []ASTNode
	if p.currentToken.Type == NOCAP {
		p.eat(NOCAP)
		elseBlock = p.parseBlock()
	}

	return &IfStmt{Condition: condition, ThenBlock: thenBlock, ElseBlock: elseBlock}
}

func (p *Parser) parseWhileStmt() ASTNode {
	p.eat(BUSSIN)
	p.eat(LPAREN)
	condition := p.parseExpression()
	p.eat(RPAREN)
	body := p.parseBlock()
	return &WhileStmt{Condition: condition, Body: body}
}

func (p *Parser) parseSigmaFunc() ASTNode {
	p.eat(SIGMA)
	name := p.currentToken.Value
	p.eat(IDENTIFIER)
	p.eat(LPAREN)
	params := []string{}
	if p.currentToken.Type == IDENTIFIER {
		params = append(params, p.currentToken.Value)
		p.eat(IDENTIFIER)
		for p.currentToken.Type == COMMA {
			p.eat(COMMA)
			params = append(params, p.currentToken.Value)
			p.eat(IDENTIFIER)
		}
	}
	p.eat(RPAREN)
	body := p.parseBlock()
	return &SigmaFunc{Name: name, Params: params, Body: body}
}

func (p *Parser) parseBetaCallStmt() ASTNode {
	p.eat(BETA)
	name := p.currentToken.Value
	p.eat(IDENTIFIER)
	p.eat(LPAREN)
	args := []ASTNode{}
	if p.currentToken.Type != RPAREN {
		args = append(args, p.parseExpression())
		for p.currentToken.Type == COMMA {
			p.eat(COMMA)
			args = append(args, p.parseExpression())
		}
	}
	p.eat(RPAREN)
	p.eat(OHIO)
	return &BetaCall{Name: name, Args: args}
}

func (p *Parser) parseAlphaReturn() ASTNode {
	p.eat(ALPHA)
	value := p.parseExpression()
	p.eat(OHIO)
	return &AlphaReturn{Value: value}
}

func (p *Parser) parseVarDeclNoOhio() ASTNode {
	p.eat(SKIBIDI)
	name := p.currentToken.Value
	p.eat(IDENTIFIER)
	if p.currentToken.Type == RIZZ {
		p.eat(RIZZ)
	} else {
		panic(fmt.Sprintf("Expected 'rizz' after variable name, got %s at line %d", p.currentToken.Value, p.currentToken.Line))
	}
	value := p.parseExpression()
	return &VarDecl{Name: name, Value: value}
}

func (p *Parser) parseAssignmentNoOhio() ASTNode {
	name := p.currentToken.Value
	p.eat(IDENTIFIER)
	if p.currentToken.Type == RIZZ {
		p.eat(RIZZ)
	} else {
		panic(fmt.Sprintf("Expected 'rizz' after variable name, got %s at line %d", p.currentToken.Value, p.currentToken.Line))
	}
	value := p.parseExpression()
	return &Assignment{Name: name, Value: value}
}

func (p *Parser) parseForStmt() ASTNode {
	p.eat(FOR)
	p.eat(LPAREN)
	var init ASTNode
	if p.currentToken.Type == SKIBIDI {
		init = p.parseVarDeclNoOhio()
	} else if p.currentToken.Type == IDENTIFIER {
		init = p.parseAssignmentNoOhio()
	} else {
		init = nil
	}
	p.eat(SEMICOLON)
	cond := p.parseExpression()
	p.eat(SEMICOLON)
	var post ASTNode
	if p.currentToken.Type == IDENTIFIER {
		post = p.parseAssignmentNoOhio()
	} else {
		post = nil
	}
	p.eat(RPAREN)
	body := p.parseBlock()
	return &ForStmt{Init: init, Condition: cond, Post: post, Body: body}
}

func (p *Parser) Parse() *Program {
	statements := []ASTNode{}

	for p.currentToken.Type != EOF {
		stmt := p.parseStatement()
		statements = append(statements, stmt)
	}

	return &Program{Statements: statements}
}

// Interpreter
type callFrame struct {
	variables   map[string]interface{}
	returnValue interface{}
	returned    bool
}

type Interpreter struct {
	globals      map[string]interface{}
	functions    map[string]*SigmaFunc
	callStack    []*callFrame
	inputScanner *bufio.Scanner
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		globals:      make(map[string]interface{}),
		functions:    make(map[string]*SigmaFunc),
		callStack:    []*callFrame{{variables: make(map[string]interface{})}},
		inputScanner: bufio.NewScanner(os.Stdin),
	}
}

func (i *Interpreter) currentFrame() *callFrame {
	return i.callStack[len(i.callStack)-1]
}

func (i *Interpreter) getVar(name string) (interface{}, bool) {
	for idx := len(i.callStack) - 1; idx >= 0; idx-- {
		if val, ok := i.callStack[idx].variables[name]; ok {
			return val, true
		}
	}
	if val, ok := i.globals[name]; ok {
		return val, true
	}
	return nil, false
}

func (i *Interpreter) setVar(name string, value interface{}) {
	for idx := len(i.callStack) - 1; idx >= 0; idx-- {
		if _, ok := i.callStack[idx].variables[name]; ok {
			i.callStack[idx].variables[name] = value
			return
		}
	}
	i.currentFrame().variables[name] = value
}

func (i *Interpreter) evaluateExpression(node ASTNode) interface{} {
	switch n := node.(type) {
	case *NumberLiteral:
		return n.Value
	case *StringLiteral:
		return n.Value
	case *BoolLiteral:
		return n.Value
	case *Identifier:
		if val, exists := i.getVar(n.Name); exists {
			return val
		}
		panic(fmt.Sprintf("Undefined variable: %s", n.Name))
	case *BinaryOp:
		left := i.evaluateExpression(n.Left)
		right := i.evaluateExpression(n.Right)

		switch n.Operator {
		case "+":
			return i.add(left, right)
		case "-":
			return i.toFloat(left) - i.toFloat(right)
		case "*":
			return i.toFloat(left) * i.toFloat(right)
		case "/":
			rightVal := i.toFloat(right)
			if rightVal == 0 {
				panic("Division by zero")
			}
			return i.toFloat(left) / rightVal
		case "==":
			return i.equals(left, right)
		case "<":
			return i.toFloat(left) < i.toFloat(right)
		case ">":
			return i.toFloat(left) > i.toFloat(right)
		case "<=":
			return i.toFloat(left) <= i.toFloat(right)
		case ">=":
			return i.toFloat(left) >= i.toFloat(right)
		case "%":
			return float64(int64(i.toFloat(left)) % int64(i.toFloat(right)))
		case "&&":
			return i.toBool(left) && i.toBool(right)
		case "||":
			return i.toBool(left) || i.toBool(right)
		}
	case *InputExpr:
		fmt.Print("")
		if i.inputScanner.Scan() {
			return i.inputScanner.Text()
		}
		return ""
	case *BetaCall:
		// Built-in functions
		if n.Name == "len" {
			if len(n.Args) != 1 {
				panic("len expects 1 argument")
			}
			arg := i.evaluateExpression(n.Args[0])
			switch v := arg.(type) {
			case string:
				return float64(len(v))
			default:
				panic("len expects a string argument")
			}
		} else if n.Name == "abs" {
			if len(n.Args) != 1 {
				panic("abs expects 1 argument")
			}
			arg := i.toFloat(i.evaluateExpression(n.Args[0]))
			return math.Abs(arg)
		} else if n.Name == "str" {
			if len(n.Args) != 1 {
				panic("str expects 1 argument")
			}
			arg := i.evaluateExpression(n.Args[0])
			return i.toString(arg)
		}
		// User-defined function call
		fn, ok := i.functions[n.Name]
		if !ok {
			panic(fmt.Sprintf("Undefined function: %s", n.Name))
		}
		if len(fn.Params) != len(n.Args) {
			panic(fmt.Sprintf("Function %s expects %d args, got %d", n.Name, len(fn.Params), len(n.Args)))
		}
		frame := &callFrame{variables: make(map[string]interface{})}
		for idx, param := range fn.Params {
			frame.variables[param] = i.evaluateExpression(n.Args[idx])
		}
		i.callStack = append(i.callStack, frame)
		for _, stmt := range fn.Body {
			i.executeStatement(stmt)
			if frame.returned {
				break
			}
		}
		i.callStack = i.callStack[:len(i.callStack)-1]
		return frame.returnValue
	}
	panic(fmt.Sprintf("Unknown expression type: %T", node))
}

func (i *Interpreter) add(left, right interface{}) interface{} {
	if leftStr, ok := left.(string); ok {
		return leftStr + i.toString(right)
	}
	if rightStr, ok := right.(string); ok {
		return i.toString(left) + rightStr
	}
	return i.toFloat(left) + i.toFloat(right)
}

func (i *Interpreter) equals(left, right interface{}) bool {
	if leftStr, ok := left.(string); ok {
		if rightStr, ok := right.(string); ok {
			return leftStr == rightStr
		}
		return false
	}
	return i.toFloat(left) == i.toFloat(right)
}

func (i *Interpreter) toFloat(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
		return 0
	case bool:
		if v {
			return 1
		}
		return 0
	}
	return 0
}

func (i *Interpreter) toString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case float64:
		if v == float64(int64(v)) {
			return fmt.Sprintf("%.0f", v)
		}
		return fmt.Sprintf("%g", v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	}
	return fmt.Sprintf("%v", val)
}

func (i *Interpreter) toBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return v
	case float64:
		return v != 0
	case string:
		return v != ""
	}
	return false
}

func (i *Interpreter) executeStatement(stmt ASTNode) {
	switch s := stmt.(type) {
	case *VarDecl:
		value := i.evaluateExpression(s.Value)
		i.setVar(s.Name, value)
	case *Assignment:
		value := i.evaluateExpression(s.Value)
		i.setVar(s.Name, value)
	case *PrintStmt:
		value := i.evaluateExpression(s.Value)
		fmt.Println(i.toString(value))
	case *IfStmt:
		condition := i.evaluateExpression(s.Condition)
		if i.toBool(condition) {
			i.pushScope()
			for _, stmt := range s.ThenBlock {
				i.executeStatement(stmt)
				if i.currentFrame().returned {
					break
				}
			}
			i.popScope()
		} else if s.ElseBlock != nil {
			i.pushScope()
			for _, stmt := range s.ElseBlock {
				i.executeStatement(stmt)
				if i.currentFrame().returned {
					break
				}
			}
			i.popScope()
		}
	case *WhileStmt:
		for i.toBool(i.evaluateExpression(s.Condition)) {
			i.pushScope()
			for _, stmt := range s.Body {
				i.executeStatement(stmt)
				if i.currentFrame().returned {
					break
				}
			}
			i.popScope()
			if i.currentFrame().returned {
				break
			}
		}
	case *ForStmt:
		i.pushScope()
		if s.Init != nil {
			i.executeStatement(s.Init)
		}
		for i.toBool(i.evaluateExpression(s.Condition)) {
			i.pushScope()
			for _, stmt := range s.Body {
				i.executeStatement(stmt)
				if i.currentFrame().returned {
					break
				}
			}
			i.popScope()
			if i.currentFrame().returned {
				break
			}
			if s.Post != nil {
				i.executeStatement(s.Post)
			}
		}
		i.popScope()
	case *SigmaFunc:
		i.functions[s.Name] = s
	case *BetaCall:
		i.evaluateExpression(s)
	case *AlphaReturn:
		val := i.evaluateExpression(s.Value)
		i.currentFrame().returnValue = val
		i.currentFrame().returned = true
	}
}

func (i *Interpreter) pushScope() {
	i.callStack = append(i.callStack, &callFrame{variables: make(map[string]interface{})})
}

func (i *Interpreter) popScope() {
	if len(i.callStack) > 1 {
		i.callStack = i.callStack[:len(i.callStack)-1]
	}
}

func (i *Interpreter) Execute(program *Program) {
	for _, stmt := range program.Statements {
		i.executeStatement(stmt)
	}
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("🚽 Skibidi Programming Language v1.0")
		fmt.Println("Usage:")
		fmt.Println("  ./skibidi run <filename.skibidi>  - Run a Skibidi program (Linux/Mac)")
		fmt.Println("  .\\skibidi.exe run <filename.skibidi>  - Run a Skibidi program (Windows)")
		fmt.Println("  ./skibidi -i                      - Interactive mode (Linux/Mac)")
		fmt.Println("  .\\skibidi.exe -i                  - Interactive mode (Windows)")
		fmt.Println("  ./skibidi help                    - Show this help")
		fmt.Println("\n📚 Skibidi Keywords:")
		fmt.Println("  skibidi x rizz 5 ohio     - declare variable")
		fmt.Println("  x rizz 10 ohio            - assign variable")
		fmt.Println("  gyatt x ohio              - print variable")
		fmt.Println("  cap (x > 5) { ... }       - if statement")
		fmt.Println("  nocap { ... }             - else statement")
		fmt.Println("  bussin (x < 10) { ... }   - while loop")
		fmt.Println("  bruh this is a comment    - comment")
		fmt.Println("  ohio                      - end statement")
		return
	}

	command := os.Args[1]

	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please specify a file to run")
			fmt.Println("Usage: skibidi run <filename.skibidi>")
			return
		}

		filename := os.Args[2]
		if !strings.HasSuffix(filename, ".skibidi") {
			fmt.Printf("⚠️  Warning: File '%s' doesn't have .skibidi extension\n", filename)
		}

		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("❌ Error reading file '%s': %v\n", filename, err)
			return
		}

		fmt.Printf("🚀 Running Skibidi program: %s\n", filename)
		fmt.Println("" + strings.Repeat("=", 40))
		runSkibidi(string(content))
		fmt.Println("" + strings.Repeat("=", 40))
		fmt.Println("✅ Program execution completed!")

	case "-i", "interactive":
		runInteractive()

	case "help", "-h", "--help":
		fmt.Println("🚽 Skibidi Programming Language v1.0")
		fmt.Println("\n📚 Commands:")
		fmt.Println("  skibidi run <file>    - Run a Skibidi program")
		fmt.Println("  skibidi -i           - Start interactive mode")
		fmt.Println("  skibidi help         - Show this help")
		fmt.Println("\n🔧 Example Usage:")
		fmt.Println("  skibidi run hello.skibidi")
		fmt.Println("  skibidi -i")

	default:
		fmt.Printf("❌ Unknown command: %s\n", command)
		fmt.Println("Use 'skibidi help' for usage information")
	}
}

func runInteractive() {
	fmt.Println("🚽 Skibidi Interactive Mode v2.0 🚽")
	fmt.Println("Type :help for commands. Type 'exit' or :exit to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	interpreter := NewInterpreter()

	var inputLines []string
	var openBraces int

	for {
		prompt := "skibidi> "
		if openBraces > 0 {
			prompt = "... "
		}
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println("Goodbye! Stay sigma! 🗿")
			break
		}
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if openBraces == 0 && strings.HasPrefix(trimmed, ":") {
			replCmd := strings.ToLower(strings.TrimSpace(trimmed[1:]))
			switch replCmd {
			case "exit":
				fmt.Println("Goodbye! Stay sigma! 🗿")
				return
			case "help":
				fmt.Println("Available commands: :help, :vars, :funcs, :exit")
				continue
			case "vars":
				fmt.Println("Variables:")
				for k, v := range interpreter.currentFrame().variables {
					fmt.Printf("  %s = %v\n", k, v)
				}
				continue
			case "funcs":
				fmt.Println("Functions:")
				for k := range interpreter.functions {
					fmt.Printf("  %s\n", k)
				}
				continue
			default:
				fmt.Println("Unknown command. Type :help for help.")
				continue
			}
		}

		if trimmed == "exit" {
			fmt.Println("Goodbye! Stay sigma! 🗿")
			break
		}
		if trimmed == "" && openBraces == 0 {
			continue
		}

		openBraces += strings.Count(line, "{")
		openBraces -= strings.Count(line, "}")
		inputLines = append(inputLines, line)

		if openBraces > 0 {
			continue
		}

		input := strings.Join(inputLines, "\n")
		inputLines = nil
		openBraces = 0

		trimmedInput := strings.TrimSpace(input)
		if !strings.HasSuffix(trimmedInput, "ohio") && !strings.Contains(trimmedInput, "{") {
			input += " ohio"
		}

		// Try to parse as expression first, then as statement
		triedExpr := false
		var exprResult interface{}
		func() {
			defer func() {
				if r := recover(); r != nil {
					triedExpr = false
				} else {
					triedExpr = true
				}
			}()
			lexer := NewLexer(input)
			parser := NewParser(lexer)
			exprResult = interpreter.evaluateExpression(parser.parseExpression())
			triedExpr = true
		}()
		if triedExpr {
			fmt.Println(interpreter.toString(exprResult))
			continue
		}
		// If not an expression, try as statement
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Skibidi Error: %v\n", r)
				}
			}()
			lexer2 := NewLexer(input)
			parser2 := NewParser(lexer2)
			program := parser2.Parse()
			interpreter.Execute(program)
		}()
	}
}

func runSkibidi(code string) {
	interpreter := NewInterpreter()
	runSkibidiInterpreter(code, interpreter)
}

func runSkibidiInterpreter(code string, interpreter *Interpreter) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Skibidi Error: %v\n", r)
		}
	}()

	lexer := NewLexer(code)
	parser := NewParser(lexer)
	program := parser.Parse()
	interpreter.Execute(program)
}
