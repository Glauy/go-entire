package main

type Post struct {
	Name    string
	Address string
}

type Service interface {
	ListPosts() ([]*Post, error)
}

func ListPosts(svc Service) ([]*Post, error) {
	return svc.ListPosts()
}

// not_recommend
// 不被推荐的写法，因为，function依赖了，数据库连接信息，不方便单测
// func ListPosts(client *grpc.ClientConn) ([]*Post, error) {
// 	return client.ListPosts()
// }
