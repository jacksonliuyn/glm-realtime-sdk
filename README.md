# 智谱 Realtime Python Low Level SDK

## How to use

1. 安装依赖

安装poetry

```bash
pip install poetry
```

安装依赖

```bash
poetry install
```

启动开发环境

```bash
poetry shell
```

2. 设置环境变量

apikey 获取方式：https://www.bigmodel.cn/ 
注册开发者账号, 创建apikey

```bash
export ZHIPU_API_KEY=your_api_key
```

3. 运行示例

```bash

python samples/low_level_sample.py <audio_file_path>

#example: cd samples && python low_level_sample.py input/arc-easy-q237-tts.wav
```
