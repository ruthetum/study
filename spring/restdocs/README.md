# REST Docs
- REST Docs는 테스트 코드 기반으로 RESTful API 문서를 돕는 도구
- Asciidoctor를 이용해서 HTML 등 다양한 포맷으로 문서를 자동으로 출력
- REST Docs의 장점이자 다른 문서화 도구의 차이점은 테스트 코드 기반으로 문서를 작성
- API Spec과 문서화를 위한 테스트 코드가 일치하지 않으면 테스트 빌드를 실패하게 되어 테스트 코드로 검증된 문서를 보장할 수 있음

## Swagger vs REST Docs
### Swagger
```java
@ApiOperation(
        value = "members 조회",
        response = "Member.class",
        responseContainer = "List"
)
@ApiResponses(
        value = {
                @ApiResponse(code = 200, message = "응답 성공")
                @ApiResponse(code = 400, message = "응답 실패")
        }
)
@GetMapping
public List<Member> getMembers(
        @ApiParam(name = "page", value = "page", defaultValue = "0")
        @RequestParam Integer page;
        @ApiParam(name = "size", value = "size", defaultValue = "10")
        @RequestParam Integer size;
) {
        final List<Member> members = new ArrayList<>();
        members.add(new Member("heedong123@naver.com", "heedong"));
        return members;
}
```

### REST Docs
```java
@GetMapping
public PageResponse<MemberResponse> getMembers(
        @PageableDefault(sort = "id", direction = Direction.DESC) Pageable pageable        
) {
        return new PageResponse(memberRepository, findAll(pageable).map(MemeberResponse::new));
}
```
