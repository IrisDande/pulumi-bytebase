// File: provider/cmd/pulumi-resource-bytebase/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	bytebase "github.com/IrisDande/pulumi-bytebase/provider/pkg/bytebase"
	"github.com/IrisDande/pulumi-bytebase/provider/pkg/version"
	p "github.com/pulumi/pulumi-go-provider"
)

// Serve the provider against Pulumi's Provider protocol.
func main() {
	trimmedVersion := strings.TrimPrefix(version.Version, "v")
	err := p.RunProvider("bytebase", trimmedVersion, bytebase.Provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
