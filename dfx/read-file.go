package dfx

import (
	"io"

	"github.com/go-gota/gota/dataframe"
)

func ReadCSV(r io.Reader, options ...dataframe.LoadOption) (*DataframeX, error) {
	dfx := dataframe.ReadCSV(r, options...)
	if dfx.Err != nil {
		return nil, dfx.Error()
	}
	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func ReadJson(r io.Reader, options ...dataframe.LoadOption) (*DataframeX, error) {
	dfx := dataframe.ReadJSON(r, options...)
	if dfx.Err != nil {
		return nil, dfx.Error()
	}
	return &DataframeX{
		Dataframe: dfx,
	}, nil
}

func ReadHTML(r io.Reader, options ...dataframe.LoadOption) ([]DataframeX, error) {
	dfx := dataframe.ReadHTML(r, options...)

	resp := []DataframeX{}
	for _, val := range dfx {
		if val.Err != nil {
			return nil, val.Error()
		}
		resp = append(resp, DataframeX{
			Dataframe: val,
		})
	}

	return resp, nil
}
