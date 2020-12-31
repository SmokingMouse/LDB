package buffer_manager

import (
	"testing"
)

func Equal(t *testing.T, v1 LruValue, v2 LruValue) {
	if v1 != v2 {
		t.Fatalf("Expected %v,but get %v", v1, v2)
	}
}

func TestBasicSetAndGet(t *testing.T) {
	lru := NewLru(2)
	lru.Set(20, 20)
	if v := lru.Get(20); v != 20 {
		t.Fatalf("should get %d,but get %d", 20, v)
	}
	lru.Set(30, 30)
	if lenth := lru.li.lenth; lenth != 2 {
		t.Fatalf("the len should be %d,but get %d", 2, lenth)
	}
}

func TestOut(t *testing.T) {
	lru := NewLru(1)
	lru.Set(20, 20)
	lru.Set(30, 30)
	if v := lru.Get(20); v != nil {
		t.Fatalf("should be nil,but get %v", v)
	}
	if v := lru.Get(30); v != 30 {
		t.Fatalf("should get 30,but get %d", v)
	}
}

func TestUpdate(t *testing.T) {
	lru := NewLru(3)
	lru.Set(20, 20)
	Equal(t, lru.Get(20), 20)
	lru.Set(20, 30)
	Equal(t, lru.Get(20), 30)
}

func TestLatest(t *testing.T) {
	lru := NewLru(2)
	lru.Set(20, 20)
	lru.Set(30, 30)
	lru.Get(20)
	lru.Set(40, 40)
	Equal(t, lru.Get(30), nil)
	Equal(t, lru.Get(20), 20)
	lru.Set(50, 50)
	Equal(t, lru.Get(20), 20)
	Equal(t, lru.Get(40), nil)
}
