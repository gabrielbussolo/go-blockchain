package node

import "net/url"

type Nodes struct {
	nodes map[string]bool
}

func New() *Nodes {
	return &Nodes{nodes: make(map[string]bool)}
}

func (n *Nodes) Add(address string) {
	parse, err := url.Parse(address)
	if err != nil {
		panic("problem to parse the url")
	}
	host := parse.Host
	n.nodes[host] = true
}
