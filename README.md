# instapay-go-client


Go Client for InstaPay

인스타페이 프로토콜을 위한 클라이언트 프로그램

세 가지 쓰레드로 구성됨.
: 이더리움 결제채널 이벤트 받는 쓰레드.
: Grpc 서버 (다른 사용자와 통신할 때 사용)
: 웹 서버 (사용자 정보 확인 결제 채널 정보 및 결제 요청 인터페이스)


//디렉토리 별 설명
config => 이더리움, 계정 정보 관련 셋업
router => 사용자 인터페이스 관련 웹 경로 정보
ex) templates/channels/list
controller => 컨트롤러 로직(웹 관련)
db => 몽고디비 연결(몽고디비는 결제 채널 데이터

