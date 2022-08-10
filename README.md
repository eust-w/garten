# garten

## 简介
garten是一个ssh cmd、http、api、bash cmd 编排工具，使用json/yaml写编排内容，使用garten管理运行编排内容

garten 支持 有限循环、超时停止、定时执行、判断运行、支持热加载(最小颗粒度为块)，实际是从管道中执行

garten 允许 暂停执行，在当前位置停止，等待允许继续执行后再执行

garten 支持 通过yaml 或者 json的形式导入 http api 提升使用能力

garten 支持 通过yaml 或者 json 的形式 导入一些简单的 命令处理，用来简化编排用例的编写

garten 支持顺序cluster和并行cluster两种集群，默认为并行集群，集群可以嵌套，一般是需要操作的一些机器或者机器的集合

garten 集群返回值 以原有的格式返回


