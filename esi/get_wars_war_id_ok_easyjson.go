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

func easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetWarsWarIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetWarsWarIdOkList, 0, 1)
			} else {
				*out = GetWarsWarIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetWarsWarIdOk
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
func easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetWarsWarIdOkList) {
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
func (v GetWarsWarIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetWarsWarIdOk) {
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
		case "id":
			out.Id = int32(in.Int32())
		case "declared":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Declared).UnmarshalJSON(data))
			}
		case "started":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Started).UnmarshalJSON(data))
			}
		case "retracted":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Retracted).UnmarshalJSON(data))
			}
		case "finished":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Finished).UnmarshalJSON(data))
			}
		case "mutual":
			out.Mutual = bool(in.Bool())
		case "open_for_allies":
			out.OpenForAllies = bool(in.Bool())
		case "aggressor":
			(out.Aggressor).UnmarshalEasyJSON(in)
		case "defender":
			easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi2(in, &out.Defender)
		case "allies":
			if in.IsNull() {
				in.Skip()
				out.Allies = nil
			} else {
				in.Delim('[')
				if out.Allies == nil {
					if !in.IsDelim(']') {
						out.Allies = make([]GetWarsWarIdAlly, 0, 8)
					} else {
						out.Allies = []GetWarsWarIdAlly{}
					}
				} else {
					out.Allies = (out.Allies)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetWarsWarIdAlly
					easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi3(in, &v4)
					out.Allies = append(out.Allies, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetWarsWarIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Int32(int32(in.Id))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"declared\":")
		out.Raw((in.Declared).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"started\":")
		out.Raw((in.Started).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"retracted\":")
		out.Raw((in.Retracted).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"finished\":")
		out.Raw((in.Finished).MarshalJSON())
	}
	if in.Mutual {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"mutual\":")
		out.Bool(bool(in.Mutual))
	}
	if in.OpenForAllies {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"open_for_allies\":")
		out.Bool(bool(in.OpenForAllies))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"aggressor\":")
		(in.Aggressor).MarshalEasyJSON(out)
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"defender\":")
		easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi2(out, in.Defender)
	}
	if len(in.Allies) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"allies\":")
		if in.Allies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Allies {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi3(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWarsWarIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetWarsWarIdAlly) {
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
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
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
func easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetWarsWarIdAlly) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
	out.RawByte('}')
}
func easyjson45a5fe98DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetWarsWarIdDefender) {
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
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "ships_killed":
			out.ShipsKilled = int32(in.Int32())
		case "isk_destroyed":
			out.IskDestroyed = float32(in.Float32())
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
func easyjson45a5fe98EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetWarsWarIdDefender) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
	if in.ShipsKilled != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ships_killed\":")
		out.Int32(int32(in.ShipsKilled))
	}
	if in.IskDestroyed != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isk_destroyed\":")
		out.Float32(float32(in.IskDestroyed))
	}
	out.RawByte('}')
}
