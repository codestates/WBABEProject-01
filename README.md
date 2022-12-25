# WBABEProject-01

#기술 스택
golang, go version go1.18.1 linux/amd64

# 프로젝트 이름

띵동주문이요, 온라인 주문 시스템(Online Ordering System)

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

해당 메뉴 정보를 가지고 DB에 메뉴 정보를 등록한다

![Untitled](https://user-images.githubusercontent.com/71590785/209467971-549731f6-8bd0-4d1a-baf6-0fb0e800dc43.png)

updateMenu - 피주문자

( 파라미터 : 메뉴이름, 가격, 수량, 원산지, 맵기, 주문가능여부, 메뉴판 표시 여부 ) 

해당 메뉴정보를 가진 메뉴를 DB에서 찾고 메뉴 정보를 변경한다

![Untitled 1](https://user-images.githubusercontent.com/71590785/209468018-2517162f-c051-4dce-8b17-9884413f0022.png)

deleteMenu - 피주문자

( 파라미터 : 메뉴이름 )

해당 메뉴이름을 메뉴DB에서 지운다 ( 뷰 플래그를 활용해 isVisible 값을 false로 변경 )

![Untitled 2](https://user-images.githubusercontent.com/71590785/209468033-7f496c8d-9c9e-49ae-9519-f7729802b962.png)

getMenu - 주문자

(파라미터 : X )

메뉴DB에 들어 있는 모든 메뉴 정보를 조회한다

![Untitled 3](https://user-images.githubusercontent.com/71590785/209468037-f254aa77-9011-4c70-8430-d557301bcfc8.png)

sortMenu - 주문자

( 파라미터 : X )

메뉴 정보에 들어 있는 메뉴들을 평점 순으로 정렬해 조회한다.

![Untitled 4](https://user-images.githubusercontent.com/71590785/209468046-eff735e1-34e2-40bb-a818-a6ff7f753855.png)

Order

insertOrder - 주문자

(파라미터 : 유저 핸드폰 번호, 유저 주소, 주문할 메뉴 이름, 메뉴 수량 )

해당 유저 정보와 메뉴 정보를 가지고 주문DB에 주문을 추가한다

![Untitled 5](https://user-images.githubusercontent.com/71590785/209468049-6be73e8c-2f02-4972-9200-a23c5bf028ac.png)

updateOrderState - 피주문자

( 파라미터 : 주문 시간, 주문 번호, 주문 상태 )

주문 시간과 주문 번호(Unique value)에 해당하는 주문을 찾고 주문 상태를 변경한다

![Untitled 6](https://user-images.githubusercontent.com/71590785/209468057-c379203a-97ec-4d7e-adcd-746aba34cfc3.png)

addOrderMenu - 주문자

( 파라미터 : 유저 핸드폰 번호, 유저 주소, 주문 메뉴 이름, 주문 수량 )

해당 유저의 주문이 아직 배달 직전이라면 주문에 해당 주문 메뉴를 추가하고 배달 이후라면 새로운 주문을 만든다

![Untitled 7](https://user-images.githubusercontent.com/71590785/209468061-ed2b2eff-11db-4abf-8e6f-1ea2b538588f.png)

getOrderByUser - 주문자

( 파라미터 : 유저 주소, 유저 핸드폰 번호 )

해당 유저의 모든 주문 내역을 조회한다

![Untitled 8](https://user-images.githubusercontent.com/71590785/209468072-df2abada-a3f6-4f7a-a4cc-c97e5ebd3274.png)

Review

getReview - 주문자

( 파라미터 : 메뉴 이름 )

해당 메뉴의 모든 리뷰와 평균 평점을 조회한다.

![Untitled 9](https://user-images.githubusercontent.com/71590785/209468078-5e40c915-573b-4302-ad4d-7a6fb3d05307.png)

writeReview

( 파라미터 : 메뉴 이름, 주문 번호 및 시간, 리뷰 내용, 유저 주소 및 핸드폰 번호 )

해당 유저 주문한 주문들 중 특정 주문의 메뉴 내용에 관해서 평점과 리뷰 내용을 작성한다

![Untitled 10](https://user-images.githubusercontent.com/71590785/209468085-59f737d4-2fea-4ab9-92d3-6684577085c1.png)

# DB 설계

## DB구조

### DB는 tMenu, tOrder, tReview 테이블로 나눠서 구성했다
![db4](https://user-images.githubusercontent.com/71590785/209468436-d45a535d-9f91-4185-afd8-0f7201f23a44.PNG)

tMenu
![db1](https://user-images.githubusercontent.com/71590785/209468202-ee281e8c-92db-4b70-9e34-4de45b020956.PNG)

tMenu 테이블은 메뉴의 정보가 들어가 있는 테이블로 메뉴 이름, 가격, 수량, 맵기 정도, 원산지, 주문 가능여부, 메뉴판 보임 여부 필드가 들어가 있다

tOrder
![db2](https://user-images.githubusercontent.com/71590785/209468275-f7c0b415-1302-4431-bec2-b83e25ff22cb.PNG)

tOrder 테이블에는 주문의 정보가 들어가 있는 테이블로 메뉴정보, 주문 시간 및 주문 번호, 주문 상태, 유저 정보 필드가 들어가 있다. 유저 정보는 핸드폰 번호, 주소 구조체로 작성하고 있다

tReview
![db3](https://user-images.githubusercontent.com/71590785/209468364-b7ec65b6-434e-4389-a738-edac00e1cf5f.PNG)

tReview 테이블에는 리뷰 정보가 들어가 있는 테이블로 유저 정보 및 주문을 구분하기 위한 주문 시간 및 번호, 메뉴 이름, 평점 과 리뷰 필드가 들어가 있다 
