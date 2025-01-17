# Process Logger 进程监测工具

用于监测进程运行时的资源消耗情况，包括 CPU、内存用于分析程序稳定性。

使用方法：

```bash
plog -p <pid> -o <output> -t 1000
```

参数如下：

```
Usage plog:

  plog -p 745 -o ./data.csv

  -i int
        采样间隔，单位毫秒 (default 1000)
  -o string
        数据存储位置 (default "./data.csv")
  -p int
        进程ID
  -v    版本信息 1.0.0
```


例如：

```bash
./plog -p 5751 -o /tmp/plog.log 
```

```csv
CPU(%), Mem(KB)
0.00, 6072
0.00, 6072
0.00, 6072
0.00, 6072
0.00, 6072
```