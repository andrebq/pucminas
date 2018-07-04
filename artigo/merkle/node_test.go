package merkle_test

import (
	"reflect"
	"testing"

	"github.com/andrebq/pucminas/artigo/merkle"
)

func TestNodeHash(t *testing.T) {
	node := merkle.NewNode()

	withValue := node.Add([]byte("hello"), []byte("world"))

	if withValue == nil {
		t.Fatal("withValue shouldn't be nil")
	}

	value, ok := withValue.Get([]byte("hello"))
	if !ok {
		t.Fatal("value should exist")
	}
	if !reflect.DeepEqual(value.Bytes(nil), []byte("world")) {
		t.Fatal("invalid value, should be world")
	}

	_, ok = node.Get([]byte("hello"))
	if ok {
		t.Fatal("node should be immutable")
	}

	if withValue.Hash() == node.Hash() {
		t.Fatal("something really bad happened")
	}
}

func TestNodeWalk(t *testing.T) {
	node := merkle.NewNode()
	withValue := node.Add([]byte("hello"), []byte("world"))

	paths := make(map[string]*merkle.Node)

	withValue.Walk(func(p merkle.Bytes, n *merkle.Node) (bool, error) {
		paths[p.String()] = n
		return true, nil
	})

	for k, v := range paths {
		t.Logf("%v: %v", k, v.Export())
	}

	if len(paths) != 6 {
		t.Fatal("should have 5 items afther the final scan")
	}
}
