// Copyright 2016 Eric Wollesen <ericw at xmtp dot net>
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

package setup

import (
	"strings"

	"xmtp.net/xmtpbot/urls"
	"xmtp.net/xmtpbot/urls/json"
	"xmtp.net/xmtpbot/urls/memory"
)

func NewStore() urls.Store {
	return NewStoreFromFilename(*urls.StoreFilename)
}

func NewStoreFromFilename(filename string) urls.Store {
	switch strings.ToLower(*urls.StoreType) {
	case "json":
		return json.New(filename)
	case "memory":
	default:
	}

	return memory.New()
}
