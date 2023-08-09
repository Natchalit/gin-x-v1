package sqlx

import (
	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
	"github.com/Natchalit/gin-x-v1/tox"
	"github.com/Natchalit/gin-x-v1/validx"
)

func (m *Map) getCase(col string) interface{} {
	if validx.IsSnakeCase(col) {
		if v, ok := (*m)[col]; ok {
			return v
		}
	} else {
		return m.getCase(caseconvert.ToSnake(col))
	}
	return nil
}

func (m *Map) String(col string) string {
	if m == nil {
		return tox.STRING
	}

	return tox.String(m.getCase(col))
}

func (m *Map) StringPtr(col string) *string {
	if m == nil {
		return &tox.STRING
	}

	return tox.StringPtr(m.getCase(col))
}

func (m *Map) Int(col string) int {
	if m == nil {
		return tox.INT
	}

	return tox.Int(m.getCase(col))
}

func (m *Map) IntPtr(col string) *int {
	if m == nil {
		return &tox.INT
	}

	return tox.IntPtr(m.getCase(col))
}

func (m *Map) Int64(col string) int64 {
	if m == nil {
		return tox.INT64
	}

	return tox.Int64(m.getCase(col))
}

func (m *Map) Float64(col string) float64 {
	if m == nil {
		return tox.FLOAT64
	}

	return tox.Float64(m.getCase(col))
}

func (m *Map) Float64Ptr(col string) *float64 {
	if m == nil {
		return &tox.FLOAT64
	}

	return tox.Float64Ptr(m.getCase(col))
}

func (m *Map) Bool(col string) bool {
	if m == nil {
		return tox.FALSE
	}

	return tox.Bool(m.getCase(col))
}

func (m *Map) BoolPtr(col string) *bool {
	if m == nil {
		return &tox.FALSE
	}

	return tox.BoolPtr(m.getCase(col))
}

func DataMapToRow(mapData []map[string]interface{}) *Row {
	dm := Row{}
	chkDm := map[string]bool{}
	for k, element := range mapData {
		dm.Rows = append(dm.Rows, element)
		if k == 0 {
			for k := range mapData[0] {
				if !chkDm[k] && k != `x0` {
					chkDm[k] = true
					dm.Columns = append(dm.Columns, caseconvert.ToSnake(k))
				}
			}
		}
	}

	return &dm
}
