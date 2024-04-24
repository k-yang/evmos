// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package transfer

import (
	ibccallbacks "github.com/cosmos/ibc-go/modules/apps/callbacks/types"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	"github.com/evmos/evmos/v17/x/ibc/transfer/keeper"
)

var _ ibccallbacks.CallbacksCompatibleModule = IBCModule{}

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	*ibctransfer.IBCModule
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	transferModule := ibctransfer.NewIBCModule(*k.Keeper)
	return IBCModule{
		IBCModule: &transferModule,
	}
}
