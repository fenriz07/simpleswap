package types

import (
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fenriz07/simpleswap/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestGetFeeToReceive(t *testing.T) {
	lp := LiquidityProvider{
		Amount: 1430000000,
	}

	fee := lp.GetFeeToReceive(12400000, 5000000000)

	fmt.Println(fee)
}

func TestMsgSwap_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSwap
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSwap{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSwap{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
