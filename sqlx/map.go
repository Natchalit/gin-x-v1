package sqlx

import (
	"strings"
	"sync"

	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
	"github.com/Natchalit/gin-x-v1/tox"
	"github.com/Natchalit/gin-x-v1/validx"
)

var (
	_LockMap = &sync.Mutex{}
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

		buffVal := map[string]interface{}{}
		for index, element := range element {
			buffVal[caseconvert.ToSnake(index)] = element
		}
		dm.Rows = append(dm.Rows, buffVal)

		if k == 0 {
			for k := range mapData[0] {
				k = caseconvert.ToSnake(k)
				if k == `x0` {
					continue
				}
				if !chkDm[k] {
					chkDm[k] = true
					dm.Columns = append(dm.Columns, k)
				}
			}
		}
	}

	for _, element := range dm.Rows {
		element.DeleteKey(`x0`)
	}

	return &dm
}

func RowToDataMap(row *Row) *[]map[string]interface{} {
	dataMap := []map[string]interface{}{}

	// Create a map to track column positions in the Rows
	colIndexMap := map[string]int{}
	for idx, col := range row.Columns {
		colIndexMap[col] = idx
	}

	for _, rowValues := range row.Rows {
		rowMap := map[string]interface{}{}
		for colName, colValue := range rowValues {
			if _, exists := colIndexMap[colName]; exists {
				rowMap[colName] = colValue
			}
		}
		dataMap = append(dataMap, rowMap)
	}

	return &dataMap
}

// หาค่าคอลัมล์ case-insensitivity , snake-case
func (m *Map) Get(col string) any {
	if m == nil {
		return nil
	}
	val := m.getCase(col)
	return val
}

func (m *Map) Set(col string, val any) {
	if m != nil {
		m.setCase(col, val)
	}
}

func (s *Map) setCase(col string, val any) {
	_LockMap.Lock()
	defer _LockMap.Unlock()
	// หาตามชื่อคอลัมล์ case-sensitivity.
	if _, ok := (*s)[col]; ok {
		(*s)[col] = val
		return
	}
	// หาตามชื่อคอลัมล์ case-lower
	colLower := strings.ToLower(col)
	if colLower != col {
		if _, ok := (*s)[colLower]; ok {
			(*s)[colLower] = val
			return
		}
	}
	// หาตามชื่อคอลัมล์ case-upper
	colUpper := strings.ToUpper(col)
	if colUpper != col {
		if _, ok := (*s)[colUpper]; ok {
			(*s)[colUpper] = val
			return
		}
	}
	// SnakeCase
	colSnake := caseconvert.ToSnake(col)
	if colSnake != col && colSnake != colLower {
		if _, ok := (*s)[colSnake]; ok {
			(*s)[colSnake] = val
			return
		}
	}
	// เพิ่มคอลัมล์
	(*s)[col] = val
}

func (m *Map) DeleteKey(key string) {
	if m == nil {
		return
	}
	_LockMap.Lock()
	defer _LockMap.Unlock()
	for k := range *m {
		if strings.EqualFold(k, key) {
			delete(*m, k)
		}
	}
}

func (m *Row) LenCols() int {
	if m == nil {
		return 0
	}

	if len(m.Rows) == 0 {
		return 0
	}

	return len(m.Rows[0])
}
