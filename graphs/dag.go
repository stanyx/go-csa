package graphs

import "fmt"

type Link struct {
	Source *Node
	Target *Node
	Cost   int
}

func (l *Link) String() string {
	return fmt.Sprintf("(%d) - %s", l.Cost, l.Target.Name)
}

type Node struct {
	Name   string
	Parent *Node
	Links  []*Link
}

func (n *Node) String() string {
	t := ""

	for _, l := range n.Links {
		t += fmt.Sprintf("[ %s ]", l) + ", "
	}

	return t
}

type Graph struct {
	nodes map[string]*Node
}

func (gr Graph) String() string {
	t := ""
	for _, n := range gr.nodes {
		t += "\n" + fmt.Sprintf("(%s: %s)", n.Name, n)
	}

	return t
}

func (gr Graph) AddNode(name string) {
	gr.nodes[name] = &Node{
		Name:  name,
		Links: make([]*Link, 0),
	}
}

func (gr Graph) AddLink(sourceName, targetName string, cost int) error {
	source, ok := gr.nodes[sourceName]
	if !ok {
		return fmt.Errorf("node with name %s not found", sourceName)
	}

	target, ok := gr.nodes[targetName]
	if !ok {
		return fmt.Errorf("node with name %s not found", targetName)
	}

	link := &Link{
		Source: source,
		Target: target,
		Cost:   cost,
	}

	source.Links = append(source.Links, link)

	return nil
}

func (gr Graph) MinLink(nodes []*Node, costs map[string]int64) *Node {
	//TODO - выбор правильного минимального стартового значения
	minCost := int64(1000000)
	minNode := nodes[0]
	for _, n := range nodes {
		nodeCost := costs[n.Name]
		if nodeCost < minCost {
			minCost = nodeCost
			minNode = n
		}
	}
	return minNode
}

func (gr Graph) Path(source string, target string) []string {

	costs := map[string]int64{}
	parents := map[string]string{}
	searched := map[string]bool{}
	nodes := []*Node{}

	for _, n := range gr.nodes {
		if n.Name != source {
			nodes = append(nodes, n)
		}
		for _, l := range n.Links {
			if l.Source.Name == source {
				costs[l.Target.Name] = int64(l.Cost)
				parents[l.Target.Name] = l.Source.Name
			}
		}
	}

	for len(nodes) > 0 {

		nextNode := gr.MinLink(nodes, costs)
		nextNodeCost := costs[nextNode.Name]

		for _, l := range nextNode.Links {
			sumCost, ok := costs[l.Target.Name]
			curCost := nextNodeCost + int64(l.Cost)
			if !ok {
				costs[l.Target.Name] = int64(curCost)
				continue
			}
			if curCost < sumCost {
				costs[l.Target.Name] = curCost
				parents[l.Target.Name] = nextNode.Name
			}
		}

		searched[nextNode.Name] = true
		newNodes := []*Node{}
		for _, n := range nodes {
			visited := searched[n.Name]
			if nextNode.Name != n.Name && !visited {
				newNodes = append(newNodes, n)
			}
		}

		nodes = newNodes
	}
	n := target
	path := []string{n}
	for n != source {
		n = parents[n]
		path = append([]string{n}, path...)
	}
	return path
}

func New() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}
