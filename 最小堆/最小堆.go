package main

import "fmt"

type Node struct {
	Value int
}

func Init(nodes []Node) {
	for i := len(nodes)/2 - 1; i >= 0; i -- 
	{
		down(nodes,i,len(nodes))
	}
}

func down(nodes []Node, i, n int) {
	parent := i
	left_child := 2 * i + 1
	right_child := 2 * i + 2
	temp := nodes[parent].Value
	for {
		if left_child < n {
			if right_child < n {
				if nodes[left_child].Value < nodes[right_child].Value {
					if nodes[left_child].Value < temp {
						nodes[parent].Value = nodes[left_child].Value
						nodes[left_child].Value = temp
						parent = left_child
						left_child = left_child*2+1
						right_child = left_child*2+2
					} else {
						break
					}
				} else {
					if nodes[right_child].Value < temp {
						nodes[parent].Value = nodes[right_child].Value
						nodes[right_child].Value = temp
						parent = right_child
						left_child = right_child*2+1
						right_child = right_child*2+2
					} else {
						break
					}
				}
			} else {
				if nodes[left_child].Value < temp {
					nodes[parent].Value = nodes[left_child].Value
					nodes[left_child].Value = temp
					parent = left_child
					left_child = left_child*2+1
					right_child = left_child*2+2
				} else {
					break
				}
			}
		} else {
			break
		}
	}
}
func up(nodes []Node, j int) {
	child := j
	parent := (j-1)/2
	for{
		if child == 0 {
			break
		}
		if nodes[parent].Value <= nodes[child].Value {
			break
		}
		temp := nodes[child].Value
		nodes[child].Value = nodes[parent].Value
		nodes[parent].Value = temp
		child = parent
		parent = (parent-1)/2
	}
}
func Pop(nodes []Node) (Node, []Node) {
	min := nodes[0]
	nodes = nodes[1:len(nodes)]
	Init(nodes)
	return min, nodes
}

func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes,node)
	up(nodes,len(nodes)-1)
	return nodes
}
func Remove(nodes []Node, node Node) []Node {
	for i := 0; i < len(nodes); i ++ {
		if nodes[i].Value == node.Value {
			nodes[i].Value = nodes[len(nodes)-1].Value
			nodes = nodes[0:len(nodes)-1]
			Init(nodes)
			break
		}
	}
	return nodes
}

func main() {
	nodes := []Node {
		Node{3},
		Node{13},
		Node{17},
		Node{2},
		Node{5},
		Node{8},
		Node{22},
		Node{1},
		Node{65},
		Node{50},
		Node{77},
		Node{99},
	}
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	Init(nodes)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	push_node := Node{0}
	nodes = Push(push_node,nodes)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	node_min,nodes := Pop(nodes)
	fmt.Printf("%d\n",node_min.Value)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	remove_node := Node{5}
	nodes = Remove(nodes,remove_node)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
}