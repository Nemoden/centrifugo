package libcentrifugo

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/centrifugal/centrifugo/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/pquerna/ffjson/ffjson"
)

func TestMessage(t *testing.T) {
	msg := newMessage(Channel("test"), []byte("{}"), "", nil)
	assert.Equal(t, msg.Channel, Channel("test"))
	msgBytes, err := json.Marshal(msg)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, strings.Contains(string(msgBytes), "\"channel\":\"test\""))
	assert.Equal(t, true, strings.Contains(string(msgBytes), "\"data\":{}"))
	assert.Equal(t, false, strings.Contains(string(msgBytes), "\"client\":\"\"")) // empty field must be omitted
	assert.Equal(t, true, strings.Contains(string(msgBytes), "\"timestamp\":"))
	assert.Equal(t, true, strings.Contains(string(msgBytes), "\"uid\":"))
}

var testMessagePayload string = "{\"foo\":\"bar\"}"

func BenchmarkMsgMarshalJSON(b *testing.B) {
	msg := newMessage(Channel("test"), []byte(testMessagePayload), "", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := newClientMessage()
		resp.Body = msg
		_, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMsgMarshalFFJSON(b *testing.B) {
	msg := newMessage(Channel("test"), []byte(testMessagePayload), "", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := newClientMessage()
		resp.Body = msg
		_, err := ffjson.Marshal(resp)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMsgMarshalFFJSONPool(b *testing.B) {
	msg := newMessage(Channel("test"), []byte(testMessagePayload), "", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := newClientMessage()
		resp.Body = msg
		bytes, err := ffjson.Marshal(resp)
		if err != nil {
			panic(err)
		}
		ffjson.Pool(bytes)
	}
}
