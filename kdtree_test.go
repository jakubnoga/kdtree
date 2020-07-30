package kdtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNearestNeighbour(t *testing.T) {
	tree := Create([][]uint32{{2, 3}, {5, 4}, {9, 6}, {4, 7}, {8, 1}, {7, 2}}, 0)

	cases := [][]uint32{
		{10, 10}, {0, 4}, {3, 0}, {0, 10}, {6, 4},
	}
	expected := [][]uint32{
		{9, 6}, {2, 3}, {2, 3}, {4, 7}, {5, 4},
	}

	for idx, val := range cases {
		t.Run(fmt.Sprintf("%v", val), testCaseRunnerProvider(tree, val, expected[idx]))
	}
}

func testCaseRunnerProvider(tree *KdTree, testCase []uint32, expected []uint32) func(t *testing.T) {
	return func(t *testing.T) {
		best := tree.NearestNeighbour(testCase)
		if !reflect.DeepEqual(best.Point, expected) {
			t.Logf("Expected %v but got %v", expected, best.Point)
			t.Fail()
		}
	}
}

func TestKdTree_NearestNeighbour(t *testing.T) {
	type args struct {
		point []uint32
	}
	tests := []struct {
		name string
		tree *KdTree
		args args
		want *KdTree
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.NearestNeighbour(tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KdTree.NearestNeighbour() = %v, want %v", got, tt.want)
			}
		})
	}
}
