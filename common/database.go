package common

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var DB *mgo.Collection
func MongoInit(biaoming string) *mgo.Collection {
	dailInfo := &mgo.DialInfo{
		Addrs:          []string{"134.175.182.249"},
		Direct:         false,
		Timeout:        0,
		FailFast:       false,
		Database:       "admin",
		ReplicaSetName: "",
		Source:         "",
		Service:        "",
		ServiceHost:    "",
		Mechanism:      "SCRAM-SHA-1",
		Username:       "jshangpin",
		Password:       "jshangpin",
		PoolLimit:      0,
		DialServer:     nil,
		Dial:           nil,
	}
	session, err := mgo.DialWithInfo(dailInfo)
	if err != nil {
		fmt.Printf("mgo dail error[%s]\n", err.Error())
		err_handler(err)
	}
	//defer session.Close()
	// set mode
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mkall").C(biaoming)
	return c
}

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}