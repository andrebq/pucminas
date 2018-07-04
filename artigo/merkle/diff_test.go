package merkle_test

import (
	"reflect"
	"testing"

	"github.com/andrebq/pucminas/artigo/merkle"
)

func TestDiff(t *testing.T) {
	hello := merkle.NewNode().Add([]byte("hello"), []byte("hello")).Add([]byte("same"), []byte("same"))
	olleh := merkle.NewNode().Add([]byte("hello"), []byte("olleh")).Add([]byte("same"), []byte("same"))

	var times int
	err := merkle.Diff(hello, olleh, func(k merkle.Bytes, a, b merkle.Bytes) (bool, error) {
		times++
		if !reflect.DeepEqual(a.Bytes(nil), []byte("hello")) {
			t.Errorf("value of a should be hello got %v", a.String())
		}
		if !reflect.DeepEqual(b.Bytes(nil), []byte("olleh")) {
			t.Errorf("value of b should be olleh got %v", a.String())
		}

		if k.String() != "hello" {
			t.Errorf("key should be hello got %v", k.String())
		}

		return true, nil
	})

	if times != 1 {
		t.Fatal("should call only once")
	}

	if err != nil {
		t.Fatal(err)
	}
}
