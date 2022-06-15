package main

import "testing"




func TestArea(t *testing.T) {


	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {

	// 		t.Errorf("got %g want %g", got, want)

	// 	}

	// }

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want:314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}


	for _, tt := range areaTests {

		got := tt.shape.Area()
		want := tt.want

		if got != want {

			t.Errorf("%#v Got %g want %g", tt.shape,  got, want)
		}
		
	}

	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkArea(t, rectangle, 100.0)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
	
	

}


// func TestArea(t *testing.T) {
// 	reactangle := Rectangle{5.0, 5.0}
// 	got :=  Area(reactangle)
// 	want := 25.0


// 	if got != want {

// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

