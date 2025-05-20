package s0133_clone_graph

// #Medium #Hash_Table #Depth_First_Search #Breadth_First_Search #Graph #Udemy_Graph
// #Top_Interview_150_Graph_General #2025_05_20_Time_0_ms_(100.00%)_Space_4.84_MB_(17.41%)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {
	return cloneGraphHelper(node, make(map[*Node]*Node))
}

func cloneGraphHelper(node *Node, processedNodes map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if processedNode, exists := processedNodes[node]; exists {
		return processedNode
	}
	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	processedNodes[node] = newNode
	for _, neighbor := range node.Neighbors {
		clonedNeighbor := cloneGraphHelper(neighbor, processedNodes)
		if clonedNeighbor != nil {
			newNode.Neighbors = append(newNode.Neighbors, clonedNeighbor)
		}
	}
	return newNode
}
