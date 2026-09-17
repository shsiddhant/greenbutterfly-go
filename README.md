# GreenButterfly

>happiness is a green butterfly . . .  
>. . . and she has flown away

A CLI for turning my scattered scripts into commands.

Go implementation using [Cobra](https://cobra.dev/).

## Usage

```shell
$ greenbutterfly-go -h

"happiness is a green butterfly...
...and she has flown away"

A CLI for turning my scattered scripts into commands.

Usage:
  greenbutterfly-go [flags]
  greenbutterfly-go [command]

Available Commands:
  completion Generate the autocompletion script for the specified shell
  days-since View days since for a list of significant dates. (aliases: ds, days_since)
  help      Help about any command

Flags:
  -h, --help      help for greenbutterfly-go
      --version   See version.

Use "greenbutterfly-go [command] --help" for more information about a command.
```

## Examples

### 1. Days Since

```
$ greenbutterfly-go days-since dates.json
 DAYS SINCE ─────────────────────────────
 1005 • The Best Day • December 17, 2023
```

**JSON Schema:**

```json
[
  {
    "title": "The Best Day",
    "date": "2023-12-17"
  },
]
```



