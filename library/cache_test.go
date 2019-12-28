package library

import (
	"testing"
)

// cache 测试
func TestSet(t *testing.T) {
	if err := SetCache("test_key", "test_value", 10); err != nil {
		t.Fatalf("Rds Set Was Wrong Err : %s", err)
	}
}

func TestGetCache(t *testing.T) {
	v := GetCache("test_key")

	if v != "test_value" {
		t.Fatalf("Rds Get Val Was Wrong Want test_value Now : %s", v)
	}
}

func TestHGet(t *testing.T) {
	HSet("test_hset", "a", 1)
	HSet("test_hset", "b", 2)
	HSet("test_hset", "c", 3)

	if b := HGet("test_hset", "b"); b != "2" {
		t.Fatalf("Rds HGet test b want 2 now %s", b)
	}

	if c := HGet("test_hst_2", "c"); c != "" {
		t.Fatalf("Rds HGet not want null now %s", c)
	}
}
