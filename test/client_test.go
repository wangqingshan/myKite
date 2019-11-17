package test

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/koding/kite"
	"testing"
)

func TestGetDocText(t *testing.T) {

	k := kite.New("exp2", "1.0.0")

	// Connect to our math kite
	mathWorker := k.NewClient("http://134.175.80.121:3636/kite")
	mathWorker.Dial()

	response, _ := mathWorker.Tell("square", 4) // call "square" method with argument 4
	fmt.Println("result:", response.MustFloat64())
}

/*@Test
public void clearOne () {
String userId = "200000024626" ;
String userIden = "qf3EAAAgEA2N-fHN7" ;

String keyB = "5001:B:" + userId ;
String keyT = "5001:T:" + userId ;
String keyC = "5001:C:" + userIden ;
memcachedRunner.getClient().delete(keyB) ;
memcachedRunner.getClient().delete(keyT) ;
memcachedRunner.getClient().delete(keyC) ;
}*/
var (
	server  = "10.230.4.136:12010"
	prekey  = "5001:B:300001089336"
	prekeyT = "5001:T:300001089336"
	//prekeyIden ="5001:C:fRRxfQEcFX0xTVBRA"
	Iden = "NfEACAwQFXEIDDQV7"
)

func TestHDMem(t *testing.T) {
	//create a handle
	mc := memcache.New(server)
	if mc == nil {
		fmt.Println("memcache New failed")
	}
	mc.Add(&memcache.Item{Key: "ceshi", Value: []byte("bluegogo")})

	it2, _ := mc.Get("ceshi")
	fmt.Println(string(it2.Value))
	mc.Delete(prekeyT)
	//mc.Delete(prekeyIden)
	//get key's value

	/*if string(it.Key) == "foo" {
		fmt.Println("value is ", string(it.Value))
	} else {
		fmt.Println("Get failed")
	}*/
}
func TestMem(t *testing.T) {
	//create a handle
	mc := memcache.New(server)
	if mc == nil {
		fmt.Println("memcache New failed")
	}

	//set key-value
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	//get key's value
	it, _ := mc.Get("foo")
	if string(it.Key) == "foo" {
		fmt.Println("value is ", string(it.Value))
	} else {
		fmt.Println("Get failed")
	}
	///Add a new key-value
	mc.Add(&memcache.Item{Key: "foo", Value: []byte("bluegogo")})
	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println("Add failed")
	} else {
		if string(it.Key) == "foo" {
			fmt.Println("Add value is ", string(it.Value))
		} else {
			fmt.Println("Get failed")
		}
	}
	//replace a key's value
	mc.Replace(&memcache.Item{Key: "foo", Value: []byte("mobike")})
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("Replace failed")
	} else {
		if string(it.Key) == "foo" {
			fmt.Println("Replace value is ", string(it.Value))
		} else {
			fmt.Println("Replace failed")
		}
	}
	//delete an exist key
	err = mc.Delete("foo")
	if err != nil {
		fmt.Println("Delete failed:", err.Error())
	}
	//incrby
	err = mc.Set(&memcache.Item{Key: "aaa", Value: []byte("1")})
	if err != nil {
		fmt.Println("Set failed :", err.Error())
	}
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("Get failed ", err.Error())
	} else {
		fmt.Println("src value is:", it.Value)
	}
	value, err := mc.Increment("aaa", 7)
	if err != nil {
		fmt.Println("Increment failed")
	} else {
		fmt.Println("after increment the value is :", value)
	}
	//decrby
	value, err = mc.Decrement("aaa", 4)
	if err != nil {
		fmt.Println("Decrement failed", err.Error())
	} else {
		fmt.Println("after decrement the value is ", value)
	}

}
