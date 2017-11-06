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

func easyjsonCeaeb3dcDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationCorporationIdMiningExtractions200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationCorporationIdMiningExtractions200OkList, 0, 1)
			} else {
				*out = GetCorporationCorporationIdMiningExtractions200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationCorporationIdMiningExtractions200Ok
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
func easyjsonCeaeb3dcEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationCorporationIdMiningExtractions200OkList) {
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
func (v GetCorporationCorporationIdMiningExtractions200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCeaeb3dcEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationCorporationIdMiningExtractions200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCeaeb3dcEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCeaeb3dcDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCeaeb3dcDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonCeaeb3dcDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationCorporationIdMiningExtractions200Ok) {
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
		case "structure_id":
			out.StructureId = int64(in.Int64())
		case "moon_id":
			out.MoonId = int32(in.Int32())
		case "extraction_start_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExtractionStartTime).UnmarshalJSON(data))
			}
		case "chunk_arrival_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ChunkArrivalTime).UnmarshalJSON(data))
			}
		case "natural_decay_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.NaturalDecayTime).UnmarshalJSON(data))
			}
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
func easyjsonCeaeb3dcEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationCorporationIdMiningExtractions200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.StructureId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"structure_id\":")
		out.Int64(int64(in.StructureId))
	}
	if in.MoonId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"moon_id\":")
		out.Int32(int32(in.MoonId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"extraction_start_time\":")
		out.Raw((in.ExtractionStartTime).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"chunk_arrival_time\":")
		out.Raw((in.ChunkArrivalTime).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"natural_decay_time\":")
		out.Raw((in.NaturalDecayTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationCorporationIdMiningExtractions200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCeaeb3dcEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationCorporationIdMiningExtractions200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCeaeb3dcEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCeaeb3dcDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCeaeb3dcDecodeGithubComAntihaxGoesiEsi1(l, v)
}
