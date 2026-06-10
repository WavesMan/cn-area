# cn-area

中国行政区划数据包 — TypeScript / Python / Go 三语言统一 API

基于 2,852 条省/市/区三级行政区划数据，开箱即用。

## 安装

```bash
# TypeScript
npm install cn-area

# Python
pip install cn-area

# Go
go get github.com/WavesMan/cn-area/go
```

## 使用

### TypeScript

```ts
import { provinces, cities, districts, lookup, flatten, search } from 'cn-area'

provinces()        // 34 个省级行政区
cities('15')       // 内蒙古各地级市
districts('1501')  // 呼和浩特各区县
lookup('110101')   // { provinceName: '北京市', districtName: '东城区', ... }
search('东城区')    // 按名称反查，精确匹配
search('朝阳')     // 模糊匹配，返回所有包含"朝阳"的记录
flatten()          // 2,852 条扁平记录
```

### Python

```python
from cn_area import provinces, cities, districts, lookup, flatten, search

provinces()        # list[Area]
cities('15')       # list[Area]
districts('1501')  # list[Area]
lookup('110101')   # AreaRecord
search('东城区')    # 按名称反查，精确匹配
search('朝阳')     # 模糊匹配
flatten()          # list[AreaRecord] (2,852 条)
```

### Go

```go
import cnarea "github.com/WavesMan/cn-area/go"

cnarea.Provinces()       // []Province
cnarea.Cities("15")      // []City
cnarea.Districts("1501") // []District
cnarea.Lookup("110101")  // (*Record, bool)
cnarea.Search("东城区")   // 按名称反查，精确匹配
cnarea.Search("朝阳")    // 模糊匹配
cnarea.Flatten()         // []Record (2,852 条)
```

## API

| 函数 | 说明 |
|------|------|
| `provinces()` | 获取全部 34 个省级行政区 |
| `cities(provinceCode)` | 按省查询地级市（直辖市返回空） |
| `districts(cityCode)` | 按市查询区县 |
| `lookup(code)` | 按行政区划码精确定位 |
| `search(name)` | 按地区名称反查（精确优先，模糊兜底） |
| `flatten()` | 返回全部 2,852 条扁平记录 |

## 数据说明

- 数据来源：国家统计局行政区划代码
- 覆盖范围：34 省级 / 333 地级 / 2,845 区县级
- 特殊结构：直辖市（北京/上海/天津/重庆）跳过地级层，直接到区县

## 开发

```bash
# 生成代码
make generate

# 验证生成文件一致性
make check

# 运行全部测试
make test

# 构建全部包
make build
```

## License

MIT
