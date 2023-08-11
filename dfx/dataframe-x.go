package dfx

import (
	"github.com/Natchalit/gin-x-v1/ginx"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func LoadMaps(maps []map[string]interface{}) (*DataframeX, error) {
	dfx := dataframe.LoadMaps(maps)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error [%s]`, dfx.Error())
	}

	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func (df *DataframeX) Select(indexes dataframe.SelectIndexes) (*DataframeX, error) {

	dfx := df.Dataframe.Select(indexes)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error select [%s]`, dfx.Error())
	}

	return df, nil
}

func (df *DataframeX) Maps() []map[string]interface{} {
	return df.Dataframe.Maps()
}

func (df *DataframeX) GroupBy(cols ...string) (*GroupX, error) {

	if len(cols) == 0 {
		return nil, ginx.BadRequest(`not found group by`)
	}

	dfx := df.Dataframe.GroupBy(cols...)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error [%s]`, dfx.Err)
	}

	return &GroupX{
		Group: dfx,
	}, nil
}

func (df *DataframeX) Arrange(order ...dataframe.Order) (*DataframeX, error) {

	dfx := df.Dataframe.Arrange(order...)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error arrange [%s]`, dfx.Error())
	}

	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func (df *DataframeX) InnerJoin(b dataframe.DataFrame, keys ...string) (*DataframeX, error) {

	dfx := df.Dataframe.InnerJoin(b, keys...)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error inner join [%s]`, dfx.Error())
	}

	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func (df *DataframeX) LeftJoin(b dataframe.DataFrame, keys ...string) (*DataframeX, error) {

	dfx := df.Dataframe.LeftJoin(b, keys...)
	if dfx.Err != nil {
		return nil, ginx.BadRequest(`error left join [%s]`, dfx.Error())
	}

	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func (df *DataframeX) Describe() *DataframeX {
	dfx := df.Dataframe.Describe()

	return &DataframeX{
		Dataframe: dfx,
	}
}

func (df *DataframeX) Dims() (int, int) {
	return df.Dataframe.Dims()
}

func (df *DataframeX) Types() []series.Type {
	return df.Dataframe.Types()
}

func (df *DataframeX) Names() []string {
	return df.Dataframe.Names()
}

func (df *DataframeX) LengthRow() int {
	return df.Dataframe.Nrow()
}

func (df *DataframeX) LengthCol() int {
	return df.Dataframe.Ncol()
}

func (df *DataframeX) Col(colName string) (*SeriesX, error) {
	col := df.Dataframe.Col(colName)
	if col.Err != nil {
		return nil, ginx.BadRequest(`error [%s]`, col.Error())
	}

	return &SeriesX{
		Series: col,
	}, nil
}

/*
	df.Dataframe.CBind()
*/
// func (df *DataframeX) () {

// }
// func (df *DataframeX) () {

// }
// func (df *DataframeX) () {

// }
// func (df *DataframeX) () {

// }
// func (df *DataframeX) () {

// }
// func (df *DataframeX) () {

// }
