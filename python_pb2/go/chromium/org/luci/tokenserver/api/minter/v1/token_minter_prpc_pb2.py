# Generated by the pRPC protocol buffer compiler plugin.  DO NOT EDIT!
# source: go.chromium.org/luci/tokenserver/api/minter/v1/token_minter.proto

import base64
import zlib

from google.protobuf import descriptor_pb2

# Includes description of the go.chromium.org/luci/tokenserver/api/minter/v1/token_minter.proto and all of its transitive
# dependencies. Includes source code info.
FILE_DESCRIPTOR_SET = descriptor_pb2.FileDescriptorSet()
FILE_DESCRIPTOR_SET.ParseFromString(zlib.decompress(base64.b64decode(
    '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')))
_INDEX = {
    f.name: {
      'descriptor': f,
      'services': {s.name: s for s in f.service},
    }
    for f in FILE_DESCRIPTOR_SET.file
}


TokenMinterServiceDescription = {
  'file_descriptor_set': FILE_DESCRIPTOR_SET,
  'file_descriptor': _INDEX[u'go.chromium.org/luci/tokenserver/api/minter/v1/token_minter.proto']['descriptor'],
  'service_descriptor': _INDEX[u'go.chromium.org/luci/tokenserver/api/minter/v1/token_minter.proto']['services'][u'TokenMinter'],
}