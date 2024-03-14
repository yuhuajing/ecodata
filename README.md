# scanEVMBlockData

1. 启动mongodb数据库 dock容器

> docker run -d --name my-mongodb   -e MONGO_INITDB_ROOT_USERNAME=root   -e MONGO_INITDB_ROOT_PASSWORD=password   -p 27017:27017   mongo:latest

通过 mongoDB Compass 软件连接本地 数据库, 创建数据库： drugtra

2. 启动后端ot
> go run ./scanblockdata.go

3. 前端界面

> http://localhost:3000/




