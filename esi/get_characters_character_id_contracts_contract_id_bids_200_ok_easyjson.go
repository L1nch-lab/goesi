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

func easyjsonDeaef821DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdContractsContractIdBids200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdContractsContractIdBids200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdContractsContractIdBids200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdContractsContractIdBids200Ok
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
func easyjsonDeaef821EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdContractsContractIdBids200OkList) {
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
func (v GetCharactersCharacterIdContractsContractIdBids200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDeaef821EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdContractsContractIdBids200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDeaef821EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdContractsContractIdBids200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDeaef821DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdContractsContractIdBids200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDeaef821DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonDeaef821DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdContractsContractIdBids200Ok) {
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
		case "bid_id":
			out.BidId = int32(in.Int32())
		case "bidder_id":
			out.BidderId = int32(in.Int32())
		case "date_bid":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateBid).UnmarshalJSON(data))
			}
		case "amount":
			out.Amount = float32(in.Float32())
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
func easyjsonDeaef821EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdContractsContractIdBids200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BidId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bid_id\":")
		out.Int32(int32(in.BidId))
	}
	if in.BidderId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bidder_id\":")
		out.Int32(int32(in.BidderId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"date_bid\":")
		out.Raw((in.DateBid).MarshalJSON())
	}
	if in.Amount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"amount\":")
		out.Float32(float32(in.Amount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdContractsContractIdBids200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDeaef821EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdContractsContractIdBids200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDeaef821EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdContractsContractIdBids200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDeaef821DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdContractsContractIdBids200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDeaef821DecodeGithubComAntihaxGoesiEsi1(l, v)
}
