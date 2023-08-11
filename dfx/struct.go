package dfx

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type DataframeX struct {
	Dataframe dataframe.DataFrame
}

type GroupX struct {
	Group *dataframe.Groups
}

type SeriesX struct {
	Series series.Series
}
