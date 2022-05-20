package mo

import "fmt"

var eitherShouldBeLeftOrRight = fmt.Errorf("either should be Left or Right")
var eitherMissingLeftValue = fmt.Errorf("no such Left value")
var eitherMissingRightValue = fmt.Errorf("no such Right value")

// Left builds the left side of the Either struct, as opposed to the Right side.
func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{
		isLeft:  true,
		isRight: false,
		left:    value,
	}
}

// Right builds the right side of the Either struct, as opposed to the Left side.
func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{
		isLeft:  false,
		isRight: true,
		right:   value,
	}
}

// Either respresents a value of 2 possible types.
// An instance of Either is an instance of either A or B.
type Either[L any, R any] struct {
	isLeft  bool
	isRight bool

	left  L
	right R
}

// IsLeft returns true if Either is an instance of Left.
func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

// IsRight returns true if Either is an instance of Right.
func (e Either[L, R]) IsRight() bool {
	return e.isRight
}

// Left returns left value of a Either struct.
func (e Either[L, R]) Left() (L, bool) {
	if e.isLeft {
		return e.left, true
	}
	return empty[L](), false
}

// Right returns right value of a Either struct.
func (e Either[L, R]) Right() (R, bool) {
	if e.isRight {
		return e.right, true
	}
	return empty[R](), false
}

// MustLeft returns left value of a Either struct or panics.
func (e Either[L, R]) MustLeft() L {
	if !e.isLeft {
		panic(eitherMissingLeftValue)
	}

	return e.left
}

// MustRight returns right value of a Either struct or panics.
func (e Either[L, R]) MustRight() R {
	if !e.isRight {
		panic(eitherMissingRightValue)
	}

	return e.right
}

// LeftOrElse returns left value of a Either struct or fallback.
func (e Either[L, R]) LeftOrElse(fallback L) L {
	if e.isLeft {
		return e.left
	}

	return fallback
}

// RightOrElse returns right value of a Either struct or fallback.
func (e Either[L, R]) RightOrElse(fallback R) R {
	if e.isRight {
		return e.right
	}

	return fallback
}

// LeftOrEmpty returns left value of a Either struct or empty value.
func (e Either[L, R]) LeftOrEmpty() L {
	if e.isLeft {
		return e.left
	}

	return empty[L]()
}

// RightOrEmpty returns right value of a Either struct or empty value.
func (e Either[L, R]) RightOrEmpty() R {
	if e.isRight {
		return e.right
	}

	return empty[R]()
}

// Swap returns the left value in Right and vice versa.
func (e Either[L, R]) Swap() Either[R, L] {
	if e.isLeft {
		return Right[R, L](e.left)
	}

	return Left[R, L](e.right)
}

// ForEach executes the given side-effecting function, depending of value is Left or Right.
func (e Either[L, R]) ForEach(leftCb func(L), rightCb func(R)) {
	if e.isLeft {
		leftCb(e.left)
	} else if e.isRight {
		rightCb(e.right)
	}
}

// Match executes the given function, depending of value is Left or Right, and returns result.
func (e Either[L, R]) Match(onLeft func(L) Either[L, R], onRight func(R) Either[L, R]) Either[L, R] {
	if e.isLeft {
		return onLeft(e.left)
	} else if e.isRight {
		return onRight(e.right)
	}

	panic(eitherShouldBeLeftOrRight)
}

// MapLeft executes the given function, if Either is of type Left, and returns result.
func (e Either[L, R]) MapLeft(mapper func(L) Either[L, R]) Either[L, R] {
	if e.isLeft {
		return mapper(e.left)
	} else if e.isRight {
		return Right[L, R](e.right)
	}

	panic(eitherShouldBeLeftOrRight)
}

// MapRight executes the given function, if Either is of type Right, and returns result.
func (e Either[L, R]) MapRight(mapper func(R) Either[L, R]) Either[L, R] {
	if e.isLeft {
		return Left[L, R](e.left)
	} else if e.isRight {
		return mapper(e.right)
	}

	panic(eitherShouldBeLeftOrRight)
}
