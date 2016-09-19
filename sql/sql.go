// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"time"
)

const (
	// Parsing format for date times
	RFCDate = "2006-01-02"
	// Parsing format for normal times
	RFCNorm = time.RFC1123
	// Parsing format for json date times
	RFCTime = time.RFC3339
	// Parsing format for json nanosecond times
	RFCNano = time.RFC3339Nano
	// Parsing format for readable text format date times
	RFCText = "2006-01-02 15:04:05.999999999 -0700 MST"
)