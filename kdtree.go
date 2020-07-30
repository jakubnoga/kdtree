package kdtree

import (
	"encoding/json"
	"sort"
)

type KdTree struct {
	Left, Right *KdTree
	Dim         int
	Point       []uint32
}

func Create(points [][]uint32, depth int) *KdTree {
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

func (tree *KdTree) NearestNeighbour(point []uint32) *KdTree {
	var dist, bestNorm uint32
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
	treeNorm := tree.norm(point)

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

func (tree *KdTree) distance(point []uint32, dim int) uint32 {
	x1 := tree.Point[dim]
	x2 := point[dim]

	d := x1 - x2
	if d >= 0 {
		return d 
	} 
	
	return -d	
}

func (tree *KdTree) norm(point []uint32) uint32 {
	var sum uint32 = 0
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
