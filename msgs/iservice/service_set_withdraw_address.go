package iservice

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgSetWithdrawAddress struct {
		Owner           string `bson:"owner" yaml:"owner"`
		WithdrawAddress string `bson:"withdraw_address" yaml:"withdraw_address"`
	}
)

func (m *DocMsgSetWithdrawAddress) GetType() string {
	return MsgTypeSetWithdrawAddress
}

func (m *DocMsgSetWithdrawAddress) BuildMsg(v interface{}) {
	msg := v.(MsgSetWithdrawAddress)

	m.Owner = msg.Owner.String()
	m.WithdrawAddress = msg.WithdrawAddress.String()
}

func (m *DocMsgSetWithdrawAddress) HandleTxMsg(msg MsgSetWithdrawAddress) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Owner, m.WithdrawAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
