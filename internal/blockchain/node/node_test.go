package node

import (
	"testing"
)

func TestNodes_Add(t *testing.T) {
	expected := "localhost:8080"
	n := New()
	n.Add("http://localhost:8080/")
	if n.nodes[expected] != true {
		t.Errorf("the addres wasnt added to the node list")
	}
}
