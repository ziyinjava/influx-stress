# influx-stress
influxdb 压力测试工具

### Research
由于 influxdb 官方提供的压力测试工具不能满足写入速率测试的要求，故开发此工具。

### Install
本项目需要 go module 支持
> git clone https://github.com/bemyth/influx-stress.git  
> cd cmd/influx-stress  
> go build   
> go install

### Start

> influx-stress --config=${path}  

### 说明

* 配置项的并发数量，为每个表的并发数量，每个表的数据独立发送并计算速率  
* 本工具的输出为平均写入速率
* 可通过配置 并发量 和 批大小，进一步测试influxdb服务性能
