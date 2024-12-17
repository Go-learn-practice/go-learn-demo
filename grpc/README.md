# Protobuf `message` 消息类型的定义及规则

## 1. 基本类型定义
> Protobuf 提供了一些内建的基本数据类型，用于定义 message 字段

| 类型         | 说明                       | 示例                            |
|------------|--------------------------|-------------------------------|
| `int32`    | 有符号的 32 位整型              | `int32 age = 1;`              |
| `int64`    | 有符号的 64 位整型              | `int64 id = 2;`               |
| `uint32`   | 无符号的 32 位整型              | `uint32 count = 3;`           |
| `uint64`   | 无符号的 64 位整型              | `uint64 timestamp = 4;`       |
| `sint32`   | 有符号的 32 位整型，使用 ZigZag 编码 | `sint32 delta = 5;`           |
| `sint64`   | 有符号的 64 位整型，使用 ZigZag 编码 | `sint64 offset = 6;`          |
| `fixed32`  | 32 位无符号整型，使用小端字节序编码      | `fixed32 crc = 7;`            |
| `fixed64`  | 64 位无符号整型，使用小端字节序编码      | `fixed64 hash = 8;`           |
| `sfixed32` | 32 位有符号整型，使用小端字节序编码      | `sfixed32 delta_fixed = 9;`   |
| `sfixed64` | 64 位有符号整型，使用小端字节序编码      | `sfixed64 offset_fixed = 10;` |
| `float`    | 单精度浮点数                   | `float pi = 11;`              |
| `double`   | 双精度浮点数                   | `double e = 12;`              |
| `bool`     | 布尔类型                     | `bool is_valid = 13;`         |
| `string`   | 字符串类型                    | `string name = 14;`           |
| `bytes`    | 字节数组类型                   | `bytes data = 15;`            |

**示例：基本数据类型定义**
```protobuf
syntax = "proto3";

message Person {
  int32 age = 1;
  int64 id = 2;
  uint32 count = 3;
  uint64 timestamp = 4;
  sint32 delta = 5;
  sint64 offset = 6;
  fixed32 crc = 7;
  fixed64 hash = 8;
}
```

## 2. 枚举类型
> Protobuf 支持枚举类型，用于表示一组固定的常量

**规则：**
- 枚举值必须是整数，并且必须唯一。
- 枚举的第一个值通常是 0，因为未设置的字段会默认赋值为 0

**示例：枚举定义**
```protobuf
syntax = "proto3";

enum Color {
  RED = 0;
  GREEN = 1;
  BLUE = 2;
}

message Person {
  Color favorite_color = 1;
  string name = 2;
}
```

## 3. 嵌套类型
> 可以在一个 message 中嵌套另一个 message，用于组织复杂的数据结构

**示例：嵌套类型**
```protobuf
syntax = "proto3";

message Address {
  string street = 1;
  string city = 2;
  string state = 3;
  string zip = 4;
}

message Person {
  string name = 1;
  int32 age = 2;
  Address address = 3;
}
```
## 4. 数组（Repeated 字段）
> Protobuf 支持数组类型，用于表示一组相同类型的元素

**规则：**
- repeated 字段可以包含 0 个或多个元素。
- 默认情况下，repeated 字段是有序的。

**示例：Repeated 字段**
```protobuf
syntax = "proto3";

message Person {
  repeated string phone_numbers = 1;
  repeated string emails = 2;
}
```

## 5. Map 类型（键值对）
> Protobuf 支持定义 map 类型来表示键值对结构

**规则：**
- map<key_type, value_type> 用于定义映射。
- key_type 只能是基本数据类型：int32、int64、uint32、uint64、bool 和 string。
- value_type 可以是任何合法的类型（包括 message）

**示例：Map 类型**
```protobuf
syntax = "proto3";
message Person {
  map<string, string> metadata = 1;
}
```

## 6. 可选字段（optional）
> 在 proto3 中，optional 关键字可以显式表示字段是否设置过

**规则：**
- 如果字段是 `optional`，可以检测字段是否被赋值（设置）
- 只能用于`基本类型`和 `enum`类型

**示例：Optional 字段**
```protobuf
syntax = "proto3";
message Person {
  optional string name = 1;
  optional int32 age = 2;
}
```

## 7. 字段规则
**字段编号：**

- 字段编号必须是正整数。
- 编号范围为 1 到 2^29-1（1 到 536870911）。
- 1-15 是高效范围，16-2047 用于较大数据。
- 字段编号不能重复。

**字段标识符：**

- 字段名需符合小写加下划线命名规范（snake_case）。

## 8. 保留字段（Reserved）
> 为了避免冲突，Protobuf 允许使用 reserved 关键字保留字段编号或字段名称

**示例：保留字段**
```protobuf
syntax = "proto3";

message Person {
  reserved 3, 15; // 保留字段编号 3 和 15
  reserved "name", "phone_number"; // 保留字段名称 "name" 和 "phone_number"
}
```