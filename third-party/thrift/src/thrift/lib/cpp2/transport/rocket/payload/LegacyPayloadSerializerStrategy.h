/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

#include <thrift/lib/cpp2/transport/rocket/PayloadUtils.h>
#include <thrift/lib/cpp2/transport/rocket/payload/PayloadSerializerStrategy.h>

namespace apache::thrift::rocket {

/**
 * Strategy that delegates to the PayloadUtils.h header free functions.
 */
class LegacyPayloadSerializerStrategy final
    : public PayloadSerializerStrategy<LegacyPayloadSerializerStrategy> {
 public:
  LegacyPayloadSerializerStrategy() : PayloadSerializerStrategy(*this) {}

  template <class T>
  FOLLY_ERASE folly::Try<T> unpackAsCompressed(
      rocket::Payload&& payload, bool useBinary) {
    return ::apache::thrift::rocket::unpackAsCompressed<T>(
        std::move(payload), useBinary);
  }

  template <class T>
  FOLLY_ERASE folly::Try<T> unpack(rocket::Payload&& payload, bool useBinary) {
    return ::apache::thrift::rocket::unpack<T>(std::move(payload), useBinary);
  }

  template <typename T>
  FOLLY_ERASE std::unique_ptr<folly::IOBuf> packCompact(T&& data) {
    return ::apache::thrift::rocket::packCompact<T>(std::forward<T>(data));
  }

  template <typename T>
  size_t unpackCompact(T& output, const folly::IOBuf* buffer) {
    return ::apache::thrift::rocket::unpackCompact<T>(output, buffer);
  }

  template <class PayloadType>
  FOLLY_ERASE rocket::Payload pack(
      PayloadType&& payload, folly::AsyncTransport* transport) {
    return ::apache::thrift::rocket::pack(
        std::forward<PayloadType>(payload), transport);
  }
};

} // namespace apache::thrift::rocket
