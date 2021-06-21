package tree

import (
	"fmt"
	"sort"
)

// Record represents 1 web forum post
type Record struct {
	ID     int
	Parent int
}

// Node ties a post to its child posts
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree from root post through all children
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	tree := make([]Node, len(records))

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, v := range records {
		if i != v.ID || v.ID < v.Parent || (v.ID > 0 && v.ID == v.Parent) {
			return nil, fmt.Errorf("invalid list of records for %v", v)
		}
		tree[i] = Node{ID: i}
		if i > 0 {
			tree[v.Parent].Children = append(tree[v.Parent].Children, &tree[i])
		}

	}
	return &tree[0], nil
}
