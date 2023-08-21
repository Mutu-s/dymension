package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/dymensionxyz/dymension/x/liquidity/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetLiquidityPoolKey() {
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetPoolKey(9))
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPoolKey(0))
}

func (s *keysTestSuite) TestGetLiquidityPoolByReserveAccIndexKey() {
	len20acc, err := types.GetReserveAcc("poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", false)
	s.Require().NoError(err)
	len32acc, err := types.GetReserveAcc("poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", true)
	s.Require().NoError(err)

	s.Require().Equal([]byte{0x12}, types.GetPoolByReserveAccIndexKey(nil))
	s.Require().Equal([]byte{0x12, 0x14, 0xd3, 0x5a, 0xc, 0xc1, 0x6e, 0xe5, 0x98, 0xf9, 0xb, 0x4, 0x4c, 0xe2, 0x96, 0xa4, 0x5, 0xba, 0x9c, 0x38, 0x1e, 0x38}, types.GetPoolByReserveAccIndexKey(len20acc))
	s.Require().Equal([]byte{0x12, 0x20, 0x87, 0xec, 0x7d, 0x8f, 0xca, 0xee, 0xb0, 0xaa, 0x2, 0x1d, 0xc7, 0xd0, 0x69, 0xb, 0x1e, 0xb8, 0xfb, 0x3e, 0x8e, 0xb1, 0x22, 0x7f, 0x78, 0xae, 0x6c, 0x5e, 0x8a, 0x96, 0xc6, 0x7, 0xc4, 0x98}, types.GetPoolByReserveAccIndexKey(len32acc))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchKey() {
	s.Require().Equal([]byte{0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPoolBatchKey(10))
	s.Require().Equal([]byte{0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPoolBatchKey(0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchDepositMsgsPrefix() {
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPoolBatchDepositMsgStatesPrefix(10))
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPoolBatchDepositMsgStatesPrefix(0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchWithdrawMsgsPrefix() {
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPoolBatchWithdrawMsgsPrefix(10))
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPoolBatchWithdrawMsgsPrefix(0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchSwapMsgsPrefix() {
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPoolBatchSwapMsgStatesPrefix(10))
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPoolBatchSwapMsgStatesPrefix(0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchDepositMsgIndex() {
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		types.GetPoolBatchDepositMsgStateIndexKey(10, 10))
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		types.GetPoolBatchDepositMsgStateIndexKey(0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchWithdrawMsgIndex() {
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		types.GetPoolBatchWithdrawMsgStateIndexKey(10, 10))
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		types.GetPoolBatchWithdrawMsgStateIndexKey(0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchSwapMsgIndex() {
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		types.GetPoolBatchSwapMsgStateIndexKey(10, 10))
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		types.GetPoolBatchSwapMsgStateIndexKey(0, 0))
}
