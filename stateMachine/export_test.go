package stateMachine

import "github.com/TerraDharitri/drt-go-bridge-eth/core"

// GetCurrentStep -
func (sm *stateMachine) GetCurrentStepIdentifier() core.StepIdentifier {
	return sm.currentStep.Identifier()
}
