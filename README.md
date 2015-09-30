Ruid: a really unique id
================

There are six identifiers defined here: Ruid, Ruid2, Ruid3, Huid, Tuid, and Tuid64.


A Huid is a really unique id. It is very fast to generate, and is base64-decodable to be human readable.

A Ruid is a really unique id. It is very fast to generate, and is an opaque identifier. It is the SHA1 hash of a Huid.

A Ruid2 is a really unique id, that starts with a Huid and then adds 100 bytes from /dev/unrandom and then SHA512 hashes it.

A Ruid3 is the same as a Ruid2, but encoded using only the characters a-z0-9, hence in base 36. This is easier to convey over the phone or through channels that don't like '-' or '='. It is always 109 bytes long, including the 9 bytes prefix 'ruid_v03_'. Example: 'ruid_v03_0zgo9tsz3r505khm87w7cgebdt3nrkbg773zvnyu32i14z4vkokqkr2y9jn34zeyxtj4tecoczgl2i1rmcs7yfngybp8zsybqbni'. The 100 bytes are a consequence of base36 encoding a 512 bit SHA512 hash, as Ceil(512 * log(2,36)) == 100.

A Tuid is a transparent version of a Huid.

A Luid64 is a fast, 128-bit random-number based unique identifier that is guaranteed to be all-in less than 64 bytes. It is ascii/base-36 encoded. Example: 07sak24k4n7onqcz0hclmy8gi-1-73377. The last two parts after the first dash are non-random, enabling some tracking of sequence number and origin. They are a sequence number (no more than 6 bytes) and a process id.

Command line versions of the generators are available in cmd/. Make will install them.

The bytes in a Ruid always start with `ruid_v` and
are followed by two digits of version identifier and an '_'
underscore before the variable portion starts.

~~~
ruid_v01_kyPC3GgPLHh1qZ_-F-12Ow8N9m4=
ruid_v01_QvU2ZIHwrsYbyj7UMFg5ZSVXj0w=
~~~

Likewise, a Huid will always start with `huid_v`.

~~~
huid_v01_fHRtOjIwMTQtMTEtMjBUMTc6NDQ6NDcuNDQ0NzU1NDE5LTA4OjAwfHBpZDowMDAwMDEwODc3fGxvYzpFNHlvWHBpZmlsa2ctVWNsM1dSemg5LWdOSFU9fHNlcTowMDAwMDAwMDAwMDAwMDAwMDAwM3w=
huid_v01_fHRtOjIwMTQtMTEtMjBUMTc6NDQ6NDcuNDQ0NzYwMTczLTA4OjAwfHBpZDowMDAwMDEwODc3fGxvYzpFNHlvWHBpZmlsa2ctVWNsM1dSemg5LWdOSFU9fHNlcTowMDAwMDAwMDAwMDAwMDAwMDAwNHw=
~~~

The huid will base64 decode (use the url safe version) to lines similar to these:

~~~
|tm:2014-11-20T17:39:06.824687644-08:00|pid:0000010801|loc:E4yoXpifilkg-Ucl3WRzh9-gNHU=|seq:00000000000000000003|
|tm:2014-11-20T17:39:06.824691496-08:00|pid:0000010801|loc:E4yoXpifilkg-Ucl3WRzh9-gNHU=|seq:00000000000000000004|
~~~

where the 'loc' is an opaque sha1 hash of the important uniqueLocation 
parameter to NewRuidGen(). It is important to make uniqueLocation
as hard to duplicate as possible, to allow sequence resolution.

The unique string after the prefix is URL-safe base64
encoded. The Huid will un-base64 to
a human readable, informative sequence line. The Ruid and
will base64 decode to an opaque identifier, a
SHA1 hash. If you use both, instantiate seperate Generators,
as Huid and Ruid utilize the same sequence counter.

Brief benchmarks on my laptop suggest that it costs less than 4 usec to
generate a Ruid or a Huid. At this rate you could generate 250K/second/core.

~~~
$ go test -v -bench .
...
BenchmarkRuid	 1000000	      2716 ns/op
BenchmarkHuid	 1000000	      3031 ns/op
~~~

use notes
------------
~~~
// NewRuidGen(): uniqueLocation should be a byte
// sequence that is unique to this specific physical location.
// Suggestions: a hardware
// mac address, your external ip address, the traceroute out
// a known distant location on the public internet.
// The uniqueLocation parameter in the NewRuidGen() call
// should be as unique as possible.
func NewRuidGen(uniqueLocation string) *RuidGen
~~~
