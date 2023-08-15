package net

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

/*
▄▀█ █▀▀ ▀█▀ █ █░█ ▄▀█ ▀█▀ █ █▀█ █▄░█   █▀▀ █░█ █▄░█ █▀▀ ▀█▀ █ █▀█ █▄░█ █▀
█▀█ █▄▄ ░█░ █ ▀▄▀ █▀█ ░█░ █ █▄█ █░▀█   █▀░ █▄█ █░▀█ █▄▄ ░█░ █ █▄█ █░▀█ ▄█

we are using a dict to store the funtions. each function will be stored under a key that is a string so that each node can have a key to a actavation function
all
*/
// dict to map funtions for the Node obj, this is done for dinamic calls to diffrent func.
// all func must be plan func and not methods with the parma(x float64, n Node) or cannot be placed inside the map
var Node_Activation_Functions map[string]func(float64, *Node) float64 = map[string]func(float64, *Node) float64{
	"NodeLU":    NodeLU,
	"ReLU":      ReLU,
	"Log":       Log,
	"ExpE":      ExpE,
	"ExpSqrtE":  ExpSqrtE,
	"ExpPi":     ExpPi,
	"ExpSqrtPi": ExpSqrtPi,
	"Collatz":   Collatz,
}

// rectified linear: takes a float64 as var x and returns 0 if input < 0 else returns var x
func ReLU(x float64, n *Node) float64 {
	if x < 0 {
		n.output = 0
		return 0
	} else {
		n.output = x
		return x
	}
}

// natral log: takes a float64 as var x and returns 0 if input < 0 else returns var log10(x)
func Log(x float64, n *Node) float64 {
	n.output = math.Log10(x)
	return n.output
}

// raw exponentiation via math.E
func ExpE(x float64, n *Node) float64 {
	n.output = math.Pow(math.E, x)
	return n.output
}

func NodeLU(x float64, n *Node) float64 {
	n.output = x
	return n.output
}

//experamental activation

// raw exponentiation via math.SqrtE
func ExpSqrtE(x float64, n *Node) float64 {
	n.output = math.Pow(math.SqrtE, x)
	return n.output
}

// raw exponentiation via PI
func ExpPi(x float64, n *Node) float64 {
	n.output = math.Pow(math.Pi, x)
	return n.output
}

// raw exponentiation via PI
func ExpSqrtPi(x float64, n *Node) float64 {
	n.output = math.Pow(math.SqrtPi, x)
	return n.output
}

// the hailstone algorithom or Collatz conjecture https://en.wikipedia.org/wiki/Collatz_conjecture
func Collatz(x float64, n *Node) float64 {

	// we half to tuncate the val so that the val can be an int then at the end of the funtion the floating point desapales are added to the val
	newx := math.Round(x)
	dif := x - newx

	if math.Mod(newx, 2) == 0 {
		newx = newx / 2
	} else {
		newx = 3*newx + 1
	}

	return newx + dif

}

var Layer_Activation_Functions map[string]func(*mat.VecDense, Layer) = map[string]func(x *mat.VecDense, L Layer){
	"NA":      NA,
	"SoftMax": SoftMax,
}

// Layer_based
// soft max

// this function is intededed to be usesd for the output of the forward pass so
// that the NN can be traned with the prodiction of how corect the model is
func SoftMax(x *mat.VecDense, L Layer) {
	for i := 0; i < x.Len(); i++ {
		x.SetVec(i, math.Pow(math.E, x.AtVec(i)))
	}
	normalize(0, x, x.Norm(1))
}

func normalize(i int, x *mat.VecDense, norm float64) {
	if i < x.Len() {
		x.SetVec(i, x.AtVec(i)/norm)
		i++
		normalize(i, x, norm)
	}
}

func NA(x *mat.VecDense, L Layer) {
}
