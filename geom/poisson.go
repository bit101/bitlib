// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
)

// PoissonDiskSampling returns a poisson disk sampling of points
func PoissonDiskSampling(x, y, width, height, radius float64, tries int) PointList {
	points := NewPointList()
	active := NewPointList()
	cellSize := math.Floor(radius / math.Sqrt(2))
	cols := int(math.Ceil(width/cellSize) + 1)
	rows := int(math.Ceil(height/cellSize) + 1)
	grid := make([][]*Point, cols)

	for i := 0; i < cols; i++ {
		col := make([]*Point, rows)
		grid[i] = col
	}

	p0 := RandomPointInRect(0, 0, width, height)
	insertPoissonPoint(grid, cellSize, p0)
	points.Add(p0)
	active.Add(p0)

	for len(active) > 0 {
		index := random.IntRange(0, len(active))
		p := active[index]
		found := false
		for i := 0; i < tries; i++ {
			angle := random.Angle()
			r := random.FloatRange(radius, radius*2)
			pNew := NewPoint(
				p.X+math.Cos(angle)*r,
				p.Y+math.Sin(angle)*r,
			)

			if isValidPoint(grid, cols, rows, pNew, cellSize, width, height, radius) {
				insertPoissonPoint(grid, cellSize, pNew)
				points.Add(pNew)
				active.Add(pNew)
				found = true
				break
			}
		}
		if !found {
			active = append(active[:index], active[index+1:]...)
		}

	}

	points.Translate(x, y)
	return points
}

func insertPoissonPoint(grid [][]*Point, cellSize float64, point *Point) {
	x := int(point.X / cellSize)
	y := int(point.Y / cellSize)
	grid[x][y] = point
}

func isValidPoint(grid [][]*Point, cols, rows int, point *Point, cellSize, width, height, radius float64) bool {
	if point.X >= width || point.X < 0 || point.Y >= height || point.Y < 0 {
		return false
	}

	x := int(math.Floor(point.X / cellSize))
	y := int(math.Floor(point.Y / cellSize))
	i0 := blmath.Max(x-2, 0)
	i1 := blmath.Min(x+2, cols-1)
	j0 := blmath.Max(y-2, 0)
	j1 := blmath.Min(y+2, rows-1)

	for i := i0; i <= i1; i++ {
		for j := j0; j <= j1; j++ {
			if grid[i][j] != nil {
				dist := grid[i][j].Distance(point)
				if dist < radius {
					return false
				}
			}
		}
	}
	return true
}
