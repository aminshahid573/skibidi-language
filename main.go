package main

import (
	"bufio"
	"fmt"
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
	}

	if (l.peek() >= 'a' && l.peek() <= 'z') || (l.peek() >= 'A' && l.peek() <= 'Z') {
		identifier := l.readIdentifier()
		if tokenType, exists := keywords[identifier]; exists {
			return Token{tokenType, identifier, l.line}
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
	return p.parseComparison()
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

	for p.currentToken.Type == MULTIPLY || p.currentToken.Type == DIVIDE {
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
	} else if token.Type == IDENTIFIER {
		p.eat(IDENTIFIER)
		return &Identifier{Name: token.Value}
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

func (p *Parser) Parse() *Program {
	statements := []ASTNode{}

	for p.currentToken.Type != EOF {
		stmt := p.parseStatement()
		statements = append(statements, stmt)
	}

	return &Program{Statements: statements}
}

// Interpreter
type Interpreter struct {
	variables map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{variables: make(map[string]interface{})}
}

func (i *Interpreter) evaluateExpression(node ASTNode) interface{} {
	switch n := node.(type) {
	case *NumberLiteral:
		return n.Value
	case *StringLiteral:
		return n.Value
	case *Identifier:
		if val, exists := i.variables[n.Name]; exists {
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
		}
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
		i.variables[s.Name] = value
	case *Assignment:
		value := i.evaluateExpression(s.Value)
		i.variables[s.Name] = value
	case *PrintStmt:
		value := i.evaluateExpression(s.Value)
		fmt.Println(i.toString(value))
	case *IfStmt:
		condition := i.evaluateExpression(s.Condition)
		if i.toBool(condition) {
			for _, stmt := range s.ThenBlock {
				i.executeStatement(stmt)
			}
		} else if s.ElseBlock != nil {
			for _, stmt := range s.ElseBlock {
				i.executeStatement(stmt)
			}
		}
	case *WhileStmt:
		for i.toBool(i.evaluateExpression(s.Condition)) {
			for _, stmt := range s.Body {
				i.executeStatement(stmt)
			}
		}
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
	fmt.Println("🚽 Skibidi Interactive Mode - Type 'exit' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	interpreter := NewInterpreter()

	for {
		fmt.Print("skibidi> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Goodbye! Stay sigma! 🗿")
			break
		}

		if input == "" {
			continue
		}

		// Add ohio if missing for single statements
		if !strings.HasSuffix(input, "ohio") && !strings.Contains(input, "{") {
			input += " ohio"
		}

		runSkibidiInterpreter(input, interpreter)
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
