package node

import "net/url"

type Nodes struct {
	node map[string]bool
}

func New() *Nodes {
	return &Nodes{node: make(map[string]bool)}
}

func (n *Nodes) Add(address string) {
	parse, err := url.Parse(address)
	if err != nil {
		panic("problem to parse the url")
	}
	host := parse.Host
	n.node[host] = true
}

func (n *Nodes) GetAll() []string {
	keys := make([]string, len(n.node))
	i := 0
	for k := range n.node {
		keys[i] = k
		i++
	}
	return keys
}
