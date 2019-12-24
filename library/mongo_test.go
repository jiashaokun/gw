package library

import (
	"gw/util"
	"testing"

	"gw/backend/admin"

	"go.mongodb.org/mongo-driver/bson"
)

var id = "4ab01a89-4509-40cb-8afb-34ab8552ae88"

func TestAdd(t *testing.T) {
	now := util.GetTime()
	b := admin.Add{
		Id:         id,
		Name:       "Test_one",
		Path:       "/add/test",
		To:         "http://baidu.com/1/test,http://baidu.com/2/test,http://baidu.com/3/test",
		Dns:        1,
		CacheTime:  200,
		Timeout:    20,
		Decay:      1,
		DecayTime:  50,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := Add("wg", b); err != nil {
		t.Fatalf("MongoDB Test Add Was Wrong Err Was %s", err)
	}
}

func TestFindOne(t *testing.T) {
	b := &admin.Add{}
	if err := FindOne("wg", bson.M{"id": id}, b); err != nil {
		t.Fatalf("MongoDB Test FindOne Was Wrong Err Was %s", err)
	}

	if b.Name != "Test_one" {
		t.Fatalf("Mongo FindOne Was Wrong Name Want Test_one now %s", b.Name)
	}
}

func TestDel(t *testing.T) {
	del := bson.M{"id": id}
	if err := Del("wg", del); err != nil {
		t.Fatal("MongoDB Del Was Err")
	}
}
