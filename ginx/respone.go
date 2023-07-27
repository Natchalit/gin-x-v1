package ginx

func Response(data map[string]any, isWeb bool, path *string) map[string]any {
	return map[string]any{
		`data`:  data,
		`isWeb`: isWeb,
		`path`:  path,
	}
}
