package registry

// 服务抽象
type Service struct {
	Name  string  `json:"name"`  // 服务的名称
	Nodes []*Node `json:"nodes"`  // 服务的节点
}

// 服务节点的抽象
type Node struct {
	Id     string `json:"id"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}
