package utils

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

func Execute(format string, args ...any) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Minute)
	defer cancel()
	cmd := fmt.Sprintf(format, args...)
	output, err := exec.CommandContext(ctx, "/bin/sh", "-ce", cmd).CombinedOutput()
	if err != nil {
		return "", errors.Errorf("error executing %s: %s: %s", cmd, err, string(output))
	}
	return string(output), nil
}
