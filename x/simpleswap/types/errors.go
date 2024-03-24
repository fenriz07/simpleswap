package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/simpleswap module sentinel errors
var (
	ErrInvalidSigner             = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                    = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrStableCoinNotFound        = sdkerrors.Register(ModuleName, 1102, "stable coin not found")
	ErrStableIsNotAvailable      = sdkerrors.Register(ModuleName, 1103, "stable coin is not available")
	ErrInvalidAccount            = sdkerrors.Register(ModuleName, 1104, "invalid account")
	ErrAmoutHasToBeAnInt         = sdkerrors.Register(ModuleName, 1105, "amount has to be an int")
	ErrACoinIsValidate           = sdkerrors.Register(ModuleName, 1106, "just a coin is validate")
	ErrSendCoin                  = sdkerrors.Register(ModuleName, 1107, "error sending coins")
	ErrDoesNotHasBalance         = sdkerrors.Register(ModuleName, 1108, "error account does not has balance")
	ErrLiquidityProviderNotFound = sdkerrors.Register(ModuleName, 1109, "error liquidity provider not found")
	ErrParsingStringToInt        = sdkerrors.Register(ModuleName, 1110, "error parsing string to int")
)
