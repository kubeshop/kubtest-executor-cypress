package main

import (
	"fmt"
	"os"

	"github.com/kubeshop/testkube-executor-cypress/pkg/runner"
	"github.com/kubeshop/testkube/pkg/executor/agent"
	"github.com/kubeshop/testkube/pkg/executor/output"
)

func main() {
	r, err := runner.NewCypressRunner(os.Getenv("DEPENDENCY_MANAGER"))
	if err != nil {
		output.PrintError(os.Stderr, fmt.Errorf("could not initialize runner: %w", err))
		os.Exit(1)
	}

	if r.Params.CloudMode {
		defer r.GRPCConn.Close()
	}

	agent.Run(r, os.Args)
}
