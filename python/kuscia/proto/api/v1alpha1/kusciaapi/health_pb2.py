# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/health.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0kuscia/proto/api/v1alpha1/kusciaapi/health.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"I\n\rHealthRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\"\x8a\x01\n\x0eHealthResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x45\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x37.kuscia.proto.api.v1alpha1.kusciaapi.HealthResponseData\"#\n\x12HealthResponseData\x12\r\n\x05ready\x18\x01 \x01(\x08\x32\x83\x01\n\rHealthService\x12r\n\x07healthZ\x12\x32.kuscia.proto.api.v1alpha1.kusciaapi.HealthRequest\x1a\x33.kuscia.proto.api.v1alpha1.kusciaapi.HealthResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.kusciaapi.health_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _globals['_HEALTHREQUEST']._serialized_start=129
  _globals['_HEALTHREQUEST']._serialized_end=202
  _globals['_HEALTHRESPONSE']._serialized_start=205
  _globals['_HEALTHRESPONSE']._serialized_end=343
  _globals['_HEALTHRESPONSEDATA']._serialized_start=345
  _globals['_HEALTHRESPONSEDATA']._serialized_end=380
  _globals['_HEALTHSERVICE']._serialized_start=383
  _globals['_HEALTHSERVICE']._serialized_end=514
# @@protoc_insertion_point(module_scope)
