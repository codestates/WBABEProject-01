# WBABEProject-01

# 프로젝트 개요

언택트 시대에 급증하고 있는 온라인 주문 시스템은 이미 생활전반에 그 영향을 끼치고 있는 상황에, 가깝게는 배달어플, 매장에는 키오스크, 식당에는 패드를 이용한 메뉴 주문까지 그 사용범위가 점점 확대되어 가고 있습니다. 이런 시대에 해당 시스템을 이해, 경험하고 각 단계별 프로세스를 이해하여 구현함으로써 서비스 구축에 경험을 쌓고, golang의 이해를 돕습니다.

# 프로젝트 목표
1. 학습자는 주문자/피주문자의 역할에서 필수적인 기능을 도출, 구현합니다.
2. 학습자는 해당 시스템에 대해 요구사항을 접수하고 주문자와 피주문자 입장에서 필요한 기능을 도출하여, 기능을 확장하고 주문 서비스를 원할하게 지원할수 있는 기능을 구현합니다.
3. 주문자는 신뢰있는 주문과 배달까지를 원합니다. 또, 피주문자는 주문내역을 관리하고 원할한 서비스가 제공되어야 합니다.

## 프로젝트 설치 및 실행

```
git clone git@github.com:codestates/WBABEProject-01.git Project
cd Project
go mod tidy

and

go run main.go
```

## API 명세서

```
  메뉴
  
 /menu/insertMenu -> 메뉴 추가
 /menu/deleteMenu -> 메뉴 삭제 ( inVisible 플래그 활성화)
 /menu/getMenu    -> 메뉴 조회
 /menu/updateMenu -> 메뉴 업데이트
 /menu/sortMenu   -> 메뉴 평점 순 정렬
 
  주문
 
 /order/insertOrder     -> 주문 하기
 /order/addOrderMenu    -> 주문 메뉴 추가
 /order/getOrderByUser  -> 유저별 주문 내역 조회
 /order/updateOrderState  ->주문 상태 변경   ( ex, 요리 중 -> 배달 중 )
 
 
  리뷰 및 평점
 /review/getReview    -> 메뉴별 리뷰 조회
 /review/writeReview  -> 메뉴별 리뷰 작성
```

### API 테스트

/menu/insert  -> 메뉴 추가

```
Give an example
```

