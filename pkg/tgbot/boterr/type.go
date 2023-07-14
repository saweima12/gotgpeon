package boterr

const (
	DeleteNotFound = "Bad Request: message to delete not found"
)

func IsNotFound(err error) bool {
	return err.Error() == DeleteNotFound
}
