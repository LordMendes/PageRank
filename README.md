# PageRank
## Project Overview
This project is a simple web page ranker written in Go. It uses a link matrix to represent the relationships between different web pages, and a ranker to calculate the relevance of each page based on the link endorsements it receives from other pages.

## Files
### linkMatrix.go
This file defines the LinkMatrix interface and its implementation. The LinkMatrix is a data structure that represents the relationships between different web pages. It provides methods to add a link from one page to another, get all links from a page, compute a link endorsement matrix, generate a link endorsement matrix from HTML content, and get the endorsement matrix.

### main.go
This file contains the main function that runs the program. It reads HTML content from a given file path, creates a new link matrix, generates the link endorsement matrix from the HTML content, gets the endorsement matrix, and adds the endorsement matrix to the ranker.

### ranker.go
This file defines the Ranker interface and its implementation. The Ranker is a data structure that calculates the relevance of each page based on the link endorsements it receives from other pages. It provides methods to compute a probability matrix, add an endorsement matrix, print the probability matrix, and sort the probability matrix.

## How to Run
Execute the `generator.py`  script to generate random html files

To run the program, simply execute the main.go file. The program will read the HTML content from the specified file path, create a link matrix, generate a link endorsement matrix, get the endorsement matrix, add the endorsement matrix to the ranker, compute a probability matrix, and print the sorted probability matrix.

## Dependencies
This project uses the golang.org/x/net/html package to parse HTML content.

## Note
This project is a simple demonstration of how to rank web pages based on link endorsements. It does not handle many real-world complexities and is not intended for use in a production environment.

## TODO
- Refactor so it don't look like a draft
- Add search query handling
- Paralelize
- Create a crawler and try to do the real stuffstuff
- Use limits, instead of a specific N number of rounds
