package config

import (
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type bytebaseProviderConfig struct {
	APIKey string `pulumi:"APIKey,optional" provider:"secret"`
}

func (c *bytebaseProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.APIKey, "The API token for bytebase.")
}

func (c *bytebaseProviderConfig) Configure(ctx p.Context) error {
	ctx.Logf(diag.Debug, "Configuring bytebase provider")
	if c.APIKey == "" {
		APIKey, exists := os.LookupEnv("bytebase_API_KEY")
		if exists {
			c.APIKey = APIKey
			return nil
		}
		return fmt.Errorf("API key is required")
	}
	return nil
}
