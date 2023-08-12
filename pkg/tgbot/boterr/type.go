package boterr

const (
	DeleteNotFound = "Bad Request: message to delete not found"
	CantbeDelete   = "Bad Request: message can't be deleted"
)

func IsNotFound(err error) bool {
	return err.Error() == DeleteNotFound
}

func IsCantBeDelete(err error) bool {
	return err.Error() == CantbeDelete
}
