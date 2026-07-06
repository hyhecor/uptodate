# uptodate

`uptodate` is a CLI tool that checks whether a binary is up to date
based on changes in source files.

It helps reduce unnecessary builds in Docker, CI, and Makefile workflows.

- 🇰🇷 한국어: [README.ko.md](README.ko.md)  
- 🇺🇸 English: [README.md](README.md)

## TL;DR

```bash
uptodate . -b ./bin/app -e go,mod,sum || echo "build command here"
```

## Install

```bash
go install github.com/hyhecor/uptodate
```

## 📖 Usage

### Basic

```bash
uptodate . -b ./bin/foo
```

### Extension filter (recommended)

```bash
uptodate . -b ./bin/foo -e go
```

```bash
uptodate . -b ./bin/foo -e go,mod,sum
```

## 🔍 How it works

1. Walk source-root
1. Filter files by extension
1. Compare modification time with binary
1. Decide rebuild necessity

## Exit Codes

| Code | Meaning |
|------|--------|
| 0    | up to date |
| 1    | rebuild required |
| 2    | error |

## 🧪 Examples

### Docker

```bash
uptodate . -b ./bin/app -e go || docker build -t app .
```

### CI

```bash
uptodate . -b ./bin/app -e go,mod,sum || go build
```

## Build-time Configuration

`uptodate` can be customized at install/build time using build flags.

This allows you to define default behaviors without changing runtime commands.

### Example (Go build)

```bash
go build -ldflags "-X main.exts=go,mod,sum" -o uptodate
```

### Example (Go install)

```
go install -ldflags "-X main.exts=go,mod,sum" github.com/hyhecor/uptodate@latest
```

### Available Build Flags

- exts
  - Default file extensions used for filtering
  - Comma-separated values (e.g. go,mod,sum)

### Behavior

- Build flags define default values at compile time
- Runtime flags (--ext, --all) can override build-time configuration
- If neither build-time nor runtime --ext is set, all files are included
