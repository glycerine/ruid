package ruid

import (
	"fmt"
	"strings"
	"testing"

	cv "github.com/glycerine/goconvey/convey"
	"github.com/twinj/uuid"
)

func TestRuid(t *testing.T) {

	ruidGen := NewRuidGen("put unique location string here")
	cv.Convey("Given we generate two Ruids() very quickly", t, func() {
		cv.Convey("Then they should be unique, and start with 'ruid_v'", func() {
			r1 := ruidGen.Ruid()
			r2 := ruidGen.Ruid()
			fmt.Printf("\n r1 = '%s'\n", r1)
			fmt.Printf("\n r2 = '%s'\n", r2)
			cv.So(strings.HasPrefix(r1, `ruid_v`), cv.ShouldEqual, true)
			cv.So(strings.HasPrefix(r2, `ruid_v`), cv.ShouldEqual, true)
			cv.So(r1, cv.ShouldNotEqual, r2)
		})

		cv.Convey("And two Huids() they should be unique, and start with 'huid_v'", func() {
			h1 := ruidGen.Huid()
			h2 := ruidGen.Huid()
			fmt.Printf("\n h1 = '%s'\n", h1)
			fmt.Printf("\n h2 = '%s'\n", h2)
			cv.So(strings.HasPrefix(h1, `huid_v`), cv.ShouldEqual, true)
			cv.So(strings.HasPrefix(h2, `huid_v`), cv.ShouldEqual, true)
			cv.So(h1, cv.ShouldNotEqual, h2)
		})
	})
}

func BenchmarkRuid(b *testing.B) {

	myExternalIP := "my example location"
	ruidGen := NewRuidGen(myExternalIP)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ruidGen.Ruid()
	}
}

func BenchmarkHuid(b *testing.B) {

	myExternalIP := "my example location"
	ruidGen := NewRuidGen(myExternalIP)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ruidGen.Huid()
	}
}

func BenchmarkUUID4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.NewV4()
	}
}

func BenchmarkRuid2(b *testing.B) {

	myExternalIP := "my example location"
	ruidGen := NewRuidGen(myExternalIP)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ruidGen.Ruid2()
	}
}

func BenchmarkLuid64(b *testing.B) {

	myExternalIP := "my example location"
	ruidGen := NewRuidGen(myExternalIP)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ruidGen.Luid64()
	}
}
