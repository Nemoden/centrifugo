package libcentrifugo

import (
	"strconv"
	"time"

	"github.com/centrifugal/centrifugo/libcentrifugo/raw"
	"github.com/nats-io/nuid"
)

func newClientInfo(user UserID, client ConnID, defaultInfo *raw.Raw, channelInfo *raw.Raw) ClientInfo {
	return ClientInfo{
		User:        string(user),
		Client:      string(client),
		DefaultInfo: defaultInfo,
		ChannelInfo: channelInfo,
	}
}

func newMessage(ch Channel, data []byte, client ConnID, info *ClientInfo) Message {
	raw := raw.Raw(data)
	return Message{
		UID:       nuid.Next(),
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		Info:      info,
		Channel:   string(ch),
		Data:      &raw,
		Client:    string(client),
	}
}

func newJoinLeave(ch Channel, info ClientInfo) *JoinLeaveMessage {
	return &JoinLeaveMessage{
		Channel: string(ch),
		Data:    info,
	}
}
