package dfx

import (
	"github.com/Natchalit/gin-x-v1/errorx"
	"github.com/go-gota/gota/dataframe"
)

type DataframeX struct {
	Dataframe dataframe.DataFrame
}

func (df *DataframeX) Filter(filters ...dataframe.F) (*DataframeX, error) {
	dfx := df.Dataframe.Filter(filters...)
	if dfx.Err != nil {
		return nil, errorx.BadRequest(`error filter : %s `, dfx.Err.Error())
	}

	df.Dataframe = dfx

	return df, nil
}
