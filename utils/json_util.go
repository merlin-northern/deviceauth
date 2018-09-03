// Copyright 2018 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package utils

import (
	"encoding/json"
)

// JsonSort decodes and re-encodes a json string to get a lexical sort on json keys.
// An error is returned if it's not in fact json.
func JsonSort(what string) (string, error) {

	var dec map[string]interface{}
	err := json.Unmarshal([]byte(what), &dec)
	if err != nil {
		return "", err
	}

	enc, err := json.Marshal(dec)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
