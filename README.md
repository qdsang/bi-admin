BI Admin
------------------
Golang + Vue = BI Dashboard

> A Data Analysis Board. 

# Project Guidelines
[Online Documentation](http://www/)

## Roadmap

- [x] 实验室 sql测试 历史管理
- [ ] 定时任务 邮件和webhook
- [ ] 定时任务 sql数据快照
- [ ] 数据源 表管理重构
- [ ] 数据源 新增视图表
- [ ] 项目管理
- [ ] 图表 重构组件
- [ ] 图表 组件配置
- [ ] 图表 数据缓存功能
- [ ] 图表 新增文字模板类型
- [ ] 图表 js计算字段
- [ ] 面板 定制css
- [ ] 面板 自定义参数
- [ ] 权限管理 功能验证 表列权限
- [ ] 图表 高级模式 支持自定义sql


## 启动

```
cd web
npm i
npm run serve


cd server
go run main.go

```

## 依赖项目

前端基于 vue-data-board
- https://github.com/dongsuo/vue-data-board

后端基于 gin-vue-admin
- https://github.com/flipped-aurora/gin-vue-admin

