(*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

(** This function never returns; it always calls [Exit.exit] *)
val main :
  ClientEnv.client_check_env ->
  ServerConfig.t ->
  ServerLocalConfig.t ->
  init_proc_stack:string list option ->
  'a
