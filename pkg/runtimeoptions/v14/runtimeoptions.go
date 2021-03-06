// Copyright 2018 The containerd Authors.
// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package v14 contains the runtimeoptions proto for containerd 1.4 and earlier.
// The package for runtimeoptions proto changed from "cri.runtimeoptions.v1" to
// "runtimeoptions.v1" in 1.5, So keep both versions until 1.4 doesn't need to
// be supported anymore.
package v14
