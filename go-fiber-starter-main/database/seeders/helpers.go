package seeders

func ptrBool(b bool) *bool        { return &b }
func ptrInt(i int) *int           { return &i }
func ptrFloat(f float64) *float64 { return &f }
func ptrString(s string) *string  { return &s }
