package library

import (
	"fmt"
	"testing"

	"gw/backend"
	"gw/util"

	"go.mongodb.org/mongo-driver/bson"
)

var id = "4ab01a89-4509-40cb-8afb-34ab8552ae88"

func TestAdd(t *testing.T) {
	now := util.GetTime()
	b := backend.MongoInfo{
		Id:         id,
		Name:       "Test_one",
		Path:       "/add/test",
		To:         "http://baidu.com/1/test,http://baidu.com/2/test,http://baidu.com/3/test",
		Method:     "GET",
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
	b := &backend.MongoInfo{}
	if err := FindOne("wg", bson.M{"path": "/add/test"}, b); err != nil {
		t.Fatalf("MongoDB Test FindOne Was Wrong Err Was %s", err)
	}

	fmt.Println(b)

	if b.Name != "Test_one" {
		t.Fatalf("Mongo FindOne Was Wrong Name Want Test_one now %s", b.Name)
	}
}

func TestFindAll(t *testing.T) {
	cux, err := FindAll("wg")
	if err != nil {
		t.Fatalf("MongoDB Test FindAll Was Wrong Err Was %s", err)
	}

	fmt.Println(cux[0])
	if cux[0].Id != "4ab01a89-4509-40cb-8afb-34ab8552ae88" {
		t.Fatalf("Mongo FindAll Was Wrong Id Want 4ab01a89-4509-40cb-8afb-34ab8552ae88 %s", cux[0].Id)
	}
}

func TestDel(t *testing.T) {
	del := bson.M{"id": id}
	if err := Del("wg", del); err != nil {
		t.Fatal("MongoDB Del Was Err")
	}
}
