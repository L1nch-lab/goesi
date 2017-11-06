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

func easyjson1febafc2DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdMedalsIssued200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdMedalsIssued200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdMedalsIssued200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdMedalsIssued200Ok
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
func easyjson1febafc2EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdMedalsIssued200OkList) {
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
func (v GetCorporationsCorporationIdMedalsIssued200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1febafc2EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdMedalsIssued200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1febafc2EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdMedalsIssued200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1febafc2DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdMedalsIssued200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1febafc2DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson1febafc2DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdMedalsIssued200Ok) {
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
		case "medal_id":
			out.MedalId = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "issuer_id":
			out.IssuerId = int32(in.Int32())
		case "issued_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.IssuedAt).UnmarshalJSON(data))
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
func easyjson1febafc2EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdMedalsIssued200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MedalId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"medal_id\":")
		out.Int32(int32(in.MedalId))
	}
	if in.CharacterId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"character_id\":")
		out.Int32(int32(in.CharacterId))
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	if in.Status != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		out.String(string(in.Status))
	}
	if in.IssuerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"issuer_id\":")
		out.Int32(int32(in.IssuerId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"issued_at\":")
		out.Raw((in.IssuedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdMedalsIssued200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1febafc2EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdMedalsIssued200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1febafc2EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdMedalsIssued200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1febafc2DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdMedalsIssued200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1febafc2DecodeGithubComAntihaxGoesiEsi1(l, v)
}
