package entities

import "fmt"

// Either é uma estrutura genérica que pode conter um valor do tipo L (esquerda) ou R (direita).
type Either[L any, R any] struct {
	left  *L
	right *R
}

// Left cria um Either representando um erro ou valor "esquerda".
func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{left: &value}
}

// Right cria um Either representando um valor de sucesso ou "direita".
func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{right: &value}
}

// IsLeft verifica se o Either contém um valor "esquerda".
func (e Either[L, R]) IsLeft() bool {
	return e.left != nil
}

// IsRight verifica se o Either contém um valor "direita".
func (e Either[L, R]) IsRight() bool {
	return e.right != nil
}

// LeftValue retorna o valor "esquerda" (erro) se estiver presente.
// Retorna um erro se tentar acessar o lado errado.
func (e Either[L, R]) LeftValue() (*L, error) {
	if !e.IsLeft() {
		return nil, fmt.Errorf("attempted to access LeftValue, but Either contains a RightValue")
	}
	return e.left, nil
}

// RightValue retorna o valor "direita" (sucesso) se estiver presente.
// Retorna um erro se tentar acessar o lado errado.
func (e Either[L, R]) RightValue() (*R, error) {
	if !e.IsRight() {
		return nil, fmt.Errorf("attempted to access RightValue, but Either contains a LeftValue")
	}
	return e.right, nil
}

// String retorna uma representação de string para depuração.
func (e Either[L, R]) String() string {
	if e.IsLeft() {
		return fmt.Sprintf("Either.Left(%v)", *e.left)
	}
	if e.IsRight() {
		return fmt.Sprintf("Either.Right(%v)", *e.right)
	}
	return "Either.Empty()"
}
