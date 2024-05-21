package main

import (
	"strings"

	"golang.org/x/net/html"
)

// LinkMatrix is an interface for a link matrix between websites
type LinkMatrix interface {
	// AddLink adds a link from one website to another
	AddLink(from string, to string)
	// GetLinks returns a list of websites linked from the given website
	GetLinks(from string) []string
	// Compute a matrix of link endorsement, where the source website gives 1/n endorsement to each of the target websites it links to
	ComputeLinkEndorcementMatrix() endorcementMatrix
	// GenerateLinkEndorcementMatrixFromHTML generates the link endorsement matrix from HTML content
	GenerateLinkEndorcementMatrixFromHTML(html string, fileName string)
	// GetEndorcementMatrix returns the endorsement matrix
	GetEndorcementMatrix() endorcementMatrix
	// Compute generates the link endorsement matrix from HTML content
	Compute(html string, fileName string)
}

type endorcementMatrix map[string]map[string]float64

type linkMatrix struct {
	links             map[string][]string
	endorcementMatrix endorcementMatrix
}

// NewLinkMatrix creates a new LinkMatrix
func NewLinkMatrix() LinkMatrix {
	return &linkMatrix{
		links:             make(map[string][]string),
		endorcementMatrix: make(endorcementMatrix),
	}
}

func (lm *linkMatrix) AddLink(from string, to string) {
	lm.links[from] = append(lm.links[from], to)
}

func (lm *linkMatrix) GetLinks(from string) []string {
	return lm.links[from]
}

func (lm *linkMatrix) GenerateLinkEndorcementMatrixFromHTML(htmlContent string, fileName string) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		panic("failed to parse HTML")
	}

	from := fileName
	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					to := attr.Val
					if from != "" && from != to {
						lm.AddLink(from, to)
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLinks(c)
		}
	}

	findLinks(doc)
}

func (lm *linkMatrix) ComputeLinkEndorcementMatrix() endorcementMatrix {
	matrix := make(endorcementMatrix)
	for from, tos := range lm.links {
		matrix[from] = make(map[string]float64)
		for _, to := range tos {
			matrix[from][to] = 1 / float64(len(tos))
		}
	}
	lm.endorcementMatrix = matrix
	return matrix
}

func (lm *linkMatrix) GetEndorcementMatrix() endorcementMatrix {
	return lm.endorcementMatrix
}

func (lm *linkMatrix) Compute(htmlContent string, fileName string) {
	lm.GenerateLinkEndorcementMatrixFromHTML(htmlContent, fileName)
	lm.ComputeLinkEndorcementMatrix()
}
