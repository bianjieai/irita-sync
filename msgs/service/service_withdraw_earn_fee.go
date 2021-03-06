package service

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgWithdrawEarnedFees struct {
		Owner    string
		Provider string
	}
)

func (m *DocMsgWithdrawEarnedFees) GetType() string {
	return MsgTypeWithdrawEarnedFees
}

func (m *DocMsgWithdrawEarnedFees) BuildMsg(v interface{}) {
	msg := v.(MsgWithdrawEarnedFees)

	m.Owner = msg.Owner.String()
	m.Provider = msg.Provider.String()
}

func (m *DocMsgWithdrawEarnedFees) HandleTxMsg(msg MsgWithdrawEarnedFees) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Owner.String(), msg.Provider.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
