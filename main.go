package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

// getHTMLContent reads the HTML content from a given file path and returns it as a string
func getHTMLContent(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	var htmlContent string
	for {
		buffer := make([]byte, 100)
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			break
		}
		if n == 0 {
			break
		}
		htmlContent += string(buffer[:n])
	}
	return htmlContent
}

// run linkMatrix.go using one of websites on ~/websites/
func main() {
	//get the html content of the website
	ranker := NewRanker()
	var webNetworkSize = 50

	for i := 1; i <= webNetworkSize; i++ {
		filename := fmt.Sprintf("website%d.html", i)
		path := path.Join("websites", filename)
		htmlContent := getHTMLContent(path)
		//create a new link matrix
		lm := NewLinkMatrix()
		//generate the link endorsement matrix from the html content
		lm.Compute(htmlContent, filename)
		//get the endorsement matrix
		endorcementMatrix := lm.GetEndorcementMatrix()
		//add the endorsement matrix to the ranker
		ranker.AddEndorcementMatrix(endorcementMatrix)
	}
	matrix := ranker.ComputeProbabilityMatrix(100)
	ranker.SortedProbabilityMatrix()
	// print the sum of the probabilities of all the websites
	sum := 0.0
	for _, probability := range matrix {
		sum += probability
	}
	fmt.Println("Sum of probabilities:", sum)
}
