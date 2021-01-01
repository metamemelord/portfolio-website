package model

// go:generate enumer -trimprefix=MetadataItemType -type=MetadataItemType -transform=snake -json -sql
type MetadataItemType int

const (
	MetadataItemTypeNone MetadataItemType = iota
	MetadataItemTypeQueryParameter
)
