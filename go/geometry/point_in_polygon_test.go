package geometry

import "testing"

var complicatedConvexPolygon = Polygon{
	[]Point{
		Pt(77.33, 323.47),
		Pt(92.67, 252.47),
		Pt(71.33, 134.00),
		Pt(91.17, 81.63),
		Pt(172.00, 134.26),
		Pt(215.67, 220.83),
		Pt(227.52, 326.30),
		Pt(185.45, 329.04),
		Pt(179.41, 274.41),
		Pt(125.39, 165.09),
		Pt(135.38, 279.78),
		Pt(145.08, 348.72),
		Pt(202.50, 360.61),
		Pt(267.33, 341.40),
		Pt(249.20, 135.92),
		Pt(313.06, 128.43),
		Pt(332.93, 218.95),
		Pt(300.79, 307.47),
		Pt(333.25, 362.54),
		Pt(387.17, 79.68),
		Pt(357.00, 49.96),
		Pt(266.67, 84.53),
		Pt(58.33, 29.53),
		Pt(24.33, 181.50),
		Pt(35.33, 330.49),
	}}

var unitSquare = Polygon{[]Point{Pt(0, 0), Pt(0, 1), Pt(1, 1), Pt(1, 0)}}

var pipTests = []struct {
	pt   Point
	poly Polygon
	in   bool
}{
	// Simple triangle, with multiple vertex orderings.
	{Pt(0, 0), Polygon{[]Point{Pt(1, 2), Pt(2, 3), Pt(3, 1)}}, false},
	{Pt(3, 3), Polygon{[]Point{Pt(1, 2), Pt(2, 3), Pt(3, 1)}}, false},
	{Pt(-1, -1), Polygon{[]Point{Pt(1, 2), Pt(2, 3), Pt(3, 1)}}, false},
	{Pt(2, 2), Polygon{[]Point{Pt(1, 2), Pt(2, 3), Pt(3, 1)}}, true},
	{Pt(2, 2), Polygon{[]Point{Pt(2, 3), Pt(3, 1), Pt(1, 2)}}, true},
	{Pt(2, 2), Polygon{[]Point{Pt(3, 1), Pt(1, 2), Pt(2, 3)}}, true},
	{Pt(2, 2), Polygon{[]Point{Pt(3, 1), Pt(2, 3), Pt(1, 2)}}, true},

	// Horizontal edge tests, using the unit square at origin.
	{Pt(0.5, 0.5), unitSquare, true},
	{Pt(0, 1.01), unitSquare, false},
	{Pt(1.01, 0), unitSquare, false},
	{Pt(-1, 1), unitSquare, false},
	{Pt(-1, 0), unitSquare, false},
	{Pt(2, 1), unitSquare, false},
	{Pt(2, 0), unitSquare, false},
	{Pt(0.5, 0), unitSquare, true},  // on right-to-left horizontal edge
	{Pt(0.5, 1), unitSquare, false}, // on left-to-right horizontal edge
	{Pt(0, 0.5), unitSquare, true},  // on left edge
	{Pt(1, 0.5), unitSquare, false}, // on right edge

	// Convex polygon.
	{Pt(1.33, 181.50), complicatedConvexPolygon, false},
	{Pt(173.67, 205.83), complicatedConvexPolygon, false},
	{Pt(290.06, 195.43), complicatedConvexPolygon, false},
	{Pt(389.25, 202.54), complicatedConvexPolygon, false},
	{Pt(156.38, 290.78), complicatedConvexPolygon, true},
	{Pt(217.00, 110.26), complicatedConvexPolygon, true},
	{Pt(357.00, 67.96), complicatedConvexPolygon, true},
	{Pt(48.67, 254.47), complicatedConvexPolygon, true},

	// Point in the center of a very small square.
	{Pt(10000, 10000), Polygon{[]Point{
		Pt(9999.9999999999, 9999.9999999999),
		Pt(9999.9999999999, 10000.0000000001),
		Pt(10000.0000000001, 10000.0000000001),
		Pt(10000.0000000001, 9999.9999999999),
	}}, true},
}

func TestPointInPolygon(t *testing.T) {
	for i, tt := range pipTests {
		if PointInPolygon(tt.pt, tt.poly) != tt.in {
			t.Errorf("#%02d: expected PointInPolygon() => %v", i, tt.in)
		}
	}
}
