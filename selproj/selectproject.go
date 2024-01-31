// Package main is the main package of the selproj tool.
package main

import (
	"context"
	"errors"
	"strings"
)

// errNoSuitableProjectFound is the error returned when no suitable project is found.
var errNoSuitableProjectFound = errors.New("no suitable project found")

// selectProject selects a project with the given prefix.
func selectProject(ctx context.Context, client AivenClient, prefix string) (string, error) {
	projects, err := client.ProjectList(ctx)
	if err != nil {
		return "", err
	}

	for _, project := range projects {
		if !strings.HasPrefix(project.Name, prefix) {
			continue
		}

		services, err := client.ServiceList(ctx, project.Name)
		if err != nil {
			return "", err
		}

		if len(services) != 0 {
			continue
		}

		vpcs, err := client.VPCList(ctx, project.Name)
		if err != nil {
			return "", err
		}

		if len(vpcs) != 0 {
			continue
		}

		return strings.TrimPrefix(project.Name, prefix), nil
	}

	return "", errNoSuitableProjectFound
}
