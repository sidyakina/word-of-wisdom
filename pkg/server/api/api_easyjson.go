// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

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

func easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi(in *jlexer.Lexer, out *Message) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Type":
			out.Type = string(in.String())
		case "Payload":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Payload).UnmarshalJSON(data))
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
func easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"Payload\":"
		out.RawString(prefix)
		out.Raw((in.Payload).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi(l, v)
}
func easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi1(in *jlexer.Lexer, out *GetQuoteResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "quote":
			out.Quote = string(in.String())
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
func easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi1(out *jwriter.Writer, in GetQuoteResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"quote\":"
		out.RawString(prefix[1:])
		out.String(string(in.Quote))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetQuoteResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetQuoteResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetQuoteResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetQuoteResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi1(l, v)
}
func easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi2(in *jlexer.Lexer, out *ChallengeResponsePayload) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "answer":
			out.Answer = string(in.String())
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
func easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi2(out *jwriter.Writer, in ChallengeResponsePayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"answer\":"
		out.RawString(prefix[1:])
		out.String(string(in.Answer))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChallengeResponsePayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChallengeResponsePayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChallengeResponsePayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChallengeResponsePayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi2(l, v)
}
func easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi3(in *jlexer.Lexer, out *ChallengeRequestPayload) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "challenge":
			out.Challenge = string(in.String())
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
func easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi3(out *jwriter.Writer, in ChallengeRequestPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"challenge\":"
		out.RawString(prefix[1:])
		out.String(string(in.Challenge))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChallengeRequestPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChallengeRequestPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSidyakinaWordOfWisdomPkgServerApi3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChallengeRequestPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChallengeRequestPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSidyakinaWordOfWisdomPkgServerApi3(l, v)
}
