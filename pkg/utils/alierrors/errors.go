/*
Copyright 2024 The CloudPilot AI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package alierrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
)

func IsNotFound(err error) bool {
	var sdkError *tea.SDKError
	if errors.As(err, &sdkError) {
		if *sdkError.StatusCode == http.StatusNotFound {
			return true
		}
	}

	return false
}

func WithRequestID(requestID string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("requestId: %s, %w", requestID, err)
}
