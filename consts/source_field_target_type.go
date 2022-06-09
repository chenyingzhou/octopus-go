package consts

type SourceFieldTargetType string

const (
	SourceFieldTargetTypeBoolExists   SourceFieldTargetType = "BOOL_EXISTS"
	SourceFieldTargetTypeBoolMatchAny SourceFieldTargetType = "BOOL_MATCH_ANY"
	SourceFieldTargetTypeBoolMatchAll SourceFieldTargetType = "BOOL_MATCH_ALL"
	SourceFieldTargetTypeNumberCount  SourceFieldTargetType = "NUMBER_COUNT"
	SourceFieldTargetTypeNumberSum    SourceFieldTargetType = "NUMBER_SUM"
	SourceFieldTargetTypeNumberMax    SourceFieldTargetType = "NUMBER_MAX"
	SourceFieldTargetTypeNumberMin    SourceFieldTargetType = "NUMBER_MIN"
)
