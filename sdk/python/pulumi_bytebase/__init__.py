# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .random import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "bytebase",
  "mod": "index",
  "fqn": "pulumi_bytebase",
  "classes": {
   "bytebase:index:Random": "Random"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "bytebase",
  "token": "pulumi:providers:bytebase",
  "fqn": "pulumi_bytebase",
  "class": "Provider"
 }
]
"""
)