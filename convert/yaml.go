package convert

import (
	"encoding/json"
	"io"
	"math"

	goyaml "github.com/richmondwang/yaml"

	"github.com/richmondwang/yj/yaml"
)

type YAML struct {
	FloatStrings bool
	JSONKeys     bool
	EscapeHTML   bool
}

func (YAML) String() string {
	return "YAML"
}

func (y YAML) Encode(w io.Writer, in interface{}) error {
	_enc := goyaml.NewEncoder(w)
	_enc.SetLineLength(-1)
	encoder := &yaml.Encoder{
		EncodeYAML: _enc.Encode,
	}
	if y.FloatStrings {
		encoder.NaN = "NaN"
		encoder.PosInf = "Infinity"
		encoder.NegInf = "-Infinity"
	}
	if y.JSONKeys {
		encoder.KeyUnmarshal = json.Unmarshal
	}
	return encoder.YAML(in)
}

func (y YAML) Decode(r io.Reader) (interface{}, error) {
	decoder := &yaml.Decoder{
		DecodeYAML: goyaml.NewDecoder(r).Decode,
		KeyMarshal: (&yaml.JSON{EscapeHTML: y.EscapeHTML}).Marshal,

		NaN:    (*float64)(nil),
		PosInf: math.MaxFloat64,
		NegInf: -math.MaxFloat64,
	}

	if y.FloatStrings {
		decoder.NaN = "NaN"
		decoder.PosInf = "Infinity"
		decoder.NegInf = "-Infinity"
	}
	return decoder.JSON()
}
