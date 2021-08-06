// +build gofuzzbeta

package reser_test

import (
	"testing"
)

// just small test to showcase dev.fuzz
// https://blog.golang.org/fuzz-beta

func Fuzz_PrefixPolySerializer(f *testing.F) {
	s := makePrefixPolySerializer(nil)
	f.Fuzz(func(t *testing.T, data []byte) {
		res, err := s.PolyDeserialize(data)
		if err != nil {
			t.Skip()
		}
		_ = res
	})
}
