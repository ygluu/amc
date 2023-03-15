// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: login_svc.proto

#ifndef PROTOBUF_INCLUDED_login_5fsvc_2eproto
#define PROTOBUF_INCLUDED_login_5fsvc_2eproto

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
#define PROTOBUF_INTERNAL_EXPORT_protobuf_login_5fsvc_2eproto 

namespace protobuf_login_5fsvc_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[3];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
}  // namespace protobuf_login_5fsvc_2eproto
namespace sproto {
class GetSessionReq;
class GetSessionReqDefaultTypeInternal;
extern GetSessionReqDefaultTypeInternal _GetSessionReq_default_instance_;
class OnLogin;
class OnLoginDefaultTypeInternal;
extern OnLoginDefaultTypeInternal _OnLogin_default_instance_;
class UserLoginReq;
class UserLoginReqDefaultTypeInternal;
extern UserLoginReqDefaultTypeInternal _UserLoginReq_default_instance_;
}  // namespace sproto
namespace google {
namespace protobuf {
template<> ::sproto::GetSessionReq* Arena::CreateMaybeMessage<::sproto::GetSessionReq>(Arena*);
template<> ::sproto::OnLogin* Arena::CreateMaybeMessage<::sproto::OnLogin>(Arena*);
template<> ::sproto::UserLoginReq* Arena::CreateMaybeMessage<::sproto::UserLoginReq>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace sproto {

// ===================================================================

class UserLoginReq : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:sproto.UserLoginReq) */ {
 public:
  UserLoginReq();
  virtual ~UserLoginReq();

  UserLoginReq(const UserLoginReq& from);

  inline UserLoginReq& operator=(const UserLoginReq& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  UserLoginReq(UserLoginReq&& from) noexcept
    : UserLoginReq() {
    *this = ::std::move(from);
  }

  inline UserLoginReq& operator=(UserLoginReq&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const UserLoginReq& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const UserLoginReq* internal_default_instance() {
    return reinterpret_cast<const UserLoginReq*>(
               &_UserLoginReq_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(UserLoginReq* other);
  friend void swap(UserLoginReq& a, UserLoginReq& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline UserLoginReq* New() const final {
    return CreateMaybeMessage<UserLoginReq>(NULL);
  }

  UserLoginReq* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<UserLoginReq>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const UserLoginReq& from);
  void MergeFrom(const UserLoginReq& from);
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
  void InternalSwap(UserLoginReq* other);
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

  // string GateAddr = 1;
  void clear_gateaddr();
  static const int kGateAddrFieldNumber = 1;
  const ::std::string& gateaddr() const;
  void set_gateaddr(const ::std::string& value);
  #if LANG_CXX11
  void set_gateaddr(::std::string&& value);
  #endif
  void set_gateaddr(const char* value);
  void set_gateaddr(const char* value, size_t size);
  ::std::string* mutable_gateaddr();
  ::std::string* release_gateaddr();
  void set_allocated_gateaddr(::std::string* gateaddr);

  // string Name = 3;
  void clear_name();
  static const int kNameFieldNumber = 3;
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

  // string Password = 4;
  void clear_password();
  static const int kPasswordFieldNumber = 4;
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

  // uint64 GateClient = 2;
  void clear_gateclient();
  static const int kGateClientFieldNumber = 2;
  ::google::protobuf::uint64 gateclient() const;
  void set_gateclient(::google::protobuf::uint64 value);

  // @@protoc_insertion_point(class_scope:sproto.UserLoginReq)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr gateaddr_;
  ::google::protobuf::internal::ArenaStringPtr name_;
  ::google::protobuf::internal::ArenaStringPtr password_;
  ::google::protobuf::uint64 gateclient_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_login_5fsvc_2eproto::TableStruct;
};
// -------------------------------------------------------------------

class GetSessionReq : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:sproto.GetSessionReq) */ {
 public:
  GetSessionReq();
  virtual ~GetSessionReq();

  GetSessionReq(const GetSessionReq& from);

  inline GetSessionReq& operator=(const GetSessionReq& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  GetSessionReq(GetSessionReq&& from) noexcept
    : GetSessionReq() {
    *this = ::std::move(from);
  }

  inline GetSessionReq& operator=(GetSessionReq&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const GetSessionReq& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const GetSessionReq* internal_default_instance() {
    return reinterpret_cast<const GetSessionReq*>(
               &_GetSessionReq_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(GetSessionReq* other);
  friend void swap(GetSessionReq& a, GetSessionReq& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline GetSessionReq* New() const final {
    return CreateMaybeMessage<GetSessionReq>(NULL);
  }

  GetSessionReq* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<GetSessionReq>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const GetSessionReq& from);
  void MergeFrom(const GetSessionReq& from);
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
  void InternalSwap(GetSessionReq* other);
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

  // string UserName = 1;
  void clear_username();
  static const int kUserNameFieldNumber = 1;
  const ::std::string& username() const;
  void set_username(const ::std::string& value);
  #if LANG_CXX11
  void set_username(::std::string&& value);
  #endif
  void set_username(const char* value);
  void set_username(const char* value, size_t size);
  ::std::string* mutable_username();
  ::std::string* release_username();
  void set_allocated_username(::std::string* username);

  // @@protoc_insertion_point(class_scope:sproto.GetSessionReq)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr username_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_login_5fsvc_2eproto::TableStruct;
};
// -------------------------------------------------------------------

class OnLogin : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:sproto.OnLogin) */ {
 public:
  OnLogin();
  virtual ~OnLogin();

  OnLogin(const OnLogin& from);

  inline OnLogin& operator=(const OnLogin& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  OnLogin(OnLogin&& from) noexcept
    : OnLogin() {
    *this = ::std::move(from);
  }

  inline OnLogin& operator=(OnLogin&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const OnLogin& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const OnLogin* internal_default_instance() {
    return reinterpret_cast<const OnLogin*>(
               &_OnLogin_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    2;

  void Swap(OnLogin* other);
  friend void swap(OnLogin& a, OnLogin& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline OnLogin* New() const final {
    return CreateMaybeMessage<OnLogin>(NULL);
  }

  OnLogin* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<OnLogin>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const OnLogin& from);
  void MergeFrom(const OnLogin& from);
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
  void InternalSwap(OnLogin* other);
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

  // string Name = 2;
  void clear_name();
  static const int kNameFieldNumber = 2;
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

  // string OwnerAddr = 4;
  void clear_owneraddr();
  static const int kOwnerAddrFieldNumber = 4;
  const ::std::string& owneraddr() const;
  void set_owneraddr(const ::std::string& value);
  #if LANG_CXX11
  void set_owneraddr(::std::string&& value);
  #endif
  void set_owneraddr(const char* value);
  void set_owneraddr(const char* value, size_t size);
  ::std::string* mutable_owneraddr();
  ::std::string* release_owneraddr();
  void set_allocated_owneraddr(::std::string* owneraddr);

  // string Token = 5;
  void clear_token();
  static const int kTokenFieldNumber = 5;
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

  // string GateAddr = 6;
  void clear_gateaddr();
  static const int kGateAddrFieldNumber = 6;
  const ::std::string& gateaddr() const;
  void set_gateaddr(const ::std::string& value);
  #if LANG_CXX11
  void set_gateaddr(::std::string&& value);
  #endif
  void set_gateaddr(const char* value);
  void set_gateaddr(const char* value, size_t size);
  ::std::string* mutable_gateaddr();
  ::std::string* release_gateaddr();
  void set_allocated_gateaddr(::std::string* gateaddr);

  // bool Broadcast = 1;
  void clear_broadcast();
  static const int kBroadcastFieldNumber = 1;
  bool broadcast() const;
  void set_broadcast(bool value);

  // uint32 NameHash = 3;
  void clear_namehash();
  static const int kNameHashFieldNumber = 3;
  ::google::protobuf::uint32 namehash() const;
  void set_namehash(::google::protobuf::uint32 value);

  // uint64 GateClient = 7;
  void clear_gateclient();
  static const int kGateClientFieldNumber = 7;
  ::google::protobuf::uint64 gateclient() const;
  void set_gateclient(::google::protobuf::uint64 value);

  // @@protoc_insertion_point(class_scope:sproto.OnLogin)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr name_;
  ::google::protobuf::internal::ArenaStringPtr owneraddr_;
  ::google::protobuf::internal::ArenaStringPtr token_;
  ::google::protobuf::internal::ArenaStringPtr gateaddr_;
  bool broadcast_;
  ::google::protobuf::uint32 namehash_;
  ::google::protobuf::uint64 gateclient_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::protobuf_login_5fsvc_2eproto::TableStruct;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// UserLoginReq

// string GateAddr = 1;
inline void UserLoginReq::clear_gateaddr() {
  gateaddr_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& UserLoginReq::gateaddr() const {
  // @@protoc_insertion_point(field_get:sproto.UserLoginReq.GateAddr)
  return gateaddr_.GetNoArena();
}
inline void UserLoginReq::set_gateaddr(const ::std::string& value) {
  
  gateaddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.UserLoginReq.GateAddr)
}
#if LANG_CXX11
inline void UserLoginReq::set_gateaddr(::std::string&& value) {
  
  gateaddr_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.UserLoginReq.GateAddr)
}
#endif
inline void UserLoginReq::set_gateaddr(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  gateaddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.UserLoginReq.GateAddr)
}
inline void UserLoginReq::set_gateaddr(const char* value, size_t size) {
  
  gateaddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.UserLoginReq.GateAddr)
}
inline ::std::string* UserLoginReq::mutable_gateaddr() {
  
  // @@protoc_insertion_point(field_mutable:sproto.UserLoginReq.GateAddr)
  return gateaddr_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* UserLoginReq::release_gateaddr() {
  // @@protoc_insertion_point(field_release:sproto.UserLoginReq.GateAddr)
  
  return gateaddr_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UserLoginReq::set_allocated_gateaddr(::std::string* gateaddr) {
  if (gateaddr != NULL) {
    
  } else {
    
  }
  gateaddr_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), gateaddr);
  // @@protoc_insertion_point(field_set_allocated:sproto.UserLoginReq.GateAddr)
}

// uint64 GateClient = 2;
inline void UserLoginReq::clear_gateclient() {
  gateclient_ = GOOGLE_ULONGLONG(0);
}
inline ::google::protobuf::uint64 UserLoginReq::gateclient() const {
  // @@protoc_insertion_point(field_get:sproto.UserLoginReq.GateClient)
  return gateclient_;
}
inline void UserLoginReq::set_gateclient(::google::protobuf::uint64 value) {
  
  gateclient_ = value;
  // @@protoc_insertion_point(field_set:sproto.UserLoginReq.GateClient)
}

// string Name = 3;
inline void UserLoginReq::clear_name() {
  name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& UserLoginReq::name() const {
  // @@protoc_insertion_point(field_get:sproto.UserLoginReq.Name)
  return name_.GetNoArena();
}
inline void UserLoginReq::set_name(const ::std::string& value) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.UserLoginReq.Name)
}
#if LANG_CXX11
inline void UserLoginReq::set_name(::std::string&& value) {
  
  name_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.UserLoginReq.Name)
}
#endif
inline void UserLoginReq::set_name(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.UserLoginReq.Name)
}
inline void UserLoginReq::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.UserLoginReq.Name)
}
inline ::std::string* UserLoginReq::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:sproto.UserLoginReq.Name)
  return name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* UserLoginReq::release_name() {
  // @@protoc_insertion_point(field_release:sproto.UserLoginReq.Name)
  
  return name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UserLoginReq::set_allocated_name(::std::string* name) {
  if (name != NULL) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:sproto.UserLoginReq.Name)
}

// string Password = 4;
inline void UserLoginReq::clear_password() {
  password_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& UserLoginReq::password() const {
  // @@protoc_insertion_point(field_get:sproto.UserLoginReq.Password)
  return password_.GetNoArena();
}
inline void UserLoginReq::set_password(const ::std::string& value) {
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.UserLoginReq.Password)
}
#if LANG_CXX11
inline void UserLoginReq::set_password(::std::string&& value) {
  
  password_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.UserLoginReq.Password)
}
#endif
inline void UserLoginReq::set_password(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.UserLoginReq.Password)
}
inline void UserLoginReq::set_password(const char* value, size_t size) {
  
  password_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.UserLoginReq.Password)
}
inline ::std::string* UserLoginReq::mutable_password() {
  
  // @@protoc_insertion_point(field_mutable:sproto.UserLoginReq.Password)
  return password_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* UserLoginReq::release_password() {
  // @@protoc_insertion_point(field_release:sproto.UserLoginReq.Password)
  
  return password_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void UserLoginReq::set_allocated_password(::std::string* password) {
  if (password != NULL) {
    
  } else {
    
  }
  password_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), password);
  // @@protoc_insertion_point(field_set_allocated:sproto.UserLoginReq.Password)
}

// -------------------------------------------------------------------

// GetSessionReq

// string UserName = 1;
inline void GetSessionReq::clear_username() {
  username_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& GetSessionReq::username() const {
  // @@protoc_insertion_point(field_get:sproto.GetSessionReq.UserName)
  return username_.GetNoArena();
}
inline void GetSessionReq::set_username(const ::std::string& value) {
  
  username_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.GetSessionReq.UserName)
}
#if LANG_CXX11
inline void GetSessionReq::set_username(::std::string&& value) {
  
  username_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.GetSessionReq.UserName)
}
#endif
inline void GetSessionReq::set_username(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  username_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.GetSessionReq.UserName)
}
inline void GetSessionReq::set_username(const char* value, size_t size) {
  
  username_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.GetSessionReq.UserName)
}
inline ::std::string* GetSessionReq::mutable_username() {
  
  // @@protoc_insertion_point(field_mutable:sproto.GetSessionReq.UserName)
  return username_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* GetSessionReq::release_username() {
  // @@protoc_insertion_point(field_release:sproto.GetSessionReq.UserName)
  
  return username_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void GetSessionReq::set_allocated_username(::std::string* username) {
  if (username != NULL) {
    
  } else {
    
  }
  username_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), username);
  // @@protoc_insertion_point(field_set_allocated:sproto.GetSessionReq.UserName)
}

// -------------------------------------------------------------------

// OnLogin

// bool Broadcast = 1;
inline void OnLogin::clear_broadcast() {
  broadcast_ = false;
}
inline bool OnLogin::broadcast() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.Broadcast)
  return broadcast_;
}
inline void OnLogin::set_broadcast(bool value) {
  
  broadcast_ = value;
  // @@protoc_insertion_point(field_set:sproto.OnLogin.Broadcast)
}

// string Name = 2;
inline void OnLogin::clear_name() {
  name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& OnLogin::name() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.Name)
  return name_.GetNoArena();
}
inline void OnLogin::set_name(const ::std::string& value) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.OnLogin.Name)
}
#if LANG_CXX11
inline void OnLogin::set_name(::std::string&& value) {
  
  name_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.OnLogin.Name)
}
#endif
inline void OnLogin::set_name(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.OnLogin.Name)
}
inline void OnLogin::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.OnLogin.Name)
}
inline ::std::string* OnLogin::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:sproto.OnLogin.Name)
  return name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* OnLogin::release_name() {
  // @@protoc_insertion_point(field_release:sproto.OnLogin.Name)
  
  return name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OnLogin::set_allocated_name(::std::string* name) {
  if (name != NULL) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:sproto.OnLogin.Name)
}

// uint32 NameHash = 3;
inline void OnLogin::clear_namehash() {
  namehash_ = 0u;
}
inline ::google::protobuf::uint32 OnLogin::namehash() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.NameHash)
  return namehash_;
}
inline void OnLogin::set_namehash(::google::protobuf::uint32 value) {
  
  namehash_ = value;
  // @@protoc_insertion_point(field_set:sproto.OnLogin.NameHash)
}

// string OwnerAddr = 4;
inline void OnLogin::clear_owneraddr() {
  owneraddr_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& OnLogin::owneraddr() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.OwnerAddr)
  return owneraddr_.GetNoArena();
}
inline void OnLogin::set_owneraddr(const ::std::string& value) {
  
  owneraddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.OnLogin.OwnerAddr)
}
#if LANG_CXX11
inline void OnLogin::set_owneraddr(::std::string&& value) {
  
  owneraddr_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.OnLogin.OwnerAddr)
}
#endif
inline void OnLogin::set_owneraddr(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  owneraddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.OnLogin.OwnerAddr)
}
inline void OnLogin::set_owneraddr(const char* value, size_t size) {
  
  owneraddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.OnLogin.OwnerAddr)
}
inline ::std::string* OnLogin::mutable_owneraddr() {
  
  // @@protoc_insertion_point(field_mutable:sproto.OnLogin.OwnerAddr)
  return owneraddr_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* OnLogin::release_owneraddr() {
  // @@protoc_insertion_point(field_release:sproto.OnLogin.OwnerAddr)
  
  return owneraddr_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OnLogin::set_allocated_owneraddr(::std::string* owneraddr) {
  if (owneraddr != NULL) {
    
  } else {
    
  }
  owneraddr_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), owneraddr);
  // @@protoc_insertion_point(field_set_allocated:sproto.OnLogin.OwnerAddr)
}

// string Token = 5;
inline void OnLogin::clear_token() {
  token_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& OnLogin::token() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.Token)
  return token_.GetNoArena();
}
inline void OnLogin::set_token(const ::std::string& value) {
  
  token_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.OnLogin.Token)
}
#if LANG_CXX11
inline void OnLogin::set_token(::std::string&& value) {
  
  token_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.OnLogin.Token)
}
#endif
inline void OnLogin::set_token(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  token_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.OnLogin.Token)
}
inline void OnLogin::set_token(const char* value, size_t size) {
  
  token_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.OnLogin.Token)
}
inline ::std::string* OnLogin::mutable_token() {
  
  // @@protoc_insertion_point(field_mutable:sproto.OnLogin.Token)
  return token_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* OnLogin::release_token() {
  // @@protoc_insertion_point(field_release:sproto.OnLogin.Token)
  
  return token_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OnLogin::set_allocated_token(::std::string* token) {
  if (token != NULL) {
    
  } else {
    
  }
  token_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), token);
  // @@protoc_insertion_point(field_set_allocated:sproto.OnLogin.Token)
}

// string GateAddr = 6;
inline void OnLogin::clear_gateaddr() {
  gateaddr_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& OnLogin::gateaddr() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.GateAddr)
  return gateaddr_.GetNoArena();
}
inline void OnLogin::set_gateaddr(const ::std::string& value) {
  
  gateaddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:sproto.OnLogin.GateAddr)
}
#if LANG_CXX11
inline void OnLogin::set_gateaddr(::std::string&& value) {
  
  gateaddr_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:sproto.OnLogin.GateAddr)
}
#endif
inline void OnLogin::set_gateaddr(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  gateaddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:sproto.OnLogin.GateAddr)
}
inline void OnLogin::set_gateaddr(const char* value, size_t size) {
  
  gateaddr_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:sproto.OnLogin.GateAddr)
}
inline ::std::string* OnLogin::mutable_gateaddr() {
  
  // @@protoc_insertion_point(field_mutable:sproto.OnLogin.GateAddr)
  return gateaddr_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* OnLogin::release_gateaddr() {
  // @@protoc_insertion_point(field_release:sproto.OnLogin.GateAddr)
  
  return gateaddr_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OnLogin::set_allocated_gateaddr(::std::string* gateaddr) {
  if (gateaddr != NULL) {
    
  } else {
    
  }
  gateaddr_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), gateaddr);
  // @@protoc_insertion_point(field_set_allocated:sproto.OnLogin.GateAddr)
}

// uint64 GateClient = 7;
inline void OnLogin::clear_gateclient() {
  gateclient_ = GOOGLE_ULONGLONG(0);
}
inline ::google::protobuf::uint64 OnLogin::gateclient() const {
  // @@protoc_insertion_point(field_get:sproto.OnLogin.GateClient)
  return gateclient_;
}
inline void OnLogin::set_gateclient(::google::protobuf::uint64 value) {
  
  gateclient_ = value;
  // @@protoc_insertion_point(field_set:sproto.OnLogin.GateClient)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace sproto

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_INCLUDED_login_5fsvc_2eproto
