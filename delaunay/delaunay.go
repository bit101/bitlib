// Package delaunay does triangulation
package delaunay

import (
	"math"

	"github.com/bit101/bitlib/geom"
)

// Triangulate does the Triangulation and returns a list of triangles.
func Triangulate(points geom.PointList) geom.TriangleList {
	superTri := getSuperTri(points)
	triangulation := geom.NewTriangleList()
	triangulation.Add(superTri)

	for _, p := range points {
		badTriangles := getBadTriangles(triangulation, p)
		polygon := getPolygon(badTriangles)
		triangulation = cullBadTriangles(triangulation, badTriangles)
		triangulation = addTriangles(p, polygon, triangulation)
	}
	triangulation = removeOuter(superTri, triangulation)
	return triangulation
}

// TriangulateEdges does the Triangulation and returns a list of segments.
func TriangulateEdges(points geom.PointList) geom.SegmentList {
	triangles := Triangulate(points)
	return triangles.Edges()
}

// getBadTriangles gets a list of trianges whose CircumCircle contains the given point
func getBadTriangles(triangulation geom.TriangleList, p *geom.Point) geom.TriangleList {
	badTriangles := geom.NewTriangleList()
	for _, t := range triangulation {
		c := t.CircumCircle()
		if c.Contains(p) {
			badTriangles.Add(t)
		}
	}
	return badTriangles
}

// cullBadTriangles culls the bad triangles from the triangulation
func cullBadTriangles(triangulation, badTriangles geom.TriangleList) geom.TriangleList {
	for _, t := range badTriangles {
		triangulation = removeTriangle(triangulation, t)
	}
	return triangulation
}

// getPolygon gets a unique list of edges from the bad triangle list
func getPolygon(badTriangles geom.TriangleList) geom.SegmentList {
	polygon := geom.NewSegmentList()
	for _, t := range badTriangles {
		for _, e := range t.Edges() {
			if !hasSharedEdge(badTriangles, t, e) {
				polygon = addEdgeToPolygon(polygon, e)
			}
		}
	}
	return polygon
}

// addEdgeToPolygon adds an edge to the polygon confirming it's unique
func addEdgeToPolygon(list geom.SegmentList, edge *geom.Segment) geom.SegmentList {
	for _, p := range list {
		if p.Equals(edge) {
			return list
		}
	}
	return append(list, edge)
}

// removeOuter removes any triangles that contain points in the original super triangle
func removeOuter(superTri *geom.Triangle, triangulation geom.TriangleList) geom.TriangleList {
	removals := geom.NewTriangleList()
	for _, t := range triangulation {
		for _, p := range t.Points() {
			if p.Equals(superTri.PointA) || p.Equals(superTri.PointB) || p.Equals(superTri.PointC) {
				removals = append(removals, t)
			}
		}
	}
	for _, t := range removals {
		for i := 0; i < len(triangulation); i++ {
			if triangulation[i].Equals(t) {
				triangulation = append(triangulation[:i], triangulation[i+1:]...)
			}
		}
	}
	return triangulation
}

// removeTriangle removes a triangle from the triangulation
func removeTriangle(triangulation geom.TriangleList, triangle *geom.Triangle) geom.TriangleList {
	for i, t := range triangulation {
		if t.Equals(triangle) {
			triangulation = append(triangulation[:i], triangulation[i+1:]...)
			break
		}
	}
	return triangulation
}

// addTriangles adds a new triangle to the triangulation based on an edge + a point
func addTriangles(p *geom.Point, polygon geom.SegmentList, triangulation geom.TriangleList) geom.TriangleList {
	for _, e := range polygon {
		tri := geom.NewTriangle(e.PointA.X, e.PointA.Y, e.PointB.X, e.PointB.Y, p.X, p.Y)
		triangulation = append(triangulation, tri)
	}
	return triangulation
}

// hasSharedEdge checks if a given edge of a triangle is shared with any other triangles in the list
func hasSharedEdge(triangles geom.TriangleList, triangle *geom.Triangle, edge *geom.Segment) bool {
	for _, t := range triangles {
		if t != triangle {
			for _, e := range t.Edges() {
				if edge.Equals(e) {
					return true
				}
			}
		}
	}
	return false
}

// getSuperTri returns a triangle that contains the given area.
func getSuperTri(points geom.PointList) *geom.Triangle {
	minX := math.MaxFloat64
	maxX := -math.MaxFloat64
	minY := math.MaxFloat64
	maxY := -math.MaxFloat64
	for _, p := range points {
		minX = math.Min(minX, p.X)
		maxX = math.Max(maxX, p.X)
		minY = math.Min(minY, p.Y)
		maxY = math.Max(maxY, p.Y)
	}
	x := minX
	y := minY
	w := maxX - minX
	h := maxY - minY
	return geom.NewTriangle(
		x-w*0.1, y-h,
		x-w*0.1, y+h*2,
		x+w*1.7, y+h*0.5,
	)
}

func contains(tri *geom.Triangle, points geom.PointList) bool {
	for _, p := range points {
		if !tri.Contains(p) {
			return false
		}
	}
	return true
}
