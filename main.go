package main

import (
	"fmt"
	"os"
	"strconv"
)

// expressão = número ( operador número )*
// operador = "+" | "-"
// número = um ou mais dígitos

type Parser struct {
	input string
	pos   int
}

func (p *Parser) skipSpaces() {
	for p.pos < len(p.input) && p.input[p.pos] == ' ' {
		p.pos++
	}
}

func (p *Parser) parseNumber() int {
	if p.pos >= len(p.input) || !isDigit(p.input[p.pos]) {
		panic("esperado número na posição " + strconv.Itoa(p.pos))
	}

	inicio := p.pos
	for p.pos < len(p.input) && isDigit(p.input[p.pos]) {
		p.pos++
	}

	numero, err := strconv.Atoi(p.input[inicio:p.pos])
	if err != nil {
		panic("número inválido: " + p.input[inicio:p.pos])
	}
	return numero
}

func (p *Parser) parseExpression() int {
	p.skipSpaces()
	resultado := p.parseNumber()

	for p.pos < len(p.input) {
		p.skipSpaces()
		if p.pos >= len(p.input) {
			break
		}

		operador := p.input[p.pos]
		if operador != '+' && operador != '-' {
			panic("operador inválido na posição " + strconv.Itoa(p.pos))
		}
		p.pos++

		p.skipSpaces()
		numero := p.parseNumber()

		if operador == '+' {
			resultado += numero
		} else {
			resultado -= numero
		}
	}

	return resultado
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	if len(os.Args) != 2 {
		panic("uso: main <expressão>")
	}

	p := Parser{input: os.Args[1]}
	fmt.Println(p.parseExpression())
}
