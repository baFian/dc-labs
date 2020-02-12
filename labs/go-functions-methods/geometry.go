// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"math"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
    
)
type Point struct{ x, y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) X()float64{
	return p.x;
}
func (p Point) Y()float64{
	return p.y;
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}


func main(){
	
	
	sides, _ := strconv.Atoi(os.Args[1])
	Vertex :=  Path{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i:=0;i<sides;i++{
		
		x:=r1.Float64()*5
		y:=r1.Float64()*5
		point :=Point{x, y}
		Vertex = append(Vertex, point)
		
	}
	fmt.Print(Vertex)
		
	fmt.Print(Vertex.Distance())
	
}



//!-path