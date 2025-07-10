package models

type InternalError string

const (
	ResourceReadEception      InternalError = "ResourceReadException"
	ResourceCreateException   InternalError = "ResourceCreateException"
	ResourceUpdateException   InternalError = "ResourceUpdateException"
	ResourceDeleteException   InternalError = "ResourceDeleteException"
	ResourceNotFoundException InternalError = "ResourceNotFoundException"
	IllegalArgumentException  InternalError = "IllegalArgumentException"
)
