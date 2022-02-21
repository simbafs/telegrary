# Telegrary
Telegrary = Telegram + diary

Telegrary 是一個 Telegram 機器人，讓你可以在 Telegram 上管理日記。同時 Telegrary 提供了一個 CLI 界面，讓你可以在終端機管理日記。

# CLI Usage
## Telegram Bot（毛坯屋，堪用但不好用）
把你的 Telegram bot token 放在 `~/.config/telegrary.toml` 或 `./telegrary.toml`，像下面這樣：

```toml
token = "dafjskdsajflkdsajflkdsjflkjdsalkf"
```

## Bot Commands
| Command                                   | Description     |
| :--------------                           | :-------------- |
| `/help`                                   | print help text |
| `/read [[[year], month], day]`            | read diary      |
| `/write [[[year], month], day] <content>` | write diary     |


然後執行命令 `telegrary bot`

## 終端機寫日記
命令 `telegrary [[[year] month] day]` 會用你喜歡的編輯器 (`$EDITOR`) 打開日記，內容會存在目錄 `~/.local/share/telegrary` 下面。你可以在 `telegrary.toml` 中加入 `root = path/to/directory` 改變預設目錄。
