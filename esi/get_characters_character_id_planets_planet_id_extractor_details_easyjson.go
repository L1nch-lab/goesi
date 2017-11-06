// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7d2da406DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7d2da406EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d2da406EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d2da406EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d2da406DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d2da406DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson7d2da406DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "heads":
			if in.IsNull() {
				in.Skip()
				out.Heads = nil
			} else {
				in.Delim('[')
				if out.Heads == nil {
					if !in.IsDelim(']') {
						out.Heads = make([]GetCharactersCharacterIdPlanetsPlanetIdHead, 0, 5)
					} else {
						out.Heads = []GetCharactersCharacterIdPlanetsPlanetIdHead{}
					}
				} else {
					out.Heads = (out.Heads)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdPlanetsPlanetIdHead
					(v4).UnmarshalEasyJSON(in)
					out.Heads = append(out.Heads, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "product_type_id":
			out.ProductTypeId = int32(in.Int32())
		case "cycle_time":
			out.CycleTime = int32(in.Int32())
		case "head_radius":
			out.HeadRadius = float32(in.Float32())
		case "qty_per_cycle":
			out.QtyPerCycle = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7d2da406EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Heads) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"heads\":")
		if in.Heads == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Heads {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.ProductTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"product_type_id\":")
		out.Int32(int32(in.ProductTypeId))
	}
	if in.CycleTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cycle_time\":")
		out.Int32(int32(in.CycleTime))
	}
	if in.HeadRadius != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"head_radius\":")
		out.Float32(float32(in.HeadRadius))
	}
	if in.QtyPerCycle != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"qty_per_cycle\":")
		out.Int32(int32(in.QtyPerCycle))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d2da406EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d2da406EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d2da406DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d2da406DecodeGithubComAntihaxGoesiEsi1(l, v)
}
