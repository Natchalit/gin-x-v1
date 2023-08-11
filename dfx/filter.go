package dfx

import (
	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
	"github.com/Natchalit/gin-x-v1/ginx"
	"github.com/go-gota/gota/dataframe"
)

func (df *DataframeX) Filter(filters ...dataframe.F) (*DataframeX, error) {

	if len(filters) == 0 {
		return nil, ginx.BadRequest(`no filter`)
	}
	for _, f := range filters {
		buf := caseconvert.ToSnake(f.Colname)
		f.Colname = buf
	}

	dfx := df.Dataframe.Filter(filters...)

	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error [%s : %s] `, dfx.Error(), filters[0].Colname)
	}

	df.Dataframe = dfx

	return df, nil
}

// agg { AND , OR }
func (df *DataframeX) FilterAggregation(agg dataframe.Aggregation, filters ...dataframe.F) (*DataframeX, error) {

	if len(filters) == 0 {
		return nil, ginx.BadRequest(`no filter`)
	}

	for _, f := range filters {
		f.Colname = caseconvert.ToSnake(f.Colname)
	}

	dfx := df.Dataframe.FilterAggregation(agg, filters...)

	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error [%s : %s] `, dfx.Error(), filters[0].Colname)
	}

	df.Dataframe = dfx

	return df, nil
}

func (grp *GroupX) FilterAggregation(typs []dataframe.AggregationType, colnames []string) (*DataframeX, error) {

	dfx := grp.Group.Aggregation(typs, colnames)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error group by [%s]`, dfx.Error())
	}

	return &DataframeX{
		Dataframe: dfx,
	}, nil
}
