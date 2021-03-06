# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v3/api_proto/projects.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google_proto.google.api import field_behavior_pb2 as google__proto_dot_google_dot_api_dot_field__behavior__pb2
from google_proto.google.api import resource_pb2 as google__proto_dot_google_dot_api_dot_resource__pb2
from api.v3.api_proto import project_objects_pb2 as api_dot_v3_dot_api__proto_dot_project__objects__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/v3/api_proto/projects.proto',
  package='monorail.v3',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x1f\x61pi/v3/api_proto/projects.proto\x12\x0bmonorail.v3\x1a,google_proto/google/api/field_behavior.proto\x1a&google_proto/google/api/resource.proto\x1a&api/v3/api_proto/project_objects.proto\"q\n\x19ListIssueTemplatesRequest\x12-\n\x06parent\x18\x01 \x01(\tB\x1d\xfa\x41\x17\n\x15\x61pi.crbug.com/Project\xe0\x41\x02\x12\x11\n\tpage_size\x18\x02 \x01(\x05\x12\x12\n\npage_token\x18\x03 \x01(\t\"d\n\x1aListIssueTemplatesResponse\x12-\n\ttemplates\x18\x01 \x03(\x0b\x32\x1a.monorail.v3.IssueTemplate\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"<\n\x13ListProjectsRequest\x12\x11\n\tpage_size\x18\x01 \x01(\x05\x12\x12\n\npage_token\x18\x02 \x01(\t\"W\n\x14ListProjectsResponse\x12&\n\x08projects\x18\x01 \x03(\x0b\x32\x14.monorail.v3.Project\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t2\xca\x01\n\x08Projects\x12g\n\x12ListIssueTemplates\x12&.monorail.v3.ListIssueTemplatesRequest\x1a\'.monorail.v3.ListIssueTemplatesResponse\"\x00\x12U\n\x0cListProjects\x12 .monorail.v3.ListProjectsRequest\x1a!.monorail.v3.ListProjectsResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[google__proto_dot_google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,google__proto_dot_google_dot_api_dot_resource__pb2.DESCRIPTOR,api_dot_v3_dot_api__proto_dot_project__objects__pb2.DESCRIPTOR,])




_LISTISSUETEMPLATESREQUEST = _descriptor.Descriptor(
  name='ListIssueTemplatesRequest',
  full_name='monorail.v3.ListIssueTemplatesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='parent', full_name='monorail.v3.ListIssueTemplatesRequest.parent', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\027\n\025api.crbug.com/Project\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_size', full_name='monorail.v3.ListIssueTemplatesRequest.page_size', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='monorail.v3.ListIssueTemplatesRequest.page_token', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=174,
  serialized_end=287,
)


_LISTISSUETEMPLATESRESPONSE = _descriptor.Descriptor(
  name='ListIssueTemplatesResponse',
  full_name='monorail.v3.ListIssueTemplatesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='templates', full_name='monorail.v3.ListIssueTemplatesResponse.templates', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='monorail.v3.ListIssueTemplatesResponse.next_page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=289,
  serialized_end=389,
)


_LISTPROJECTSREQUEST = _descriptor.Descriptor(
  name='ListProjectsRequest',
  full_name='monorail.v3.ListProjectsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='page_size', full_name='monorail.v3.ListProjectsRequest.page_size', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='monorail.v3.ListProjectsRequest.page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=391,
  serialized_end=451,
)


_LISTPROJECTSRESPONSE = _descriptor.Descriptor(
  name='ListProjectsResponse',
  full_name='monorail.v3.ListProjectsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='projects', full_name='monorail.v3.ListProjectsResponse.projects', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='monorail.v3.ListProjectsResponse.next_page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=453,
  serialized_end=540,
)

_LISTISSUETEMPLATESRESPONSE.fields_by_name['templates'].message_type = api_dot_v3_dot_api__proto_dot_project__objects__pb2._ISSUETEMPLATE
_LISTPROJECTSRESPONSE.fields_by_name['projects'].message_type = api_dot_v3_dot_api__proto_dot_project__objects__pb2._PROJECT
DESCRIPTOR.message_types_by_name['ListIssueTemplatesRequest'] = _LISTISSUETEMPLATESREQUEST
DESCRIPTOR.message_types_by_name['ListIssueTemplatesResponse'] = _LISTISSUETEMPLATESRESPONSE
DESCRIPTOR.message_types_by_name['ListProjectsRequest'] = _LISTPROJECTSREQUEST
DESCRIPTOR.message_types_by_name['ListProjectsResponse'] = _LISTPROJECTSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListIssueTemplatesRequest = _reflection.GeneratedProtocolMessageType('ListIssueTemplatesRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTISSUETEMPLATESREQUEST,
  __module__ = 'api.v3.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ListIssueTemplatesRequest)
  ))
_sym_db.RegisterMessage(ListIssueTemplatesRequest)

ListIssueTemplatesResponse = _reflection.GeneratedProtocolMessageType('ListIssueTemplatesResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTISSUETEMPLATESRESPONSE,
  __module__ = 'api.v3.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ListIssueTemplatesResponse)
  ))
_sym_db.RegisterMessage(ListIssueTemplatesResponse)

ListProjectsRequest = _reflection.GeneratedProtocolMessageType('ListProjectsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTPROJECTSREQUEST,
  __module__ = 'api.v3.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ListProjectsRequest)
  ))
_sym_db.RegisterMessage(ListProjectsRequest)

ListProjectsResponse = _reflection.GeneratedProtocolMessageType('ListProjectsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTPROJECTSRESPONSE,
  __module__ = 'api.v3.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ListProjectsResponse)
  ))
_sym_db.RegisterMessage(ListProjectsResponse)


_LISTISSUETEMPLATESREQUEST.fields_by_name['parent']._options = None

_PROJECTS = _descriptor.ServiceDescriptor(
  name='Projects',
  full_name='monorail.v3.Projects',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=543,
  serialized_end=745,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListIssueTemplates',
    full_name='monorail.v3.Projects.ListIssueTemplates',
    index=0,
    containing_service=None,
    input_type=_LISTISSUETEMPLATESREQUEST,
    output_type=_LISTISSUETEMPLATESRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListProjects',
    full_name='monorail.v3.Projects.ListProjects',
    index=1,
    containing_service=None,
    input_type=_LISTPROJECTSREQUEST,
    output_type=_LISTPROJECTSRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PROJECTS)

DESCRIPTOR.services_by_name['Projects'] = _PROJECTS

# @@protoc_insertion_point(module_scope)
