package reser_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/teawithsand/reser"
)

func DoTestPolyEnDeCoder[T any](
	t *testing.T,
	encoderFactroy func(w io.Writer) (e reser.PolyEncoder[T], err error),
	decoderFactory func(r io.Reader) (d reser.PolyDecoder[T], err error),
) {

}

type A struct {
	N int
}
type B struct {
	A int
}
type C struct {
	Q int
}

func FuzzPrefixPolyUnmarshaler(f *testing.F) {
	reg := reser.NewTypeTagRegistry[string]()
	reg.AddMappingSimple("a", A{})
	reg.AddMappingSimple("b", B{})
	reg.AddMappingSimple("c", C{})

	fac := reser.PrefixPolyDecoderFactroy[string, any]{
		TagRegistry: reg,
		TagDecoderFactory: func(r io.Reader) (d reser.PolyDecoder[string], err error) {
			ef := reser.SinglePolyDecoderFactory[string]{
				DecoderFactory: func(r io.Reader) (dec reser.Decoder, err error) {
					dec = json.NewDecoder(r)
					return
				},
				Factory: func() *string {
					var s string
					return &s
				},
			}
			d, err = ef.NewDecoder(r)
			if err != nil {
				return
			}
			return
		},
		DataDecoderFactory: func(r io.Reader) (d reser.Decoder, err error) {
			d = json.NewDecoder(r)
			return
		},
	}

	f.Add([]byte(`"a"{"N":12}`))

	f.Fuzz(func(t *testing.T, data []byte) {
		d, err := fac.NewDecoder(bytes.NewReader(data))
		if err != nil {
			return
		}

		v, err := d.PolyDecode()
		if err != nil {
			return
		}

		panic(v)
	})
}
