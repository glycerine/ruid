package ruid

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// Ruid: a really unique id, very fast to generate, opaque identifier.
// Huid: a really unique id, very fast to generate, decodable to be human readable.
// Ruid2: mostly random, opaque and unguessable really-unique id.
// Tuid: a transparent id, showing what goes into a huid before the reversible base64 encode.
// Ruid3: mostly random, opaque and unguessable really-unique id. base36 encoded using characters [a-z0-9].

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

	return fmt.Sprintf("huid_v%02d_%s", 1, base64.URLEncoding.EncodeToString([]byte(huid)))
}

// transparent version
func (r *RuidGen) Tuid() string {
	tm := time.Now()
	r.counter++

	huid := fmt.Sprintf("|tm:%s|pid:%010d|loc:%s|seq:%020d|",
		tm.Format(time.RFC3339Nano),
		r.pid,
		r.base64uniqLoc,
		r.counter)

	return fmt.Sprintf("huid_v%02d_%s", 1, huid)
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
	ruid := fmt.Sprintf("ruid_v%02d_%s", 1, base64.URLEncoding.EncodeToString(res[:]))

	return ruid
}

// Ruid2 adds randomness from /dev/urandom and uses a Sha512 hash instead of Sha1.
// It is really opaque.
func (r *RuidGen) Ruid2() string {

	huid2 := r.huid2core()

	res := sha512.Sum512([]byte(huid2))
	ruid2 := fmt.Sprintf("ruid_v%02d_%s", 2, base64.URLEncoding.EncodeToString(res[:]))

	return ruid2
}

func getRandomBytes(c int) []byte {
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func (r *RuidGen) huid2core() string {
	// generate random bytes
	randomBytes := getRandomBytes(400)

	// use between 50 - 100 of them. At either end. And in the middle.
	w1 := binary.LittleEndian.Uint64(randomBytes[:8]) % 50
	w2 := binary.LittleEndian.Uint64(randomBytes[8:16]) % 50
	w3 := binary.LittleEndian.Uint64(randomBytes[16:24]) % 50
	randomBytes = randomBytes[24:]

	ran1 := randomBytes[0 : 50+w1]
	ran2 := randomBytes[100 : 150+w2]
	ran3 := randomBytes[200 : 250+w3]

	tm := time.Now()
	r.counter++

	huid2 := fmt.Sprintf("%s|tm:%s|pid:%010d|%s|loc:%s|seq:%d|%s",
		ran1,
		tm.Format(time.RFC3339Nano),
		r.pid,
		ran2,
		r.base64uniqLoc,
		r.counter,
		ran3,
	)
	return huid2
}

// Ruid3 is a Ruid2 whose suffix is base36 encoded using only the letters a-z and the
// numbers 0-9, and ensure it can be used in places that don't allow characters
// like '-' or '=' in identifiers.
func (r *RuidGen) Ruid3() string {

	huid2 := r.huid2core()

	res := sha512.Sum512([]byte(huid2))
	b := EncodeBytesBase36(res[:])

	ruid3 := fmt.Sprintf("ruid_v%02d_%s", 3, string(b))
	return ruid3
}
