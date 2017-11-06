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

func easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetKillmailsKillmailIdKillmailHashOkList, 0, 1)
			} else {
				*out = GetKillmailsKillmailIdKillmailHashOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetKillmailsKillmailIdKillmailHashOk
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
func easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashOkList) {
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
func (v GetKillmailsKillmailIdKillmailHashOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashOk) {
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
		case "killmail_id":
			out.KillmailId = int32(in.Int32())
		case "killmail_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.KillmailTime).UnmarshalJSON(data))
			}
		case "victim":
			easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi2(in, &out.Victim)
		case "attackers":
			if in.IsNull() {
				in.Skip()
				out.Attackers = nil
			} else {
				in.Delim('[')
				if out.Attackers == nil {
					if !in.IsDelim(']') {
						out.Attackers = make([]GetKillmailsKillmailIdKillmailHashAttacker, 0, 1)
					} else {
						out.Attackers = []GetKillmailsKillmailIdKillmailHashAttacker{}
					}
				} else {
					out.Attackers = (out.Attackers)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetKillmailsKillmailIdKillmailHashAttacker
					(v4).UnmarshalEasyJSON(in)
					out.Attackers = append(out.Attackers, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "moon_id":
			out.MoonId = int32(in.Int32())
		case "war_id":
			out.WarId = int32(in.Int32())
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
func easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.KillmailId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"killmail_id\":")
		out.Int32(int32(in.KillmailId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"killmail_time\":")
		out.Raw((in.KillmailTime).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"victim\":")
		easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi2(out, in.Victim)
	}
	if len(in.Attackers) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"attackers\":")
		if in.Attackers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Attackers {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.SolarSystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"solar_system_id\":")
		out.Int32(int32(in.SolarSystemId))
	}
	if in.MoonId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"moon_id\":")
		out.Int32(int32(in.MoonId))
	}
	if in.WarId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"war_id\":")
		out.Int32(int32(in.WarId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashVictim) {
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
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "damage_taken":
			out.DamageTaken = int32(in.Int32())
		case "ship_type_id":
			out.ShipTypeId = int32(in.Int32())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetKillmailsKillmailIdKillmailHashItem1, 0, 1)
					} else {
						out.Items = []GetKillmailsKillmailIdKillmailHashItem1{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v7 GetKillmailsKillmailIdKillmailHashItem1
					easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi3(in, &v7)
					out.Items = append(out.Items, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "position":
			easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi4(in, &out.Position)
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
func easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashVictim) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CharacterId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"character_id\":")
		out.Int32(int32(in.CharacterId))
	}
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
	if in.FactionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"faction_id\":")
		out.Int32(int32(in.FactionId))
	}
	if in.DamageTaken != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"damage_taken\":")
		out.Int32(int32(in.DamageTaken))
	}
	if in.ShipTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ship_type_id\":")
		out.Int32(int32(in.ShipTypeId))
	}
	if len(in.Items) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"items\":")
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Items {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi3(out, v9)
			}
			out.RawByte(']')
		}
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"position\":")
		easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi4(out, in.Position)
	}
	out.RawByte('}')
}
func easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashPosition) {
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
		case "x":
			out.X = float32(in.Float32())
		case "y":
			out.Y = float32(in.Float32())
		case "z":
			out.Z = float32(in.Float32())
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
func easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Float32(float32(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Float32(float32(in.Y))
	}
	if in.Z != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"z\":")
		out.Float32(float32(in.Z))
	}
	out.RawByte('}')
}
func easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItem1) {
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
		case "item_type_id":
			out.ItemTypeId = int32(in.Int32())
		case "quantity_destroyed":
			out.QuantityDestroyed = int64(in.Int64())
		case "quantity_dropped":
			out.QuantityDropped = int64(in.Int64())
		case "singleton":
			out.Singleton = int32(in.Int32())
		case "flag":
			out.Flag = int32(in.Int32())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetKillmailsKillmailIdKillmailHashItem, 0, 2)
					} else {
						out.Items = []GetKillmailsKillmailIdKillmailHashItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v10 GetKillmailsKillmailIdKillmailHashItem
					easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi5(in, &v10)
					out.Items = append(out.Items, v10)
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
func easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItem1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item_type_id\":")
		out.Int32(int32(in.ItemTypeId))
	}
	if in.QuantityDestroyed != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity_destroyed\":")
		out.Int64(int64(in.QuantityDestroyed))
	}
	if in.QuantityDropped != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity_dropped\":")
		out.Int64(int64(in.QuantityDropped))
	}
	if in.Singleton != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"singleton\":")
		out.Int32(int32(in.Singleton))
	}
	if in.Flag != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"flag\":")
		out.Int32(int32(in.Flag))
	}
	if len(in.Items) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"items\":")
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Items {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi5(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonB5ebbfcbDecodeGithubComAntihaxGoesiEsi5(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItem) {
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
		case "item_type_id":
			out.ItemTypeId = int32(in.Int32())
		case "quantity_destroyed":
			out.QuantityDestroyed = int64(in.Int64())
		case "quantity_dropped":
			out.QuantityDropped = int64(in.Int64())
		case "singleton":
			out.Singleton = int32(in.Int32())
		case "flag":
			out.Flag = int32(in.Int32())
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
func easyjsonB5ebbfcbEncodeGithubComAntihaxGoesiEsi5(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item_type_id\":")
		out.Int32(int32(in.ItemTypeId))
	}
	if in.QuantityDestroyed != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity_destroyed\":")
		out.Int64(int64(in.QuantityDestroyed))
	}
	if in.QuantityDropped != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity_dropped\":")
		out.Int64(int64(in.QuantityDropped))
	}
	if in.Singleton != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"singleton\":")
		out.Int32(int32(in.Singleton))
	}
	if in.Flag != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"flag\":")
		out.Int32(int32(in.Flag))
	}
	out.RawByte('}')
}
