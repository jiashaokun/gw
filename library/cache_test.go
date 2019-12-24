package library

import "testing"

// cache 测试
func TestSet(t *testing.T) {
	if err := Set("test_key", "test_value", 10); err != nil {
		t.Fatalf("Rds Set Was Wrong Err : %s", err)
	}
}

func TestGet(t *testing.T) {
	v, err := Get("test_key")
	if err != nil {
		t.Fatalf("Rds Get Was Wrong Err : %s", err)
	}

	if v != "test_value" {
		t.Fatalf("Rds Get Val Was Wrong Want test_value Now : %s", v)
	}
}
