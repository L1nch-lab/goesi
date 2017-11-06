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

func easyjsonF642f465DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostUiOpenwindowNewmailNewMailList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostUiOpenwindowNewmailNewMailList, 0, 1)
			} else {
				*out = PostUiOpenwindowNewmailNewMailList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostUiOpenwindowNewmailNewMail
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
func easyjsonF642f465EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostUiOpenwindowNewmailNewMailList) {
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
func (v PostUiOpenwindowNewmailNewMailList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642f465EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUiOpenwindowNewmailNewMailList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642f465EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMailList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642f465DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMailList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642f465DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonF642f465DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostUiOpenwindowNewmailNewMail) {
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
		case "subject":
			out.Subject = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "recipients":
			if in.IsNull() {
				in.Skip()
				out.Recipients = nil
			} else {
				in.Delim('[')
				if out.Recipients == nil {
					if !in.IsDelim(']') {
						out.Recipients = make([]int32, 0, 16)
					} else {
						out.Recipients = []int32{}
					}
				} else {
					out.Recipients = (out.Recipients)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Recipients = append(out.Recipients, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "to_mailing_list_id":
			out.ToMailingListId = int32(in.Int32())
		case "to_corp_or_alliance_id":
			out.ToCorpOrAllianceId = int32(in.Int32())
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
func easyjsonF642f465EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostUiOpenwindowNewmailNewMail) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Subject != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"subject\":")
		out.String(string(in.Subject))
	}
	if in.Body != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"body\":")
		out.String(string(in.Body))
	}
	if len(in.Recipients) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"recipients\":")
		if in.Recipients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Recipients {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	if in.ToMailingListId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"to_mailing_list_id\":")
		out.Int32(int32(in.ToMailingListId))
	}
	if in.ToCorpOrAllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"to_corp_or_alliance_id\":")
		out.Int32(int32(in.ToCorpOrAllianceId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUiOpenwindowNewmailNewMail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642f465EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUiOpenwindowNewmailNewMail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642f465EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642f465DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642f465DecodeGithubComAntihaxGoesiEsi1(l, v)
}
