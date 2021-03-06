package msg
//
//import (
//	. "github.com/bianjieai/irita-sync/msgs"
//	"github.com/bianjieai/irita-sync/models"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//)
//
//type DocTxMsgAddLiquidity struct {
//	MaxToken     models.Coin `bson:"max_token"`      // coin to be deposited as liquidity with an upper bound for its amount
//	ExactIrisAmt string      `bson:"exact_iris_amt"` // exact amount of native asset being add to the liquidity pool
//	MinLiquidity string      `bson:"min_liquidity"`  // lower bound UNI sender is willing to accept for deposited coins
//	Deadline     int64       `bson:"deadline"`
//	Sender       string      `bson:"sender"`
//}
//
//func (doctx *DocTxMsgAddLiquidity) GetType() string {
//	return TxTypeAddLiquidity
//}
//
//func (doctx *DocTxMsgAddLiquidity) BuildMsg(txMsg interface{}) {
//	msg := txMsg.(MsgAddLiquidity)
//	doctx.Sender = msg.Sender.String()
//	doctx.MinLiquidity = msg.MinLiquidity.String()
//	doctx.ExactIrisAmt = msg.ExactStandardAmt.String()
//	doctx.Deadline = msg.Deadline
//	doctx.MaxToken = models.BuildDocCoin(msg.MaxToken)
//}
//
//func (m *DocTxMsgAddLiquidity) HandleTxMsg(msg sdk.Msg) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, m.Sender)
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
//
//type DocTxMsgRemoveLiquidity struct {
//	MinToken          string      `bson:"min_token"`          // coin to be withdrawn with a lower bound for its amount
//	WithdrawLiquidity models.Coin `bson:"withdraw_liquidity"` // amount of UNI to be burned to withdraw liquidity from a reserve pool
//	MinIrisAmt        string      `bson:"min_iris_amt"`       // minimum amount of the native asset the sender is willing to accept
//	Deadline          int64       `bson:"deadline"`
//	Sender            string      `bson:"sender"`
//}
//
//func (doctx *DocTxMsgRemoveLiquidity) GetType() string {
//	return TxTypeRemoveLiquidity
//}
//
//func (doctx *DocTxMsgRemoveLiquidity) BuildMsg(txMsg interface{}) {
//	msg := txMsg.(MsgRemoveLiquidity)
//	doctx.Sender = msg.Sender.String()
//	doctx.MinIrisAmt = msg.MinStandardAmt.String()
//	doctx.MinToken = msg.MinToken.String()
//	doctx.Deadline = msg.Deadline
//	doctx.WithdrawLiquidity = models.BuildDocCoin(msg.WithdrawLiquidity)
//}
//func (m *DocTxMsgRemoveLiquidity) HandleTxMsg(msg sdk.Msg) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, m.Sender)
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
//
//type DocTxMsgSwapOrder struct {
//	Input      Input  `bson:"input"`        // the amount the sender is trading
//	Output     Output `bson:"output"`       // the amount the sender is receiving
//	Deadline   int64  `bson:"deadline"`     // deadline for the transaction to still be considered valid
//	IsBuyOrder bool   `bson:"is_buy_order"` // boolean indicating whether the order should be treated as a buy or sell
//}
//
//type Input struct {
//	Address string      `bson:"address"`
//	Coin    models.Coin `bson:"coin"`
//}
//
//type Output struct {
//	Address string      `bson:"address"`
//	Coin    models.Coin `bson:"coin"`
//}
//
//func (doctx *DocTxMsgSwapOrder) GetType() string {
//	return TxTypeSwapOrder
//}
//
//func (doctx *DocTxMsgSwapOrder) BuildMsg(txMsg interface{}) {
//	msg := txMsg.(MsgSwapOrder)
//	doctx.Deadline = msg.Deadline
//	doctx.IsBuyOrder = msg.IsBuyOrder
//	doctx.Input = Input{
//		Address: msg.Input.Address.String(),
//		Coin:    models.BuildDocCoin(msg.Input.Coin),
//	}
//	doctx.Output = Output{
//		Address: msg.Output.Address.String(),
//		Coin:    models.BuildDocCoin(msg.Output.Coin),
//	}
//}
//func (m *DocTxMsgSwapOrder) HandleTxMsg(msg sdk.Msg) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, m.Output.Address, m.Input.Address)
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
