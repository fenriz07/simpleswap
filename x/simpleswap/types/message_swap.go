package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	eth   = "eth"
	weth  = "weth"
	steth = "steth"
)

var _ sdk.Msg = &MsgSwap{}

func NewMsgSwap(creator string, amount sdk.Coins, denom string) *MsgSwap {
	return &MsgSwap{
		Creator: creator,
		Amount:  amount,
		Denom:   denom,
	}
}

func (msg *MsgSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Amount) > 1 {
		return errorsmod.Wrapf(ErrACoinIsValidate, "")
	}

	coin := msg.Amount[0]

	if coin.Denom != eth && coin.Denom != weth && coin.Denom != steth {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "the coin denom has to be %s %s %s your denom was: %s", eth, weth, steth, coin.Denom)
	}

	return nil
}
