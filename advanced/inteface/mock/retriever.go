package mock

type Retriever struct {
	Contents string
}

func (r Retriever) String() string {
	//TODO implement me
	panic("implement me")
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
