# gorm

## 简介

- 全功能 ORM
- 关联 (Has One，Has Many，Belongs To，Many To Many，多态，单表继承)
- Create，Save，Update，Delete，Find 中钩子方法
- 支持 Preload、Joins 的预加载
- 事务，嵌套事务，Save Point，Rollback To Saved Point
- Context，预编译模式，DryRun 模式
- 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUD
- SQL 构建器，Upsert，数据库锁，Optimizer/Index/Comment Hint，命名参数，子查询
- 复合主键，索引，约束
- Auto Migration
- 自定义 Logger
- 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…
- 每个特性都经过了测试的重重考验
- 开发者友好

## 结构介绍

### Config

```golang
type Config struct {
  SkipDefaultTransaction   bool//是否启动事务操作，默认启动(false)
  NamingStrategy           schema.Namer//替换默认命名策略(NamingStrategy,此默认策略也可部分配置)，实现schema.Namer接口
  Logger                   logger.Interface//替换默认log
  NowFunc                  func() time.Time//更新创建时间函数
  DryRun                   bool//生成 SQL 但不执行
  PrepareStmt              bool//查看 Session 获取详情
  DisableNestedTransaction bool//查看 Session 获取详情
  AllowGlobalUpdate        bool//查看 Session 获取详情
  DisableAutomaticPing     bool//在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性
  DisableForeignKeyConstraintWhenMigrating bool//在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束
}
```

### DB

### Session

## 参考资料

1. [doc](https://gorm.io/zh_CN/docs/index.html)
2. [GORM中文文档](https://jasperxu.github.io/gorm-zh/)
