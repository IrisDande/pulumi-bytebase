package index

import (
	"fmt"
	"net/http"

	"github.com/IrisDande/pulumi-bytebase/provider/pkg/bytebase/client"
	"github.com/IrisDande/pulumi-bytebase/provider/pkg/bytebase/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type LookupbytebaseIndex struct{}

func (*LookupbytebaseIndex) Call(ctx p.Context, args LookupbytebaseIndexArgs) (LookupbytebaseIndexResult, error) {
	bytebaseConfig := infer.GetConfig[config.bytebaseProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    bytebaseConfig.APIKey,
		},
	}
	bytebaseClient, err := client.NewClientWithResponses("https://api.bytebase.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return LookupbytebaseIndexResult{}, err
	}
	resp, err := bytebaseClient.DescribeIndexWithResponse(ctx, args.IndexName)
	ctx.Logf(diag.Debug, "DescribeIndexWithResponse: %v", resp.Status())
	if err != nil {
		ctx.Logf(diag.Error, "DescribeIndexWithResponse: %v", resp.Status())
		return LookupbytebaseIndexResult{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return LookupbytebaseIndexResult{}, fmt.Errorf("DescribeIndexWithResponse: %v", resp.Status())
	}
	return LookupbytebaseIndexResult{
		bytebaseIndexArgs: bytebaseIndexArgs{
			IndexName:      resp.JSON200.Name,
			IndexMetric:    IndexMetric(resp.JSON200.Metric),
			IndexDimension: resp.JSON200.Dimension,
			IndexSpec: bytebaseSpec{
				Serverless: bytebaseServerlessSpec{
					Cloud:  ServerlessSpecCloud(resp.JSON200.Spec.Serverless.Cloud),
					Region: resp.JSON200.Spec.Serverless.Region,
				},
			},
		},
		IndextStatus: resp.JSON200.Status.Ready,
		IndexHost:    resp.JSON200.Host,
	}, nil
}

type LookupbytebaseIndexArgs struct {
	IndexName string `pulumi:"name"`
}

func (g *LookupbytebaseIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&g.IndexName, "The name of the bytebase index.")
}

type LookupbytebaseIndexResult struct {
	bytebaseIndexArgs
	IndexHost    string `pulumi:"host,omitempty"`
	IndextStatus bool   `pulumi:"status,omitempty"`
}

func (g *LookupbytebaseIndexResult) Annotate(a infer.Annotator) {
	a.Describe(&g.IndexSpec, "Describe how the index should be deployed.")
	a.Describe(&g.IndexHost, "The host of the index.")
	a.Describe(&g.IndextStatus, "The status of the index.")
}
