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

func easyjsonE446b04DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdMailLabelsLabelList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdMailLabelsLabelList, 0, 2)
			} else {
				*out = PostCharactersCharacterIdMailLabelsLabelList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdMailLabelsLabel
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
func easyjsonE446b04EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdMailLabelsLabelList) {
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
func (v PostCharactersCharacterIdMailLabelsLabelList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE446b04EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdMailLabelsLabelList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE446b04EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdMailLabelsLabelList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE446b04DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdMailLabelsLabelList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE446b04DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonE446b04DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdMailLabelsLabel) {
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
		case "name":
			out.Name = string(in.String())
		case "color":
			out.Color = string(in.String())
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
func easyjsonE446b04EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdMailLabelsLabel) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Color != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"color\":")
		out.String(string(in.Color))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdMailLabelsLabel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE446b04EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdMailLabelsLabel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE446b04EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdMailLabelsLabel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE446b04DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdMailLabelsLabel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE446b04DecodeGithubComAntihaxGoesiEsi1(l, v)
}
