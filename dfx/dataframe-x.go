package dfx

import (
	"github.com/Natchalit/gin-x-v1/errorx"
	"github.com/go-gota/gota/dataframe"
)

func LoadMaps(maps []map[string]interface{}) (*DataframeX, error) {
	dfx := dataframe.LoadMaps(maps)
	if dfx.Err != nil {
		return nil, errorx.BadRequest(`error load maps : %s`, dfx.Error())
	}

	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func (df *DataframeX) Filter(filters ...dataframe.F) (*DataframeX, error) {
	dfx := df.Dataframe.Filter(filters...)
	if dfx.Err != nil {
		return nil, errorx.BadRequest(`error filter : %s `, dfx.Error())
	}

	df.Dataframe = dfx

	return df, nil
}

func (df *DataframeX) Select(indexes dataframe.SelectIndexes) (*DataframeX, error) {
	dfx := df.Dataframe.Select(indexes)
	if dfx.Err != nil {
		return nil, errorx.BadRequest(`error select : %s`, dfx.Error())
	}

	return df, nil
}

func (df *DataframeX) Map() []map[string]interface{} {
	return df.Dataframe.Maps()
}
