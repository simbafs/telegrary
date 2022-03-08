# Telegrary
Telegrary = Telegram + diary

Telegrary is a diary manage tool with build in telegram bot, which let you edit your diary in telegram(not finish yet)

# Installation
```
go install github.com/simba-fs/telegrary@latest
```

# Dependency
* tree
* git
* go 1.17+ (I beleve 1.16 also work, but I develop it under 1.17)

# CLI Usage
## Telegram Bot(not finish yet)
Put Telegram bot token in config file(introduced below)   
Execute command `telegrary bot`

## Bot Commands
| Command                                   | Description                                                 |
| :---                                      | :---                                                        |
| `/help`                                   | print help text                                             |
| `/read [[[year], month], day]`            | read diary                                                  |
| `/write [[[year], month], day] <content>` | write diary                                                 |
| `/tree`                                   | list all notes in tree form(This depend on `tree` CLI tool) |

> It doesn't support Markdown syntax in `/write` besause I still can't get origin message from Telegram bot. I wish I will solve it in the future

## Write diary in terminal
Executing command `telegrary [[[year] month] day]` will open diary with your editor(`$EDITOR`), which is stored in `~/.local/share/telegrary`, or you can change it with a config file.

## Git
Telegrary use git to version control. After editing, running bot, telegrary will auto execute `git add` and `git commit`. If `git_repo` is set in config file, it will execute `git push`.

## Config
Executing command `telegrary config` to open the nearest config file(file in the config file path list). Below are the supported fields: 

| Field    | Type   | Description                                               |
| :---:    | :---:  | :---                                                      |
| token    | string | Telegram Bot Token, Optional                              |
| root     | string | where notes stored, default = `~/.lcoal/telegrary`        |
| git      | string | path to git exec, default = `git`                         |
| git_sign | bool   | if you want to use gpg sign when commit, default = `true` |
| git_repo | string | path to remote git repository                             |

> config file path list: `~/.config/telegrary.toml`, `./telegrary.toml`
