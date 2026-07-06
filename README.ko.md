# uptodate

- 🇺🇸 English: [README.en.md](README.md)
 
## TL;DR

```bash
uptodate . -b ./bin/app -e go,mod,sum || echo "build command here"
```

## 설치

```bash
go install github.com/hyhecor/uptodate
```

## 📖 사용법

### 기본 사용

```bash
uptodate . -b ./bin/foo
```

### 확장자 필터 (권장)

```bash
uptodate . -b ./bin/foo -e go
```

```bash
uptodate . -b ./bin/foo -e go,mod,sum
```

## 🔍 동작 방식
1. source-root 아래 파일 탐색
1. 지정된 확장자만 필터링
1. binary보다 최신 파일이 있는지 확인
1. 결과 반환

## Exit Codes

| Code | Meaning |
|------|--------|
| 0    | up to date |
| 1    | rebuild required |
| 2    | error |

## 🧪 예시

### Docker

```bash
uptodate . -b ./bin/app -e go || docker build -t app .
```

### CI

```bash
uptodate . -b ./bin/app -e go,mod,sum || go build
```
