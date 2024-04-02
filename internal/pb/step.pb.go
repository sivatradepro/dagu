// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: internal/proto/step.proto

package pb

import (
	any1 "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description    string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Variables      []string        `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty"`
	Dir            string          `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	ExecutorConfig *ExecutorConfig `protobuf:"bytes,5,opt,name=executor_config,json=executorConfig,proto3" json:"executor_config,omitempty"`
	CmdWithArgs    string          `protobuf:"bytes,6,opt,name=cmd_with_args,json=cmdWithArgs,proto3" json:"cmd_with_args,omitempty"`
	Command        string          `protobuf:"bytes,7,opt,name=command,proto3" json:"command,omitempty"`
	Script         string          `protobuf:"bytes,8,opt,name=script,proto3" json:"script,omitempty"`
	Stdout         string          `protobuf:"bytes,9,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr         string          `protobuf:"bytes,10,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Output         string          `protobuf:"bytes,11,opt,name=output,proto3" json:"output,omitempty"`
	Args           []string        `protobuf:"bytes,12,rep,name=args,proto3" json:"args,omitempty"`
	Depends        []string        `protobuf:"bytes,13,rep,name=depends,proto3" json:"depends,omitempty"`
	ContinueOn     *ContinueOn     `protobuf:"bytes,14,opt,name=continue_on,json=continueOn,proto3" json:"continue_on,omitempty"`
	RetryPolicy    *RetryPolicy    `protobuf:"bytes,15,opt,name=retry_policy,json=retryPolicy,proto3" json:"retry_policy,omitempty"`
	RepeatPolicy   *RepeatPolicy   `protobuf:"bytes,16,opt,name=repeat_policy,json=repeatPolicy,proto3" json:"repeat_policy,omitempty"`
	MailOnError    bool            `protobuf:"varint,17,opt,name=mail_on_error,json=mailOnError,proto3" json:"mail_on_error,omitempty"`
	Preconditions  []*Condition    `protobuf:"bytes,18,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
	SignalOnStop   string          `protobuf:"bytes,19,opt,name=signal_on_stop,json=signalOnStop,proto3" json:"signal_on_stop,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_step_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_step_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_internal_proto_step_proto_rawDescGZIP(), []int{0}
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Step) GetVariables() []string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *Step) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *Step) GetExecutorConfig() *ExecutorConfig {
	if x != nil {
		return x.ExecutorConfig
	}
	return nil
}

func (x *Step) GetCmdWithArgs() string {
	if x != nil {
		return x.CmdWithArgs
	}
	return ""
}

func (x *Step) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Step) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

func (x *Step) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *Step) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *Step) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Step) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Step) GetDepends() []string {
	if x != nil {
		return x.Depends
	}
	return nil
}

func (x *Step) GetContinueOn() *ContinueOn {
	if x != nil {
		return x.ContinueOn
	}
	return nil
}

func (x *Step) GetRetryPolicy() *RetryPolicy {
	if x != nil {
		return x.RetryPolicy
	}
	return nil
}

func (x *Step) GetRepeatPolicy() *RepeatPolicy {
	if x != nil {
		return x.RepeatPolicy
	}
	return nil
}

func (x *Step) GetMailOnError() bool {
	if x != nil {
		return x.MailOnError
	}
	return false
}

func (x *Step) GetPreconditions() []*Condition {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

func (x *Step) GetSignalOnStop() string {
	if x != nil {
		return x.SignalOnStop
	}
	return ""
}

type ExecutorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Config map[string]*any1.Any `protobuf:"bytes,2,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExecutorConfig) Reset() {
	*x = ExecutorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_step_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorConfig) ProtoMessage() {}

func (x *ExecutorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_step_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorConfig.ProtoReflect.Descriptor instead.
func (*ExecutorConfig) Descriptor() ([]byte, []int) {
	return file_internal_proto_step_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutorConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExecutorConfig) GetConfig() map[string]*any1.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

type ContinueOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure bool `protobuf:"varint,1,opt,name=failure,proto3" json:"failure,omitempty"`
	Skipped bool `protobuf:"varint,2,opt,name=skipped,proto3" json:"skipped,omitempty"`
}

func (x *ContinueOn) Reset() {
	*x = ContinueOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_step_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinueOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinueOn) ProtoMessage() {}

func (x *ContinueOn) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_step_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinueOn.ProtoReflect.Descriptor instead.
func (*ContinueOn) Descriptor() ([]byte, []int) {
	return file_internal_proto_step_proto_rawDescGZIP(), []int{2}
}

func (x *ContinueOn) GetFailure() bool {
	if x != nil {
		return x.Failure
	}
	return false
}

func (x *ContinueOn) GetSkipped() bool {
	if x != nil {
		return x.Skipped
	}
	return false
}

type RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit    int32              `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *RetryPolicy) Reset() {
	*x = RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_step_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy) ProtoMessage() {}

func (x *RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_step_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return file_internal_proto_step_proto_rawDescGZIP(), []int{3}
}

func (x *RetryPolicy) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RetryPolicy) GetInterval() *duration.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

type RepeatPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repeat   bool               `protobuf:"varint,1,opt,name=repeat,proto3" json:"repeat,omitempty"`
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *RepeatPolicy) Reset() {
	*x = RepeatPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_step_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatPolicy) ProtoMessage() {}

func (x *RepeatPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_step_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatPolicy.ProtoReflect.Descriptor instead.
func (*RepeatPolicy) Descriptor() ([]byte, []int) {
	return file_internal_proto_step_proto_rawDescGZIP(), []int{4}
}

func (x *RepeatPolicy) GetRepeat() bool {
	if x != nil {
		return x.Repeat
	}
	return false
}

func (x *RepeatPolicy) GetInterval() *duration.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition string `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Expected  string `protobuf:"bytes,2,opt,name=expected,proto3" json:"expected,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_step_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_step_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_internal_proto_step_proto_rawDescGZIP(), []int{5}
}

func (x *Condition) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Condition) GetExpected() string {
	if x != nil {
		return x.Expected
	}
	return ""
}

var File_internal_proto_step_proto protoreflect.FileDescriptor

var file_internal_proto_step_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4,
	0x05, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x3f, 0x0a,
	0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22,
	0x0a, 0x0d, 0x63, 0x6d, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6d, 0x64, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x5f, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x65, 0x4f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x4f, 0x6e, 0x12,
	0x36, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x39, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x6e, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6c, 0x4f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x24, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4f,
	0x6e, 0x53, 0x74, 0x6f, 0x70, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x4f, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x40, 0x0a, 0x0a, 0x43, 0x6f, 0x6e,
	0x74, 0x69, 0x6e, 0x75, 0x65, 0x4f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x22, 0x5a, 0x0a, 0x0b, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x5d, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x12,
	0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x45, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_step_proto_rawDescOnce sync.Once
	file_internal_proto_step_proto_rawDescData = file_internal_proto_step_proto_rawDesc
)

func file_internal_proto_step_proto_rawDescGZIP() []byte {
	file_internal_proto_step_proto_rawDescOnce.Do(func() {
		file_internal_proto_step_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_step_proto_rawDescData)
	})
	return file_internal_proto_step_proto_rawDescData
}

var file_internal_proto_step_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_proto_step_proto_goTypes = []interface{}{
	(*Step)(nil),              // 0: models.Step
	(*ExecutorConfig)(nil),    // 1: models.ExecutorConfig
	(*ContinueOn)(nil),        // 2: models.ContinueOn
	(*RetryPolicy)(nil),       // 3: models.RetryPolicy
	(*RepeatPolicy)(nil),      // 4: models.RepeatPolicy
	(*Condition)(nil),         // 5: models.Condition
	nil,                       // 6: models.ExecutorConfig.ConfigEntry
	(*duration.Duration)(nil), // 7: google.protobuf.Duration
	(*any1.Any)(nil),          // 8: google.protobuf.Any
}
var file_internal_proto_step_proto_depIdxs = []int32{
	1, // 0: models.Step.executor_config:type_name -> models.ExecutorConfig
	2, // 1: models.Step.continue_on:type_name -> models.ContinueOn
	3, // 2: models.Step.retry_policy:type_name -> models.RetryPolicy
	4, // 3: models.Step.repeat_policy:type_name -> models.RepeatPolicy
	5, // 4: models.Step.preconditions:type_name -> models.Condition
	6, // 5: models.ExecutorConfig.config:type_name -> models.ExecutorConfig.ConfigEntry
	7, // 6: models.RetryPolicy.interval:type_name -> google.protobuf.Duration
	7, // 7: models.RepeatPolicy.interval:type_name -> google.protobuf.Duration
	8, // 8: models.ExecutorConfig.ConfigEntry.value:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_internal_proto_step_proto_init() }
func file_internal_proto_step_proto_init() {
	if File_internal_proto_step_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_step_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_step_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_step_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinueOn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_step_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_step_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_step_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_step_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_step_proto_goTypes,
		DependencyIndexes: file_internal_proto_step_proto_depIdxs,
		MessageInfos:      file_internal_proto_step_proto_msgTypes,
	}.Build()
	File_internal_proto_step_proto = out.File
	file_internal_proto_step_proto_rawDesc = nil
	file_internal_proto_step_proto_goTypes = nil
	file_internal_proto_step_proto_depIdxs = nil
}
