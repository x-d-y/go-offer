package main

func main() {
	buildNode([][]int{{7, 5}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
}

func buildNode(nums [][]int) *Node {
	nodes := make([]*Node, 0)

	for _, v := range nums {
		node := new(Node)
		node.Val = v[0]
		nodes = append(nodes, node)
	}

	nodes = append(nodes, nil)

	for i, v := range nums {
		nodes[i].Next = nodes[i+1]
		nodes[i].Random = nodes[v[1]]
	}

	a := copyRandomList(nodes[0])
	return a
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil{
		return nil
	}
	m := make(map[*Node]*Node)
	n := new(Node)
	res := n
	m[head] = n
	for head != nil {
		n.Val = head.Val
		if _, ok := m[head.Next]; !ok {
			var next *Node
			if head.Next != nil {
				next = new(Node)
			}
			m[head.Next] = next
		}
		n.Next = m[head.Next]
		if _, ok := m[head.Random]; !ok {
			var random *Node
			if head.Random != nil {
				random = new(Node)
			}
			m[head.Random] = random
		}
		n.Random = m[head.Random]
		n = n.Next
		head = head.Next
	}
	return res
}
