
?y
google/api/http.proto
google.api"y
Http*
rules (2.google.api.HttpRuleRrulesE
fully_decode_reserved_expansion (RfullyDecodeReservedExpansion"?
HttpRule
selector (	Rselector
get (	H Rget
put (	H Rput
post (	H Rpost
delete (	H Rdelete
patch (	H Rpatch7
custom (2.google.api.CustomHttpPatternH Rcustom
body (	Rbody#
response_body (	RresponseBodyE
additional_bindings (2.google.api.HttpRuleRadditionalBindingsB	
pattern";
CustomHttpPattern
kind (	Rkind
path (	RpathBj
com.google.apiB	HttpProtoPZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations??GAPIJ?t
 ?
?
 2? Copyright 2019 Google LLC.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.





 
	
 

 X
	
 X

 "
	

 "

 *
	
 *

 '
	
 '

 "
	
$ "
?
  *? Defines the HTTP configuration for an API service. It contains a list of
 [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
 to one or more HTTP REST API methods.



 
?
  !? A list of HTTP configuration rules that apply to individual API methods.

 **NOTE:** All service configuration rules follow "last one wins" order.


  !


  !

  !

  !
?
 )+? When set to true, URL path parameters will be fully URI-decoded except in
 cases of single segment matches in reserved expansion, where "%2F" will be
 left encoded.

 The default behavior is to not decode RFC 6570 reserved characters in multi
 segment matches.


 )!

 )

 )&

 ))*
?S
? ??S # gRPC Transcoding

 gRPC Transcoding is a feature for mapping between a gRPC method and one or
 more HTTP REST endpoints. It allows developers to build a single API service
 that supports both gRPC APIs and REST APIs. Many systems, including [Google
 APIs](https://github.com/googleapis/googleapis),
 [Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
 Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
 and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
 and use it for large scale production services.

 `HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
 how different portions of the gRPC request message are mapped to the URL
 path, URL query parameters, and HTTP request body. It also controls how the
 gRPC response message is mapped to the HTTP response body. `HttpRule` is
 typically specified as an `google.api.http` annotation on the gRPC method.

 Each mapping specifies a URL path template and an HTTP method. The path
 template may refer to one or more fields in the gRPC request message, as long
 as each field is a non-repeated field with a primitive (non-message) type.
 The path template controls how fields of the request message are mapped to
 the URL path.

 Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get: "/v1/{name=messages/*}"
         };
       }
     }
     message GetMessageRequest {
       string name = 1; // Mapped to URL path.
     }
     message Message {
       string text = 1; // The resource content.
     }

 This enables an HTTP REST to gRPC mapping as below:

 HTTP | gRPC
 -----|-----
 `GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`

 Any fields in the request message which are not bound by the path template
 automatically become HTTP query parameters if there is no HTTP request body.
 For example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get:"/v1/messages/{message_id}"
         };
       }
     }
     message GetMessageRequest {
       message SubMessage {
         string subfield = 1;
       }
       string message_id = 1; // Mapped to URL path.
       int64 revision = 2;    // Mapped to URL query parameter `revision`.
       SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
     }

 This enables a HTTP JSON to RPC mapping as below:

 HTTP | gRPC
 -----|-----
 `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
 `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
 "foo"))`

 Note that fields which are mapped to URL query parameters must have a
 primitive type or a repeated primitive type or a non-repeated message type.
 In the case of a repeated type, the parameter can be repeated in the URL
 as `...?param=A&param=B`. In the case of a message type, each field of the
 message is mapped to a separate parameter, such as
 `...?foo.a=A&foo.b=B&foo.c=C`.

 For HTTP methods that allow a request body, the `body` field
 specifies the mapping. Consider a REST update method on the
 message resource collection:

     service Messaging {
       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "message"
         };
       }
     }
     message UpdateMessageRequest {
       string message_id = 1; // mapped to the URL
       Message message = 2;   // mapped to the body
     }

 The following HTTP JSON to RPC mapping is enabled, where the
 representation of the JSON in the request body is determined by
 protos JSON encoding:

 HTTP | gRPC
 -----|-----
 `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
 "123456" message { text: "Hi!" })`

 The special name `*` can be used in the body mapping to define that
 every field not bound by the path template should be mapped to the
 request body.  This enables the following alternative definition of
 the update method:

     service Messaging {
       rpc UpdateMessage(Message) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "*"
         };
       }
     }
     message Message {
       string message_id = 1;
       string text = 2;
     }


 The following HTTP JSON to RPC mapping is enabled:

 HTTP | gRPC
 -----|-----
 `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
 "123456" text: "Hi!")`

 Note that when using `*` in the body mapping, it is not possible to
 have HTTP parameters, as all fields not bound by the path end in
 the body. This makes this option more rarely used in practice when
 defining REST APIs. The common usage of `*` is in custom methods
 which don't use the URL at all for transferring data.

 It is possible to define multiple HTTP methods for one RPC by using
 the `additional_bindings` option. Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
           get: "/v1/messages/{message_id}"
           additional_bindings {
             get: "/v1/users/{user_id}/messages/{message_id}"
           }
         };
       }
     }
     message GetMessageRequest {
       string message_id = 1;
       string user_id = 2;
     }

 This enables the following two alternative HTTP JSON to RPC mappings:

 HTTP | gRPC
 -----|-----
 `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
 `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
 "123456")`

 ## Rules for HTTP mapping

 1. Leaf request fields (recursive expansion nested messages in the request
    message) are classified into three categories:
    - Fields referred by the path template. They are passed via the URL path.
    - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They are passed via the HTTP
      request body.
    - All other fields are passed via the URL query parameters, and the
      parameter name is the field path in the request message. A repeated
      field can be represented as multiple query parameters under the same
      name.
  2. If [HttpRule.body][google.api.HttpRule.body] is "*", there is no URL query parameter, all fields
     are passed via URL path and HTTP request body.
  3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP request body, all
     fields are passed via URL path and URL query parameters.

 ### Path template syntax

     Template = "/" Segments [ Verb ] ;
     Segments = Segment { "/" Segment } ;
     Segment  = "*" | "**" | LITERAL | Variable ;
     Variable = "{" FieldPath [ "=" Segments ] "}" ;
     FieldPath = IDENT { "." IDENT } ;
     Verb     = ":" LITERAL ;

 The syntax `*` matches a single URL path segment. The syntax `**` matches
 zero or more URL path segments, which must be the last part of the URL path
 except the `Verb`.

 The syntax `Variable` matches part of the URL path as specified by its
 template. A variable template must not contain other variables. If a variable
 matches a single path segment, its template may be omitted, e.g. `{var}`
 is equivalent to `{var=*}`.

 The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
 contains any reserved character, such characters should be percent-encoded
 before the matching.

 If a variable contains exactly one path segment, such as `"{var}"` or
 `"{var=*}"`, when such a variable is expanded into a URL path on the client
 side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
 server side does the reverse decoding. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{var}`.

 If a variable contains multiple path segments, such as `"{var=foo/*}"`
 or `"{var=**}"`, when such a variable is expanded into a URL path on the
 client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
 The server side does the reverse decoding, except "%2F" and "%2f" are left
 unchanged. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{+var}`.

 ## Using gRPC API Service Configuration

 gRPC API Service Configuration (service config) is a configuration language
 for configuring a gRPC service to become a user-facing product. The
 service config is simply the YAML representation of the `google.api.Service`
 proto message.

 As an alternative to annotating your proto file, you can configure gRPC
 transcoding in your service config YAML files. You do this by specifying a
 `HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
 effect as the proto annotation. This can be particularly useful if you
 have a proto that is reused in multiple services. Note that any transcoding
 specified in the service config will override any matching transcoding
 configuration in the proto.

 Example:

     http:
       rules:
         # Selects a gRPC method and applies HttpRule to it.
         - selector: example.v1.Messaging.GetMessage
           get: /v1/messages/{message_id}/{sub.subfield}

 ## Special notes

 When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
 proto to JSON conversion must follow the [proto3
 specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

 While the single segment variable follows the semantics of
 [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
 Expansion, the multi segment variable **does not** follow RFC 6570 Section
 3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
 does not expand special characters like `?` and `#`, which would lead
 to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
 for multi segment variables.

 The path variables **must not** refer to any repeated or mapped field,
 because client libraries are not capable of handling such variable expansion.

 The path variables **must not** capture the leading "/" character. The reason
 is that the most common use case "{var}" does not capture the leading "/"
 character. For consistency, all path variables must share the same behavior.

 Repeated message fields must not be mapped to URL query parameters, because
 no client library can support such complicated mapping.

 If an API needs to use a JSON array for request or response body, it can map
 the request or response body to a repeated field. However, some gRPC
 Transcoding implementations may not support this feature.


?
?
 ? Selects a method to which this rule applies.

 Refer to [selector][google.api.DocumentationRule.selector] for syntax details.


 ??

 ?

 ?	

 ?
?
 ??? Determines the URL pattern is matched by this rules. This pattern can be
 used with any of the {get|put|post|delete|patch} methods. A custom method
 can be defined using the 'custom' field.


 ?
\
?N Maps to HTTP GET. Used for listing and getting information about
 resources.


?


?

?
@
?2 Maps to HTTP PUT. Used for replacing a resource.


?


?

?
X
?J Maps to HTTP POST. Used for creating a resource or performing an action.


?


?

?
B
?4 Maps to HTTP DELETE. Used for deleting a resource.


?


?

?
A
?3 Maps to HTTP PATCH. Used for updating a resource.


?


?

?
?
?!? The custom pattern is used for specifying an HTTP method that is not
 included in the `pattern` field, such as HEAD, or "*" to leave the
 HTTP method unspecified for this rule. The wild-card rule is useful
 for services that provide content to Web (HTML) clients.


?

?

? 
?
?? The name of the request field whose value is mapped to the HTTP request
 body, or `*` for mapping all request fields not captured by the path
 pattern to the HTTP body, or omitted for not having any HTTP request body.

 NOTE: the referred field must be present at the top-level of the request
 message type.


??

?

?	

?
?
?? Optional. The name of the response field whose value is mapped to the HTTP
 response body. When omitted, the entire response message will be used
 as the HTTP response body.

 NOTE: The referred field must be present at the top-level of the response
 message type.


??

?

?	

?
?
	?-? Additional HTTP bindings for the selector. Nested bindings must
 not contain an `additional_bindings` field themselves (that is,
 the nesting may only be one level deep).


	?


	?

	?'

	?*,
G
? ?9 A custom pattern is used for defining custom HTTP verb.


?
2
 ?$ The name of this custom HTTP verb.


 ??

 ?

 ?	

 ?
5
?' The path matched by this custom verb.


??

?

?	

?bproto3
??
 google/protobuf/descriptor.protogoogle.protobuf"M
FileDescriptorSet8
file (2$.google.protobuf.FileDescriptorProtoRfile"?
FileDescriptorProto
name (	Rname
package (	Rpackage

dependency (	R
dependency+
public_dependency
 (RpublicDependency'
weak_dependency (RweakDependencyC
message_type (2 .google.protobuf.DescriptorProtoRmessageTypeA
	enum_type (2$.google.protobuf.EnumDescriptorProtoRenumTypeA
service (2'.google.protobuf.ServiceDescriptorProtoRserviceC
	extension (2%.google.protobuf.FieldDescriptorProtoR	extension6
options (2.google.protobuf.FileOptionsRoptionsI
source_code_info	 (2.google.protobuf.SourceCodeInfoRsourceCodeInfo
syntax (	Rsyntax"?
DescriptorProto
name (	Rname;
field (2%.google.protobuf.FieldDescriptorProtoRfieldC
	extension (2%.google.protobuf.FieldDescriptorProtoR	extensionA
nested_type (2 .google.protobuf.DescriptorProtoR
nestedTypeA
	enum_type (2$.google.protobuf.EnumDescriptorProtoRenumTypeX
extension_range (2/.google.protobuf.DescriptorProto.ExtensionRangeRextensionRangeD

oneof_decl (2%.google.protobuf.OneofDescriptorProtoR	oneofDecl9
options (2.google.protobuf.MessageOptionsRoptionsU
reserved_range	 (2..google.protobuf.DescriptorProto.ReservedRangeRreservedRange#
reserved_name
 (	RreservedNamez
ExtensionRange
start (Rstart
end (Rend@
options (2&.google.protobuf.ExtensionRangeOptionsRoptions7
ReservedRange
start (Rstart
end (Rend"|
ExtensionRangeOptionsX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
FieldDescriptorProto
name (	Rname
number (RnumberA
label (2+.google.protobuf.FieldDescriptorProto.LabelRlabel>
type (2*.google.protobuf.FieldDescriptorProto.TypeRtype
	type_name (	RtypeName
extendee (	Rextendee#
default_value (	RdefaultValue
oneof_index	 (R
oneofIndex
	json_name
 (	RjsonName7
options (2.google.protobuf.FieldOptionsRoptions"?
Type
TYPE_DOUBLE

TYPE_FLOAT

TYPE_INT64
TYPE_UINT64

TYPE_INT32
TYPE_FIXED64
TYPE_FIXED32
	TYPE_BOOL
TYPE_STRING	

TYPE_GROUP

TYPE_MESSAGE

TYPE_BYTES
TYPE_UINT32
	TYPE_ENUM
TYPE_SFIXED32
TYPE_SFIXED64
TYPE_SINT32
TYPE_SINT64"C
Label
LABEL_OPTIONAL
LABEL_REQUIRED
LABEL_REPEATED"c
OneofDescriptorProto
name (	Rname7
options (2.google.protobuf.OneofOptionsRoptions"?
EnumDescriptorProto
name (	Rname?
value (2).google.protobuf.EnumValueDescriptorProtoRvalue6
options (2.google.protobuf.EnumOptionsRoptions]
reserved_range (26.google.protobuf.EnumDescriptorProto.EnumReservedRangeRreservedRange#
reserved_name (	RreservedName;
EnumReservedRange
start (Rstart
end (Rend"?
EnumValueDescriptorProto
name (	Rname
number (Rnumber;
options (2!.google.protobuf.EnumValueOptionsRoptions"?
ServiceDescriptorProto
name (	Rname>
method (2&.google.protobuf.MethodDescriptorProtoRmethod9
options (2.google.protobuf.ServiceOptionsRoptions"?
MethodDescriptorProto
name (	Rname

input_type (	R	inputType
output_type (	R
outputType8
options (2.google.protobuf.MethodOptionsRoptions0
client_streaming (:falseRclientStreaming0
server_streaming (:falseRserverStreaming"?	
FileOptions!
java_package (	RjavaPackage0
java_outer_classname (	RjavaOuterClassname5
java_multiple_files
 (:falseRjavaMultipleFilesD
java_generate_equals_and_hash (BRjavaGenerateEqualsAndHash:
java_string_check_utf8 (:falseRjavaStringCheckUtf8S
optimize_for	 (2).google.protobuf.FileOptions.OptimizeMode:SPEEDRoptimizeFor

go_package (	R	goPackage5
cc_generic_services (:falseRccGenericServices9
java_generic_services (:falseRjavaGenericServices5
py_generic_services (:falseRpyGenericServices7
php_generic_services* (:falseRphpGenericServices%

deprecated (:falseR
deprecated/
cc_enable_arenas (:falseRccEnableArenas*
objc_class_prefix$ (	RobjcClassPrefix)
csharp_namespace% (	RcsharpNamespace!
swift_prefix' (	RswiftPrefix(
php_class_prefix( (	RphpClassPrefix#
php_namespace) (	RphpNamespace4
php_metadata_namespace, (	RphpMetadataNamespace!
ruby_package- (	RrubyPackageX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption":
OptimizeMode	
SPEED
	CODE_SIZE
LITE_RUNTIME*	?????J&'"?
MessageOptions<
message_set_wire_format (:falseRmessageSetWireFormatL
no_standard_descriptor_accessor (:falseRnoStandardDescriptorAccessor%

deprecated (:falseR
deprecated
	map_entry (RmapEntryX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????J	J	
"?
FieldOptionsA
ctype (2#.google.protobuf.FieldOptions.CType:STRINGRctype
packed (RpackedG
jstype (2$.google.protobuf.FieldOptions.JSType:	JS_NORMALRjstype
lazy (:falseRlazy%

deprecated (:falseR
deprecated
weak
 (:falseRweakX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption"/
CType

STRING 
CORD
STRING_PIECE"5
JSType
	JS_NORMAL 
	JS_STRING
	JS_NUMBER*	?????J"s
OneofOptionsX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
EnumOptions
allow_alias (R
allowAlias%

deprecated (:falseR
deprecatedX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????J"?
EnumValueOptions%

deprecated (:falseR
deprecatedX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
ServiceOptions%

deprecated! (:falseR
deprecatedX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
MethodOptions%

deprecated! (:falseR
deprecatedq
idempotency_level" (2/.google.protobuf.MethodOptions.IdempotencyLevel:IDEMPOTENCY_UNKNOWNRidempotencyLevelX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption"P
IdempotencyLevel
IDEMPOTENCY_UNKNOWN 
NO_SIDE_EFFECTS

IDEMPOTENT*	?????"?
UninterpretedOptionA
name (2-.google.protobuf.UninterpretedOption.NamePartRname)
identifier_value (	RidentifierValue,
positive_int_value (RpositiveIntValue,
negative_int_value (RnegativeIntValue!
double_value (RdoubleValue!
string_value (RstringValue'
aggregate_value (	RaggregateValueJ
NamePart
	name_part (	RnamePart!
is_extension (RisExtension"?
SourceCodeInfoD
location (2(.google.protobuf.SourceCodeInfo.LocationRlocation?
Location
path (BRpath
span (BRspan)
leading_comments (	RleadingComments+
trailing_comments (	RtrailingComments:
leading_detached_comments (	RleadingDetachedComments"?
GeneratedCodeInfoM

annotation (2-.google.protobuf.GeneratedCodeInfo.AnnotationR
annotationm

Annotation
path (BRpath
source_file (	R
sourceFile
begin (Rbegin
end (RendB?
com.google.protobufBDescriptorProtosHZ>github.com/golang/protobuf/protoc-gen-go/descriptor;descriptor??GPB?Google.Protobuf.ReflectionJ??
' ?
?
' 2? Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
2? Author: kenton@google.com (Kenton Varda)
  Based on original Protocol Buffers design by
  Sanjay Ghemawat, Jeff Dean, and others.

 The messages in this file describe the definitions found in .proto files.
 A valid .proto file can be translated directly to a FileDescriptorProto
 without any other information (e.g. without reading its imports).


)

* U
	
* U

+ ,
	
+ ,

, 1
	
, 1

- 7
	
%- 7

. !
	
$. !

/ 
	
/ 

3 

	3 t descriptor.proto must be optimized for speed because reflection-based
 algorithms don't work during bootstrapping.

j
 7 9^ The protocol compiler can output a FileDescriptorSet containing the .proto
 files it parses.



 7

  8(

  8


  8

  8#

  8&'
/
< Y# Describes a complete .proto file.



<
9
 =", file name, relative to root of source tree


 =


 =

 =

 =
*
>" e.g. "foo", "foo.bar", etc.


>


>

>

>
4
A!' Names of files imported by this file.


A


A

A

A 
Q
C(D Indexes of the public imported files in the dependency list above.


C


C

C"

C%'
z
F&m Indexes of the weak imported files in the dependency list.
 For Google-internal migration only. Do not use.


F


F

F 

F#%
6
I,) All top-level definitions in this file.


I


I

I'

I*+

J-

J


J

J(

J+,

K.

K


K!

K")

K,-

L.

L


L

L )

L,-

	N#

	N


	N

	N

	N!"
?

T/? This field contains optional information about the original source code.
 You may safely remove this entire field without harming runtime
 functionality of the descriptors -- the information is needed only by
 development tools.



T



T


T*


T-.
]
XP The syntax of the proto file.
 The supported values are "proto2" and "proto3".


X


X

X

X
'
\ | Describes a message type.



\

 ]

 ]


 ]

 ]

 ]

_*

_


_

_ %

_()

`.

`


`

` )

`,-

b+

b


b

b&

b)*

c-

c


c

c(

c+,

 ej

 e


  f

  f

  f

  f

  f

 g

 g

 g

 g

 g

 i/

 i

 i"

 i#*

 i-.

k.

k


k

k)

k,-

m/

m


m

m *

m-.

o&

o


o

o!

o$%
?
tw? Range of reserved tag numbers. Reserved tag numbers may not be used by
 fields or extension ranges in the same message. Reserved ranges may
 not overlap.


t


 u" Inclusive.


 u

 u

 u

 u

v" Exclusive.


v

v

v

v

x,

x


x

x'

x*+
?
	{%u Reserved field names, which may not be used by fields in the same message.
 A given name may only be reserved once.


	{


	{

	{

	{"$

~ ?


~
O
 ?:A The parser stores options it doesn't recognize here. See above.


 ?


 ?

 ?3

 ?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?
3
? ?% Describes a field within a message.


?

 ??

 ?
S
  ?C 0 is reserved for errors.
 Order is weird for historical reasons.


  ?

  ?

 ?

 ?

 ?
w
 ?g Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
 negative values are likely.


 ?

 ?

 ?

 ?

 ?
w
 ?g Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
 negative values are likely.


 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?
?
 	?? Tag-delimited aggregate.
 Group type is deprecated and not supported in proto3. However, Proto3
 implementations should still be able to parse the group wire format and
 treat group fields as unknown fields.


 	?

 	?
-
 
?" Length-delimited aggregate.


 
?

 
?
#
 ? New in version 2.


 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?
'
 ?" Uses ZigZag encoding.


 ?

 ?
'
 ?" Uses ZigZag encoding.


 ?

 ?

??

?
*
 ? 0 is reserved for errors


 ?

 ?

?

?

?

?

?

?

 ?

 ?


 ?

 ?

 ?

?

?


?

?

?

?

?


?

?

?
?
?? If type_name is set, this need not be set.  If both this and type_name
 are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.


?


?

?

?
?
? ? For message and enum types, this is the name of the type.  If the name
 starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
 rules are used to find the type (i.e. first the nested types within this
 message are searched, then within the parent, on up to the root
 namespace).


?


?

?

?
~
?p For extensions, this is the name of the type being extended.  It is
 resolved in the same manner as type_name.


?


?

?

?
?
?$? For numeric types, contains the original text representation of the value.
 For booleans, "true" or "false".
 For strings, contains the default text contents (not escaped in any way).
 For bytes, contains the C escaped value.  All bytes >= 128 are escaped.
 TODO(kenton):  Base-64 encode?


?


?

?

?"#
?
?!v If set, gives the index of a oneof in the containing type's oneof_decl
 list.  This field is a member of that oneof.


?


?

?

? 
?
?!? JSON name of this field. The value is set by protocol compiler. If the
 user has set a "json_name" option on this field, that option's value
 will be used. Otherwise, it's deduced from the field's name by converting
 it to camelCase.


?


?

?

? 

	?$

	?


	?

	?

	?"#
"
? ? Describes a oneof.


?

 ?

 ?


 ?

 ?

 ?

?$

?


?

?

?"#
'
? ? Describes an enum type.


?

 ?

 ?


 ?

 ?

 ?

?.

?


?#

?$)

?,-

?#

?


?

?

?!"
?
 ??? Range of reserved numeric values. Reserved values may not be used by
 entries in the same enum. Reserved ranges may not overlap.

 Note that this is distinct from DescriptorProto.ReservedRange in that it
 is inclusive such that it can appropriately represent the entire int32
 domain.


 ?


  ?" Inclusive.


  ?

  ?

  ?

  ?

 ?" Inclusive.


 ?

 ?

 ?

 ?
?
?0? Range of reserved numeric values. Reserved numeric values may not be used
 by enum values in the same enum declaration. Reserved ranges may not
 overlap.


?


?

?+

?./
l
?$^ Reserved enum value names, which may not be reused. A given name may only
 be reserved once.


?


?

?

?"#
1
? ?# Describes a value within an enum.


? 

 ?

 ?


 ?

 ?

 ?

?

?


?

?

?

?(

?


?

?#

?&'
$
? ? Describes a service.


?

 ?

 ?


 ?

 ?

 ?

?,

?


? 

?!'

?*+

?&

?


?

?!

?$%
0
	? ?" Describes a method of a service.


	?

	 ?

	 ?


	 ?

	 ?

	 ?
?
	?!? Input and output type names.  These are resolved in the same way as
 FieldDescriptorProto.type_name, but must refer to a message type.


	?


	?

	?

	? 

	?"

	?


	?

	?

	? !

	?%

	?


	?

	? 

	?#$
E
	?57 Identifies if client streams multiple client messages


	?


	?

	? 

	?#$

	?%4

	?.3
E
	?57 Identifies if server streams multiple server messages


	?


	?

	? 

	?#$

	?%4

	?.3
?

? ?2N ===================================================================
 Options
2? Each of the definitions above may have "options" attached.  These are
 just annotations which may cause code to be generated slightly differently
 or may contain hints for code that manipulates protocol messages.

 Clients may define custom options as extensions of the *Options messages.
 These extensions may not yet be known at parsing time, so the parser cannot
 store the values in them.  Instead it stores them in a field in the *Options
 message called uninterpreted_option. This field must have the same name
 across all *Options messages. We then use this field to populate the
 extensions when we build a descriptor, at which point all protos have been
 parsed and so all extensions are known.

 Extension numbers for custom options may be chosen as follows:
 * For options which will only be used within a single application or
   organization, or for experimental options, use field numbers 50000
   through 99999.  It is up to you to ensure that you do not use the
   same number for multiple options.
 * For options which will be published and used publicly by multiple
   independent entities, e-mail protobuf-global-extension-registry@google.com
   to reserve extension numbers. Simply provide your project name (e.g.
   Objective-C plugin) and your project website (if available) -- there's no
   need to explain how you intend to use them. Usually you only need one
   extension number. You can declare multiple options with only one extension
   number by putting them in a sub-message. See the Custom Options section of
   the docs for examples:
   https://developers.google.com/protocol-buffers/docs/proto#options
   If this turns out to be popular, a web service will be set up
   to automatically assign option numbers.



?
?

 ?#? Sets the Java package where classes generated from this .proto will be
 placed.  By default, the proto package is used, but this is often
 inappropriate because proto packages do not normally start with backwards
 domain names.



 ?



 ?


 ?


 ?!"
?

?+? If set, all the classes from the .proto file are wrapped in a single
 outer class with the given name.  This applies to both Proto1
 (equivalent to the old "--one_java_file" option) and Proto2 (where
 a .proto always translates to a single class, but you may want to
 explicitly choose the class name).



?



?


?&


?)*
?

?9? If set true, then the Java code generator will generate a separate .java
 file for each top-level message, enum, and service defined in the .proto
 file.  Thus, these types will *not* be nested inside the outer class
 named by java_outer_classname.  However, the outer class will still be
 generated to contain the file's getDescriptor() method as well as any
 top-level extensions defined in the file.



?



?


?#


?&(


?)8


?27
)

?E This option does nothing.



?



?


?-


?02


?3D


?4C
?

?<? If set true, then the Java2 code generator will generate code that
 throws an exception whenever an attempt is made to assign a non-UTF-8
 byte sequence to a string field.
 Message reflection will do the same.
 However, an extension field still accepts non-UTF-8 byte sequences.
 This option has no effect on when used with the lite runtime.



?



?


?&


?)+


?,;


?5:
L

 ??< Generated classes can be optimized for speed or code size.



 ?
D

  ?"4 Generate complete code for parsing, serialization,



  ?	


  ?
G

 ? etc.
"/ Use ReflectionOps to implement these methods.



 ?


 ?
G

 ?"7 Generate code using MessageLite and the lite runtime.



 ?


 ?


?9


?



?


?$


?'(


?)8


?27
?

?"? Sets the Go package where structs generated from this .proto will be
 placed. If omitted, the Go package will be derived from the following:
   - The basename of the package import path, if provided.
   - Otherwise, the package statement in the .proto file, if present.
   - Otherwise, the basename of the .proto file, without extension.



?



?


?


?!
?

?9? Should generic services be generated in each language?  "Generic" services
 are not specific to any particular RPC system.  They are generated by the
 main code generators in each language (without additional plugins).
 Generic services were the only kind of service generation supported by
 early versions of google.protobuf.

 Generic services are now considered deprecated in favor of using plugins
 that generate code specific to your particular RPC system.  Therefore,
 these default to false.  Old code which depends on generic services should
 explicitly set them to true.



?



?


?#


?&(


?)8


?27


?;


?



?


?%


?(*


?+:


?49


	?9


	?



	?


	?#


	?&(


	?)8


	?27



?:



?




?



?$



?')



?*9



?38
?

?0? Is this file deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for everything in the file, or it will be completely ignored; in the very
 least, this is a formalization for deprecating files.



?



?


?


?


? /


?).


?6q Enables the use of arenas for the proto messages in this file. This applies
 only to generated classes for C++.



?



?


? 


?#%


?&5


?/4
?

?)? Sets the objective c class prefix which is prepended to all objective c
 generated classes from this .proto. There is no default.



?



?


?#


?&(
I

?(; Namespace for generated classes; defaults to the package.



?



?


?"


?%'
?

?$? By default Swift generators will take the proto package and CamelCase it
 replacing '.' with underscore and use that to prefix the types/symbols
 defined. When this options is provided, they will use this value instead
 to prefix the types/symbols defined.



?



?


?


?!#
~

?(p Sets the php class prefix which is prepended to all php generated classes
 from this .proto. Default is empty.



?



?


?"


?%'
?

?%? Use this option to change the namespace of php generated classes. Default
 is empty. When this option is empty, the package name will be used for
 determining the namespace.



?



?


?


?"$
?

?.? Use this option to change the namespace of php generated metadata classes.
 Default is empty. When this option is empty, the proto file name will be used
 for determining the namespace.



?



?


?(


?+-
?

?$? Use this option to change the package of ruby generated classes. Default
 is empty. When this option is not set, the package name will be used for
 determining the ruby package.



?



?


?


?!#
|

?:n The parser stores options it doesn't recognize here.
 See the documentation for the "Options" section above.



?



?


?3


?69
?

?z Clients can define custom options in extensions of this message.
 See the documentation for the "Options" section above.



 ?


 ?


 ?


	?


	 ?


	 ?


	 ?

? ?

?
?
 ?<? Set true to use the old proto1 MessageSet wire format for extensions.
 This is provided for backwards-compatibility with the MessageSet wire
 format.  You should not use this for any other reason:  It's less
 efficient, has fewer features, and is more complicated.

 The message must be defined exactly as follows:
   message Foo {
     option message_set_wire_format = true;
     extensions 4 to max;
   }
 Note that the message cannot have any defined fields; MessageSets only
 have extensions.

 All extensions of your type must be singular messages; e.g. they cannot
 be int32s, enums, or repeated messages.

 Because this is an option, the above two restrictions are not enforced by
 the protocol compiler.


 ?


 ?

 ?'

 ?*+

 ?,;

 ?5:
?
?D? Disables the generation of the standard "descriptor()" accessor, which can
 conflict with a field of the same name.  This is meant to make migration
 from proto1 easier; new code should avoid fields named "descriptor".


?


?

?/

?23

?4C

?=B
?
?/? Is this message deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the message, or it will be completely ignored; in the very least,
 this is a formalization for deprecating messages.


?


?

?

?

?.

?(-
?
?? Whether the message is an automatically generated map entry type for the
 maps field.

 For maps fields:
     map<KeyType, ValueType> map_field = 1;
 The parsed descriptor looks like:
     message MapFieldEntry {
         option map_entry = true;
         optional KeyType key = 1;
         optional ValueType value = 2;
     }
     repeated MapFieldEntry map_field = 1;

 Implementations may choose not to generate the map_entry=true message, but
 use a native map in the target language to hold the keys and values.
 The reflection APIs in such implementions still need to work as
 if the field is a repeated message field.

 NOTE: Do not set the option in .proto files. Always use the maps syntax
 instead. The option should only be implicitly set by the proto compiler
 parser.


?


?

?

?
$
	?" javalite_serializable


	 ?

	 ?

	 ?

	?" javanano_as_lite


	?

	?

	?
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?.? The ctype option instructs the C++ code generator to use a different
 representation of the field than it normally would.  See the specific
 options below.  This option is not yet implemented in the open source
 release -- sorry, we'll try to include it in a future version!


 ?


 ?

 ?

 ?

 ?-

 ?&,

 ??

 ?

  ? Default mode.


  ?


  ?

 ?

 ?

 ?

 ?

 ?

 ?
?
?? The packed option can be enabled for repeated primitive fields to enable
 a more efficient representation on the wire. Rather than repeatedly
 writing the tag and type for each element, the entire array is encoded as
 a single length-delimited blob. In proto3, only explicit setting it to
 false will avoid using packed encoding.


?


?

?

?
?
?3? The jstype option determines the JavaScript type used for values of the
 field.  The option is permitted only for 64 bit integral and fixed types
 (int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING
 is represented as JavaScript string, which avoids loss of precision that
 can happen when a large value is converted to a floating point JavaScript.
 Specifying JS_NUMBER for the jstype causes the generated JavaScript code to
 use the JavaScript "number" type.  The behavior of the default option
 JS_NORMAL is implementation dependent.

 This option is an enum to permit additional types to be added, e.g.
 goog.math.Integer.


?


?

?

?

?2

?(1

??

?
'
 ? Use the default type.


 ?

 ?
)
? Use JavaScript strings.


?

?
)
? Use JavaScript numbers.


?

?
?
?)? Should this field be parsed lazily?  Lazy applies only to message-type
 fields.  It means that when the outer message is initially parsed, the
 inner message's contents will not be parsed but instead stored in encoded
 form.  The inner message will actually be parsed when it is first accessed.

 This is only a hint.  Implementations are free to choose whether to use
 eager or lazy parsing regardless of the value of this option.  However,
 setting this option true suggests that the protocol author believes that
 using lazy parsing on this field is worth the additional bookkeeping
 overhead typically needed to implement it.

 This option does not affect the public interface of any generated code;
 all method signatures remain the same.  Furthermore, thread-safety of the
 interface is not affected by this option; const methods remain safe to
 call from multiple threads concurrently, while non-const methods continue
 to require exclusive access.


 Note that implementations may choose not to check required fields within
 a lazy sub-message.  That is, calling IsInitialized() on the outer message
 may return true even if the inner message has missing required fields.
 This is necessary because otherwise the inner message would have to be
 parsed in order to perform the check, defeating the purpose of lazy
 parsing.  An implementation which chooses not to check required fields
 must be consistent about it.  That is, for any particular sub-message, the
 implementation must either *always* check its required fields, or *never*
 check its required fields, regardless of whether or not the message has
 been parsed.


?


?

?

?

?(

?"'
?
?/? Is this field deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for accessors, or it will be completely ignored; in the very least, this
 is a formalization for deprecating fields.


?


?

?

?

?.

?(-
?
?*1 For Google-internal migration only. Do not use.


?


?

?

?

?)

?#(
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

	?" removed jtype


	 ?

	 ?

	 ?

? ?

?
O
 ?:A The parser stores options it doesn't recognize here. See above.


 ?


 ?

 ?3

 ?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
`
 ? R Set this option to true to allow mapping different tag names to the same
 value.


 ?


 ?

 ?

 ?
?
?/? Is this enum deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the enum, or it will be completely ignored; in the very least, this
 is a formalization for deprecating enums.


?


?

?

?

?.

?(-

	?" javanano_as_lite


	 ?

	 ?

	 ?
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?/? Is this enum value deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the enum value, or it will be completely ignored; in the very least,
 this is a formalization for deprecating enum values.


 ?


 ?

 ?

 ?

 ?.

 ?(-
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?0? Is this service deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the service, or it will be completely ignored; in the very least,
 this is a formalization for deprecating services.
2? Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
   framework.  We apologize for hoarding these numbers to ourselves, but
   we were already using them long before we decided to release Protocol
   Buffers.


 ?


 ?

 ?

 ?

 ? /

 ?).
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?0? Is this method deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the method, or it will be completely ignored; in the very least,
 this is a formalization for deprecating methods.
2? Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
   framework.  We apologize for hoarding these numbers to ourselves, but
   we were already using them long before we decided to release Protocol
   Buffers.


 ?


 ?

 ?

 ?

 ? /

 ?).
?
 ??? Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
 or neither? HTTP based RPC implementation may choose GET verb for safe
 methods, and PUT verb for idempotent methods instead of the default POST.


 ?

  ?

  ?

  ?
$
 ?" implies idempotent


 ?

 ?
7
 ?"' idempotent, but may have side effects


 ?

 ?

??'

?


?

?-

?

?	&

?%
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?
?
? ?? A message representing a option the parser does not recognize. This only
 appears in options protos created by the compiler::Parser class.
 DescriptorPool resolves these when building Descriptor objects. Therefore,
 options protos in descriptor objects (e.g. returned by Descriptor::options(),
 or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
 in them.


?
?
 ??? The name of the uninterpreted option.  Each string represents a segment in
 a dot-separated name.  is_extension is true iff a segment represents an
 extension (denoted with parentheses in options specs in .proto files).
 E.g.,{ ["foo", false], ["bar.baz", true], ["qux", false] } represents
 "foo.(bar.baz).qux".


 ?


  ?"

  ?

  ?

  ?

  ? !

 ?#

 ?

 ?

 ?

 ?!"

 ?

 ?


 ?

 ?

 ?
?
?'? The value of the uninterpreted option, in whatever type the tokenizer
 identified it as during parsing. Exactly one of these should be set.


?


?

?"

?%&

?)

?


?

?$

?'(

?(

?


?

?#

?&'

?#

?


?

?

?!"

?"

?


?

?

? !

?&

?


?

?!

?$%
?
? ?j Encapsulates information about the original source file from which a
 FileDescriptorProto was generated.
2` ===================================================================
 Optional source code info


?
?
 ?!? A Location identifies a piece of source code in a .proto file which
 corresponds to a particular definition.  This information is intended
 to be useful to IDEs, code indexers, documentation generators, and similar
 tools.

 For example, say we have a file like:
   message Foo {
     optional string foo = 1;
   }
 Let's look at just the field definition:
   optional string foo = 1;
   ^       ^^     ^^  ^  ^^^
   a       bc     de  f  ghi
 We have the following locations:
   span   path               represents
   [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
   [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
   [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
   [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
   [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).

 Notes:
 - A location may refer to a repeated field itself (i.e. not to any
   particular index within it).  This is used whenever a set of elements are
   logically enclosed in a single code segment.  For example, an entire
   extend block (possibly containing multiple extension definitions) will
   have an outer location whose path refers to the "extensions" repeated
   field without an index.
 - Multiple locations may have the same path.  This happens when a single
   logical declaration is spread out across multiple places.  The most
   obvious example is the "extend" block again -- there may be multiple
   extend blocks in the same scope, each of which will have the same path.
 - A location's span is not always a subset of its parent's span.  For
   example, the "extendee" of an extension declaration appears at the
   beginning of the "extend" block and is shared by all extensions within
   the block.
 - Just because a location's span is a subset of some other location's span
   does not mean that it is a descendent.  For example, a "group" defines
   both a type and a field in a single declaration.  Thus, the locations
   corresponding to the type and field and their components will overlap.
 - Code which tries to interpret locations should probably be designed to
   ignore those that it doesn't understand, as more types of locations could
   be recorded in the future.


 ?


 ?

 ?

 ? 

 ??

 ?

?
  ?*? Identifies which part of the FileDescriptorProto was defined at this
 location.

 Each element is a field number or an index.  They form a path from
 the root FileDescriptorProto to the place where the definition.  For
 example, this path:
   [ 4, 3, 2, 7, 1 ]
 refers to:
   file.message_type(3)  // 4, 3
       .field(7)         // 2, 7
       .name()           // 1
 This is because FileDescriptorProto.message_type has field number 4:
   repeated DescriptorProto message_type = 4;
 and DescriptorProto.field has field number 2:
   repeated FieldDescriptorProto field = 2;
 and FieldDescriptorProto.name has field number 1:
   optional string name = 1;

 Thus, the above path gives the location of a field name.  If we removed
 the last element:
   [ 4, 3, 2, 7 ]
 this path refers to the whole field declaration (from the beginning
 of the label to the terminating semicolon).


  ?

  ?

  ?

  ?

  ?)

  ?(
?
 ?*? Always has exactly three or four elements: start line, start column,
 end line (optional, otherwise assumed same as start line), end column.
 These are packed into a single field for efficiency.  Note that line
 and column numbers are zero-based -- typically you will want to add
 1 to each before displaying to a user.


 ?

 ?

 ?

 ?

 ?)

 ?(
?
 ?)? If this SourceCodeInfo represents a complete declaration, these are any
 comments appearing before and after the declaration which appear to be
 attached to the declaration.

 A series of line comments appearing on consecutive lines, with no other
 tokens appearing on those lines, will be treated as a single comment.

 leading_detached_comments will keep paragraphs of comments that appear
 before (but not connected to) the current element. Each paragraph,
 separated by empty lines, will be one comment element in the repeated
 field.

 Only the comment content is provided; comment markers (e.g. //) are
 stripped out.  For block comments, leading whitespace and an asterisk
 will be stripped from the beginning of each line other than the first.
 Newlines are included in the output.

 Examples:

   optional int32 foo = 1;  // Comment attached to foo.
   // Comment attached to bar.
   optional int32 bar = 2;

   optional string baz = 3;
   // Comment attached to baz.
   // Another line attached to baz.

   // Comment attached to qux.
   //
   // Another line attached to qux.
   optional double qux = 4;

   // Detached comment for corge. This is not leading or trailing comments
   // to qux or corge because there are blank lines separating it from
   // both.

   // Detached comment for corge paragraph 2.

   optional string corge = 5;
   /* Block comment attached
    * to corge.  Leading asterisks
    * will be removed. */
   /* Block comment attached to
    * grault. */
   optional int32 grault = 6;

   // ignored detached comments.


 ?

 ?

 ?$

 ?'(

 ?*

 ?

 ?

 ?%

 ?()

 ?2

 ?

 ?

 ?-

 ?01
?
? ?? Describes the relationship between generated code and its original source
 file. A GeneratedCodeInfo message is associated with only one generated
 source file, but may contain references to different source .proto files.


?
x
 ?%j An Annotation connects some span of text in generated code to an element
 of its generating .proto file.


 ?


 ?

 ? 

 ?#$

 ??

 ?

?
  ?* Identifies the element in the original source .proto file. This field
 is formatted the same as SourceCodeInfo.Location.path.


  ?

  ?

  ?

  ?

  ?)

  ?(
O
 ?$? Identifies the filesystem path to the original source .proto.


 ?

 ?

 ?

 ?"#
w
 ?g Identifies the starting offset in bytes in the generated code
 that relates to the identified object.


 ?

 ?

 ?

 ?
?
 ?? Identifies the ending offset in bytes in the generated code that
 relates to the identified offset. The end offset should be one past
 the last relevant byte (so the length of the text = end - begin).


 ?

 ?

 ?

 ?
?
google/api/annotations.proto
google.apigoogle/api/http.proto google/protobuf/descriptor.proto:K
http.google.protobuf.MethodOptions?ʼ" (2.google.api.HttpRuleRhttpBn
com.google.apiBAnnotationsProtoPZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations?GAPIJ?
 
?
 2? Copyright (c) 2015, Google Inc.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.



	
 
	
)

 X
	
 X

 "
	

 "

 1
	
 1

 '
	
 '

 "
	
$ "
	
 

  See `HttpRule`.



 $

 &


 



 


 bproto3
?
google/api/field_behavior.proto
google.api google/protobuf/descriptor.proto*{
FieldBehavior
FIELD_BEHAVIOR_UNSPECIFIED 
OPTIONAL
REQUIRED
OUTPUT_ONLY

INPUT_ONLY
	IMMUTABLE:`
field_behavior.google.protobuf.FieldOptions? (2.google.api.FieldBehaviorRfieldBehaviorBp
com.google.apiBFieldBehaviorProtoPZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations?GAPIJ?
 N
?
 2? Copyright 2019 Google LLC.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.




	
 )

 X
	
 X

 "
	

 "

 3
	
 3

 '
	
 '

 "
	
$ "
	
 )
?
 (:? A designation of a specific field behavior (required, output only, etc.)
 in protobuf messages.

 Examples:

   string name = 1 [(google.api.field_behavior) = REQUIRED];
   State state = 1 [(google.api.field_behavior) = OUTPUT_ONLY];
   google.protobuf.Duration ttl = 1
     [(google.api.field_behavior) = INPUT_ONLY];
   google.protobuf.Timestamp expire_time = 1
     [(google.api.field_behavior) = OUTPUT_ONLY,
      (google.api.field_behavior) = IMMUTABLE];



 #


 (



 (#


 ($2


 (59
?
 1 N? An indicator of the behavior of a given field (for example, that a field
 is required in requests, or given as output but ignored as input).
 This **does not** change the behavior in protocol buffers itself; it only
 denotes the behavior and may affect how API tooling handles the field.

 Note: This enum **may** receive new values in the future.



 1
?
  3!2 Conventional default for enums. Do not use this.


  3

  3 
?
 8? Specifically denotes a field as optional.
 While all fields in protocol buffers are optional, this may be specified
 for emphasis if appropriate.


 8


 8
?
 =? Denotes a field as required.
 This indicates that the field **must** be provided as part of the request,
 and failure to do so will cause an error (usually `INVALID_ARGUMENT`).


 =


 =
?
 C? Denotes a field as output only.
 This indicates that the field is provided in responses, but including the
 field in a request does nothing (the server *must* ignore it and
 *must not* throw an error as a result of the field's presence).


 C

 C
?
 H? Denotes a field as input only.
 This indicates that the field is provided in requests, and the
 corresponding field is not included in output.


 H

 H
?
 M? Denotes a field as immutable.
 This indicates that the field may be set once in a request to create a
 resource, but may not be changed thereafter.


 M

 Mbproto3
?R
google/api/resource.proto
google.api google/protobuf/descriptor.proto"?
ResourceDescriptor
type (	Rtype
pattern (	Rpattern

name_field (	R	nameField@
history (2&.google.api.ResourceDescriptor.HistoryRhistory
plural (	Rplural
singular (	Rsingular"[
History
HISTORY_UNSPECIFIED 
ORIGINALLY_SINGLE_PATTERN
FUTURE_MULTI_PATTERN"F
ResourceReference
type (	Rtype

child_type (	R	childType:l
resource_reference.google.protobuf.FieldOptions? (2.google.api.ResourceReferenceRresourceReference:n
resource_definition.google.protobuf.FileOptions? (2.google.api.ResourceDescriptorRresourceDefinition:\
resource.google.protobuf.MessageOptions? (2.google.api.ResourceDescriptorRresourceBn
com.google.apiBResourceProtoPZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations??GAPIJ?K
 ?
?
 2? Copyright 2019 Google LLC.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.




	
 )

 
	
 

 X
	
 X

 "
	

 "

 .
	
 .

 '
	
 '

 "
	
$ "
	
  
[
 9P An annotation that describes a resource reference, see
 [ResourceReference][].



 #

 %


 


 1


 48
	
" &
}
%Dr An annotation that describes a resource definition without a corresponding
 message; see [ResourceDescriptor][].



""


%



%(


%)<


%?C
	
( ,
]
+0R An annotation that describes a resource definition, see
 [ResourceDescriptor][].



(%

+('


+


+ (


++/
? 
 ? ?? A simple descriptor of a resource type.

 ResourceDescriptor annotates a resource message (either by means of a
 protobuf annotation or use in the service config), and associates the
 resource's schema, the resource type, and the pattern of the resource name.

 Example:

     message Topic {
       // Indicates this message defines a resource schema.
       // Declares the resource type in the format of {service}/{kind}.
       // For Kubernetes resources, the format is {api group}/{kind}.
       option (google.api.resource) = {
         type: "pubsub.googleapis.com/Topic"
         name_descriptor: {
           pattern: "projects/{project}/topics/{topic}"
           parent_type: "cloudresourcemanager.googleapis.com/Project"
           parent_name_extractor: "projects/{project}"
         }
       };
     }

 The ResourceDescriptor Yaml config will look like:

    resources:
    - type: "pubsub.googleapis.com/Topic"
      name_descriptor:
        - pattern: "projects/{project}/topics/{topic}"
          parent_type: "cloudresourcemanager.googleapis.com/Project"
          parent_name_extractor: "projects/{project}"

 Sometimes, resources have multiple patterns, typically because they can
 live under multiple parents.

 Example:

     message LogEntry {
       option (google.api.resource) = {
         type: "logging.googleapis.com/LogEntry"
         name_descriptor: {
           pattern: "projects/{project}/logs/{log}"
           parent_type: "cloudresourcemanager.googleapis.com/Project"
           parent_name_extractor: "projects/{project}"
         }
         name_descriptor: {
           pattern: "folders/{folder}/logs/{log}"
           parent_type: "cloudresourcemanager.googleapis.com/Folder"
           parent_name_extractor: "folders/{folder}"
         }
         name_descriptor: {
           pattern: "organizations/{organization}/logs/{log}"
           parent_type: "cloudresourcemanager.googleapis.com/Organization"
           parent_name_extractor: "organizations/{organization}"
         }
         name_descriptor: {
           pattern: "billingAccounts/{billing_account}/logs/{log}"
           parent_type: "billing.googleapis.com/BillingAccount"
           parent_name_extractor: "billingAccounts/{billing_account}"
         }
       };
     }

 The ResourceDescriptor Yaml config will look like:

     resources:
     - type: 'logging.googleapis.com/LogEntry'
       name_descriptor:
         - pattern: "projects/{project}/logs/{log}"
           parent_type: "cloudresourcemanager.googleapis.com/Project"
           parent_name_extractor: "projects/{project}"
         - pattern: "folders/{folder}/logs/{log}"
           parent_type: "cloudresourcemanager.googleapis.com/Folder"
           parent_name_extractor: "folders/{folder}"
         - pattern: "organizations/{organization}/logs/{log}"
           parent_type: "cloudresourcemanager.googleapis.com/Organization"
           parent_name_extractor: "organizations/{organization}"
         - pattern: "billingAccounts/{billing_account}/logs/{log}"
           parent_type: "billing.googleapis.com/BillingAccount"
           parent_name_extractor: "billingAccounts/{billing_account}"

 For flexible resources, the resource name doesn't contain parent names, but
 the resource itself has parents for policy evaluation.

 Example:

     message Shelf {
       option (google.api.resource) = {
         type: "library.googleapis.com/Shelf"
         name_descriptor: {
           pattern: "shelves/{shelf}"
           parent_type: "cloudresourcemanager.googleapis.com/Project"
         }
         name_descriptor: {
           pattern: "shelves/{shelf}"
           parent_type: "cloudresourcemanager.googleapis.com/Folder"
         }
       };
     }

 The ResourceDescriptor Yaml config will look like:

     resources:
     - type: 'library.googleapis.com/Shelf'
       name_descriptor:
         - pattern: "shelves/{shelf}"
           parent_type: "cloudresourcemanager.googleapis.com/Project"
         - pattern: "shelves/{shelf}"
           parent_type: "cloudresourcemanager.googleapis.com/Folder"


 ?
c
  ??S A description of the historical or future-looking state of the
 resource pattern.


  ?
$
   ? The "unset" value.


   ?

   ?
z
  ?"j The resource originally had one pattern and launched as such, and
 additional patterns were added later.


  ?

  ? !
?
  ?? The resource has one pattern, but the API owner expects to add more
 later. (This is the inverse of ORIGINALLY_SINGLE_PATTERN, and prevents
 that from being necessary once there are multiple patterns.)


  ?

  ?
?
  ?? The resource type. It must be in the format of
 {service_name}/{resource_type_kind}. The `resource_type_kind` must be
 singular and must not include version numbers.

 Example: `storage.googleapis.com/Bucket`

 The value of the resource_type_kind must follow the regular expression
 /[A-Za-z][a-zA-Z0-9]+/. It should start with an upper case character and
 should use PascalCase (UpperCamelCase). The maximum number of
 characters allowed for the `resource_type_kind` is 100.


  ??

  ?

  ?	

  ?
?
 ?? Optional. The relative resource name pattern associated with this resource
 type. The DNS prefix of the full resource name shouldn't be specified here.

 The path pattern must follow the syntax, which aligns with HTTP binding
 syntax:

     Template = Segment { "/" Segment } ;
     Segment = LITERAL | Variable ;
     Variable = "{" LITERAL "}" ;

 Examples:

     - "projects/{project}/topics/{topic}"
     - "projects/{project}/knowledgeBases/{knowledge_base}"

 The components in braces correspond to the IDs for each resource in the
 hierarchy. It is expected that, if multiple patterns are provided,
 the same component name (e.g. "project") refers to IDs of the same
 type of resource.


 ?


 ?

 ?

 ?
?
 ?y Optional. The field on the resource that designates the resource name
 field. If omitted, this is assumed to be "name".


 ??

 ?

 ?	

 ?
?
 ?? Optional. The historical or future-looking state of the resource pattern.

 Example:

     // The InspectTemplate message originally only supported resource
     // names with organization, and project was added later.
     message InspectTemplate {
       option (google.api.resource) = {
         type: "dlp.googleapis.com/InspectTemplate"
         pattern:
         "organizations/{organization}/inspectTemplates/{inspect_template}"
         pattern: "projects/{project}/inspectTemplates/{inspect_template}"
         history: ORIGINALLY_SINGLE_PATTERN
       };
     }


 ??

 ?	

 ?


 ?
?
 ?? The plural name used in the resource name, such as 'projects' for
 the name of 'projects/{project}'. It is the same concept of the `plural`
 field in k8s CRD spec
 https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/


 ??

 ?

 ?	

 ?
?
 ?? The same concept of the `singular` field in k8s CRD spec
 https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/
 Such as "project" for the `resourcemanager.googleapis.com/Project` type.


 ??

 ?

 ?	

 ?
i
? ?[ Defines a proto annotation that describes a string field that refers to
 an API resource.


?
?
 ?? The resource type that the annotated field references.

 Example:

     message Subscription {
       string topic = 2 [(google.api.resource_reference) = {
         type: "pubsub.googleapis.com/Topic"
       }];
     }


 ??

 ?

 ?	

 ?
?
?? The resource type of a child collection that the annotated field
 references. This is useful for annotating the `parent` field that
 doesn't have a fixed resource type.

 Example:

   message ListLogEntriesRequest {
     string parent = 1 [(google.api.resource_reference) = {
       child_type: "logging.googleapis.com/LogEntry"
     };
   }


??

?

?	

?bproto3
?/
google/protobuf/timestamp.protogoogle.protobuf";
	Timestamp
seconds (Rseconds
nanos (RnanosB~
com.google.protobufBTimestampProtoPZ+github.com/golang/protobuf/ptypes/timestamp??GPB?Google.Protobuf.WellKnownTypesJ?-
 ?
?
 2? Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


 

" ;
	
%" ;

# 
	
# 

$ B
	
$ B

% ,
	
% ,

& /
	
& /

' "
	

' "

( !
	
$( !
?
 z ?? A Timestamp represents a point in time independent of any time zone
 or calendar, represented as seconds and fractions of seconds at
 nanosecond resolution in UTC Epoch time. It is encoded using the
 Proleptic Gregorian Calendar which extends the Gregorian calendar
 backwards to year one. It is encoded assuming all minutes are 60
 seconds long, i.e. leap seconds are "smeared" so that no leap second
 table is needed for interpretation. Range is from
 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z.
 By restricting to that range, we ensure that we can convert to
 and from  RFC 3339 date strings.
 See [https://www.ietf.org/rfc/rfc3339.txt](https://www.ietf.org/rfc/rfc3339.txt).

 # Examples

 Example 1: Compute Timestamp from POSIX `time()`.

     Timestamp timestamp;
     timestamp.set_seconds(time(NULL));
     timestamp.set_nanos(0);

 Example 2: Compute Timestamp from POSIX `gettimeofday()`.

     struct timeval tv;
     gettimeofday(&tv, NULL);

     Timestamp timestamp;
     timestamp.set_seconds(tv.tv_sec);
     timestamp.set_nanos(tv.tv_usec * 1000);

 Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

     FILETIME ft;
     GetSystemTimeAsFileTime(&ft);
     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
     Timestamp timestamp;
     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

 Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

     long millis = System.currentTimeMillis();

     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
         .setNanos((int) ((millis % 1000) * 1000000)).build();


 Example 5: Compute Timestamp from current time in Python.

     timestamp = Timestamp()
     timestamp.GetCurrentTime()

 # JSON Mapping

 In JSON format, the Timestamp type is encoded as a string in the
 [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
 format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
 where {year} is always expressed using four digits while {month}, {day},
 {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
 seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
 are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
 is required. A proto3 JSON serializer should always use UTC (as indicated by
 "Z") when printing the Timestamp type and a proto3 JSON parser should be
 able to accept both UTC and other timezones (as indicated by an offset).

 For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
 01:30 UTC on January 15, 2017.

 In JavaScript, one can convert a Date object to this format using the
 standard [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString]
 method. In Python, a standard `datetime.datetime` object can be converted
 to this format using [`strftime`](https://docs.python.org/2/library/time.html#time.strftime)
 with the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one
 can use the Joda Time's [`ISODateTimeFormat.dateTime()`](
 http://www.joda.org/joda-time/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime--
 ) to obtain a formatter capable of generating timestamps in this format.





 z
?
  ? Represents seconds of UTC time since Unix epoch
 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
 9999-12-31T23:59:59Z inclusive.


  z

  

  

  
?
 ?? Non-negative fractions of a second at nanosecond resolution. Negative
 second values with fractions must still have non-negative nanos values
 that count forward in time. Must be from 0 to 999,999,999
 inclusive.


 ?

 ?

 ?

 ?bproto3
??
(infra/chromeperf/pinpoint/pinpoint.protopinpointgoogle/api/annotations.protogoogle/api/field_behavior.protogoogle/api/resource.protogoogle/protobuf/timestamp.proto"?
TelemetryBenchmark
	benchmark (	R	benchmark
story (	H RstoryJ

story_tags (2).pinpoint.TelemetryBenchmark.StoryTagListH R	storyTags 
measurement (	Rmeasurement%
grouping_label (	RgroupingLabelD
	statistic (2&.pinpoint.TelemetryBenchmark.StatisticR	statistic

extra_args (	R	extraArgs-
StoryTagList

story_tags (	R	storyTags"{
	Statistic
STATISTIC_UNSPECIFIED 
NONE
MIN
MAX
MEAN
STD_DEV

MEDIAN	
PCT90	
PCT99B
story_selection"d
GTestBenchmark
	benchmark (	R	benchmark 
measurement (	Rmeasurement
test (	Rtest"g
GitilesCommit
host (	B?ARhost
project (	B?ARproject
git_hash (	B?ARgitHash"?
GerritChange
host (	B?ARhost
project (	B?ARproject
change (B?ARchange
patchset (B?ARpatchset"?
GitilesCommitRange
host (	B?ARhost
project (	B?ARproject)
start_git_hash (	B?ARstartGitHash%
end_git_hash (	B?AR
endGitHash"
	BisectionD
commit_range (2.pinpoint.GitilesCommitRangeB?ARcommitRange,
patch (2.pinpoint.GerritChangeRpatch"?

Experiment=
base_commit (2.pinpoint.GitilesCommitB?AR
baseCommit5

base_patch (2.pinpoint.GerritChangeR	basePatchD
experiment_commit (2.pinpoint.GitilesCommitRexperimentCommitF
experiment_patch (2.pinpoint.GerritChangeB?ARexperimentPatch"N
MonorailIssue
project (	B?ARproject
issue_id (B?ARissueId"?
JobSpecI
comparison_mode (2 .pinpoint.JobSpec.ComparisonModeRcomparisonMode1
comparison_magnitude (RcomparisonMagnitude
config (	B?ARconfig
target (	B?ARtarget3
	bisection (2.pinpoint.BisectionH R	bisection6

experiment (2.pinpoint.ExperimentH R
experimentO
telemetry_benchmark (2.pinpoint.TelemetryBenchmarkHRtelemetryBenchmarkC
gtest_benchmark (2.pinpoint.GTestBenchmarkHRgtestBenchmark>
monorail_issue	 (2.pinpoint.MonorailIssueRmonorailIssue"R
ComparisonMode
COMPARISON_MODE_UNSPECIFIED 
PERFORMANCE

FUNCTIONALB

job_kindB
	arguments">
ScheduleJobRequest(
job (2.pinpoint.JobSpecB?ARjob"?
Job
id (	Rid)
state (2.pinpoint.Job.StateRstate

created_by (	R	createdBy;
create_time (2.google.protobuf.TimestampR
createTimeD
last_update_time (2.google.protobuf.TimestampRlastUpdateTime,
job_spec (2.pinpoint.JobSpecRjobSpec/
cancellation_reason (	RcancellationReason"b
State
STATE_UNSPECIFIED 
PENDING
RUNNING
	SUCCEEDED

FAILED
	CANCELLED"@
GetJobRequest/
id (	B?A?A
api.pinpoint.cr.dev/JobRid"e
ListJobsRequest
	page_size (RpageSize

page_token (	R	pageToken
filter (	Rfilter"]
ListJobsResponse!
jobs (2.pinpoint.JobRjobs&
next_page_token (	RnextPageToken"`
CancelJobRequest/
id (	B?A?A
api.pinpoint.cr.dev/JobRid
reason (	B?ARreason2?
PinpointQ
ScheduleJob.pinpoint.ScheduleJobRequest.pinpoint.Job"????"/v1/jobs:jobD
GetJob.pinpoint.GetJobRequest.pinpoint.Job"????
/v1/jobs/*S
ListJobs.pinpoint.ListJobsRequest.pinpoint.ListJobsResponse"????
/v1/jobsS
	CancelJob.pinpoint.CancelJobRequest.pinpoint.Job"????"/v1/job/*:cancel:*J?i
 ?
?
 2? Copyright 2020 The Chromium Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.



	
 %
	
(
	
"
	
(
4
  :( Service definition for Pinpoint's API.



 
7
  !) Schedules a Pinpoint Job for execution.


  

  $

  /2

  
?
	  ?ʼ""p TODO(dberris): Determine if we need the client API convenience.
 option (google.api.method_signature) = "job";

7
 $*) Retrieves details about a Pinpoint Job.


 $

 $

 $%(

 %'
?
	 ?ʼ"%'"o TODO(dberris): Determine if we need the client API convenience.
 option (google.api.method_signature) = "id";

(
 -1 Lists jobs with filters.


 -

 -

 -)9

 .0

	 ?ʼ".0
'
 49 Cancels an ongoing job.


 4

 4 

 4+.

 58

	 ?ʼ"58
P
 = }D Specifies the details for data generated by a Telemetry benchmark.



 =
)
  ? The name of the benchmark.


  ?=

  ?

  ?	

  ?

  AF

  A

?
   E#? A list of story tags which specify stories that will be run. All stories
 that match any of the specified tags will be collected (unioned) in the
 run.


   E

   E

   E

   E!"

  GN

  G

 I A story to run.


 I


 I

 I
l
 M _ NOTE: We use a nested message in this proto oneof because it cannot
 include repeated fields.


 M

 M

 M
?
 Z? The name of the measurement to extract. This refers to the name of the
 histogram in a HistogramSet which we'll be making statistical analysis
 decisions on.

 See
 https://chromium.googlesource.com/catapult.git/+/HEAD/docs/how-to-write-metrics.md#histograms
 and
 https://chromium.googlesource.com/catapult/+/HEAD/docs/histogram-set-json-format.md
 for details on what metric names can be and how to select those from
 histograms produced by Telemetry benchmarks.


 ZN

 Z

 Z	

 Z
?
 _? The optional grouping label, as described in
 https://chromium.googlesource.com/catapult.git/+/HEAD/docs/how-to-write-metrics.md#reserved-names
 (as `labels`, passed as a flag to Telemetry via the --results-label flag).


 _Z

 _

 _	

 _
Z
  buL Supported statistics for when performing statistical analysis/comparisons.


  b
m
   e^ If unspecified, lets Pinpoint determine the default statistic to use
 depending on the mode.


   e

   e
]
  iN Use all the sample values when performing statistical
 analysis/comparisons.


  i

  i
?
  n? All the following values select the appropriate statistic across
 multiple attempts, and performs statistical analysis on these values,
 instead of the raw sample values.


  n

  n


  o

  o

  o


  p

  p

  p

  q

  q

  q

  r

  r


  r

  s

  s	

  s

  t

  t	

  t
?
 y? The statistic to extract. If unspecified defaults to "NONE", i.e. use all
 the sample values instead of computing a statistic.


 yu

 y

 y

 y
F
 |!9 List of extra arguments passed to the Telemetry runner.


 |


 |

 |

 | 
L
? ?> Specify the details for data generated by a GTest benchmark.


?
?
 ?? Name of the benchmark. This is provided to the GTest benchmark through the
 '--gtest-benchmark-name' flag, to identify which benchmark in the `target`
 (provided in the JobSpec) will be run.


 ??

 ?

 ?	

 ?
R
?D Name of the measurement that shows up as a chart in the Dashboard.


??

?

?	

?
Y
?K Name of the testcase in the benchmark, supplied as a filter to the GTest.


??

?

?	

?
0
? ?" A commit in a Gitiles repostory.


?
6
 ?;( The gitiles host to query for commits.


 ??

 ?

 ?	

 ?

 ?:

 ? ?9
9
?>+ The gitiles project to query for commits.


??;

?

?	

?

?=

? ?<
5
??' The hash for the specific Git commit.


??>

?

?	

?

?>

? ?=
O
? ?A A Gerrit change review, to refer to patches to apply in builds.


?
 
 ?; Gerrit hostname.


 ??

 ?

 ?	

 ?

 ?:

 ? ?9
+
?> Gerrit project in the host.


??;

?

?	

?

?=

? ?<
<
?<. A change ID against the project in the host.


??>

?

?

?

?;

? ?:
]
?>O A reference to the patch set. If unspecified this will always be the
 latest.


??<

?

?

?

?=

? ?<
?
? ?? A CommitRange consists of two git commit hashes from the same project that
 the end commit must be reachable from the start commit.


?
6
 ?;( The gitiles host to query for commits.


 ??

 ?

 ?	

 ?

 ?:

 ? ?9
9
?>+ The gitiles project to query for commits.


??;

?

?	

?

?=

? ?<
.
?E  The start of the commit range.


??>

?

?	

?

?D

? ?C
,
?C The end of the commit range.


??E

?

?	

?

?B

? ?A
?
? ?t A Bisection specifies a range through which Pinpoint will perform
 builds/runs against to find potential culprits.


?
M
 ?O? A commit range through which the bisection will be performed.


 ??

 ?

 ?!

 ?$%

 ?&N

 ? ?'M
?
?? A patch to apply to all commits in the range before building. This is
 useful if you want to apply a patch that fixes the build in a range where
 the bisection might encounter broken builds.


??O

?

?

?
?
? ?? An Experiment defines an A/B test between a base build (with an optional
 patch) and an experimental build. This is also referred to as a "tryjob" in
 Pinpoint.


?
S
 ?IE The commit and optional patch from which we'd perform the A/B Test.


 ??

 ?

 ?

 ?

 ? H

 ? ?!G

?

??I

?

?

?
?
?&? The commit and required patch which we'll compare against the base and
 optional patch. If the experiment_commit is empty, we use the base_commit
 instead.


??

?

?!

?$%

?M

??&

?

?

?"#

?$L

? ?%K

? ?

?
-
 ?> Name of the Monorail project.


 ??

 ?

 ?	

 ?

 ?=

 ? ?<
5
?>' The Issue ID in the provided project.


??>

?

?

?

?=

? ?<
s
? ?e A JobSpec defines inputs for Pinpoint to configure a job being created and
 managed by the service.


?
C
 ??3 The comparison mode used by Pinpoint in this job.


 ?
I
  ?$9 If unspecified, it's as if 'PERFORMANCE' was specified.


  ?

  ?"#
V
 ?F Performs statistical tests appropriate for performance measurements.


 ?

 ?
?
 ?? Performs statistical tests appropriate for functional success/failure
 measurements. This is used in determining whether a benchmark has failed
 or started to become flaky in a bisection or to try out whether a patch
 will cause failures for specific benchmarks/stories in experiments.


 ?

 ?
Y
 ?%K The comparison mode to perform. Defaults to 'PERFORMANCE' if unspecified.


 ??

 ?

 ? 

 ?#$
?
?"? A threshold value which determines how sensitive the comparison algorithm
 should be. This is typically a multiple of the inter-quartile range for
 the measurements, indicating how "tight" the bounds for differences must
 be when considering statistical significance.

 If unspecified, the service will determine a default comparison magnitude
 based on the empirically measured inter-quartile range.


??%

?

?	

? !
?
?=? A named configuration, representing a "bot pool" in Pinpoint.

 These are managed through the administration console on the Chromeperf UI.

 TODO(dberris): Add a reference to the luci-config files when these have
 been migrated.


??"

?

?	

?

?<

? ?;
Z
?=L The target representing a build target that will be built/run by Pinpoint.


??=

?

?	

?

?<

? ?;

 ??

 ?

?

?

?

?

?

?

?

?

??

?

?/

?

?*

?-.

?'

?

?"

?%&
.
?#  The associated Monorail issue.


??

?

?

?!"
c
	? ?U Contains the required information for creating a Job represented in
 the datastore.


	?

	 ?;

	 ??

	 ?	

	 ?


	 ?

	 ?:

	 ? ?9
3

? ?% A representation of a Pinpoint Job.



?


 ?


 ??


 ?


 ?	


 ?


 ??


 ?


  ?


  ?


  ?


 ?


 ?


 ?


 ?


 ?


 ?


 ?


 ?


 ?


 ?


 ?



 ?


 ?


 ?


 ?


?


??


?


?


?
M

?? The user that created this Job, referred to by email address.



??


?


?	


?


?,


??


?


?'


?*+


?1


??,


?


?,


?/0
@

?2 Specification provided when the Job was created.



??1


?	


?



?
L

?!> If present, the reason provided for when a job is cancelled.



??


?


?	


? 

? ?

?

 ??

 ??

 ?

 ?	

 ?

 ??

 ? ?*

 ??I

? ?

?
y
 ?k If unspecified, at most 100 Jobs will be returned.
 Max is 1000, any values above will be capped to 1000.


 ??

 ?

 ?

 ?
?
?? A page token, received from a previous `ListJobs` call.
 Provide this to retrieve the subsequent page.
 When paginating, all other parameters provided must match previous
 requests.


??

?

?	

?
?
?? A structured string specifying a filter on Job properties.
 TODO(dberris): Document this, see https://aip.dev/160 for details.


??

?

?	

?

? ?

?

 ?

 ?


 ?

 ?

 ?

?

??

?

?	

?

? ?

?

 ??

 ??

 ?

 ?	

 ?

 ??

 ? ?*

 ??I

?=

??

?

?	

?

?<

? ?;bproto3