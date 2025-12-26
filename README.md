# 가상계좌 입금 모니터링 시스템

가상계좌 입금을 실시간으로 모니터링하고 관리하는 웹 애플리케이션입니다.

## 🚀 기능

- **실시간 입금 모니터링**: 가상계좌 입금 내역 실시간 확인
- **인증 시스템**: 계좌번호와 인증키를 통한 보안 로그인
- **자동 새로고침**: 30초마다 자동으로 데이터 업데이트
- **모던 UI**: 반응형 디자인과 직관적인 사용자 인터페이스
- **INI 파일 저장**: 로컬 개발용 간편한 데이터 저장

## 🛠️ 기술 스택

### Backend
- **Go**: HTTP 서버 및 API
- **INI 파일**: 데이터 저장 (로컬 개발용)
- **CORS 지원**: 프론트엔드 연동

### Frontend
- **Svelte**: 반응형 웹 프레임워크
- **Vite**: 빠른 개발 서버
- **CSS3**: 모던 스타일링

## 📦 설치 및 실행

### 1. 저장소 클론
```bash
git clone https://github.com/Song-wh/account_monitor.git
cd account_monitor
```

### 2. 백엔드 실행
```bash
cd backend
go mod tidy
go run .
```

### 3. 프론트엔드 실행
```bash
cd frontend
npm install
npm run dev
```

## 🌐 접속 정보

- **프론트엔드**: http://localhost:7000
- **백엔드 API**: http://localhost:7001

## 🔑 테스트 계정

| 계좌번호 | 인증키 |
|---------|--------|
| 123-456-7890 | key123 |
| 987-654-3210 | key456 |
| 555-123-4567 | key789 |

## 📡 API 엔드포인트

### 인증
- `POST /api/auth` - 계좌 인증

### 입금 관리
- `POST /webhook/deposit` - 입금 웹훅 수신
- `GET /api/deposits?virtual_account={계좌번호}` - 입금 내역 조회

## 🧪 테스트

### 웹훅 테스트
```bash
curl -X POST http://localhost:7001/webhook/deposit \
  -H "Content-Type: application/json" \
  -d '{
    "virtual_account": "123-456-7890",
    "remitter_name": "테스트사용자",
    "remitter_account": "110-999-888777",
    "amount": 50000
  }'
```

### API 테스트
```bash
# 인증 테스트
curl -X POST http://localhost:7001/api/auth \
  -H "Content-Type: application/json" \
  -d '{"virtual_account":"123-456-7890","auth_key":"key123"}'

# 입금 내역 조회
curl "http://localhost:7001/api/deposits?virtual_account=123-456-7890"

# 입금 내역 조회 -- console
fetch('http://localhost:7001/webhook/deposit', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    virtual_account: "123-456-7890",
    remitter_name: "브라우저테스트2",
    remitter_account: "110-777-666555",
    amount: 25000
  })
})
.then(response => response.text())
.then(data => console.log(data));
```

## 📁 프로젝트 구조

```
account_monitor/
├── backend/
│   ├── main.go          # 메인 서버
│   ├── handlers.go      # API 핸들러
│   ├── db.go           # 데이터 관리
│   ├── config.ini      # 샘플 데이터
│   └── README.md       # 백엔드 문서
├── frontend/
│   ├── src/
│   │   ├── App.svelte      # 메인 앱
│   │   ├── components/
│   │   │   └── LoginForm.svelte  # 로그인 폼
│   │   └── routes/
│   │       └── Deposit.svelte    # 입금 내역
│   ├── package.json
│   └── vite.config.js
└── README.md
```

## 🔧 개발 환경

- Go 1.19+
- Node.js 16+
- npm 또는 yarn

## 📝 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다.

## 🤝 기여하기

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📞 문의

프로젝트에 대한 문의사항이 있으시면 이슈를 생성해주세요.
