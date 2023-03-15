// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto.proto

#ifndef PROTOBUF_INCLUDED_proto_2eproto
#define PROTOBUF_INCLUDED_proto_2eproto

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3006000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3006000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#define PROTOBUF_INTERNAL_EXPORT_protobuf_proto_2eproto 

namespace protobuf_proto_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[4];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
}  // namespace protobuf_proto_2eproto
namespace cproto {
class LoginReq;
class LoginReqDefaultTypeInternal;
extern LoginReqDefaultTypeInternal _LoginReq_default_instance_;
class LoginRes;
class LoginResDefaultTypeInternal;
extern LoginResDefaultTypeInternal _LoginRes_default_instance_;
class NewUserReq;
class NewUserReqDefaultTypeInternal;
extern NewUserReqDefaultTypeInternal _NewUserReq_default_instance_;
class NewUserRes;
class NewUserResDefaultTypeInternal;
extern NewUserResDefaultTypeInternal _NewUserRes_default_instance_;
}  // namespace cproto
namespace google {
namespace protobuf {
template<> ::cproto::LoginReq* Arena::CreateMaybeMessage<::cproto::LoginReq>(Arena*);
template<> ::cproto::LoginRes* Arena::CreateMaybeMessage<::cproto::LoginRes>(Arena*);
template<> ::cproto::NewUserReq* Arena::CreateMaybeMessage<::cproto::NewUserReq>(Arena*);
template<> ::cproto::NewUserRes* Arena::CreateMaybeMessage<::cproto::NewUserRes>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace cproto {

// ===================================================================

class NewUserReq : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cproto.NewUserReq) */ {
 public:
  NewUserReq();
  virtual ~NewUserReq();

  NewUserReq(const NewUserReq& from);

  inline NewUserReq& operator=(const NewUserReq& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  NewUserReq(NewUserReq&& from) noexcept
    : NewUserReq() {
    *this = ::std::move(from);
  }

  inline NewUserReq& operator=(NewUserReq&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const NewUserReq& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const NewUserReq* internal_default_instance() {
    return reinterpret_cast<const NewUserReq*>(
               &_NewUserReq_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(NewUserReq* other);
  friend void swap(NewUserReq& a, NewUserReq& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline NewUserReq* New() const final {
    return CreateMaybeMessage<NewUserReq>(NULL);
  }

  NewUserReq* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<NewUserReq>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const NewUserReq& from);
  void MergeFrom(const NewUserReq& from);
  void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(NewUserReq* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string Name = 1;
  void clear_name();
  static const int kNameFieldNumber = 1;
  const ::std::string& name() const;
  void set_name(const ::std::string& value);
  #if LANG_CXX11
  void set_name(::std::string&& value);
  #endif
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  ::std::string* mutable_name();
  ::std::string* release_name();
  void set_allocated_name(::std::string* name);

  // string Password = 2;
  void clear_password();
  static const int kPasswordFieldNumber = 2;
  const ::std::string& password() const;
  void set_password(const ::std::string& value);
  #if LANG_CXX11
  void set_password(::std::string&& value);
  #endif
  void set_password(const char* value);
  void set_password(const char* value, size_t size);
  ::std::string* mutable_password();
  ::std::string* release_password();
  void set_allocated_password(::std::string* password);

  // @@protoc_insertion_point(class_scope:cproto.NewUserReq)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr name_;
  ::google::protobuf::internal::ArenaStringPtr password_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_proto_2eproto::TableStruct;
};
// -------------------------------------------------------------------

class NewUserRes : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cproto.NewUserRes) */ {
 public:
  NewUserRes();
  virtual ~NewUserRes();

  NewUserRes(const NewUserRes& from);

  inline NewUserRes& operator=(const NewUserRes& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  NewUserRes(NewUserRes&& from) noexcept
    : NewUserRes() {
    *this = ::std::move(from);
  }

  inline NewUserRes& operator=(NewUserRes&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const NewUserRes& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const NewUserRes* internal_default_instance() {
    return reinterpret_cast<const NewUserRes*>(
               &_NewUserRes_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(NewUserRes* other);
  friend void swap(NewUserRes& a, NewUserRes& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline NewUserRes* New() const final {
    return CreateMaybeMessage<NewUserRes>(NULL);
  }

  NewUserRes* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<NewUserRes>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const NewUserRes& from);
  void MergeFrom(const NewUserRes& from);
  void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(NewUserRes* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string Msg = 2;
  void clear_msg();
  static const int kMsgFieldNumber = 2;
  const ::std::string& msg() const;
  void set_msg(const ::std::string& value);
  #if LANG_CXX11
  void set_msg(::std::string&& value);
  #endif
  void set_msg(const char* value);
  void set_msg(const char* value, size_t size);
  ::std::string* mutable_msg();
  ::std::string* release_msg();
  void set_allocated_msg(::std::string* msg);

  // int32 Ret = 1;
  void clear_ret();
  static const int kRetFieldNumber = 1;
  ::google::protobuf::int32 ret() const;
  void set_ret(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:cproto.NewUserRes)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr msg_;
  ::google::protobuf::int32 ret_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_proto_2eproto::TableStruct;
};
// -------------------------------------------------------------------

class LoginReq : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cproto.LoginReq) */ {
 public:
  LoginReq();
  virtual ~LoginReq();

  LoginReq(const LoginReq& from);

  inline LoginReq& operator=(const LoginReq& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  LoginReq(LoginReq&& from) noexcept
    : LoginReq() {
    *this = ::std::move(from);
  }

  inline LoginReq& operator=(LoginReq&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const LoginReq& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const LoginReq* internal_default_instance() {
    return reinterpret_cast<const LoginReq*>(
               &_LoginReq_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    2;

  void Swap(LoginReq* other);
  friend void swap(LoginReq& a, LoginReq& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline LoginReq* New() const final {
    return CreateMaybeMessage<LoginReq>(NULL);
  }

  LoginReq* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<LoginReq>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const LoginReq& from);
  void MergeFrom(const LoginReq& from);
  void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(LoginReq* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string Name = 1;
  void clear_name();
  static const int kNameFieldNumber = 1;
  const ::std::string& name() const;
  void set_name(const ::std::string& value);
  #if LANG_CXX11
  void set_name(::std::string&& value);
  #endif
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  ::std::string* mutable_name();
  ::std::string* release_name();
  void set_allocated_name(::std::string* name);

  // string Password = 2;
  void clear_password();
  static const int kPasswordFieldNumber = 2;
  const ::std::string& password() const;
  void set_password(const ::std::string& value);
  #if LANG_CXX11
  void set_password(::std::string&& value);
  #endif
  void set_password(const char* value);
  void set_password(const char* value, size_t size);
  ::std::string* mutable_password();
  ::std::string* release_password();
  void set_allocated_password(::std::string* password);

  // @@protoc_insertion_point(class_scope:cproto.LoginReq)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr name_;
  ::google::protobuf::internal::ArenaStringPtr password_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_proto_2eproto::TableStruct;
};
// -------------------------------------------------------------------

class LoginRes : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:cproto.LoginRes) */ {
 public:
  LoginRes();
  virtual ~LoginRes();

  LoginRes(const LoginRes& from);

  inline LoginRes& operator=(const LoginRes& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  LoginRes(LoginRes&& from) noexcept
    : LoginRes() {
    *this = ::std::move(from);
  }

  inline LoginRes& operator=(LoginRes&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const LoginRes& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const LoginRes* internal_default_instance() {
    return reinterpret_cast<const LoginRes*>(
               &_LoginRes_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    3;

  void Swap(LoginRes* other);
  friend void swap(LoginRes& a, LoginRes& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline LoginRes* New() const final {
    return CreateMaybeMessage<LoginRes>(NULL);
  }

  LoginRes* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<LoginRes>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const LoginRes& from);
  void MergeFrom(const LoginRes& from);
  void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(LoginRes* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string Msg = 2;
  void clear_msg();
  static const int kMsgFieldNumber = 2;
  const ::std::string& msg() const;
  void set_msg(const ::std::string& value);
  #if LANG_CXX11
  void set_msg(::std::string&& value);
  #endif
  void set_msg(const char* value);
  void set_msg(const char* value, size_t size);
  ::std::string* mutable_msg();
  ::std::string* release_msg();
  void set_allocated_msg(::std::string* msg);

  // string Token = 3;
  void clear_token();
  static const int kTokenFieldNumber = 3;
  const ::std::string& token() const;
  void set_token(const ::std::string& value);
  #if LANG_CXX11
  void set_token(::std::string&& value);
  #endif
  void set_token(const char* value);
  void set_token(const char* value, size_t size);
  ::std::string* mutable_token();
  ::std::string* release_token();
  void set_allocated_token(::std::string* token);

  // int32 Ret = 1;
  void clear_ret();
  static const int kRetFieldNumber = 1;
  ::google::protobuf::int32 ret() const;
  void set_ret(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:cproto.LoginRes)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr msg_;
  ::google::protobuf::internal::ArenaStringPtr token_;
  ::google::protobuf::int32 ret_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_proto_2eproto::TableStruct;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// NewUserReq

// string Name = 1;
inline void NewUserReq::clear_name() {
  name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& NewUserReq::name() const {
  // @@protoc_insertion_point(field_get:cproto.NewUserReq.Name)
  return name_.GetNoArena();
}
inline void NewUserReq::set_name(const ::std::string& value) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.NewUserReq.Name)
}
#if LANG_CXX11
inline void NewUserReq::set_name(::std::string&& value) {
  
  name_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.NewUserReq.Name)
}
#endif
inline void NewUserReq::set_name(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.NewUserReq.Name)
}
inline void NewUserReq::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.NewUserReq.Name)
}
inline ::std::string* NewUserReq::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:cproto.NewUserReq.Name)
  return name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* NewUserReq::release_name() {
  // @@protoc_insertion_point(field_release:cproto.NewUserReq.Name)
  
  return name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void NewUserReq::set_allocated_name(::std::string* name) {
  if (name != NULL) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:cproto.NewUserReq.Name)
}

// string Password = 2;
inline void NewUserReq::clear_password() {
  password_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& NewUserReq::password() const {
  // @@protoc_insertion_point(field_get:cproto.NewUserReq.Password)
  return password_.GetNoArena();
}
inline void NewUserReq::set_password(const ::std::string& value) {
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.NewUserReq.Password)
}
#if LANG_CXX11
inline void NewUserReq::set_password(::std::string&& value) {
  
  password_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.NewUserReq.Password)
}
#endif
inline void NewUserReq::set_password(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.NewUserReq.Password)
}
inline void NewUserReq::set_password(const char* value, size_t size) {
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.NewUserReq.Password)
}
inline ::std::string* NewUserReq::mutable_password() {
  
  // @@protoc_insertion_point(field_mutable:cproto.NewUserReq.Password)
  return password_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* NewUserReq::release_password() {
  // @@protoc_insertion_point(field_release:cproto.NewUserReq.Password)
  
  return password_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void NewUserReq::set_allocated_password(::std::string* password) {
  if (password != NULL) {
    
  } else {
    
  }
  password_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), password);
  // @@protoc_insertion_point(field_set_allocated:cproto.NewUserReq.Password)
}

// -------------------------------------------------------------------

// NewUserRes

// int32 Ret = 1;
inline void NewUserRes::clear_ret() {
  ret_ = 0;
}
inline ::google::protobuf::int32 NewUserRes::ret() const {
  // @@protoc_insertion_point(field_get:cproto.NewUserRes.Ret)
  return ret_;
}
inline void NewUserRes::set_ret(::google::protobuf::int32 value) {
  
  ret_ = value;
  // @@protoc_insertion_point(field_set:cproto.NewUserRes.Ret)
}

// string Msg = 2;
inline void NewUserRes::clear_msg() {
  msg_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& NewUserRes::msg() const {
  // @@protoc_insertion_point(field_get:cproto.NewUserRes.Msg)
  return msg_.GetNoArena();
}
inline void NewUserRes::set_msg(const ::std::string& value) {
  
  msg_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.NewUserRes.Msg)
}
#if LANG_CXX11
inline void NewUserRes::set_msg(::std::string&& value) {
  
  msg_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.NewUserRes.Msg)
}
#endif
inline void NewUserRes::set_msg(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  msg_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.NewUserRes.Msg)
}
inline void NewUserRes::set_msg(const char* value, size_t size) {
  
  msg_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.NewUserRes.Msg)
}
inline ::std::string* NewUserRes::mutable_msg() {
  
  // @@protoc_insertion_point(field_mutable:cproto.NewUserRes.Msg)
  return msg_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* NewUserRes::release_msg() {
  // @@protoc_insertion_point(field_release:cproto.NewUserRes.Msg)
  
  return msg_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void NewUserRes::set_allocated_msg(::std::string* msg) {
  if (msg != NULL) {
    
  } else {
    
  }
  msg_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), msg);
  // @@protoc_insertion_point(field_set_allocated:cproto.NewUserRes.Msg)
}

// -------------------------------------------------------------------

// LoginReq

// string Name = 1;
inline void LoginReq::clear_name() {
  name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& LoginReq::name() const {
  // @@protoc_insertion_point(field_get:cproto.LoginReq.Name)
  return name_.GetNoArena();
}
inline void LoginReq::set_name(const ::std::string& value) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.LoginReq.Name)
}
#if LANG_CXX11
inline void LoginReq::set_name(::std::string&& value) {
  
  name_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.LoginReq.Name)
}
#endif
inline void LoginReq::set_name(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.LoginReq.Name)
}
inline void LoginReq::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.LoginReq.Name)
}
inline ::std::string* LoginReq::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:cproto.LoginReq.Name)
  return name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LoginReq::release_name() {
  // @@protoc_insertion_point(field_release:cproto.LoginReq.Name)
  
  return name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LoginReq::set_allocated_name(::std::string* name) {
  if (name != NULL) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:cproto.LoginReq.Name)
}

// string Password = 2;
inline void LoginReq::clear_password() {
  password_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& LoginReq::password() const {
  // @@protoc_insertion_point(field_get:cproto.LoginReq.Password)
  return password_.GetNoArena();
}
inline void LoginReq::set_password(const ::std::string& value) {
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.LoginReq.Password)
}
#if LANG_CXX11
inline void LoginReq::set_password(::std::string&& value) {
  
  password_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.LoginReq.Password)
}
#endif
inline void LoginReq::set_password(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.LoginReq.Password)
}
inline void LoginReq::set_password(const char* value, size_t size) {
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.LoginReq.Password)
}
inline ::std::string* LoginReq::mutable_password() {
  
  // @@protoc_insertion_point(field_mutable:cproto.LoginReq.Password)
  return password_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LoginReq::release_password() {
  // @@protoc_insertion_point(field_release:cproto.LoginReq.Password)
  
  return password_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LoginReq::set_allocated_password(::std::string* password) {
  if (password != NULL) {
    
  } else {
    
  }
  password_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), password);
  // @@protoc_insertion_point(field_set_allocated:cproto.LoginReq.Password)
}

// -------------------------------------------------------------------

// LoginRes

// int32 Ret = 1;
inline void LoginRes::clear_ret() {
  ret_ = 0;
}
inline ::google::protobuf::int32 LoginRes::ret() const {
  // @@protoc_insertion_point(field_get:cproto.LoginRes.Ret)
  return ret_;
}
inline void LoginRes::set_ret(::google::protobuf::int32 value) {
  
  ret_ = value;
  // @@protoc_insertion_point(field_set:cproto.LoginRes.Ret)
}

// string Msg = 2;
inline void LoginRes::clear_msg() {
  msg_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& LoginRes::msg() const {
  // @@protoc_insertion_point(field_get:cproto.LoginRes.Msg)
  return msg_.GetNoArena();
}
inline void LoginRes::set_msg(const ::std::string& value) {
  
  msg_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.LoginRes.Msg)
}
#if LANG_CXX11
inline void LoginRes::set_msg(::std::string&& value) {
  
  msg_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.LoginRes.Msg)
}
#endif
inline void LoginRes::set_msg(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  msg_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.LoginRes.Msg)
}
inline void LoginRes::set_msg(const char* value, size_t size) {
  
  msg_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.LoginRes.Msg)
}
inline ::std::string* LoginRes::mutable_msg() {
  
  // @@protoc_insertion_point(field_mutable:cproto.LoginRes.Msg)
  return msg_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LoginRes::release_msg() {
  // @@protoc_insertion_point(field_release:cproto.LoginRes.Msg)
  
  return msg_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LoginRes::set_allocated_msg(::std::string* msg) {
  if (msg != NULL) {
    
  } else {
    
  }
  msg_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), msg);
  // @@protoc_insertion_point(field_set_allocated:cproto.LoginRes.Msg)
}

// string Token = 3;
inline void LoginRes::clear_token() {
  token_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& LoginRes::token() const {
  // @@protoc_insertion_point(field_get:cproto.LoginRes.Token)
  return token_.GetNoArena();
}
inline void LoginRes::set_token(const ::std::string& value) {
  
  token_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:cproto.LoginRes.Token)
}
#if LANG_CXX11
inline void LoginRes::set_token(::std::string&& value) {
  
  token_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:cproto.LoginRes.Token)
}
#endif
inline void LoginRes::set_token(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  token_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:cproto.LoginRes.Token)
}
inline void LoginRes::set_token(const char* value, size_t size) {
  
  token_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:cproto.LoginRes.Token)
}
inline ::std::string* LoginRes::mutable_token() {
  
  // @@protoc_insertion_point(field_mutable:cproto.LoginRes.Token)
  return token_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* LoginRes::release_token() {
  // @@protoc_insertion_point(field_release:cproto.LoginRes.Token)
  
  return token_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void LoginRes::set_allocated_token(::std::string* token) {
  if (token != NULL) {
    
  } else {
    
  }
  token_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), token);
  // @@protoc_insertion_point(field_set_allocated:cproto.LoginRes.Token)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------

// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace cproto

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_INCLUDED_proto_2eproto
