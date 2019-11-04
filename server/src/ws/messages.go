package ws

import (
	"encoding/base64"
	"encoding/json"

	"github.com/SergeyShpak/owngame/server/src/types"
)

func NewMsgParticipants(participants []types.Participant) ([]byte, error) {
	msg := types.WSMsgParticipant{
		Participants: participants,
	}
	wsMsg, err := newWsMsg(types.WSMessageTypeParticipants, msg)
	if err != nil {
		return nil, err
	}
	return wsMsg, nil
}

func newWsMsg(t types.WSMessageType, msg interface{}) ([]byte, error) {
	msgB, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	msgS := base64.StdEncoding.EncodeToString(msgB)
	wsMsg := types.WSMessage{
		Type:    string(t),
		Message: msgS,
	}
	wsMsgB, err := json.Marshal(wsMsg)
	if err != nil {
		return nil, err
	}
	return wsMsgB, nil
}
