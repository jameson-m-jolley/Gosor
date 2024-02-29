package net

import (
	"math"
)

//this will be the implamentation of traing the network

type Trainer struct {
	inputs           [][]float64
	outputs          [][]float64
	network          *Gorebrum
	hotvector        int
	rate_of_change   float64
	algorithm        string
	Traing_algoithms map[string]func(*Trainer)
}

// traing functions
func no_mutation(T *Trainer) {
	println("Warring: no changes where made to the network while traning plese use set_algorithm to train the network")
}

func rand_mutation(T *Trainer) {
	print("this is not finished")
}

// this creates a new training obj for training a model on a data set
// doing it this way will give is freedom to implement different training strategies
func New_trainer(network *Gorebrum) *Trainer {
	this := Trainer{
		network:   network,
		algorithm: "",
		inputs:    [][]float64{{}},
		outputs:   [][]float64{{}},
		Traing_algoithms: map[string]func(*Trainer){
			"":              no_mutation,
			"rand_mutation": rand_mutation,
		},
	}
	return &this
}

func (T *Trainer) set_algorithm(name string, fn func(*Trainer)) {
	T.Traing_algoithms[name] = fn
}

func (T *Trainer) compute_rate_of_change() float64 {
	loss1 := Compute_loss(T.outputs[len(T.outputs)-1][T.hotvector])
	loss2 := Compute_loss(T.outputs[len(T.outputs)][T.hotvector])

	return (loss2 - loss1) / 1

}

func Compute_loss(value float64) float64 {
	return -(math.Log(value))
}

func (T *Trainer) Compute_Training_Pass() {

}
