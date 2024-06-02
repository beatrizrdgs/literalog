package cerrors

var ErrInternal = New(500, map[string][]any{"internal": {"server error"}})
