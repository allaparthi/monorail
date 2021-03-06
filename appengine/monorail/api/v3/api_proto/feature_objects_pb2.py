# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v3/api_proto/feature_objects.proto

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
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from api.v3.api_proto import issue_objects_pb2 as api_dot_v3_dot_api__proto_dot_issue__objects__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/v3/api_proto/feature_objects.proto',
  package='monorail.v3',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n&api/v3/api_proto/feature_objects.proto\x12\x0bmonorail.v3\x1a,google_proto/google/api/field_behavior.proto\x1a&google_proto/google/api/resource.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a$api/v3/api_proto/issue_objects.proto\"\xac\x03\n\x07Hotlist\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x19\n\x0c\x64isplay_name\x18\x02 \x01(\tB\x03\xe0\x41\x02\x12)\n\x05owner\x18\x03 \x01(\tB\x1a\xfa\x41\x14\n\x12\x61pi.crbug.com/User\xe0\x41\x02\x12(\n\x07\x65\x64itors\x18\x04 \x03(\tB\x17\xfa\x41\x14\n\x12\x61pi.crbug.com/User\x12\x14\n\x07summary\x18\x05 \x01(\tB\x03\xe0\x41\x02\x12\x18\n\x0b\x64\x65scription\x18\x06 \x01(\tB\x03\xe0\x41\x02\x12\x36\n\x0f\x64\x65\x66\x61ult_columns\x18\x07 \x03(\x0b\x32\x1d.monorail.v3.IssuesListColumn\x12<\n\x0fhotlist_privacy\x18\x08 \x01(\x0e\x32#.monorail.v3.Hotlist.HotlistPrivacy\"J\n\x0eHotlistPrivacy\x12\x1f\n\x1bHOTLIST_PRIVACY_UNSPECIFIED\x10\x00\x12\x0b\n\x07PRIVATE\x10\x01\x12\n\n\x06PUBLIC\x10\x02:1\xea\x41.\n\x15\x61pi.crbug.com/Hotlist\x12\x15hotlists/{hotlist_id}\"\x90\x02\n\x0bHotlistItem\x12\x0c\n\x04name\x18\x01 \x01(\t\x12*\n\x05issue\x18\x02 \x01(\tB\x1b\xfa\x41\x15\n\x13\x61pi.crbug.com/Issue\xe0\x41\x05\x12\x11\n\x04rank\x18\x03 \x01(\rB\x03\xe0\x41\x03\x12)\n\x05\x61\x64\x64\x65r\x18\x04 \x01(\tB\x1a\xfa\x41\x14\n\x12\x61pi.crbug.com/User\xe0\x41\x03\x12\x34\n\x0b\x63reate_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x03\xe0\x41\x03\x12\x0c\n\x04note\x18\x06 \x01(\t:E\xea\x41\x42\n\x19\x61pi.crbug.com/HotlistItem\x12%hotlists/{hotlist_id}/items/{item_id}b\x06proto3')
  ,
  dependencies=[google__proto_dot_google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,google__proto_dot_google_dot_api_dot_resource__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,api_dot_v3_dot_api__proto_dot_issue__objects__pb2.DESCRIPTOR,])



_HOTLIST_HOTLISTPRIVACY = _descriptor.EnumDescriptor(
  name='HotlistPrivacy',
  full_name='monorail.v3.Hotlist.HotlistPrivacy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HOTLIST_PRIVACY_UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIVATE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PUBLIC', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=516,
  serialized_end=590,
)
_sym_db.RegisterEnumDescriptor(_HOTLIST_HOTLISTPRIVACY)


_HOTLIST = _descriptor.Descriptor(
  name='Hotlist',
  full_name='monorail.v3.Hotlist',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='monorail.v3.Hotlist.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='display_name', full_name='monorail.v3.Hotlist.display_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='owner', full_name='monorail.v3.Hotlist.owner', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\024\n\022api.crbug.com/User\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='editors', full_name='monorail.v3.Hotlist.editors', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\024\n\022api.crbug.com/User'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='summary', full_name='monorail.v3.Hotlist.summary', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='monorail.v3.Hotlist.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_columns', full_name='monorail.v3.Hotlist.default_columns', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hotlist_privacy', full_name='monorail.v3.Hotlist.hotlist_privacy', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _HOTLIST_HOTLISTPRIVACY,
  ],
  serialized_options=_b('\352A.\n\025api.crbug.com/Hotlist\022\025hotlists/{hotlist_id}'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=213,
  serialized_end=641,
)


_HOTLISTITEM = _descriptor.Descriptor(
  name='HotlistItem',
  full_name='monorail.v3.HotlistItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='monorail.v3.HotlistItem.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='issue', full_name='monorail.v3.HotlistItem.issue', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\025\n\023api.crbug.com/Issue\340A\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rank', full_name='monorail.v3.HotlistItem.rank', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\003'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='adder', full_name='monorail.v3.HotlistItem.adder', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\024\n\022api.crbug.com/User\340A\003'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='monorail.v3.HotlistItem.create_time', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\003'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='note', full_name='monorail.v3.HotlistItem.note', index=5,
      number=6, type=9, cpp_type=9, label=1,
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
  serialized_options=_b('\352AB\n\031api.crbug.com/HotlistItem\022%hotlists/{hotlist_id}/items/{item_id}'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=644,
  serialized_end=916,
)

_HOTLIST.fields_by_name['default_columns'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUESLISTCOLUMN
_HOTLIST.fields_by_name['hotlist_privacy'].enum_type = _HOTLIST_HOTLISTPRIVACY
_HOTLIST_HOTLISTPRIVACY.containing_type = _HOTLIST
_HOTLISTITEM.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['Hotlist'] = _HOTLIST
DESCRIPTOR.message_types_by_name['HotlistItem'] = _HOTLISTITEM
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Hotlist = _reflection.GeneratedProtocolMessageType('Hotlist', (_message.Message,), dict(
  DESCRIPTOR = _HOTLIST,
  __module__ = 'api.v3.api_proto.feature_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.Hotlist)
  ))
_sym_db.RegisterMessage(Hotlist)

HotlistItem = _reflection.GeneratedProtocolMessageType('HotlistItem', (_message.Message,), dict(
  DESCRIPTOR = _HOTLISTITEM,
  __module__ = 'api.v3.api_proto.feature_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.HotlistItem)
  ))
_sym_db.RegisterMessage(HotlistItem)


_HOTLIST.fields_by_name['display_name']._options = None
_HOTLIST.fields_by_name['owner']._options = None
_HOTLIST.fields_by_name['editors']._options = None
_HOTLIST.fields_by_name['summary']._options = None
_HOTLIST.fields_by_name['description']._options = None
_HOTLIST._options = None
_HOTLISTITEM.fields_by_name['issue']._options = None
_HOTLISTITEM.fields_by_name['rank']._options = None
_HOTLISTITEM.fields_by_name['adder']._options = None
_HOTLISTITEM.fields_by_name['create_time']._options = None
_HOTLISTITEM._options = None
# @@protoc_insertion_point(module_scope)
