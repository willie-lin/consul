// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/overload/v3/overload.proto

package envoy_config_overload_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResourceMonitor struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ResourceMonitor_HiddenEnvoyDeprecatedConfig
	//	*ResourceMonitor_TypedConfig
	ConfigType           isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResourceMonitor) Reset()         { *m = ResourceMonitor{} }
func (m *ResourceMonitor) String() string { return proto.CompactTextString(m) }
func (*ResourceMonitor) ProtoMessage()    {}
func (*ResourceMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{0}
}

func (m *ResourceMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMonitor.Unmarshal(m, b)
}
func (m *ResourceMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMonitor.Marshal(b, m, deterministic)
}
func (m *ResourceMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMonitor.Merge(m, src)
}
func (m *ResourceMonitor) XXX_Size() int {
	return xxx_messageInfo_ResourceMonitor.Size(m)
}
func (m *ResourceMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMonitor proto.InternalMessageInfo

func (m *ResourceMonitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
}

type ResourceMonitor_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type ResourceMonitor_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ResourceMonitor_HiddenEnvoyDeprecatedConfig) isResourceMonitor_ConfigType() {}

func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *ResourceMonitor) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ResourceMonitor_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *ResourceMonitor) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResourceMonitor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResourceMonitor_HiddenEnvoyDeprecatedConfig)(nil),
		(*ResourceMonitor_TypedConfig)(nil),
	}
}

type ThresholdTrigger struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThresholdTrigger) Reset()         { *m = ThresholdTrigger{} }
func (m *ThresholdTrigger) String() string { return proto.CompactTextString(m) }
func (*ThresholdTrigger) ProtoMessage()    {}
func (*ThresholdTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{1}
}

func (m *ThresholdTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThresholdTrigger.Unmarshal(m, b)
}
func (m *ThresholdTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThresholdTrigger.Marshal(b, m, deterministic)
}
func (m *ThresholdTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThresholdTrigger.Merge(m, src)
}
func (m *ThresholdTrigger) XXX_Size() int {
	return xxx_messageInfo_ThresholdTrigger.Size(m)
}
func (m *ThresholdTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_ThresholdTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_ThresholdTrigger proto.InternalMessageInfo

func (m *ThresholdTrigger) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Trigger struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TriggerOneof:
	//	*Trigger_Threshold
	TriggerOneof         isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{2}
}

func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (m *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(m, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,proto3,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (m *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := m.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Trigger) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Trigger_Threshold)(nil),
	}
}

type OverloadAction struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Triggers             []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OverloadAction) Reset()         { *m = OverloadAction{} }
func (m *OverloadAction) String() string { return proto.CompactTextString(m) }
func (*OverloadAction) ProtoMessage()    {}
func (*OverloadAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{3}
}

func (m *OverloadAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadAction.Unmarshal(m, b)
}
func (m *OverloadAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadAction.Marshal(b, m, deterministic)
}
func (m *OverloadAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadAction.Merge(m, src)
}
func (m *OverloadAction) XXX_Size() int {
	return xxx_messageInfo_OverloadAction.Size(m)
}
func (m *OverloadAction) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadAction.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadAction proto.InternalMessageInfo

func (m *OverloadAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OverloadAction) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type OverloadManager struct {
	RefreshInterval      *duration.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	ResourceMonitors     []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors,proto3" json:"resource_monitors,omitempty"`
	Actions              []*OverloadAction  `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverloadManager) Reset()         { *m = OverloadManager{} }
func (m *OverloadManager) String() string { return proto.CompactTextString(m) }
func (*OverloadManager) ProtoMessage()    {}
func (*OverloadManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{4}
}

func (m *OverloadManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadManager.Unmarshal(m, b)
}
func (m *OverloadManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadManager.Marshal(b, m, deterministic)
}
func (m *OverloadManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadManager.Merge(m, src)
}
func (m *OverloadManager) XXX_Size() int {
	return xxx_messageInfo_OverloadManager.Size(m)
}
func (m *OverloadManager) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadManager.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadManager proto.InternalMessageInfo

func (m *OverloadManager) GetRefreshInterval() *duration.Duration {
	if m != nil {
		return m.RefreshInterval
	}
	return nil
}

func (m *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if m != nil {
		return m.ResourceMonitors
	}
	return nil
}

func (m *OverloadManager) GetActions() []*OverloadAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceMonitor)(nil), "envoy.config.overload.v3.ResourceMonitor")
	proto.RegisterType((*ThresholdTrigger)(nil), "envoy.config.overload.v3.ThresholdTrigger")
	proto.RegisterType((*Trigger)(nil), "envoy.config.overload.v3.Trigger")
	proto.RegisterType((*OverloadAction)(nil), "envoy.config.overload.v3.OverloadAction")
	proto.RegisterType((*OverloadManager)(nil), "envoy.config.overload.v3.OverloadManager")
}

func init() {
	proto.RegisterFile("envoy/config/overload/v3/overload.proto", fileDescriptor_1b32999822dfa297)
}

var fileDescriptor_1b32999822dfa297 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x1c, 0xcd, 0x39, 0xfd, 0x7b, 0xa1, 0x34, 0x58, 0x95, 0xea, 0xb6, 0x50, 0xa5, 0x91, 0xa0, 0x29,
	0x6a, 0x6d, 0x94, 0xc0, 0x40, 0x16, 0xd4, 0xa3, 0x88, 0x82, 0x54, 0x51, 0x99, 0xee, 0xe6, 0x1a,
	0x5f, 0x1c, 0x4b, 0xee, 0x9d, 0x75, 0x3e, 0x5b, 0xcd, 0xc6, 0xc8, 0xdc, 0x91, 0xaf, 0xc0, 0x88,
	0xc4, 0xc0, 0x8e, 0xc4, 0xca, 0x37, 0x61, 0x44, 0x9d, 0x90, 0xef, 0xce, 0x89, 0xe2, 0x28, 0x01,
	0x4f, 0x27, 0xbf, 0xdf, 0x7b, 0xbf, 0xf7, 0x9e, 0xcf, 0x70, 0x9f, 0xd0, 0x8c, 0x0d, 0x9d, 0x1e,
	0xa3, 0xfd, 0x30, 0x70, 0x58, 0x46, 0x78, 0xc4, 0xb0, 0xef, 0x64, 0x9d, 0xd1, 0xd9, 0x8e, 0x39,
	0x13, 0xcc, 0xb4, 0xe4, 0xa0, 0xad, 0x06, 0xed, 0x11, 0x98, 0x75, 0xb6, 0xb7, 0x02, 0xc6, 0x82,
	0x88, 0x38, 0x72, 0xee, 0x32, 0xed, 0x3b, 0x98, 0x0e, 0x15, 0x69, 0x7b, 0xb7, 0x0c, 0xf9, 0x29,
	0xc7, 0x22, 0x64, 0x54, 0xe3, 0xf7, 0xcb, 0x78, 0x22, 0x78, 0xda, 0x13, 0x1a, 0x7d, 0x90, 0xfa,
	0x31, 0x76, 0x30, 0xa5, 0x4c, 0x48, 0x52, 0xe2, 0x24, 0x02, 0x8b, 0x34, 0xd1, 0xf0, 0xde, 0x14,
	0x9c, 0x11, 0x9e, 0x84, 0x8c, 0x86, 0x34, 0xd0, 0x23, 0x9b, 0x19, 0x8e, 0x42, 0x1f, 0x0b, 0xe2,
	0x14, 0x07, 0x05, 0x34, 0x6f, 0x0c, 0xb8, 0xee, 0x92, 0x84, 0xa5, 0xbc, 0x47, 0xce, 0x18, 0x0d,
	0x05, 0xe3, 0xe6, 0x0e, 0x5c, 0xa0, 0xf8, 0x8a, 0x58, 0xa0, 0x01, 0x5a, 0xab, 0x68, 0xf9, 0x16,
	0x2d, 0x70, 0xa3, 0x01, 0x5c, 0xf9, 0xd2, 0xbc, 0x84, 0xbb, 0x83, 0xd0, 0xf7, 0x09, 0xf5, 0x64,
	0x0f, 0x9e, 0x4f, 0x62, 0x4e, 0x7a, 0x58, 0x10, 0xdf, 0x53, 0x95, 0x58, 0x46, 0x03, 0xb4, 0x6a,
	0xed, 0x4d, 0x5b, 0x45, 0xb2, 0x8b, 0x48, 0xf6, 0x7b, 0x19, 0x09, 0x19, 0x16, 0x38, 0xad, 0xb8,
	0x3b, 0x4a, 0xe4, 0x55, 0xae, 0x71, 0x32, 0x92, 0x78, 0x29, 0x15, 0xcc, 0xe7, 0xf0, 0x8e, 0x18,
	0xc6, 0x63, 0xc5, 0xaa, 0x54, 0xdc, 0x98, 0x52, 0x3c, 0xa6, 0xc3, 0xd3, 0x8a, 0x5b, 0x93, 0xb3,
	0x8a, 0xda, 0x7d, 0xfa, 0xf9, 0xc7, 0xa7, 0x5d, 0x07, 0x1e, 0xcd, 0xf8, 0x48, 0x6d, 0x1c, 0xc5,
	0x03, 0x6c, 0x97, 0x12, 0xa3, 0x35, 0x58, 0x53, 0xa3, 0x5e, 0xae, 0xd5, 0xbc, 0x86, 0xf5, 0x8b,
	0x01, 0x27, 0xc9, 0x80, 0x45, 0xfe, 0x05, 0x0f, 0x83, 0x80, 0x70, 0xf3, 0x08, 0x2e, 0x66, 0x38,
	0x4a, 0x55, 0x2b, 0x00, 0x6d, 0xde, 0xa2, 0x0d, 0xd3, 0xdc, 0xaa, 0xc8, 0xe7, 0xf7, 0x8b, 0x83,
	0x8a, 0x7e, 0x5c, 0x35, 0xd5, 0x7d, 0x96, 0xfb, 0x78, 0x02, 0xed, 0xf9, 0x3e, 0xca, 0x5b, 0x9a,
	0xdf, 0x00, 0x5c, 0x2e, 0x36, 0xce, 0xfd, 0x0c, 0x6f, 0xe1, 0xaa, 0x28, 0xc8, 0xba, 0xf1, 0xc7,
	0xb3, 0x96, 0x75, 0xa6, 0xf6, 0x9c, 0x56, 0xdc, 0x31, 0xbd, 0x7b, 0x98, 0x7b, 0xdd, 0x87, 0x0f,
	0xff, 0xe1, 0x55, 0x51, 0xd1, 0x06, 0x5c, 0x13, 0xea, 0xe8, 0x31, 0x4a, 0x58, 0xdf, 0xac, 0xfe,
	0x41, 0xa0, 0xf9, 0x05, 0xc0, 0xbb, 0xef, 0x34, 0xe5, 0xb8, 0x97, 0xdf, 0xc2, 0xf9, 0xfe, 0x5f,
	0xc3, 0x15, 0xad, 0x92, 0x58, 0x46, 0xa3, 0xda, 0xaa, 0xb5, 0xf7, 0xe6, 0xd8, 0xd7, 0xab, 0x57,
	0x6e, 0xd1, 0xe2, 0x0d, 0x30, 0x56, 0x80, 0x3b, 0x22, 0x77, 0x3b, 0xb9, 0x79, 0x1b, 0x1e, 0xce,
	0x37, 0x3f, 0x69, 0xad, 0xf9, 0xd5, 0x80, 0xeb, 0xc5, 0xab, 0x33, 0x4c, 0x71, 0x5e, 0xf7, 0x09,
	0xac, 0x73, 0xd2, 0xcf, 0x3b, 0xf1, 0x42, 0x2a, 0x08, 0xcf, 0x70, 0x24, 0xad, 0xd7, 0xda, 0x5b,
	0x53, 0x17, 0xef, 0x44, 0xff, 0xbd, 0xee, 0xba, 0xa6, 0xbc, 0xd1, 0x0c, 0xf3, 0x03, 0xbc, 0xc7,
	0xf5, 0xe5, 0xf2, 0xae, 0xd4, 0xed, 0x2a, 0x02, 0x1e, 0xcc, 0x0e, 0x58, 0xbe, 0x8f, 0xe3, 0xa0,
	0x75, 0x3e, 0x09, 0x25, 0x26, 0x82, 0xcb, 0x58, 0xa6, 0x48, 0xac, 0xaa, 0xd4, 0x6d, 0xcd, 0xd6,
	0x9d, 0x8c, 0xed, 0x16, 0xc4, 0xff, 0xfc, 0x4b, 0x4a, 0x0d, 0xa1, 0xe3, 0xef, 0x1f, 0x7f, 0xfe,
	0x5a, 0x32, 0xea, 0x06, 0x7c, 0x14, 0x32, 0xb5, 0x34, 0xe6, 0xec, 0x7a, 0x38, 0x73, 0x3f, 0x5a,
	0x2b, 0x24, 0xce, 0xf3, 0xe6, 0xce, 0xc1, 0xe5, 0x92, 0xac, 0xb0, 0xf3, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x43, 0x57, 0x6d, 0xa2, 0x6e, 0x05, 0x00, 0x00,
}