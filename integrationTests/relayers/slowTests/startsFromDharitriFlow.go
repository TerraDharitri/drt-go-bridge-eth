//go:build slow

package slowTests

import (
	"fmt"
	"testing"

	"github.com/TerraDharitri/drt-go-bridge-eth/integrationTests/relayers/slowTests/framework"
)

type startsFromDharitrIFlow struct {
	testing.TB
	setup        *framework.TestSetup
	ethToDrtDone bool
	drtToEthDone bool
	tokens       []framework.TestTokenParams
}

func (flow *startsFromDharitrIFlow) process() (finished bool) {
	if len(flow.tokens) == 0 {
		return true
	}
	if flow.drtToEthDone && flow.ethToDrtDone {
		return true
	}

	isTransferDoneFromDharitrI := flow.setup.IsTransferDoneFromDharitrI(flow.tokens...)
	if !flow.drtToEthDone && isTransferDoneFromDharitrI {
		flow.drtToEthDone = true
		log.Info(fmt.Sprintf(framework.LogStepMarker, "DharitrI->PeerChain transfer finished, now sending back to DharitrI..."))

		flow.setup.PeerChainHandler.SendFromPeerChainToDharitrI(flow.setup.Ctx, flow.setup.DharitriHandler.TestCallerAddress, flow.tokens...)
	}
	if !flow.drtToEthDone {
		// return here, no reason to check downwards
		return false
	}

	isTransferDoneFromPeerChain := flow.setup.IsTransferDoneFromPeerChain(flow.tokens...)
	if !flow.ethToDrtDone && isTransferDoneFromPeerChain {
		flow.ethToDrtDone = true
		log.Info(fmt.Sprintf(framework.LogStepMarker, "DharitrI<->PeerChain from DharitrI transfers done"))
		return true
	}

	return false
}

func (flow *startsFromDharitrIFlow) areTokensFullyRefunded() bool {
	if len(flow.tokens) == 0 {
		return true
	}
	if !flow.ethToDrtDone {
		return false // regular flow is not completed
	}

	return flow.setup.IsTransferDoneFromPeerChainWithRefund(flow.tokens...)
}
