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

package goldendataset

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/model/pdata"
)

func TestGenDefault(t *testing.T) {
	md := MetricsFromCfg(DefaultCfg())
	require.Equal(t, 1, md.MetricCount())
	require.Equal(t, 1, md.DataPointCount())
	rms := md.ResourceMetrics()
	rm := rms.At(0)
	resource := rm.Resource()
	rattrs := resource.Attributes()
	rattrs.Len()
	require.Equal(t, 1, rattrs.Len())
	val, _ := rattrs.Get("resource-attr-name-0")
	require.Equal(t, "resource-attr-val-0", val.StringVal())
	ilms := rm.InstrumentationLibraryMetrics()
	require.Equal(t, 1, ilms.Len())
	ms := ilms.At(0).Metrics()
	require.Equal(t, 1, ms.Len())
	pdm := ms.At(0)
	require.Equal(t, "metric_0", pdm.Name())
	require.Equal(t, "my-md-description", pdm.Description())
	require.Equal(t, "my-md-units", pdm.Unit())

	require.Equal(t, pdata.MetricDataTypeGauge, pdm.DataType())
	pts := pdm.Gauge().DataPoints()
	require.Equal(t, 1, pts.Len())
	pt := pts.At(0)

	require.Equal(t, 1, pt.LabelsMap().Len())
	ptLabels, _ := pt.LabelsMap().Get("pt-label-key-0")
	require.Equal(t, "pt-label-val-0", ptLabels)

	require.EqualValues(t, 940000000000000000, pt.StartTimestamp())
	require.EqualValues(t, 940000000000000042, pt.Timestamp())
	require.EqualValues(t, 1, pt.IntVal())
}

func TestDoubleHistogramFunctions(t *testing.T) {
	pt := pdata.NewHistogramDataPoint()
	setDoubleHistogramBounds(pt, 1, 2, 3, 4, 5)
	require.Equal(t, 5, len(pt.ExplicitBounds()))
	require.Equal(t, 5, len(pt.BucketCounts()))

	addDoubleHistogramVal(pt, 1)
	require.EqualValues(t, 1, pt.Count())
	require.EqualValues(t, 1, pt.Sum())
	require.EqualValues(t, 1, pt.BucketCounts()[0])

	addDoubleHistogramVal(pt, 2)
	require.EqualValues(t, 2, pt.Count())
	require.EqualValues(t, 3, pt.Sum())
	require.EqualValues(t, 1, pt.BucketCounts()[1])

	addDoubleHistogramVal(pt, 2)
	require.EqualValues(t, 3, pt.Count())
	require.EqualValues(t, 5, pt.Sum())
	require.EqualValues(t, 2, pt.BucketCounts()[1])
}

func TestGenDoubleHistogram(t *testing.T) {
	cfg := DefaultCfg()
	cfg.MetricDescriptorType = pdata.MetricDataTypeHistogram
	cfg.PtVal = 2
	md := MetricsFromCfg(cfg)
	pts := getMetric(md).Histogram().DataPoints()
	pt := pts.At(0)
	buckets := pt.BucketCounts()
	require.Equal(t, 5, len(buckets))
	require.EqualValues(t, 2, buckets[2])
}

func TestGenDoubleGauge(t *testing.T) {
	cfg := DefaultCfg()
	cfg.MetricDescriptorType = pdata.MetricDataTypeGauge
	md := MetricsFromCfg(cfg)
	metric := getMetric(md)
	pts := metric.Gauge().DataPoints()
	require.Equal(t, 1, pts.Len())
	pt := pts.At(0)
	require.EqualValues(t, float64(1), pt.IntVal())
}

func getMetric(md pdata.Metrics) pdata.Metric {
	return md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0)
}
