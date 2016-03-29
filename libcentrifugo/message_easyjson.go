// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package libcentrifugo

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_decode_github_com_centrifugal_centrifugo_libcentrifugo_Message(in *jlexer.Lexer, out *Message) {
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
		case "uid":
			out.UID = MessageID(in.String())
		case "timestamp":
			out.Timestamp = string(in.String())
		case "info":
			if in.IsNull() {
				in.Skip()
				out.Info = nil
			} else {
				out.Info = new(ClientInfo)
				easyjson_decode_github_com_centrifugal_centrifugo_libcentrifugo_ClientInfo(in, &*out.Info)
			}
		case "channel":
			out.Channel = Channel(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				out.Data = new(json.RawMessage)
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Data).UnmarshalJSON(data))
				}
			}
		case "client":
			out.Client = ConnID(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_encode_github_com_centrifugal_centrifugo_libcentrifugo_Message(out *jwriter.Writer, in *Message) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"uid\":")
	out.String(string(in.UID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"timestamp\":")
	out.String(string(in.Timestamp))
	if in.Info != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"info\":")
		if in.Info == nil {
			out.RawString("null")
		} else {
			easyjson_encode_github_com_centrifugal_centrifugo_libcentrifugo_ClientInfo(out, &*in.Info)
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"channel\":")
	out.String(string(in.Channel))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"data\":")
	if in.Data == nil {
		out.RawString("null")
	} else {
		out.Raw((*in.Data).MarshalJSON())
	}
	if in.Client != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"client\":")
		out.String(string(in.Client))
	}
	out.RawByte('}')
}
func (v *Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_encode_github_com_centrifugal_centrifugo_libcentrifugo_Message(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v *Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_encode_github_com_centrifugal_centrifugo_libcentrifugo_Message(w, v)
}
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_decode_github_com_centrifugal_centrifugo_libcentrifugo_Message(&r, v)
	return r.Error()
}
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_decode_github_com_centrifugal_centrifugo_libcentrifugo_Message(l, v)
}
func easyjson_decode_github_com_centrifugal_centrifugo_libcentrifugo_ClientInfo(in *jlexer.Lexer, out *ClientInfo) {
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
		case "user":
			out.User = UserID(in.String())
		case "client":
			out.Client = ConnID(in.String())
		case "default_info":
			if in.IsNull() {
				in.Skip()
				out.DefaultInfo = nil
			} else {
				out.DefaultInfo = new(json.RawMessage)
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DefaultInfo).UnmarshalJSON(data))
				}
			}
		case "channel_info":
			if in.IsNull() {
				in.Skip()
				out.ChannelInfo = nil
			} else {
				out.ChannelInfo = new(json.RawMessage)
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ChannelInfo).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_encode_github_com_centrifugal_centrifugo_libcentrifugo_ClientInfo(out *jwriter.Writer, in *ClientInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"user\":")
	out.String(string(in.User))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"client\":")
	out.String(string(in.Client))
	if in.DefaultInfo != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"default_info\":")
		if in.DefaultInfo == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.DefaultInfo).MarshalJSON())
		}
	}
	if in.ChannelInfo != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"channel_info\":")
		if in.ChannelInfo == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.ChannelInfo).MarshalJSON())
		}
	}
	out.RawByte('}')
}
