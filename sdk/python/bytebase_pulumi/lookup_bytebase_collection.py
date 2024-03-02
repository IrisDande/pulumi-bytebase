# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'LookupPineconeCollectionResult',
    'AwaitableLookupPineconeCollectionResult',
    'lookup_pinecone_collection',
    'lookup_pinecone_collection_output',
]

@pulumi.output_type
class LookupPineconeCollectionResult:
    """
    The result of a get operation on a Pinecone collection.
    """
    def __init__(__self__, dimension=None, environment=None, name=None, size=None, source=None, vector_count=None):
        if dimension and not isinstance(dimension, int):
            raise TypeError("Expected argument 'dimension' to be a int")
        pulumi.set(__self__, "dimension", dimension)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if vector_count and not isinstance(vector_count, int):
            raise TypeError("Expected argument 'vector_count' to be a int")
        pulumi.set(__self__, "vector_count", vector_count)

    @property
    @pulumi.getter
    def dimension(self) -> int:
        """
        The dimension of the vectors stored in each record held in the collection.
        """
        return pulumi.get(self, "dimension")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        The environment where the collection is hosted.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the collection to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of the collection in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        The name of the index to be used as the source for the collection.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="vectorCount")
    def vector_count(self) -> int:
        """
        The number of records stored in the collection.
        """
        return pulumi.get(self, "vector_count")


class AwaitableLookupPineconeCollectionResult(LookupPineconeCollectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LookupPineconeCollectionResult(
            dimension=self.dimension,
            environment=self.environment,
            name=self.name,
            size=self.size,
            source=self.source,
            vector_count=self.vector_count)


def lookup_pinecone_collection(name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLookupPineconeCollectionResult:
    """
    The result of a get operation on a Pinecone collection.


    :param str name: The name of the Pinecone collection.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('pinecone:index:lookupPineconeCollection', __args__, opts=opts, typ=LookupPineconeCollectionResult).value

    return AwaitableLookupPineconeCollectionResult(
        dimension=pulumi.get(__ret__, 'dimension'),
        environment=pulumi.get(__ret__, 'environment'),
        name=pulumi.get(__ret__, 'name'),
        size=pulumi.get(__ret__, 'size'),
        source=pulumi.get(__ret__, 'source'),
        vector_count=pulumi.get(__ret__, 'vector_count'))


@_utilities.lift_output_func(lookup_pinecone_collection)
def lookup_pinecone_collection_output(name: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LookupPineconeCollectionResult]:
    """
    The result of a get operation on a Pinecone collection.


    :param str name: The name of the Pinecone collection.
    """
    ...