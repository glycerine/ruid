Ruid: a really unique id
================

Ruid: a really unique id, Very fast to generate, opaque identifier.
Huid: a really unique id, very fast to generate, decodable to be human readable.

The bytes in a Ruid always start with `ruid_v` and
be followed by two digits of version identifier and an '_'
underscore before the variable portion starts.

ruid_v01_kyPC3GgPLHh1qZ_-F-12Ow8N9m4=
ruid_v01_QvU2ZIHwrsYbyj7UMFg5ZSVXj0w=

Likewise, a Huid will always start with `huid_v`.

huid_v01_fHRtOjIwMTQtMTEtMjBUMTc6NDQ6NDcuNDQ0NzU1NDE5LTA4OjAwfHBpZDowMDAwMDEwODc3fGxvYzpFNHlvWHBpZmlsa2ctVWNsM1dSemg5LWdOSFU9fHNlcTowMDAwMDAwMDAwMDAwMDAwMDAwM3w=

huid_v01_fHRtOjIwMTQtMTEtMjBUMTc6NDQ6NDcuNDQ0NzYwMTczLTA4OjAwfHBpZDowMDAwMDEwODc3fGxvYzpFNHlvWHBpZmlsa2ctVWNsM1dSemg5LWdOSFU9fHNlcTowMDAwMDAwMDAwMDAwMDAwMDAwNHw=

The huid will decode to lines similar to these:

|tm:2014-11-20T17:39:06.824687644-08:00|pid:0000010801|loc:E4yoXpifilkg-Ucl3WRzh9-gNHU=|seq:00000000000000000003|
|tm:2014-11-20T17:39:06.824691496-08:00|pid:0000010801|loc:E4yoXpifilkg-Ucl3WRzh9-gNHU=|seq:00000000000000000004|

where the 'loc' is an opaque sha1 hash of the parameter to NewRuidGen().

The unique string after the prefix is URL-safe base64
encoded. The Huid will un-base64 to
a human readable, informative sequence line. The Ruid and
will base64 decode to an opaque identifier, a
SHA1 hash. If you use both, instantiate seperate Generators,
as Huid and Ruid utilize the same sequence counter.

Brief benchmarks show that a cost of less than 4 usec to
generate a Ruid or a Huid. You could generate 250K/second
if you needed to.

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
~~~
