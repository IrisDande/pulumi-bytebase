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

type bytebaseCollection struct{}

type bytebaseCollectionArgs struct {
	CollectionName   string `pulumi:"name"`
	CollectionSource string `pulumi:"source"`
}

func (pc *bytebaseCollectionArgs) Annotate(a infer.Annotator) {
	a.Describe(&pc.CollectionName, "The name of the collection to be created.")
	a.Describe(&pc.CollectionSource, "The name of the index to be used as the source for the collection.")
}

func (*bytebaseCollection) Create(ctx p.Context, name string, args bytebaseCollectionArgs, preview bool) (string, bytebaseCollectionState, error) {
	bytebaseConfig := infer.GetConfig[config.bytebaseProviderConfig](ctx)

	if preview {
		ctx.Logf(diag.Debug, "Creating bytebase collection %s", args.CollectionName)
		return "", bytebaseCollectionState{
			bytebaseCollectionArgs: bytebaseCollectionArgs{
				CollectionName:   args.CollectionName,
				CollectionSource: args.CollectionSource,
			},
		}, nil
	}
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    bytebaseConfig.APIKey,
		},
	}

	bytebaseClient, err := client.NewClientWithResponses("https://api.bytebase.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return "", bytebaseCollectionState{}, fmt.Errorf("failed to create bytebase client: %w", err)
	}
	ctx.Logf(diag.Debug, "Creating bytebase collection %s", args.CollectionName)
	resp, err := bytebaseClient.CreateCollectionWithResponse(ctx, client.CreateCollectionJSONRequestBody{
		Name:   args.CollectionName,
		Source: args.CollectionSource,
	})
	if err != nil {
		ctx.Logf(diag.Debug, "Failed to create bytebase collection %s with http status code %d", args.CollectionName, resp.StatusCode())
		return "", bytebaseCollectionState{}, fmt.Errorf("failed to create bytebase collection: %w", err)
	}
	ctx.Logf(diag.Debug, "bytebase collection creaation responese: %s", string(resp.Body))

	ready := false
	for !ready {
		resp, err := bytebaseClient.DescribeCollectionWithResponse(ctx, args.CollectionName)
		if err != nil {
			return "", bytebaseCollectionState{}, fmt.Errorf("failed to describe bytebase collection: %w", err)
		}
		if resp.JSON200.Status == client.CollectionModelStatusReady {
			ready = true
		}
	}
	return args.CollectionName, bytebaseCollectionState{
		bytebaseCollectionArgs: bytebaseCollectionArgs{
			CollectionName:   args.CollectionName,
			CollectionSource: args.CollectionSource,
		},
		CollectionDimension:   resp.JSON201.Dimension,
		CollectionSize:        resp.JSON201.Size,
		CollectionVectorCount: resp.JSON201.VectorCount,
		CollectionEnvironment: resp.JSON201.Environment,
	}, nil
}

func (*bytebaseCollection) Delete(ctx p.Context, id string, args bytebaseCollectionArgs) error {
	bytebaseConfig := infer.GetConfig[config.bytebaseProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    bytebaseConfig.APIKey,
		},
	}
	bytebaseClient, err := client.NewClientWithResponses("https://api.bytebase.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return fmt.Errorf("failed to create bytebase client: %w", err)
	}
	response, err := bytebaseClient.DeleteCollectionWithResponse(ctx, args.CollectionName)
	if err != nil {
		ctx.Logf(diag.Error, "Failed to delete bytebase collection: %s with http status code: %d", args.CollectionName, response.StatusCode())
		return fmt.Errorf("failed to delete bytebase collection: %w", err)
	}
	ctx.Logf(diag.Debug, "Successfully deleted bytebase collection: %s", args.CollectionName)
	return nil
}

func (*bytebaseCollection) Read(ctx p.Context, id string, args bytebaseCollectionArgs, state bytebaseCollectionState) (canonicalID string, normalizedInputs bytebaseCollectionArgs, normalizedState bytebaseCollectionState, err error) {
	bytebaseConfig := infer.GetConfig[config.bytebaseProviderConfig](ctx)
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			APIKey:    bytebaseConfig.APIKey,
		},
	}
	bytebaseClient, err := client.NewClientWithResponses("https://api.bytebase.io", client.WithHTTPClient(httpClient))
	if err != nil {
		return id, args, state, fmt.Errorf("failed to create bytebase client: %w", err)
	}

	resp, err := bytebaseClient.DescribeCollectionWithResponse(ctx, args.CollectionName)
	if err != nil {
		if resp.JSON404 != nil {
			ctx.Logf(diag.Debug, "bytebase collection '%s' not found", args.CollectionName)
			return id, args, state, nil
		}
		return id, args, state, fmt.Errorf("error getting bytebase collection details '%s': %w", args.CollectionName, err)
	}

	state.CollectionName = resp.JSON200.Name
	state.CollectionDimension = resp.JSON200.Dimension
	state.CollectionSize = resp.JSON200.Size
	state.CollectionVectorCount = resp.JSON200.VectorCount
	state.CollectionEnvironment = resp.JSON200.Environment

	return id, args, state, nil
}

type bytebaseCollectionState struct {
	bytebaseCollectionArgs
	CollectionSize        *int64  `pulumi:"size"`
	CollectionDimension   *int32  `pulumi:"dimension"`
	CollectionVectorCount *int32  `pulumi:"vectorCount"`
	CollectionEnvironment *string `pulumi:"environment"`
}

func (pcs *bytebaseCollectionState) Annotate(a infer.Annotator) {
	a.Describe(&pcs.CollectionSize, "The size of the collection in bytes.")
	a.Describe(&pcs.CollectionDimension, "The dimension of the vectors stored in each record held in the collection.")
	a.Describe(&pcs.CollectionVectorCount, "The number of records stored in the collection.")
	a.Describe(&pcs.CollectionEnvironment, "The environment where the collection is hosted.")
}
