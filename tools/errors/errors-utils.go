package errors

func GetErrorStringOrDefault(err error, def string) string {
	if err == nil {
		return def
	}
	return err.Error()
}
