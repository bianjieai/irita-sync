package record

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/cosmos/cosmos-sdk/types"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msgData := v.(type) {
	case MsgRecordCreate:
		docMsg := DocMsgRecordCreate{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
