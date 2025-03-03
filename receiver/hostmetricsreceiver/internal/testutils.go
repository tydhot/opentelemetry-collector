// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/model/pdata"
)

func AssertContainsAttribute(t *testing.T, attr pdata.AttributeMap, key string) {
	_, ok := attr.Get(key)
	assert.True(t, ok)
}

func AssertDescriptorEqual(t *testing.T, expected pdata.Metric, actual pdata.Metric) {
	assert.Equal(t, expected.Name(), actual.Name())
	assert.Equal(t, expected.Description(), actual.Description())
	assert.Equal(t, expected.Unit(), actual.Unit())
	assert.Equal(t, expected.DataType(), actual.DataType())
}

func AssertSumMetricHasLabelValue(t *testing.T, metric pdata.Metric, index int, labelName string, expectedVal string) {
	val, ok := metric.Sum().DataPoints().At(index).Attributes().Get(labelName)
	assert.Truef(t, ok, "Missing attribute %q in metric %q", labelName, metric.Name())
	assert.Equal(t, expectedVal, val.StringVal())
}

func AssertSumMetricHasLabel(t *testing.T, metric pdata.Metric, index int, labelName string) {
	_, ok := metric.Sum().DataPoints().At(index).Attributes().Get(labelName)
	assert.Truef(t, ok, "Missing attribute %q in metric %q", labelName, metric.Name())
}

func AssertSumMetricStartTimeEquals(t *testing.T, metric pdata.Metric, startTime pdata.Timestamp) {
	ddps := metric.Sum().DataPoints()
	for i := 0; i < ddps.Len(); i++ {
		require.Equal(t, startTime, ddps.At(i).StartTimestamp())
	}
}

func AssertSameTimeStampForAllMetrics(t *testing.T, metrics pdata.MetricSlice) {
	AssertSameTimeStampForMetrics(t, metrics, 0, metrics.Len())
}

func AssertSameTimeStampForMetrics(t *testing.T, metrics pdata.MetricSlice, startIdx, endIdx int) {
	var ts pdata.Timestamp
	for i := startIdx; i < endIdx; i++ {
		metric := metrics.At(i)
		if metric.DataType() == pdata.MetricDataTypeSum {
			ddps := metric.Sum().DataPoints()
			for j := 0; j < ddps.Len(); j++ {
				if ts == 0 {
					ts = ddps.At(j).Timestamp()
				}
				require.Equalf(t, ts, ddps.At(j).Timestamp(), "metrics contained different end timestamp values")
			}
		}
	}
}
