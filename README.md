# simpleRemoteControl

设计思路:

- DDNS 客户端: 通过 DDNS 服务商将 IP 地址映射到域名
- 插件系统
- API 服务器: 接收客户端的请求，并将请求转发给插件

插件系统实现:

- 动态加载
- 运行
- 卸载
- 更新

API 端点:

/v1/auth: 验证客户端身份
/v1/plugins: 获取插件列表
/v1/plugins/{plugin_name}: 获取插件信息
/v1/plugins/{plugin_id}: 获取插件信息