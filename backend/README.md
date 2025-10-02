# 가상계좌 입금 모니터링 시스템

가상계좌 입금을 실시간으로 모니터링하고 관리하는 Go 기반 웹훅 서버입니다.

## 기능

- 가상계좌 입금 웹훅 수신
- 특정 가상계좌의 입금 내역 조회
- INI 파일 기반 데이터 저장 (로컬 개발용)
- RESTful API 엔드포인트

## 사전 요구사항

- Go 1.19 이상

## 설정

1. 의존성 설치:
   ```bash
   go mod tidy
   ```

2. 애플리케이션 실행:
   ```bash
   go run .
   ```

서버는 기본적으로 7001 포트에서 시작됩니다.

## 데이터 저장

로컬 개발 환경에서는 MySQL 대신 INI 파일(`config.ini`)을 사용하여 데이터를 저장합니다.

### INI 파일 구조
```ini
next_id = 1

[deposit_1]
virtual_account_no = 1234567890
remitter_name = 홍길동
amount = 10000.00
pg_source = TOSPAYMENTS
payload = {"virtual_account":"1234567890","remitter_name":"홍길동","amount":10000}
created_at = 2024-01-15 10:30:00
```

## API 엔드포인트

- `POST /webhook/deposit` - 가상계좌 입금 웹훅 수신
- `GET /api/deposits?virtual_account={가상계좌번호}` - 특정 가상계좌 입금 내역 조회

## 웹훅 페이로드 형식

입금 웹훅은 다음 형식의 JSON을 받습니다:
```json
{
  "virtual_account": "가상계좌번호",
  "remitter_name": "입금자명",
  "remitter_account": "송금자 계좌번호",
  "amount": 입금금액
}
```

## 응답 형식

입금 내역 조회 응답:
```json
{
  "data": [
    {
      "remitter_name": "입금자명",
      "remitter_account": "송금자 계좌번호",
      "amount": "입금금액",
      "created_at": "입금시간"
    }
  ]
}
```

## 환경 변수

현재 코드에서는 하드코딩된 포트(7001)를 사용합니다. 필요시 환경 변수로 변경 가능합니다.

## 프로덕션 환경

프로덕션 환경에서는 MySQL 데이터베이스를 사용하도록 `db.go` 파일을 수정하세요.
