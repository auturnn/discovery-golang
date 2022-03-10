package main

import (
	"math"
)

//아래는 패턴추상화 예제
//패턴 추상화에 관한 자세한 예제를 통해 어떻게 쓰일 수 있는지를 확인한다.
//수학에 관한 예제이기에 모든 것을 이해하기는 힘들지만 구조만이라도 확인하고 넘어가자..

//Func는 y=f(x) 형태로 수학에서 쓰는 함수형이다. 실수값을 하나 받아서 실수값 하나를 돌려준다.
type Func func(float64) float64

//Transform은 Func를 받아서 Func를 돌려주는 함수형이다. 이것은 함수 하나를 다른 함수로 변환하는 함수의 형태
type Transform func(Func) Func

const tolerance = 0.00001
const dx = 0.00001

//Square 함수는 제곱을 구하는 함수이다.
func Square(x float64) float64 {
	return x * x
}

// FixedPoint 함수는 어떤 함수 f를 계속해서 적용했을 때, 어떤 값으로 수렴하는 경우에 그 수렴값을 찾는 함수다.
// f(x),f(f(x)),f(f(f(x))... 이런싣으로 계속 적용했을때, 어느 한 곳으로 수렴할 때 그 값을 찾는 것이다.
// ex) 제곱근. 1이상의 숫자에 제곱근을 계속 적용시키면 1로 수렴된다.
func FixedPoint(f Func, firstGuess float64) float64 {
	//closeEnough 함수는 두 수 v1,v2가 서로 tolerance 이하로 가까워졌으면 참을 돌려준다.
	closeEnough := func(v1, v2 float64) bool {
		return math.Abs(v1-v2) < tolerance
	}

	//try 함수는 guess에 반복적으로 함수 f를 적용시키다가 그 변화가 충분히 작으면 수렴된 것으로 보고 그 값을 반환하고 종료하는 함수이다.
	var try Func
	try = func(guess float64) float64 {
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		} else {
			return try(next)
		}
	}
	return try(firstGuess)
}

//FixedPointOfTransform 함수는 아까 전에 살펴본 FixedPoint의 아이디어를 함수로 변환하여 적용시켜 수렴값을 찾는 것이다.
//transform에 어떤 함수를 수렴하는 성질을 갖는 함수로 변환시키는 함수를 넘겨주면 이 아이디어가 동작한다.
func FixedPointOfTransform(g Func, transform Transform, guess float64) float64 {
	return FixedPoint(transform(g), guess)
}

//Deriv 함수는 어떤 함수를 받아서 다른 함수를 돌려주는 형태이기 때문에 Transform의 형태를 갖고있는 함수이다.
func Deriv(g Func) Func {
	//dx는 매우 작은 숫자로, 수식을 잘 보면 각 점에서 기울기를 구하고 있으므로
	//이 함수는 g함수를 받아서 미분된 함수로 g'를 돌려주는 함수다.
	return func(x float64) float64 {
		return (g(x+dx) - g(x)) / dx
	}
}

//NewtonTransform 함수 역시 Transform의 형태이다. 뉴턴 방법은 f(x) = 0이 되는 x를 찾는 방법중 하나이다.
//NewtonTransform은 함수 g를 받아 반복적으로 적용하면 g(x)=0에 점점 다가가는 x값을 반환하는 함수를 돌려준다.
//코드의 수식은 위키백과의 수식 그대로이다.
func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		return x - (g(x) / Deriv(g)(x))
	}
}

//Sqrt 함수는 x의 제곱근을 구하는 함수이다. 그런데 제곱은 Square함수에서 보다시피 구하기 쉬운데,
//제곱근은 구하기 어렵다. 이것은 뉴턴방법을 이용하여 제곱함수인 Square의 역함수를 구하는 방법으로 제곱근을 구해낸다.
//y=Sqrt(x)의 형태로 x를 넣어서 y를 구하고 싶은데, square는 구하기 쉬우므로 Square(y)=x를 구하면 되는 것이다.
//여기서 x는 이미 주어진 값이고, y는 알고싶은 값이 된다. 우변을 좌변으로 이항하면 Square(y)-x=0이 된다.
//아래의 코드를 보면 FixedPointOfTransform을 호출하는 첫 번째 인자로 y에 대한 이 함수를 넘기고 있는데,
//이 식을 뉴턴변환하고 나서 FixedPoint 패턴으로 수렴하는 y값을 찾으면 된다.
func Sqrt(x float64) float64 {
	return FixedPointOfTransform(func(y float64) float64 {
		return Square(y) - x
	}, NewtonTransform, 1.0)
}

//뉴턴 방법을 적용할 때, 다른 함수들의 도움을 받지 않았다면 더 복잡했을 것이다.
//뉴턴방법 내에 뉴턴 변환을 한 함수를 수렴할 때까지 반복 적용한다는 아이디어를
//FixedPoint 함수가 추상화하여 만들어 놓은 아이디어를 이용하여 구현되었다.
//이 예제에서는 대부분의 함수들이 일회용으로 쓰였지만,
//이 각각의 아이디어들은 여러 번 다른 계산에 쓰일 수 있는 아이디어이기에 더욱 가치가 있다.
