package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter:=func (t testing.TB, shape Shape,want float64){
		t.Helper()
		got:=shape.Perimeter()
		if got !=want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("perimeter rectangle",func(t *testing.T) {
		rectangle:=Rectangle{Width:10.0,Height: 10.0}
		checkPerimeter(t,rectangle,40)
	})
	t.Run("Perimeter of a circle",func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkPerimeter(t,circle,want)
	
	})
	
}

func TestArea(t *testing.T) {

	checkArea:= func (t testing.TB, shape Shape,want float64)  {
		t.Helper()
		got:=shape.Area()
		if got !=want {
			t.Errorf("got %g want %g", got, want)
		}
		
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t,rectangle,72.0)
		
	})
	t.Run("Area Circle",func(t *testing.T) {
		circle := Circle{10}
		checkArea(t,circle,314.1592653589793)
	})

	

}