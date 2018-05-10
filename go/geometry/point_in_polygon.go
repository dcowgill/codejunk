package geometry

type Point struct {
	x, y float64
}

func Pt(x, y float64) Point { return Point{x, y} }
func (p Point) X() float64  { return p.x }
func (p Point) Y() float64  { return p.y }

// Polygon represents a possibly non-simple polygon.
type Polygon struct {
	vertices []Point
}

func NewPolygon(vertices []Point) Polygon { return Polygon{vertices} }
func (p Polygon) Vertex(i int) Point      { return p.vertices[i] }
func (p Polygon) NumVertices() int        { return len(p.vertices) }

// Return true if p is inside poly.
//
// Uses the even-odd rule, as described in the SVG specification: "This rule
// determines the "insideness" of a point on the canvas by drawing a ray from
// that point to infinity in any direction and counting the number of path
// segments from the given shape that the ray crosses. If this number is odd,
// the point is inside; if even, the point is outside."
//
// Not recommended for use with non-simple polygons.
func PointInPolygon(p Point, poly Polygon) bool {
	nvert := len(poly.vertices)
	odd := false
	for i, j := 0, nvert-1; i < nvert; i++ {
		v1, v2 := poly.vertices[i], poly.vertices[j]
		if (v1.y <= p.y && p.y < v2.y) || (v2.y <= p.y && p.y < v1.y) {
			if v1.x >= p.x || v2.x >= p.x {
				// Compute intersection of ray with edge (v1,v2) at y, where
				// v1=(x1,y1) and v2=(x2,y2). Solve for x:
				//
				//		slope = (y2-y1) / (x2-x1) = (y-y1) / (x-x1)
				//		(y2-y1) / (x2-x1) = (y-y1) / (x-x1)
				//		(y2-y1) / (x2-x1) * (x-x1) = (y-y1)
				//		x = (y-y1) * (x2-x1) / (y2-y1) + x1
				//
				// Point must be strictly to the right of the edge.
				if p.x < (p.y-v1.y)*(v2.x-v1.x)/(v2.y-v1.y)+v1.x {
					odd = !odd
				}
			}
		}
		j = i // j follows i
	}
	return odd
}
