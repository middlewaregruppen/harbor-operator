package util

func PtrBool(b bool) *bool {
	return &b
}

func PtrInt64(b int64) *int64 {
	return &b
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
