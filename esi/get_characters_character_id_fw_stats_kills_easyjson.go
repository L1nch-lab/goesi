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

func easyjsonC00473cbDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdFwStatsKillsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdFwStatsKillsList, 0, 5)
			} else {
				*out = GetCharactersCharacterIdFwStatsKillsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdFwStatsKills
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
func easyjsonC00473cbEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdFwStatsKillsList) {
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
func (v GetCharactersCharacterIdFwStatsKillsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC00473cbEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdFwStatsKillsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC00473cbEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdFwStatsKillsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC00473cbDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdFwStatsKillsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC00473cbDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonC00473cbDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdFwStatsKills) {
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
		case "yesterday":
			out.Yesterday = int32(in.Int32())
		case "last_week":
			out.LastWeek = int32(in.Int32())
		case "total":
			out.Total = int32(in.Int32())
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
func easyjsonC00473cbEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdFwStatsKills) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Yesterday != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"yesterday\":")
		out.Int32(int32(in.Yesterday))
	}
	if in.LastWeek != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"last_week\":")
		out.Int32(int32(in.LastWeek))
	}
	if in.Total != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"total\":")
		out.Int32(int32(in.Total))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdFwStatsKills) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC00473cbEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdFwStatsKills) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC00473cbEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdFwStatsKills) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC00473cbDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdFwStatsKills) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC00473cbDecodeGithubComAntihaxGoesiEsi1(l, v)
}
