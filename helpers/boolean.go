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
