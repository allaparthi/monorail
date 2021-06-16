# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/secrets.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/secrets.proto',
  package='monorail.secrets',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x13proto/secrets.proto\x12\x10monorail.secrets\"k\n\x13ListRequestContents\x12\x0e\n\x06parent\x18\x01 \x01(\t\x12\x11\n\tpage_size\x18\x02 \x01(\x05\x12\x10\n\x08order_by\x18\x03 \x01(\t\x12\r\n\x05query\x18\x04 \x01(\t\x12\x10\n\x08projects\x18\x05 \x03(\t\"K\n\x11PageTokenContents\x12\r\n\x05start\x18\x01 \x01(\x05\x12\'\n\x1f\x65ncrypted_list_request_contents\x18\x02 \x01(\x0c\x62\x06proto3')
)




_LISTREQUESTCONTENTS = _descriptor.Descriptor(
  name='ListRequestContents',
  full_name='monorail.secrets.ListRequestContents',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='parent', full_name='monorail.secrets.ListRequestContents.parent', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_size', full_name='monorail.secrets.ListRequestContents.page_size', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='order_by', full_name='monorail.secrets.ListRequestContents.order_by', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='query', full_name='monorail.secrets.ListRequestContents.query', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='projects', full_name='monorail.secrets.ListRequestContents.projects', index=4,
      number=5, type=9, cpp_type=9, label=3,
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
  serialized_start=41,
  serialized_end=148,
)


_PAGETOKENCONTENTS = _descriptor.Descriptor(
  name='PageTokenContents',
  full_name='monorail.secrets.PageTokenContents',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start', full_name='monorail.secrets.PageTokenContents.start', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='encrypted_list_request_contents', full_name='monorail.secrets.PageTokenContents.encrypted_list_request_contents', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=150,
  serialized_end=225,
)

DESCRIPTOR.message_types_by_name['ListRequestContents'] = _LISTREQUESTCONTENTS
DESCRIPTOR.message_types_by_name['PageTokenContents'] = _PAGETOKENCONTENTS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListRequestContents = _reflection.GeneratedProtocolMessageType('ListRequestContents', (_message.Message,), dict(
  DESCRIPTOR = _LISTREQUESTCONTENTS,
  __module__ = 'proto.secrets_pb2'
  # @@protoc_insertion_point(class_scope:monorail.secrets.ListRequestContents)
  ))
_sym_db.RegisterMessage(ListRequestContents)

PageTokenContents = _reflection.GeneratedProtocolMessageType('PageTokenContents', (_message.Message,), dict(
  DESCRIPTOR = _PAGETOKENCONTENTS,
  __module__ = 'proto.secrets_pb2'
  # @@protoc_insertion_point(class_scope:monorail.secrets.PageTokenContents)
  ))
_sym_db.RegisterMessage(PageTokenContents)


# @@protoc_insertion_point(module_scope)