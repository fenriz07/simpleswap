package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimLiquidity{}

func NewMsgClaimLiquidity(creator string, providerLiquidityId string) *MsgClaimLiquidity {
	return &MsgClaimLiquidity{
		Creator:             creator,
		ProviderLiquidityId: providerLiquidityId,
	}
}

func (msg *MsgClaimLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
