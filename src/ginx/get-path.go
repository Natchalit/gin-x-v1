package ginx

func getPath(path, relativePath string) string {

	if relativePath != "" && relativePath != "/" {
		path = path + "/" + relativePath
	}

	return path

}
