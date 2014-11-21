package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

// Ruid: a really unique id, Very fast to generate, opaque identifier.
// Huid: a really unique id, very fast to generate, decodable to be human readable.

const RuidVer = 1

func main() {
	myExternalIP := "my example location: 10.0.0.1"
	ruidGen := NewRuidGen(myExternalIP)
	fmt.Printf("%s\n", string(ruidGen.Ruid()))
	//fmt.Printf("\n\n sha1 ruid = '%s', len = %d\n", r.Sha1Ruid(), len(r.Sha1Ruid()))
}

type RuidGen struct {
	origLoc       string
	uniqLoc       [20]byte
	base64uniqLoc string
	counter       int64
	pid           int64
}

//
// NewRuidGen(): uniqueLocation should be a byte
// sequence that is unique to this specific physical location.
// Suggestions: a hardware
// mac address, your external ip address, the traceroute out
// a known distant location on the public internet.
// The uniqueLocation should be as unique as possible.
// It will be crypto hashed down to 20 bytes.
//
func NewRuidGen(uniqueLocation string) *RuidGen {

	r := &RuidGen{}
	r.uniqLoc = sha1.Sum([]byte(uniqueLocation))
	r.pid = int64(os.Getpid())
	r.base64uniqLoc = base64.URLEncoding.EncodeToString(r.uniqLoc[:])

	return r
}

func (r *RuidGen) Huid() string {
	huid := r.huidBase()
	return fmt.Sprintf("huid_v%02d_%s", RuidVer, base64.URLEncoding.EncodeToString([]byte(huid)))
}

func (r *RuidGen) huidBase() string {

	tm := time.Now()

	r.counter++

	return fmt.Sprintf("|tm:%s|pid:%010d|loc:%s|seq:%020d|",
		tm.Format(time.RFC3339Nano),
		r.pid,
		r.base64uniqLoc,
		r.counter)
}

// A Ruid applies a sha1sum to a Huid.
func (r *RuidGen) Ruid() string {

	res := sha1.Sum([]byte(r.huidBase()))
	ruid := fmt.Sprintf("ruid_v%02d_%s", RuidVer, base64.URLEncoding.EncodeToString(res[:]))

	return ruid
}
