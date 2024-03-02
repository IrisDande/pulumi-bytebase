// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bytebase

import (
	"context"
	"reflect"

	"errors"
	"github.com/IrisDande/pulumi-bytebase/sdk/go/bytebase/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type bytebaseCollection struct {
	pulumi.CustomResourceState

	// The dimension of the vectors stored in each record held in the collection.
	Dimension pulumi.IntOutput `pulumi:"dimension"`
	// The environment where the collection is hosted.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The name of the collection to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The size of the collection in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// The name of the index to be used as the source for the collection.
	Source pulumi.StringOutput `pulumi:"source"`
	// The number of records stored in the collection.
	VectorCount pulumi.IntOutput `pulumi:"vectorCount"`
}

// NewbytebaseCollection registers a new resource with the given unique name, arguments, and options.
func NewbytebaseCollection(ctx *pulumi.Context,
	name string, args *bytebaseCollectionArgs, opts ...pulumi.ResourceOption) (*bytebaseCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource bytebaseCollection
	err := ctx.RegisterResource("bytebase:index:bytebaseCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetbytebaseCollection gets an existing bytebaseCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetbytebaseCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *bytebaseCollectionState, opts ...pulumi.ResourceOption) (*bytebaseCollection, error) {
	var resource bytebaseCollection
	err := ctx.ReadResource("bytebase:index:bytebaseCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering bytebaseCollection resources.
type bytebaseCollectionState struct {
}

type bytebaseCollectionState struct {
}

func (bytebaseCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*bytebaseCollectionState)(nil)).Elem()
}

type bytebaseCollectionArgs struct {
	// The name of the collection to be created.
	Name string `pulumi:"name"`
	// The name of the index to be used as the source for the collection.
	Source string `pulumi:"source"`
}

// The set of arguments for constructing a bytebaseCollection resource.
type bytebaseCollectionArgs struct {
	// The name of the collection to be created.
	Name pulumi.StringInput
	// The name of the index to be used as the source for the collection.
	Source pulumi.StringInput
}

func (bytebaseCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bytebaseCollectionArgs)(nil)).Elem()
}

type bytebaseCollectionInput interface {
	pulumi.Input

	TobytebaseCollectionOutput() bytebaseCollectionOutput
	TobytebaseCollectionOutputWithContext(ctx context.Context) bytebaseCollectionOutput
}

func (*bytebaseCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**bytebaseCollection)(nil)).Elem()
}

func (i *bytebaseCollection) TobytebaseCollectionOutput() bytebaseCollectionOutput {
	return i.TobytebaseCollectionOutputWithContext(context.Background())
}

func (i *bytebaseCollection) TobytebaseCollectionOutputWithContext(ctx context.Context) bytebaseCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(bytebaseCollectionOutput)
}

// bytebaseCollectionArrayInput is an input type that accepts bytebaseCollectionArray and bytebaseCollectionArrayOutput values.
// You can construct a concrete instance of `bytebaseCollectionArrayInput` via:
//
//	bytebaseCollectionArray{ bytebaseCollectionArgs{...} }
type bytebaseCollectionArrayInput interface {
	pulumi.Input

	TobytebaseCollectionArrayOutput() bytebaseCollectionArrayOutput
	TobytebaseCollectionArrayOutputWithContext(context.Context) bytebaseCollectionArrayOutput
}

type bytebaseCollectionArray []bytebaseCollectionInput

func (bytebaseCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*bytebaseCollection)(nil)).Elem()
}

func (i bytebaseCollectionArray) TobytebaseCollectionArrayOutput() bytebaseCollectionArrayOutput {
	return i.TobytebaseCollectionArrayOutputWithContext(context.Background())
}

func (i bytebaseCollectionArray) TobytebaseCollectionArrayOutputWithContext(ctx context.Context) bytebaseCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(bytebaseCollectionArrayOutput)
}

// bytebaseCollectionMapInput is an input type that accepts bytebaseCollectionMap and bytebaseCollectionMapOutput values.
// You can construct a concrete instance of `bytebaseCollectionMapInput` via:
//
//	bytebaseCollectionMap{ "key": bytebaseCollectionArgs{...} }
type bytebaseCollectionMapInput interface {
	pulumi.Input

	TobytebaseCollectionMapOutput() bytebaseCollectionMapOutput
	TobytebaseCollectionMapOutputWithContext(context.Context) bytebaseCollectionMapOutput
}

type bytebaseCollectionMap map[string]bytebaseCollectionInput

func (bytebaseCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*bytebaseCollection)(nil)).Elem()
}

func (i bytebaseCollectionMap) TobytebaseCollectionMapOutput() bytebaseCollectionMapOutput {
	return i.TobytebaseCollectionMapOutputWithContext(context.Background())
}

func (i bytebaseCollectionMap) TobytebaseCollectionMapOutputWithContext(ctx context.Context) bytebaseCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(bytebaseCollectionMapOutput)
}

type bytebaseCollectionOutput struct{ *pulumi.OutputState }

func (bytebaseCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**bytebaseCollection)(nil)).Elem()
}

func (o bytebaseCollectionOutput) TobytebaseCollectionOutput() bytebaseCollectionOutput {
	return o
}

func (o bytebaseCollectionOutput) TobytebaseCollectionOutputWithContext(ctx context.Context) bytebaseCollectionOutput {
	return o
}

// The dimension of the vectors stored in each record held in the collection.
func (o bytebaseCollectionOutput) Dimension() pulumi.IntOutput {
	return o.ApplyT(func(v *bytebaseCollection) pulumi.IntOutput { return v.Dimension }).(pulumi.IntOutput)
}

// The environment where the collection is hosted.
func (o bytebaseCollectionOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *bytebaseCollection) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The name of the collection to be created.
func (o bytebaseCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *bytebaseCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The size of the collection in bytes.
func (o bytebaseCollectionOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *bytebaseCollection) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The name of the index to be used as the source for the collection.
func (o bytebaseCollectionOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *bytebaseCollection) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// The number of records stored in the collection.
func (o bytebaseCollectionOutput) VectorCount() pulumi.IntOutput {
	return o.ApplyT(func(v *bytebaseCollection) pulumi.IntOutput { return v.VectorCount }).(pulumi.IntOutput)
}

type bytebaseCollectionArrayOutput struct{ *pulumi.OutputState }

func (bytebaseCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*bytebaseCollection)(nil)).Elem()
}

func (o bytebaseCollectionArrayOutput) TobytebaseCollectionArrayOutput() bytebaseCollectionArrayOutput {
	return o
}

func (o bytebaseCollectionArrayOutput) TobytebaseCollectionArrayOutputWithContext(ctx context.Context) bytebaseCollectionArrayOutput {
	return o
}

func (o bytebaseCollectionArrayOutput) Index(i pulumi.IntInput) bytebaseCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *bytebaseCollection {
		return vs[0].([]*bytebaseCollection)[vs[1].(int)]
	}).(bytebaseCollectionOutput)
}

type bytebaseCollectionMapOutput struct{ *pulumi.OutputState }

func (bytebaseCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*bytebaseCollection)(nil)).Elem()
}

func (o bytebaseCollectionMapOutput) TobytebaseCollectionMapOutput() bytebaseCollectionMapOutput {
	return o
}

func (o bytebaseCollectionMapOutput) TobytebaseCollectionMapOutputWithContext(ctx context.Context) bytebaseCollectionMapOutput {
	return o
}

func (o bytebaseCollectionMapOutput) MapIndex(k pulumi.StringInput) bytebaseCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *bytebaseCollection {
		return vs[0].(map[string]*bytebaseCollection)[vs[1].(string)]
	}).(bytebaseCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*bytebaseCollectionInput)(nil)).Elem(), &bytebaseCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*bytebaseCollectionArrayInput)(nil)).Elem(), bytebaseCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*bytebaseCollectionMapInput)(nil)).Elem(), bytebaseCollectionMap{})
	pulumi.RegisterOutputType(bytebaseCollectionOutput{})
	pulumi.RegisterOutputType(bytebaseCollectionArrayOutput{})
	pulumi.RegisterOutputType(bytebaseCollectionMapOutput{})
}
