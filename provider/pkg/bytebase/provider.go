// File: provider/pkg/bytebase/provider.go
package provider

//go:generate oapi-codegen -generate types,client -o ./client/bytebase.gen.go -package client ./swagger/bytebase-swagger.yaml

import (
	"strings"

	"github.com/IrisDande/pulumi-bytebase/provider/pkg/bytebase/index"
	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[*index.bytebaseIndex, index.bytebaseIndexArgs, index.bytebaseIndexState](),
			infer.Resource[*index.bytebaseCollection, index.bytebaseCollectionArgs, index.bytebaseCollectionState](),
		},
		Functions: []infer.InferredFunction{
			infer.Function[*index.LookupbytebaseIndex, index.LookupbytebaseIndexArgs, index.LookupbytebaseIndexResult](),
			infer.Function[*index.LookupbytebaseCollection, index.LookupbytebaseCollectionArgs, index.LookupbytebaseCollectionResult](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"bytebase": "index",
		},
		Metadata: schema.Metadata{
			DisplayName: "bytebase",
			Description: "A Pulumi native provider for bytebase",
			Keywords: []string{
				"pulumi",
				"bytebase",
				"category/utility",
				"kind/native",
			},
			Homepage:          "https://www.bytebase.io",
			License:           "Apache-2.0",
			Repository:        "https://github.com/IrisDande/pulumi-bytebase",
			PluginDownloadURL: "github://api.github.com/IrisDande/pulumi-bytebase",
			Publisher:         "bytebase-io",
			LogoURL:           "",
			LanguageMap: map[string]any{
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/IrisDande/pulumi-bytebase/sdk/go/bytebase",
				},
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
					"rootNamespace": "bytebaseDatabase",
				},
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
					"packageName": "@bytebase-database/pulumi",
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					"packageName": "bytebase_pulumi",
				},
			},
		},
		Config: infer.Config[*config.bytebaseProviderConfig](),
	})
}

func Schema(version string) (string, error) {
	version = strings.TrimPrefix(version, "v")
	s, err := integration.NewServer("bytebase", semver.MustParse(version), Provider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
