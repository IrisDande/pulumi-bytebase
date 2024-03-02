package main

import (
	"github.com/IrisDande/pulumi-bytebase/sdk/go/bytebase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		mybytebaseIndex, err := bytebase.NewbytebaseIndex(ctx, "mybytebaseIndex", &bytebase.bytebaseIndexArgs{
			Name:   pulumi.String("example-index2"),
			Metric: bytebase.IndexMetricCosine,
			Spec: &bytebase.bytebaseSpecArgs{
				Serverless: &bytebase.bytebaseServerlessSpecArgs{
					Cloud:  bytebase.ServerlessSpecCloudAws,
					Region: pulumi.String("us-west-2"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("mybytebaseIndex", mybytebaseIndex.Name)
		ctx.Export("mybytebaseIndexHost", mybytebaseIndex.Host)

		mybytebaseIndex2, err := bytebase.NewbytebaseIndex(ctx, "mybytebaseIndex2", &bytebase.bytebaseIndexArgs{
			Name:      pulumi.String("example-index3"),
			Dimension: pulumi.Int(1536),
			Metric:    bytebase.IndexMetricCosine,
			Spec: &bytebase.bytebaseSpecArgs{
				Pod: &bytebase.bytebasePodSpecArgs{
					Environment: pulumi.String("gcp-starter"),
					MetaDataConfig: &bytebase.MetaDataConfigArgs{
						Indexed: pulumi.StringArray{
							pulumi.String("genre"),
							pulumi.String("title"),
							pulumi.String("imdb_rating"),
						},
					},
					Shards:   pulumi.Int(1),
					Replicas: pulumi.Int(1),
					Pods:     pulumi.Int(1),
					PodType:  pulumi.String("p1.x1"),
					//SourceCollection: pulumi.String("movie-embeddings"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("mybytebaseIndex2", mybytebaseIndex2.Host)

		mybytebaseIndex3, err := bytebase.NewbytebaseIndex(ctx, "mybytebaseIndex3", &bytebase.bytebaseIndexArgs{
			Name:      pulumi.String("example-index4"),
			Dimension: pulumi.Int(1536),
			Metric:    bytebase.IndexMetricCosine,
			Spec: &bytebase.bytebaseSpecArgs{
				Pod: &bytebase.bytebasePodSpecArgs{
					Environment: pulumi.String("us-east-1-aws"),
					MetaDataConfig: &bytebase.MetaDataConfigArgs{
						Indexed: pulumi.StringArray{
							pulumi.String("genre"),
							pulumi.String("title"),
							pulumi.String("imdb_rating"),
						},
					},
					Shards:   pulumi.Int(1),
					Replicas: pulumi.Int(1),
					Pods:     pulumi.Int(1),
					PodType:  pulumi.String("p1.x1"),
					//SourceCollection: pulumi.String("movie-embeddings"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("mybytebaseIndex3", mybytebaseIndex3.Host)

		return nil
	})
}
