# mc-snbt
Minecraft snbt parse

此项目是基于 `https://github.com/wlwanpan/minecraft-wrapper.git`下的snbt模块构建

这是一个基于我的世界snbt格式的数据解析, 同时目前已知支持1.12.x-1.19.x版本的snbt数据格式.

# SNBT

SNBT (Stringified Name Binary Tag) is a file format introduced by Minecraft to save its data. While there are some already exist good packages to decode those files, the server logs prints its stringified counterpart. This lightweight package is meant to decode a given SNBT to a Go struct.

SNBT (stringfied Name Binary Tag)是Minecraft引入的用于保存数据的文件格式。虽然已经有一些好的包可以解码这些文件，但服务器日志会打印其对应的字符串。这个轻量级的包旨在将给定的SNBT解码为Go结构体。

## Basic Usage 基础用法

```go
import snbt "github.com/nageslan/mc-snbt"
```

func1 语法1

```go

bytesToDecode := []byte(`{Base: 1.0d, Name: "minecraft:generic.attack_damage"}`)

bytesStruct := struct {
    Base float64
    Name string
}{}

snbt.Decode(bytesToDecode, &bytesStruct)

fmt.Printf("%+v", bytesStruct) // {Base:1 Name:minecraft:generic.attack_damage}
```

func2 用法2

```go
bytesToDecode := []byte(`{Base: 1.0d, Name: "minecraft:generic.attack_damage"}`)

var m1 = map[string]interface{}{}

snbt.Decode(bytesToDecode, &m1)

fmt.Println(m1) // map[Base:1 Name:minecraft:generic.attack_damage]
```





## Resources 资源

- https://minecraft.gamepedia.com/NBT_format
