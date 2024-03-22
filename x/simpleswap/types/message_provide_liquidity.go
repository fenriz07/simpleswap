package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProvideLiquidity{}

func NewMsgProvideLiquidity(creator string, stableCoinId string, amount string) *MsgProvideLiquidity {
	return &MsgProvideLiquidity{
		Creator:      creator,
		StableCoinId: stableCoinId,
		Amount:       amount,
	}
}

func (msg *MsgProvideLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = strconv.Atoi(msg.Amount)

	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount value (%s)", err)
	}

	return nil
}
