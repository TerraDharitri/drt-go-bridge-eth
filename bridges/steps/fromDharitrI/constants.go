package dharitritoeth

const (
	// GettingPendingBatchFromDharitrI is the step identifier for fetching the pending batch from the DharitrI chain
	GettingPendingBatchFromDharitrI = "get pending batch from DharitrI"

	// SigningProposedTransferOnPeerChain is the step identifier for signing proposed transfer
	SigningProposedTransferOnPeerChain = "sign proposed transfer"

	// WaitingForQuorumOnTransfer is the step identifier for waiting until the quorum is reached
	WaitingForQuorumOnTransfer = "wait for quorum on transfer"

	// PerformingTransfer is the step identifier for performing the transfer on PeerChain
	PerformingTransfer = "perform transfer"

	// WaitingTransferConfirmation is the step identifier for waiting the transfer confirmation on PeerChain
	WaitingTransferConfirmation = "wait transfer confirmating"

	// ResolvingSetStatusOnDharitrI is the step identifier for resolving set status on DharitrI
	ResolvingSetStatusOnDharitrI = "resolve set status"

	// ProposingSetStatusOnDharitrI is the step idetifier for proposing set status action on DharitrI
	ProposingSetStatusOnDharitrI = "propose set status"

	// SigningProposedSetStatusOnDharitrI is the step identifier for signing proposed set status action
	SigningProposedSetStatusOnDharitrI = "sign proposed set status"

	// WaitingForQuorumOnSetStatus is the step identifier for waiting until the quorum is reached
	WaitingForQuorumOnSetStatus = "wait for quorum on set status"

	// PerformingSetStatus is the step identifier for performing the set status action on DharitrI
	PerformingSetStatus = "perform set status"

	// NumSteps indicates how many steps the state machine for DharitrI -> PeerChain flow has
	NumSteps = 10
)
