# uptodate

`uptodate` is a CLI tool that checks whether a binary is up to date
based on changes in source files.

It helps reduce unnecessary builds in Docker, CI, and Makefile workflows.

- 🇰🇷 한국어: [README.ko.md](README.ko.md)  

## TL;DR

```bash
uptodate . -b ./bin/app -e go,mod,sum || echo "build command here"
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

