package library

import "testing"

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
