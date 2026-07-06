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

## 빌드 시 설정 (Build-time Configuration)

`uptodate`는 빌드 또는 설치 시점에 build flag를 통해 기본 동작을 커스터마이징할 수 있습니다.

이를 통해 실행 시 옵션을 매번 지정하지 않고도 기본 설정을 정의할 수 있습니다.

### 예시 (Go build)

```
go build -ldflags "-X main.exts=go,mod,sum" -o uptodate
```

### 예시 (Go install)

```
go install -ldflags "-X main.exts=go,mod,sum" github.com/hyhecor/uptodate@latest
```

### 사용 가능한 빌드 플래그

- `exts`
  - 파일 변경 감지에 사용할 기본 확장자 목록
  - 콤마(,)로 구분된 값 (예: `go,mod,sum`)

### 동작 방식

- 빌드 시 설정된 값은 기본값으로 사용됩니다
- 런타임 플래그(`--ext`, `--all`)가 있는 경우 이를 우선적으로 적용합니다
- 빌드 시 설정과 런타임 모두 `--ext`가 없으면 모든 파일이 대상이 됩니다
