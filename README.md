[en doc](./README_en.md)

# Telegrary
Telegrary = Telegram + diary

Telegrary 是一個 Telegram 機器人，讓你可以在 Telegram 上管理日記。同時 Telegrary 提供了一個 CLI 界面，讓你可以在終端機管理日記。

# Installation
```
go install github.com/simba-fs/telegrary@latest
```

# Dependency
* tree
* git
* go 1.17+ (I beleve 1.16 also work, but I develop it under 1.17)

# CLI Usage
## Telegram Bot
把你的 Telegram bot token 放在 `~/.config/telegrary.toml` 或 `./telegrary.toml`，並設定 secret(hashed)，像下面這樣：

```toml
token = "dafjskdsajflkdsajflkdsjflkjdsalkf"
secret = "fdshafjdafhjdjasnmalfjsdjkf"
```

然後執行命令 `telegrary bot`

## Bot Commands
| Command                                   | Description                                                 |
| :---                                      | :---                                                        |
| `/help`                                   | print help text                                             |
| `/read [[[year], month], day]`            | read diary                                                  |
| `/write [[[year], month], day] <content>` | write diary                                                 |
| `/tree`                                   | list all notes in tree form(This depend on `tree` CLI tool) |

> 現在 `/write` 還不支援 MD 語法，因為機器人還讀不到原始的文字，MD 語法會被 TG 吃掉，解決中......

## 終端機寫日記
命令 `telegrary [[[year] month] day]` 會用你喜歡的編輯器 (`$EDITOR`) 打開日記，內容會存在目錄 `~/.local/share/telegrary` 下面。你可以在 `telegrary.toml` 中加入 `root = path/to/directory` 改變預設目錄。  

## Git
Telegrary 使用 Git 對日記進行版本管理，在編輯、Bot 命令結束後都會自動執行 `git add`、`git commit`，如果設定檔中有設定 `git_repo` 的話就會執行 `git push`。  

## Config
用命令 `telegrary config` 可以開啟最近的設定檔（在路徑列表中優先序第一且存在的檔案，但是這檔案沒寫的設定可能會由其他設定檔提供）  
以下是有支援的設定

| Field    | Type   | Description                                               |
| :---:    | :---:  | :---                                                      |
| token    | string | Telegram Bot Token, Optional                              |
| root     | string | where notes stored, default = `~/.lcoal/telegrary`        |
| git      | string | path to git exec, default = `git`                         |
| git_sign | bool   | if you want to use gpg sign when commit, default = `true` |
| git_repo | string | path to remote git repository                             |

> 路徑列表：`~/.config/telegrary.toml`、`./telegrary.toml`

## Hash
在 bot 中使用密語驗證使用者身份，為了避免明碼除存密語，因此先經過 hash，你可以用指令 `telegrary hash  <secret>` 來產生 hash 過的密語
