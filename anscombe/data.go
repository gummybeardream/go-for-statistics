package anscombe

import (
	"github.com/montanaflynn/stats"
)

func DatasetOne() stats.Series {
	datasetOne := stats.Series{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}

	return datasetOne
}
