package discovery

type Post struct {
	ID          string
	OwnerID     string
	FrontPicUrl string
	BackPicUrl  string
}

type AddPost struct {
	OwnerID     string
	FrontPicUrl string
	BackPicUrl  string
}

type Page[T any] struct {
	Data   []T
	Cursor string
}
