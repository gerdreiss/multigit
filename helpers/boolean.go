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

func IfElseErr[T any](cond bool, ifTrue func() (T, error), ifFalse func() (T, error)) (T, error) {
	if cond {
		return ifTrue()
	}
	return ifFalse()
}

// Helper to wrap values in functions
func Val[T any](v T) func() T {
	return func() T { return v }
}
