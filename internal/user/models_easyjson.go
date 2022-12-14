// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package user

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

func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser(in *jlexer.Lexer, out *userCacheKey) {
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser(out *jwriter.Writer, in userCacheKey) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v userCacheKey) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v userCacheKey) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *userCacheKey) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *userCacheKey) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser1(in *jlexer.Lexer, out *updateUserRequest) {
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
		case "name":
			out.Name = string(in.String())
		case "date_of_birth":
			out.DateOfBirth = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser1(out *jwriter.Writer, in updateUserRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"date_of_birth\":"
		out.RawString(prefix)
		out.String(string(in.DateOfBirth))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v updateUserRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v updateUserRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *updateUserRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *updateUserRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser1(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser2(in *jlexer.Lexer, out *response) {
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
		case "success":
			out.Success = bool(in.Bool())
		case "body":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Body).UnmarshalJSON(data))
			}
		case "errors":
			if in.IsNull() {
				in.Skip()
				out.Errors = nil
			} else {
				in.Delim('[')
				if out.Errors == nil {
					if !in.IsDelim(']') {
						out.Errors = make([]errResponse, 0, 2)
					} else {
						out.Errors = []errResponse{}
					}
				} else {
					out.Errors = (out.Errors)[:0]
				}
				for !in.IsDelim(']') {
					var v1 errResponse
					(v1).UnmarshalEasyJSON(in)
					out.Errors = append(out.Errors, v1)
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser2(out *jwriter.Writer, in response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"success\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Success))
	}
	if len(in.Body) != 0 {
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.Raw((in.Body).MarshalJSON())
	}
	if len(in.Errors) != 0 {
		const prefix string = ",\"errors\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Errors {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser2(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser3(in *jlexer.Lexer, out *getUserResponse) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "date_of_birth":
			out.DateOfBirth = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser3(out *jwriter.Writer, in getUserResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"date_of_birth\":"
		out.RawString(prefix)
		out.String(string(in.DateOfBirth))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getUserResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getUserResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getUserResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getUserResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser3(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser4(in *jlexer.Lexer, out *errResponse) {
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
		case "code":
			out.Code = string(in.String())
		case "context":
			out.Context = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser4(out *jwriter.Writer, in errResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"context\":"
		out.RawString(prefix)
		out.String(string(in.Context))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v errResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v errResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *errResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *errResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser4(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser5(in *jlexer.Lexer, out *createUserResponse) {
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
		case "id":
			out.ID = int64(in.Int64())
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser5(out *jwriter.Writer, in createUserResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createUserResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createUserResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createUserResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createUserResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser5(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser6(in *jlexer.Lexer, out *createUserRequest) {
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
		case "name":
			out.Name = string(in.String())
		case "date_of_birth":
			out.DateOfBirth = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser6(out *jwriter.Writer, in createUserRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"date_of_birth\":"
		out.RawString(prefix)
		out.String(string(in.DateOfBirth))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createUserRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createUserRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createUserRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createUserRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser6(l, v)
}
func easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser7(in *jlexer.Lexer, out *User) {
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
		case "ID":
			out.ID = int64(in.Int64())
		case "Name":
			out.Name = string(in.String())
		case "DateOfBirth":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateOfBirth).UnmarshalJSON(data))
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
func easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser7(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"DateOfBirth\":"
		out.RawString(prefix)
		out.Raw((in.DateOfBirth).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComCrayonwowBaseServiceInternalUser7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComCrayonwowBaseServiceInternalUser7(l, v)
}
