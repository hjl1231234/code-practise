# code-practise

portfolio


## docker
1. 
   修改 docker-compose.yml
   
   - 添加了 nginx 服务配置，使用 nginx:stable 镜像
   - 配置了端口映射（80和443端口）
   - 设置了卷挂载（配置文件、静态文件和日志目录）
   - 加入了 app-network 网络，使其能与其他服务通信
   - 添加了健康检查和自动重启策略
2. 
   创建目录结构
   
   - 创建了 nginx-conf 目录（存放Nginx配置文件）
   - 创建了 nginx-html 目录（存放静态网站内容）
   - 创建了 nginx-logs 目录（存放Nginx日志）
3. 
   创建配置文件
   
   - 在 nginx-conf 目录中创建了默认配置文件 default.conf
   - 配置了基础的HTTP服务器设置
   - 添加了代理其他服务的示例配置（注释状态）
4.
   创建测试页面
   
   - 在 nginx-html 目录中创建了 index.html 测试页面
## 使用
1. 
   启动所有服务（包括新添加的 Nginx）：
   
   ```
   cd /home/hjl/web3project/
   code-practise
   docker-compose up -d
   ```
2. 
   验证 Nginx 是否正常运行：
   
   - 打开浏览器访问 http://localhost
   - 应该能看到 "Welcome to Nginx!" 的测试页面
3. 
   自定义配置：
   
   - 修改 nginx-conf/default.conf 文件可以自定义 Nginx 配置
   - 在 nginx-html 目录中放置您的网站文件
   - 查看 nginx-logs 目录可以获取访问日志和错误日志
## 扩展
如果需要使用 Nginx 代理其他服务（如 Web 应用后端），可以取消默认配置文件中代理相关的注释，并根据实际情况修改代理设置。