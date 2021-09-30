package main

type visitor interface { // visitor interface
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

//The functions will allow us to add functionality for
// squares, circles and triangles
