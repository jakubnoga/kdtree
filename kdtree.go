package kdtree

import (
	"encoding/json"
	"log"
	"math"
	"sort"
	"time"
)

type KdTree struct {
	Left, Right *KdTree
	Dim         int
	Point       []float64
}

func Create(points [][]float64, depth int) *KdTree {
	tree := new(KdTree)
	tree.Dim = depth % len(points[0])

	sort.SliceStable(points, func(i, j int) bool {
		return points[i][tree.Dim] < points[j][tree.Dim]
	})

	switch len(points) {
	case 1:
		tree.Point = points[0]
	case 2:
		tree.Point = points[1]
		tree.Left = Create(points[:1], tree.Dim+1)
	default:
		medianPoint := len(points) / 2
		tree.Left = Create(points[:medianPoint], tree.Dim+1)
		tree.Right = Create(points[medianPoint+1:], tree.Dim+1)
		tree.Point = points[medianPoint]
	}

	return tree
}

func (tree *KdTree) NearestNeighbour(point []float64) *KdTree {
	// defer duration(track("nn"))
	var dist float64
	var candidate, best, other *KdTree

	if tree.Left == nil && tree.Right == nil {
		return tree
	} else if tree.Right == nil {
		candidate = tree.Left.NearestNeighbour(point)
	} else {
		if tree.Point[tree.Dim] > point[tree.Dim] {
			other = tree.Right
			candidate = tree.Left.NearestNeighbour(point)
		} else {
			other = tree.Left
			candidate = tree.Right.NearestNeighbour(point)
		}
	}

	candidateNorm := candidate.norm(point)
	if candidateNorm == 0 {
		return candidate
	}

	treeNorm := tree.norm(point)
	if treeNorm == 0 {
		return tree
	}

	var bestNorm float64

	if candidateNorm < treeNorm {
		best = candidate
		bestNorm = candidateNorm
	} else {
		best = tree
		bestNorm = treeNorm
	}

	if other != nil {
		dist = tree.distance(point, tree.Dim)

		if bestNorm > dist {
			candidate = other.NearestNeighbour(point)

			if candidate.norm(point) < bestNorm {
				best = candidate
			}
		}
	}

	return best
}

func (tree *KdTree) distance(point []float64, dim int) float64 {
	x1 := tree.Point[dim]
	x2 := point[dim]

	return math.Abs(x1 - x2)
}

func (tree *KdTree) norm(point []float64) float64 {
	sum := 0.0
	for idx, val := range point {
		d := val - tree.Point[idx]
		sum += d * d
	}

	return sum
}

func (tree *KdTree) ToJson() string {
	marshal, err := json.Marshal(tree)
	if err != nil {
		return ""
	}

	return string(marshal)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start).Microseconds())
}
