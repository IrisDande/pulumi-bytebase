// File: provider/pkg/bytebase/index/index.go
package index

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IrisDande/pulumi-bytebase/provider/pkg/bytebase/client"
	"github.com/IrisDande/pulumi-bytebase/provider/pkg/bytebase/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type bytebaseIndex struct{}

type ServerlessSpecCloud string

type IndexMetric string

type bytebaseServerlessSpec struct {
	Cloud  ServerlessSpecCloud `pulumi:"cloud"`
	Region string              `pulumi:"region"`
}

type PodSpecPodType = string

type PodSpecReplicas = int32

type PodSpecShards = int32

type MetaDataConfig struct {
	Indexed *[]string `pulumi:"indexed,optional"`
}

type bytebasePodSpec struct {
	Environment      string          `pulumi:"environment"`
	Replicas         PodSpecReplicas `pulumi:"replicas"`
	Shards           PodSpecShards   `pulumi:"shards,optional,omitempty"`
	PodType          PodSpecPodType  `pulumi:"podType"`
	Pods             int             `pulumi:"pods,optional,omitempty"`
	MetaDataConfig   MetaDataConfig  `pulumi:"metaDataConfig,optional"`
	SourceCollection *string         `pulumi:"sourceCollection,optional"`
}

type bytebaseSpec struct {
	Serverless bytebaseServerlessSpec `pulumi:"serverless,optional,omitempty"`
	Pod        bytebasePodSpec        `pulumi:"pod,optional,omitempty"`
}

// bytebaseIndexArgs describes the configuration options available for creating or managing a bytebase index.
type bytebaseIndexArgs struct {
	/**
	 * IndexName specifies the unique name for the bytebase index. This is a mandatory field and is used to
	 * identify the index in the bytebase environment.
	 */
	IndexName string `pulumi:"name"`

	/**
	 * IndexDimension specifies the dimensions of the vectors that the index will store. It is optional and, if
	 * not provided, a default dimension may be used based on the bytebase environment settings.
	 *
	 * Defaults to 1536.
	 */
	IndexDimension client.IndexDimension `pulumi:"dimension,omitempty,optional"`

	/**
	 * IndexMetric determines the type of metric to be used for measuring distances in the vector space. This
	 * could be, for example, Euclidean or cosine distance, depending on the IndexMetric value.
	 */
	IndexMetric IndexMetric `pulumi:"metric"`

	/**
	 * IndexSpec defines the specific configuration and settings for the bytebase index.
	 */
	IndexSpec bytebaseSpec `pulumi:"spec"`
}

func (ServerlessSpecCloud) Values() []infer.EnumValue[ServerlessSpecCloud] {
	return []infer.EnumValue[ServerlessSpecCloud]{
		{Name: string(client.Aws), Value: ServerlessSpecCloud(client.Aws)},
		{Name: string(client.Azure), Value: ServerlessSpecCloud(client.Azure)},
		{Name: string(client.Gcp), Value: ServerlessSpecCloud(client.Gcp)},
	}
}

func (IndexMetric) Values() []infer.EnumValue[IndexMetric] {
	return []infer.EnumValue[IndexMetric]{
		{Name: string(client.Dotproduct), Value: IndexMetric(client.Dotproduct)},
		{Name: string(client.Cosine), Value: IndexMetric(client.Cosine)},
		{Name: string(client.Euclidean), Value: IndexMetric(client.Euclidean)},
	}
}

type bytebaseIndexState struct {
	bytebaseIndexArgs
	IndexHost string `pulumi:"host,omitempty"`
}

func (pip *bytebaseIndexState) Annotate(a infer.Annotator) {
	a.Describe(&pip.IndexHost, "The URL address where the index is hosted.")
}

func (pip *bytebaseIndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&pip.IndexName, "The name of the bytebase index.")
	a.Describe(&pip.IndexDimension, "The dimensions of the vectors in the index. Defaults to 1536.")
	a.Describe(&pip.IndexMetric, "The metric used to compute the distance between vectors.")
	a.Describe(&pip.IndexSpec, "Describe how the index should be deployed.")
}

func (pip *bytebaseServerlessSpec) Annotate(a infer.Annotator) {
	a.Describe(&pip.Cloud, "The public cloud where you would like your index hosted.")
	a.Describe(&pip.Region, "The region where you would like your index to be created. Different cloud "+
		"providers have different regions available.")
}

func (pip *bytebaseSpec) Annotate(a infer.Annotator) {
	a.Describe(&pip.Serverless, "Configuration needed to deploy a serverless index.")
	a.Describe(&pip.Pod, "Configuration needed to deploy a pod index.")
}

func (pip *bytebasePodSpec) Annotate(a infer.Annotator) {
	a.Describe(&pip.Environment, "The environment where the index is hosted.")
	a.Describe(&pip.Replicas, "The number of replicas. Replicas duplicate your index. They provide higher"+
		" availability and throughput. Replicas can be scaled up or down as your needs change.")
	a.Describe(&pip.Shards, "The number of shards. Shards split your data across multiple pods so you can"+
		" fit more data into an index.")
	a.Describe(&pip.PodType, "The type of pod to use. One of `s1`, `p1`, or `p2` appended with `.` and one"+
		" of `x1`, `x2`, `x4`, or `x8`.")
	a.Describe(&pip.Pods, "The number of pods to be used in the index. This should be equal to "+
		"`shards` x `replicas`.")
	a.Describe(&pip.MetaDataConfig, "Configuration for the behavior of bytebase's internal metadata index.")
	a.Describe(&pip.SourceCollection, "The name of the collection to be used as the source for the index.")
}

func (pim *MetaDataConfig) Annotate(a infer.Annotator) {
	a.Describe(&pim.Indexed, " Indexed By default, all metadata is indexed; to change this behavior, use this"+
		" property to specify an array of metadata fields which should be indexed.")
}

func (*bytebaseIndex) Create(ctx p.Context, name string, args bytebaseIndexArgs, preview bool) (string, bytebaseIndexState, error) {
	bytebaseConfig := infer.GetConfig[config.bytebaseProviderConfig](ctx)
	if err := utils.ValidateIndexName(args.IndexName); err != nil {
		return "", bytebaseIndexState{}, fmt.Errorf("invalid index name: %w", err)
	}

	indexDimension := args.IndexDimension
	if indexDimension == 0 {
		indexDimension = utils.IndexDimensionDefault
	}
	ctx.Logf(diag.Debug, "bytebase index dimension: %d", indexDimension)

	if preview {
		ctx.Logf(diag.Debug, "Previewing bytebase index creation: %s", args.IndexName)
		return args.IndexName, bytebaseIndexState{
			bytebaseIndexArgs: bytebaseIndexArgs{
				IndexName:      args.IndexName,
				IndexMetric:    args.IndexMetric,
				IndexDimension: indexDimension,
				IndexSpec:      args.IndexSpec,
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
		return "", bytebaseIndexState{}, fmt.Errorf("failed to create bytebase client: %w", err)
	}

	ctx.Logf(diag.Debug, "Creating bytebase index: %s", args.IndexName)

	var spec client.CreateIndexRequest_Spec

	if args.IndexSpec.Serverless != (bytebaseServerlessSpec{}) {
		ctx.Logf(diag.Debug, "Creating bytebase serverless index: %s", args.IndexName)
		spec = client.CreateIndexRequest_Spec{
			Serverless: &client.ServerlessSpec{
				Cloud:  client.ServerlessSpecCloud(args.IndexSpec.Serverless.Cloud),
				Region: args.IndexSpec.Serverless.Region,
			},
		}
	} else if args.IndexSpec.Pod != (bytebasePodSpec{}) {
		ctx.Logf(diag.Debug, "Creating bytebase pod-based index: %s", args.IndexName)
		spec = client.CreateIndexRequest_Spec{
			Pod: &client.PodSpec{
				Environment: args.IndexSpec.Pod.Environment,
				MetadataConfig: &struct {
					Indexed *[]string `json:"indexed,omitempty"`
				}{
					Indexed: args.IndexSpec.Pod.MetaDataConfig.Indexed,
				},
				Pods:             args.IndexSpec.Pod.Pods,
				PodType:          args.IndexSpec.Pod.PodType,
				Replicas:         args.IndexSpec.Pod.Replicas,
				Shards:           args.IndexSpec.Pod.Shards,
				SourceCollection: args.IndexSpec.Pod.SourceCollection,
			},
		}
	}

	response, err := bytebaseClient.CreateIndexWithResponse(context.Background(), client.CreateIndexJSONRequestBody{
		Dimension: indexDimension,
		Metric:    client.IndexMetric(args.IndexMetric),
		Name:      args.IndexName,
		Spec:      spec,
	})
	if err != nil {
		ctx.Logf(diag.Error, "Failed to create bytebase index: %s with http status code: %d", args.IndexName, response.StatusCode())
		ctx.Logf(diag.Error, "Please run the Pulumi command with the `-d` flag to see the full error message")
		return "", bytebaseIndexState{}, fmt.Errorf("failed to create bytebase index: %w", err)
	}
	ctx.Logf(diag.Debug, "bytebase index creation response: %s", string(response.Body))

	ready := false
	for !ready {
		ctx.Logf(diag.Debug, "Waiting for bytebase index: %s to be ready", args.IndexName)
		response, err := bytebaseClient.DescribeIndexWithResponse(context.Background(), args.IndexName)
		if err != nil {
			ctx.Logf(diag.Error, "Failed to get bytebase index: %s with http error code: %d", args.IndexName, response.StatusCode())
			return "", bytebaseIndexState{}, fmt.Errorf("failed to get bytebase index: %w", err)
		}
		if response.StatusCode() != http.StatusOK {
			ctx.Logf(diag.Error, "Failed to get bytebase index: %s with http error code: %d", args.IndexName, response.StatusCode())
			return "", bytebaseIndexState{}, fmt.Errorf("failed to get bytebase index: %s", args.IndexName)
		}
		if response.JSON200.Status.Ready {
			ready = true
		}
	}

	return args.IndexName, bytebaseIndexState{
		bytebaseIndexArgs{
			IndexName:      args.IndexName,
			IndexMetric:    args.IndexMetric,
			IndexDimension: args.IndexDimension,
			IndexSpec:      args.IndexSpec,
		},
		response.JSON201.Host,
	}, nil
}

func (pi *bytebaseIndex) Delete(ctx p.Context, id string, state bytebaseIndexState) error {
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

	response, err := bytebaseClient.DeleteIndexWithResponse(context.Background(), state.IndexName)
	if err != nil {
		ctx.Logf(diag.Error, "Failed to delete bytebase index: %s with http error code: %d", state.IndexName, response.StatusCode())
		return fmt.Errorf("error deleting bytebase index '%s': %w", state.IndexName, err)
	}
	ctx.Logf(diag.Debug, "Successfully deleted bytebase index: %s", state.IndexName)
	return nil
}

func (pi *bytebaseIndex) Read(ctx p.Context, id string, args bytebaseIndexArgs, state bytebaseIndexState) (canonicalID string, normalizedInputs bytebaseIndexArgs, normalizedState bytebaseIndexState, err error) {
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

	indexDetails, err := bytebaseClient.DescribeIndexWithResponse(context.Background(), state.IndexName)
	if err != nil {
		if indexDetails.JSON404 != nil {
			ctx.Logf(diag.Debug, "bytebase index '%s' not found", state.IndexName)
			return id, args, state, nil
		}
		return id, args, state, fmt.Errorf("error getting bytebase index details '%s': %w", state.IndexName, err)
	}

	state.IndexName = indexDetails.JSON200.Name
	state.IndexDimension = indexDetails.JSON200.Dimension
	state.IndexMetric = IndexMetric(indexDetails.JSON200.Metric)
	state.IndexHost = indexDetails.JSON200.Host
	return id, args, state, nil
}
