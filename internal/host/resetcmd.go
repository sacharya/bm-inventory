package host

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/filanov/bm-inventory/models"
)

type resetCmd struct {
	baseCmd
}

func NewResetCmd(log logrus.FieldLogger) *resetCmd {
	return &resetCmd{
		baseCmd: baseCmd{log: log},
	}
}

func (h *resetCmd) GetStep(ctx context.Context, host *models.Host) (*models.Step, error) {
	step := &models.Step{
		StepType: models.StepTypeReset,
		Command:  "systemctl",
		Args: []string{
			"restart", "agent",
		},
	}
	return step, nil
}
