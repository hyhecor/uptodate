# uptodate

`uptodate`는 소스 파일 변경 여부를 확인하여  
대상 바이너리가 다시 빌드되어야 하는지 판단하는 CLI 도구입니다.

Docker, CI, Makefile 환경에서 불필요한 빌드를 줄이기 위해 사용합니다.

- 🇰🇷 한국어: [README.ko.md](README.ko.md)  
- 🇺🇸 English: [README.md](README.md)
 
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
