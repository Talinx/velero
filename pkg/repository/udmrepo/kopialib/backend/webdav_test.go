/*
Copyright the Velero contributors.

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

package backend

import (
	"context"
	"testing"

	"github.com/kopia/kopia/repo/blob/webdav"
	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/velero/pkg/repository/udmrepo"
)

func TestWebDAVSetup(t *testing.T) {
	testCases := []struct {
		name            string
		flags           map[string]string
		expectedOptions webdav.Options
		expectedErr     string
	}{
		{
			name: "must have URL",
			flags: map[string]string{
				udmrepo.StoreOptionsWebDAVUsername: "fakeUsername",
				udmrepo.StoreOptionsWebDAVPassword: "fake password",
			},
			expectedErr: "key " + udmrepo.StoreOptionsWebDAVURL + " not found",
		},
		{
			name: "must have username",
			flags: map[string]string{
				udmrepo.StoreOptionsWebDAVURL:      "fake-url",
				udmrepo.StoreOptionsWebDAVPassword: "fake password",
			},
			expectedErr: "key " + udmrepo.StoreOptionsWebDAVUsername + " not found",
		},
		{
			name: "must have password",
			flags: map[string]string{
				udmrepo.StoreOptionsWebDAVURL:      "fake-url",
				udmrepo.StoreOptionsWebDAVUsername: "fakeUsername",
			},
			expectedErr: "key " + udmrepo.StoreOptionsWebDAVPassword + " not found",
		},
		{
			name: "with minimum required flags",
			flags: map[string]string{
				udmrepo.StoreOptionsWebDAVURL:      "fake-url",
				udmrepo.StoreOptionsWebDAVUsername: "fakeUsername",
				udmrepo.StoreOptionsWebDAVPassword: "fake password",
			},
			expectedOptions: webdav.Options{
				URL:      "fake-url",
				Username: "fakeUsername",
				Password: "fake password",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			webDAVFlags := WebDAVBackend{}

			err := webDAVFlags.Setup(context.Background(), tc.flags)

			if tc.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
