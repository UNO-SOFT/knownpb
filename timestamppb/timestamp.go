// Copyright 2019, 2021 Tamás Gulácsi
//
// SPDX-License-Identifier: Apache-2.0

package timestamppb

import "github.com/godror/knownpb/timestamppb"

// Timestamp is a wrapped timestamppb.Timestamp with proper JSON and XML marshaling (as string date).
type Timestamp = timestamppb.Timestamp
