package ethtodharitri

const (
	// GettingPendingBatchFromPeerChain is the step identifier for fetching the pending batch from the PeerChain chain
	GettingPendingBatchFromPeerChain = "get pending batch from peer chain"

	// ProposingTransferOnDharitrI is the step identifier for proposing transfer on DharitrI
	ProposingTransferOnDharitrI = "propose transfer"

	// SigningProposedTransferOnDharitrI is the step identifier for signing proposed transfer
	SigningProposedTransferOnDharitrI = "sign proposed transfer"

	// WaitingForQuorum is the step identifier for waiting until the quorum is reached
	WaitingForQuorum = "wait for quorum"

	// PerformingActionID is the step identifier for performing the ActionID on DharitrI
	PerformingActionID = "perform action"

	// NumSteps indicates how many steps the state machine for PeerChain -> DharitrI flow has
	NumSteps = 5
)
