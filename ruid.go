package ruid

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

// Ruid: a really unique id, Very fast to generate, opaque identifier.
// Huid: a really unique id, very fast to generate, decodable to be human readable.

const RuidVer = 1

type RuidGen struct {
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
// The uniqueLocation parameter in the NewRuidGen() call
// should be as unique as possible.
//
// RuidGen has methods Ruid() and Huid() to generate
// Ruid and Huid respectively.
//
func NewRuidGen(uniqueLocation string) *RuidGen {

	r := &RuidGen{}
	uniqLoc := sha1.Sum([]byte(uniqueLocation))
	r.pid = int64(os.Getpid())
	r.base64uniqLoc = base64.URLEncoding.EncodeToString(uniqLoc[:])

	return r
}

func (r *RuidGen) Huid() string {
	tm := time.Now()
	r.counter++

	huid := fmt.Sprintf("|tm:%s|pid:%010d|loc:%s|seq:%020d|",
		tm.Format(time.RFC3339Nano),
		r.pid,
		r.base64uniqLoc,
		r.counter)

	return fmt.Sprintf("huid_v%02d_%s", RuidVer, base64.URLEncoding.EncodeToString([]byte(huid)))
}

// A Ruid applies a sha1sum to a Huid.
func (r *RuidGen) Ruid() string {

	tm := time.Now()
	r.counter++

	huid := fmt.Sprintf("|tm:%s|pid:%010d|loc:%s|seq:%020d|",
		tm.Format(time.RFC3339Nano),
		r.pid,
		r.base64uniqLoc,
		r.counter)

	res := sha1.Sum([]byte(huid))
	ruid := fmt.Sprintf("ruid_v%02d_%s", RuidVer, base64.URLEncoding.EncodeToString(res[:]))

	return ruid
}

// Ruid2 adds a random number from /dev/urandom and uses a Sha512 hash instead of Sha1.
func (r *RuidGen) Ruid2() string {

	randomBytes := getRandomBytes()

	tm := time.Now()
	r.counter++

	huid := fmt.Sprintf("|tm:%s|pid:%010d|loc:%s|seq:%020d|random:%x",
		tm.Format(time.RFC3339Nano),
		r.pid,
		r.base64uniqLoc,
		r.counter,
		randomBytes)

	//res := sha1.Sum([]byte(huid))
	res := sha512.Sum512([]byte(huid))
	ruid := fmt.Sprintf("ruid_v%02d_%s", 2, base64.URLEncoding.EncodeToString(res[:]))

	return ruid
}

func getRandomBytes() []byte {
	c := 100
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
