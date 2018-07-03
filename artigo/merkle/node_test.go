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
	if !reflect.DeepEqual(value, []byte("world")) {
		t.Fatal("invalid value, should be world")
	}

	_, ok = node.Get([]byte("hello"))
	if ok {
		t.Fatal("node should be immutable")
	}

	if withValue.Hash() == node.Hash() {
		t.Fatal("something really bad happened")
	}

	println("withValue: ", withValue.Export().String())
}
