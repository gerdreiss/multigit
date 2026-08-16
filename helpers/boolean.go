/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package helpers

func IfElse[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

func IfElseLazy[T any](cond bool, ifTrue func() T, ifFalse func() T) T {
	if cond {
		return ifTrue()
	}
	return ifFalse()
}
