package main

import "fmt"

// Ranker is an interface that will rank the relevance of a website to a given search query based on the all endorcementMatrix
type probabilityMatrix map[string]float64

type Ranker interface {
	ComputeProbabilityMatrix(rounds int) probabilityMatrix
	AddEndorcementMatrix(endorcementMatrix endorcementMatrix)
	PrintProbabilityMatrix()
	SortedProbabilityMatrix()
}

type ranker struct {
	probabilityMatrix     probabilityMatrix
	endorcementMatrixList []endorcementMatrix
}

// NewRanker creates a new Ranker
func NewRanker() Ranker {
	return &ranker{
		probabilityMatrix: make(probabilityMatrix),
	}
}

func (r *ranker) ComputeProbabilityMatrix(rounds int) probabilityMatrix {
	// initialize the probability matrix with 1/n for each website
	for _, endorcementMatrix := range r.endorcementMatrixList {
		for from := range endorcementMatrix {
			r.probabilityMatrix[from] = 1 / float64(len(r.endorcementMatrixList))
		}
	}

	// multiply the endorcement matrix by the probablity matrix for the given number of rounds
	for i := 0; i < rounds; i++ {
		newProbabilityMatrix := make(probabilityMatrix)
		for from, tos := range r.probabilityMatrix {
			for _, endorcementMatrix := range r.endorcementMatrixList {
				for to, endorsement := range endorcementMatrix[from] {
					newProbabilityMatrix[to] += tos * endorsement
				}
			}
		}
		r.probabilityMatrix = newProbabilityMatrix
	}

	return r.probabilityMatrix
}

func (r *ranker) AddEndorcementMatrix(endorcementMatrix endorcementMatrix) {
	r.endorcementMatrixList = append(r.endorcementMatrixList, endorcementMatrix)
}

// PrintProbabilityMatrix prints the probability matrix a X Y table where X is the website that links to Y and Y is the website that is linked to
func (r *ranker) PrintProbabilityMatrix() {
	for from, tos := range r.probabilityMatrix {
		fmt.Printf("%s -> %f\n", from, tos)
	}
}

func (r *ranker) SortedProbabilityMatrix() {
	// sort the probability matrix by the probability
	type websiteProbability struct {
		website     string
		probability float64
	}
	var websiteProbabilities []websiteProbability
	for website, probability := range r.probabilityMatrix {
		websiteProbabilities = append(websiteProbabilities, websiteProbability{website: website, probability: probability})
	}
	// sort the websiteProbabilities by the probability
	for i := 0; i < len(websiteProbabilities); i++ {
		for j := i + 1; j < len(websiteProbabilities); j++ {
			if websiteProbabilities[i].probability < websiteProbabilities[j].probability {
				websiteProbabilities[i], websiteProbabilities[j] = websiteProbabilities[j], websiteProbabilities[i]
			}
		}
	}
	// print the sorted websiteProbabilities
	for _, websiteProbability := range websiteProbabilities {
		fmt.Printf("%s -> %f\n", websiteProbability.website, websiteProbability.probability)
	}
}
