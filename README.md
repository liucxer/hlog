# hlog

# 设计特点
* 只支持ERROR, WARNING, INFO三个等级, 使用更加简单，更加严格
```
ERROR: 用于错误信息记录，此类日志必须处理
WARNING: 用于保留监控日志，用户在性能优化，分析统计的时候关注
INFO: 用于调试信息记录，仅用于调试使用
```
* 只能输出到stdout。规范了日志输出的唯一途径。
* 日志带有颜色，利于本地调试

# 测试示例
![alt 测试样例](https://github.com/liucxer/hlog/log.jpg)