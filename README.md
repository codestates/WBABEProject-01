# WBABEProject-01

# 기술 스택
golang, go version go1.18.1 linux/amd64

DB : mongoDB

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

### Swagger

![image](https://user-images.githubusercontent.com/71590785/209706178-2c72cea1-0124-430c-9c03-23e8854ca905.png)


### API 테스트

insertMenu - 피주문자
( 파라미터 : 매뉴이름, 가격, 수량, 원산지, 맵기, 주문가능여부, 며뉴판 표시 여부 )

![image](https://user-images.githubusercontent.com/71590785/209692027-e091aa03-da4a-464e-9e4b-74026a35302f.png)

해당 메뉴 정보를 가지고 DB에 메뉴 정보를 등록한다

결과

![image](https://user-images.githubusercontent.com/71590785/209691834-07e23d14-52c8-415b-b913-0652ab290c65.png)


updateMenu - 피주문자

( 파라미터 : 메뉴이름, 가격, 수량, 원산지, 맵기, 주문가능여부, 메뉴판 표시 여부 ) 

![image](https://user-images.githubusercontent.com/71590785/209692215-bd32680d-5483-4fe6-b883-ab890cf716c9.png)

해당 메뉴정보를 가진 메뉴를 DB에서 찾고 메뉴 정보를 변경한다

결과
  
![image](https://user-images.githubusercontent.com/71590785/209692264-d81b897f-2bed-486c-9544-6cd92acb060c.png)


deleteMenu - 피주문자

( 파라미터 : 메뉴이름 )

해당 메뉴이름을 메뉴DB에서 지운다 ( 뷰 플래그를 활용해 isVisible 값을 false로 변경 )

![image](https://user-images.githubusercontent.com/71590785/209692501-7e5144f4-ab81-41e1-807c-9be4e6f9d647.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209692605-772b8b4b-f9a4-46c5-84b8-0504d814af79.png)


getMenu - 주문자

(파라미터 : X )

메뉴DB에 들어 있는 모든 메뉴 정보를 조회한다

![image](https://user-images.githubusercontent.com/71590785/209692696-a8c85d9f-7b98-4803-9ad7-bdb5403a8532.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209692777-6bce7615-569c-4a93-bcd1-03c14b392948.png)


sortMenu - 주문자

( 파라미터 : X )

메뉴 정보에 들어 있는 메뉴들을 평점 순으로 정렬해 조회한다.

![image](https://user-images.githubusercontent.com/71590785/209692930-cda8da5f-d9f1-4d1d-91da-6e3a25cdc4a0.png)

결과 ( 만약 리뷰 와 평점이 없는 경우 0으로 가정한다 )

![image](https://user-images.githubusercontent.com/71590785/209694788-98f41551-c060-4e09-b13c-d66c9bda9e2d.png)


Order

insertOrder - 주문자

(파라미터 : 유저 핸드폰 번호, 유저 주소, 주문할 메뉴 이름, 메뉴 수량 )

해당 유저 정보와 메뉴 정보를 가지고 주문DB에 주문을 추가한다

![image](https://user-images.githubusercontent.com/71590785/209695037-ed25948b-62d4-457f-aacc-d8b614e664f9.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209695185-4a996a37-3841-40d0-a976-b6677b24277b.png)


updateOrderState - 피주문자

( 파라미터 : 주문 시간, 주문 번호, 주문 상태 )

주문 시간과 주문 번호(Unique value)에 해당하는 주문을 찾고 주문 상태를 변경한다 ( 0: 주문 취소, 1: 접수 중, 2: 요리 중, 3: 배달 중, 4: 배달완료 )

![image](https://user-images.githubusercontent.com/71590785/209695292-96e46c43-dd7d-461b-9a13-dc5b281cea7d.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209695395-32af87e1-9208-4bde-aaaa-b9998ac3855f.png)


addOrderMenu - 주문자

( 파라미터 : 유저 핸드폰 번호, 유저 주소, 주문 메뉴 이름, 주문 수량 )

해당 유저의 주문이 아직 배달 직전이라면 주문에 해당 주문 메뉴를 추가하고 배달 이후라면 새로운 주문을 만든다

![image](https://user-images.githubusercontent.com/71590785/209703158-3e77d261-846a-4075-9e2f-707bc9fe977f.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209703366-cfa48cea-f58b-4bfe-bc5a-617cc7c75b2e.png)
( 기존 주문이 배달 중(state 3)이므로 새로운 주문으로 추가된 모습 )

  ![image](https://user-images.githubusercontent.com/71590785/209703442-7b87d79b-fe9c-4093-b5ce-be5b7238670e.png)
  ![image](https://user-images.githubusercontent.com/71590785/209703512-583176d3-c056-433d-8ae5-b17bbd6a6340.png)
  기존 주문이 아직 1 (접수 중) 상태이므로 기존 주문에 새로운 메뉴를 추가한다


getOrderByUser - 주문자

( 파라미터 : 유저 주소, 유저 핸드폰 번호 )

해당 유저의 모든 주문 내역을 조회한다

![image](https://user-images.githubusercontent.com/71590785/209703673-1817266b-49ec-421f-a1b4-1f3123b3fb5e.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209705279-a524308a-bbac-40bd-af64-fe3a504d3310.png)


Review

writeReview - 주문자

( 파라미터 : 메뉴 이름, 주문 번호 및 시간, 리뷰 내용, 유저 주소 및 핸드폰 번호 )

해당 유저 주문한 주문들 중 특정 주문의 메뉴 내용에 관해서 평점과 리뷰 내용을 작성한다

![image](https://user-images.githubusercontent.com/71590785/209705457-ead812ab-928b-450d-9241-79ff90131a98.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209705508-d88bff50-e2e6-4f8a-bec0-6a923f34bfa5.png)


getReview - 주문자

( 파라미터 : 메뉴 이름 )

해당 메뉴의 모든 리뷰와 평균 평점을 조회한다.

![image](https://user-images.githubusercontent.com/71590785/209705637-f7fb8a0d-46ef-4297-8a58-59c9945f76f1.png)

결과

![image](https://user-images.githubusercontent.com/71590785/209705839-52eccbd8-643f-4f25-bba7-69f77d1fee20.png)



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
