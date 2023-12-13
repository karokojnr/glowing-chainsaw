package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Millisecond)
	if c.cache == nil {
		t.Error("Expected cache to be initialized")
	}
}

func TestAddToCache(t *testing.T) {
	c := NewCache(time.Millisecond)

	cases := []struct {
		inputKey   string
		inputValue []byte
	}{
		{
			inputKey:   "key1",
			inputValue: []byte("value1"),
		},
		{
			inputKey:   "key2",
			inputValue: []byte("value2"),
		},
	}

	for _, cs := range cases {
		c.Add(cs.inputKey, cs.inputValue)
		actual, ok := c.Get(cs.inputKey)
		if !ok {
			t.Errorf("Expected key %s to be in cache", cs.inputKey)
			continue
		}
		if string(actual) != string(cs.inputValue) {
			t.Errorf("Expected value1, got %s", string(actual))
			continue
		}

	}

}

func TestReap(t *testing.T) {
	inverval := time.Millisecond * 10
	c := NewCache(inverval)

	keyOne := "key1"
	c.Add(keyOne, []byte("value1"))
	time.Sleep(inverval + time.Millisecond)

	_, ok := c.Get(keyOne)
	if ok {
		t.Errorf("Expected key %s to be reaped", keyOne)
	}

}

func TestReapFail(t *testing.T) {
	inverval := time.Millisecond * 10
	c := NewCache(inverval)

	keyOne := "key1"
	c.Add(keyOne, []byte("value1"))
	time.Sleep(inverval / 2)

	_, ok := c.Get(keyOne)
	if !ok {
		t.Errorf("Expected key %s to not be reaped", keyOne)
	}

}
