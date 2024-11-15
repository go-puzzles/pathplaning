# PathPlaning
PathPlaning 是一个用于路径规划的示例项目，该项目实现了一个简单的图形结构，并提供了相应的接口和实现。

## 特性
- 实现 A* 算法进行路径搜索
- 实现 BFS 算法进行路径搜索
- 支持设置障碍物
- 可视化路径搜索过程

## 安装

1. 克隆该仓库：
   ```bash
   git clone https://github.com/go-puzzles/pathplaning.git
   cd pathplaning
   ```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

## 使用
### BFS 搜索示例
```go
graph := NewSimpleGraph(6, 6)

start := &SimplePoint{X: 0, Y: 0}
goal := &SimplePoint{X: 4, Y: 3}

path, err := BFSSearch(graph, start, goal)
```

### A* 搜索示例
```go
graph := NewSimpleGraph(6, 6)

start := &SimplePoint{X: 0, Y: 0}
goal := &SimplePoint{X: 4, Y: 3}

path, err := AStarSearch(graph, start, goal)
```

### 设置障碍物
使用 `SetBlock` 方法设置图形中的障碍物。
```go
graph.SetBlock(&SimplePoint{X: 3, Y: 2})
graph.SetBlock(&SimplePoint{X: 3, Y: 4})
graph.SetBlock(&SimplePoint{X: 2, Y: 3})
```

## 运行测试
您可以使用以下命令运行测试：
```bash
go test -v .
```

## 许可证
该项目使用 MIT 许可证。有关详细信息，请参阅 [LICENSE](LICENSE) 文件。

## 贡献
欢迎任何形式的贡献！请提交问题或拉取请求。

## 联系
如有任何问题，请联系 [Hoven](https://hovenyang.best@gmail.com)。
