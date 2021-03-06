# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v3/api_proto/permissions.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google_proto.google.api import field_behavior_pb2 as google__proto_dot_google_dot_api_dot_field__behavior__pb2
from api.v3.api_proto import permission_objects_pb2 as api_dot_v3_dot_api__proto_dot_permission__objects__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/v3/api_proto/permissions.proto',
  package='monorail.v3',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\"api/v3/api_proto/permissions.proto\x12\x0bmonorail.v3\x1a,google_proto/google/api/field_behavior.proto\x1a)api/v3/api_proto/permission_objects.proto\",\n\x17GetPermissionSetRequest\x12\x11\n\x04name\x18\x01 \x01(\tB\x03\xe0\x41\x02\"3\n\x1d\x42\x61tchGetPermissionSetsRequest\x12\x12\n\x05names\x18\x01 \x03(\tB\x03\xe0\x41\x02\"U\n\x1e\x42\x61tchGetPermissionSetsResponse\x12\x33\n\x0fpermission_sets\x18\x01 \x03(\x0b\x32\x1a.monorail.v3.PermissionSet2\xda\x01\n\x0bPermissions\x12V\n\x10GetPermissionSet\x12$.monorail.v3.GetPermissionSetRequest\x1a\x1a.monorail.v3.PermissionSet\"\x00\x12s\n\x16\x42\x61tchGetPermissionSets\x12*.monorail.v3.BatchGetPermissionSetsRequest\x1a+.monorail.v3.BatchGetPermissionSetsResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[google__proto_dot_google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,api_dot_v3_dot_api__proto_dot_permission__objects__pb2.DESCRIPTOR,])




_GETPERMISSIONSETREQUEST = _descriptor.Descriptor(
  name='GetPermissionSetRequest',
  full_name='monorail.v3.GetPermissionSetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='monorail.v3.GetPermissionSetRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=140,
  serialized_end=184,
)


_BATCHGETPERMISSIONSETSREQUEST = _descriptor.Descriptor(
  name='BatchGetPermissionSetsRequest',
  full_name='monorail.v3.BatchGetPermissionSetsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='names', full_name='monorail.v3.BatchGetPermissionSetsRequest.names', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=186,
  serialized_end=237,
)


_BATCHGETPERMISSIONSETSRESPONSE = _descriptor.Descriptor(
  name='BatchGetPermissionSetsResponse',
  full_name='monorail.v3.BatchGetPermissionSetsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='permission_sets', full_name='monorail.v3.BatchGetPermissionSetsResponse.permission_sets', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=239,
  serialized_end=324,
)

_BATCHGETPERMISSIONSETSRESPONSE.fields_by_name['permission_sets'].message_type = api_dot_v3_dot_api__proto_dot_permission__objects__pb2._PERMISSIONSET
DESCRIPTOR.message_types_by_name['GetPermissionSetRequest'] = _GETPERMISSIONSETREQUEST
DESCRIPTOR.message_types_by_name['BatchGetPermissionSetsRequest'] = _BATCHGETPERMISSIONSETSREQUEST
DESCRIPTOR.message_types_by_name['BatchGetPermissionSetsResponse'] = _BATCHGETPERMISSIONSETSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetPermissionSetRequest = _reflection.GeneratedProtocolMessageType('GetPermissionSetRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETPERMISSIONSETREQUEST,
  __module__ = 'api.v3.api_proto.permissions_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.GetPermissionSetRequest)
  ))
_sym_db.RegisterMessage(GetPermissionSetRequest)

BatchGetPermissionSetsRequest = _reflection.GeneratedProtocolMessageType('BatchGetPermissionSetsRequest', (_message.Message,), dict(
  DESCRIPTOR = _BATCHGETPERMISSIONSETSREQUEST,
  __module__ = 'api.v3.api_proto.permissions_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.BatchGetPermissionSetsRequest)
  ))
_sym_db.RegisterMessage(BatchGetPermissionSetsRequest)

BatchGetPermissionSetsResponse = _reflection.GeneratedProtocolMessageType('BatchGetPermissionSetsResponse', (_message.Message,), dict(
  DESCRIPTOR = _BATCHGETPERMISSIONSETSRESPONSE,
  __module__ = 'api.v3.api_proto.permissions_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.BatchGetPermissionSetsResponse)
  ))
_sym_db.RegisterMessage(BatchGetPermissionSetsResponse)


_GETPERMISSIONSETREQUEST.fields_by_name['name']._options = None
_BATCHGETPERMISSIONSETSREQUEST.fields_by_name['names']._options = None

_PERMISSIONS = _descriptor.ServiceDescriptor(
  name='Permissions',
  full_name='monorail.v3.Permissions',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=327,
  serialized_end=545,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetPermissionSet',
    full_name='monorail.v3.Permissions.GetPermissionSet',
    index=0,
    containing_service=None,
    input_type=_GETPERMISSIONSETREQUEST,
    output_type=api_dot_v3_dot_api__proto_dot_permission__objects__pb2._PERMISSIONSET,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='BatchGetPermissionSets',
    full_name='monorail.v3.Permissions.BatchGetPermissionSets',
    index=1,
    containing_service=None,
    input_type=_BATCHGETPERMISSIONSETSREQUEST,
    output_type=_BATCHGETPERMISSIONSETSRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PERMISSIONS)

DESCRIPTOR.services_by_name['Permissions'] = _PERMISSIONS

# @@protoc_insertion_point(module_scope)
