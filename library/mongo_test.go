package library

import (
	"gw/util"
	"testing"

	"gw/backend/admin"

	"go.mongodb.org/mongo-driver/bson"
)

func TestAdd(t *testing.T) {
	now := util.GetTime()
	b := admin.Add{
		Id:         "4ab01a89-4509-40cb-8afb-34ab8552ae88",
		Name:       "Test_one",
		Path:       "/add/test",
		To:         "http://baidu.com/add/test",
		Dns:        0,
		CacheTime:  0,
		Timeout:    0,
		Decay:      0,
		DecayTime:  0,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := Add("wg", b); err != nil {
		t.Fatalf("MongoDB Test Add Was Wrong Err Was %s", err)
	}
}

func TestFindOne(t *testing.T) {
	b := &admin.Add{}
	id := "4ab01a89-4509-40cb-8afb-34ab8552ae88"
	if err := FindOne("wg", bson.M{"id": id}, b); err != nil {
		t.Fatalf("MongoDB Test FindOne Was Wrong Err Was %s", err)
	}

	if b.Name != "Test_one" {
		t.Fatalf("Mongo FindOne Was Wrong Name Want Test_one now %s", b.Name)
	}
}
