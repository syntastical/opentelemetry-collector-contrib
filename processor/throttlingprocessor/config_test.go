// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package throttlingprocessor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidateConfig(t *testing.T) {
	testCases := []struct {
		desc        string
		cfg         *Config
		expectedErr error
	}{
		{
			desc: "invalid Threshold config",
			cfg: &Config{
				Threshold:     0,
				Interval:      time.Second,
				KeyExpression: "true",
			},
			expectedErr: errInvalidThreshold,
		},
		{
			desc: "invalid Interval config",
			cfg: &Config{
				Threshold:     1,
				Interval:      0,
				KeyExpression: "true",
			},
			expectedErr: errInvalidInterval,
		},
		{
			desc: "invalid KeyExpression config",
			cfg: &Config{
				Threshold:     1,
				Interval:      time.Second,
				KeyExpression: "",
			},
			expectedErr: errInvalidKeyExpression,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.expectedErr != nil {
				require.ErrorContains(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
