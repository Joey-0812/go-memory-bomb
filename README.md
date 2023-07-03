# go-memory-bomb

构建一个golang的内存爆炸demo

使用场景：

istio中对于每次不同的请求都会产生指标，基于构建一个内存爆炸镜像用于指标老化测试。

参考[kyessenov](https://github.com/kyessenov) 的代码，未作修改。

FROM https://github.com/istio/istio/pull/44605#issuecomment-1530416653
