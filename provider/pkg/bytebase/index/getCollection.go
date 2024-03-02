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

type LookupbytebaseCollection struct{}

func (g *LookupbytebaseCollection) Annotate(a infer.Annotator) {
	a.Describe(&g, "The result of a get operation on a bytebase collection.")
}

func (*LookupbytebaseCollection) Call(ctx p.Context, args LookupbytebaseCollectionArgs) (LookupbytebaseCollectionResult, error) {
	bytebaseConfig := infer.GetConfig[config.bytebaseProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    bytebaseConfig.APIKey,
		},
	}
	bytebaseClient, err := client.NewClientWithResponses("https://api.bytebase.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return LookupbytebaseCollectionResult{}, err
	}
	resp, err := bytebaseClient.DescribeCollectionWithResponse(ctx, args.CollectionName)
	ctx.Logf(diag.Debug, "DescribeCollectionWithResponse: %v", resp.Status())
	if err != nil {
		ctx.Logf(diag.Error, "DescribeCollectionWithResponse: %v", resp.Status())
		return LookupbytebaseCollectionResult{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return LookupbytebaseCollectionResult{}, fmt.Errorf("DescribeCollectionWithResponse: %v", resp.Status())
	}
	return LookupbytebaseCollectionResult{
		bytebaseCollectionState: bytebaseCollectionState{
			bytebaseCollectionArgs: bytebaseCollectionArgs{
				CollectionName: resp.JSON200.Name,
			},
			CollectionSize:        resp.JSON200.Size,
			CollectionDimension:   resp.JSON200.Dimension,
			CollectionVectorCount: resp.JSON200.VectorCount,
			CollectionEnvironment: resp.JSON200.Environment,
		},
	}, nil
}

type LookupbytebaseCollectionArgs struct {
	CollectionName string `pulumi:"name"`
}

func (g *LookupbytebaseCollectionArgs) Annotate(a infer.Annotator) {
	a.Describe(&g.CollectionName, "The name of the bytebase collection.")
}

type LookupbytebaseCollectionResult struct {
	bytebaseCollectionState
}

func (g *LookupbytebaseCollectionResult) Annotate(a infer.Annotator) {
	a.Describe(&g.bytebaseCollectionArgs, "The result of a get operation on a bytebase collection.")
}
