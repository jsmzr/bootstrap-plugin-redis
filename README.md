# bootstrap-plugin-redis

bootstrap 系列 redis 插件，即插即用，在配置文件中配置添加相关配置后即可使用。

## 使用说明

1. 声明需要初始化的插件
2. 初始化插件
3. 获取 client

```go
import (
	"github.com/go-redis/redis/v8"
	"github.com/jsmzr/bootstrap-config/config"
    // 声明配置插件
    _ "github.com/jsmzr/bootstrap-plugin-config-yaml/yaml"
    // 声明 redis 插件
    _ "github.com/jsmzr/bootstrap-plugin-redis/redis"
    "github.com/jsmzr/bootstrap-plugin-redis/connection"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

func main() {
    // 初始化插件
    err := plugin.PostProccess()
	if err != nil {
		fmt.Println(err)
		return
	}
    // 获取 redis client
    conn := connection.GetClient()
}
```

[详见](https://github.com/jsmzr/bootstrap-plugin-example)