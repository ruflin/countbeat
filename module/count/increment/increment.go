package increment

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
)

func init() {
	if err := mb.Registry.AddMetricSet("count", "increment", New); err != nil {
		panic(err)
	}
}

type MetricSet struct{
	mb.BaseMetricSet
	counter int64
}

// New creates and returns a new instance of MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

func (m *MetricSet) Fetch() (common.MapStr, error) {

	m.counter++

	return common.MapStr{
		"counter": m.counter,
	}, nil
}
