# MSA 환경에서 REST DOCS 적용하기

## Question
- 원활한 유지보수, 클라이언트와의 협업을 위해 API 문서화가 필요
- Monolithic Application에서는 하나의 서비스, 하나의 포트니까 RestDocs를 적용하면 되지만 MSA 환경에서는 여러 서비스, 랜덤 포트인데 어떻게 이 문제를 해결할 것인가?

## Answer
1. RestDocs를 이용해 테스트 문서 만들기
2. openAPI Spec 추출하기
3. 도커를 이용해 추출한 파일을 합쳐서 Swagger UI로 보내기
4. 빌드 배포까지 자동화

### 1. RestDocs 적용
- 일단 기존에 RestDocs를 적용하는 것처럼 테스트 코드를 바탕으로 테스트 문서 생성

### 2. openAPI Spec 추출
```
테스트에서 import 수정
MockMvcRestDocumentation
---->>>
MockMvcRestDocumentationWrapper;
바꿔준다.
```

#### cf.
- https://github.com/ePages-de/restdocs-api-spec

### 3. 도커를 이용해 추출한 파일을 합쳐서 Swagger UI로 보내기
- 도커 실행하고 swagger ui 이미지를 받는다.

  `docker pull swaggerapi/swagger-ui`

- 합쳐진 json 파일을 모두 폴더로 이동

     ```
     docker run -d -p 80:8080 \
     -e URLS_PRIMARY_NAME=Swagger \
     -e URLS="[ \
     { url: 'docs/api.json', name: 'Swagger' } \2
     , { url: 'docs/auction.json', name: ‘Auction’ } \
     ]" \
     -v Josn 파일 모은 폴더:/usr/share/nginx/html/docs/ \
     swaggerapi/swagger-ui
     ```

### Reference
- https://taetaetae.github.io/posts/a-combination-of-swagger-and-spring-restdocs/
- https://youtu.be/qguXHW0s8RY
- https://velog.io/@jifrozen/Dining-together-MSA-%ED%99%98%EA%B2%BD%EC%97%90%EC%84%9C-API-%EB%AC%B8%EC%84%9C-%ED%86%B5%ED%95%A9-%EA%B4%80%EB%A6%AC-Spring-restDocs%EC%99%80-Swagger-UI-%EC%A1%B0%ED%95%A9-blx07ot7
- https://blog.naver.com/qjawnswkd/222340413113