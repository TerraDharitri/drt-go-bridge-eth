//go:build slow

package slowTests

import (
	"fmt"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/relayers/slowTests/framework"
)

type startsFromEthereumEdgecaseFlow struct {
	testing.TB
	setup        *framework.TestSetup
	ethToDrtDone bool
	drtToEthDone bool
	tokens       []framework.TestTokenParams
}

func (flow *startsFromEthereumEdgecaseFlow) process() (finished bool) {
	if len(flow.tokens) == 0 {
		return true
	}
	if flow.drtToEthDone && flow.ethToDrtDone {
		return true
	}

	isTransferDoneFromEthereum := flow.setup.IsTransferDoneFromPeerChain(flow.tokens...)
	if !flow.ethToDrtDone && isTransferDoneFromEthereum {
		flow.ethToDrtDone = true
		log.Info(fmt.Sprintf(framework.LogStepMarker, "Ethereum->DharitrI transfer finished, now sending back to Ethereum & another round from Ethereum..."))

		flow.setup.SendFromDharitriToPeerChain(flow.tokens...)
		flow.setup.PeerChainHandler.SendFromPeerChainToDharitrI(flow.setup.Ctx, flow.setup.DharitriHandler.TestCallerAddress, flow.tokens...)
	}
	if !flow.ethToDrtDone {
		// return here, no reason to check downwards
		return false
	}

	isTransferDoneFromDharitrI := flow.setup.IsTransferDoneFromDharitrI(flow.tokens...)
	if !flow.drtToEthDone && isTransferDoneFromDharitrI {
		flow.drtToEthDone = true
		log.Info(fmt.Sprintf(framework.LogStepMarker, "DharitrI<->Ethereum from Ethereum transfers done"))
		return true
	}

	return false
}
