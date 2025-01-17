<?hh
/**
 * Copyright (c) 2014, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *
 */

class A extends Exception {}

function f(?int $x): int {
  if ($x is null) {
    throw (new A());
  }

  return $x;
}
