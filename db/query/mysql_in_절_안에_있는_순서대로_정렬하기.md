# MySQL IN 절 안에 있는 순서대로 정렬하기

## 상황
- MySQL에서 IN 절 안에 조회하고 싶은 컬럼의 리스트를 넣은 후 조건으로 설정한 리스트의 순서대로 반환받고 싶을 때
- ex.
> - 조회해야 하는 Member들의 id와 반환되어야 할 순서로 정렬된 리스트를 입력받는다.
> - Member를 조회하고, 입력받은 순서에 맞게 정렬해서 Member List를 리턴한다.

```java
@Service
@RequiredArgsConstructor
public class MemberService {

    private final MemberRepository memberRepository;

    // Example 1
    public List<Member> getMembers1(List<Long> memberIds) {

        List<Member> members = memberRepository.findAllById(memberIds); // ①

        List<Member> sortedMembers = new ArrayList<>(); 
        for (Long memeberId: memberIds) {
            Optional<Member> optionalMember = members.stream()  // ②
                    .filter(m -> m.getId().equals(id))
                    .findAny();

            optionalMember.ifPresent(sortedMembers::add);       // ③
        }

        return sortedMembers;
    }

    // Example 2
    public List<Member> getMembers2(List<Long> memberIds) {
        
        List<Member> members = memberRepository.findAllById(memberIds); // ①

        Map<Long, Member> memberMap = new HashMap<>();
        members.forEach(member -> memberMap.put(member.getId(), user)); // ②

        List<Member> sortedMembers = new ArrayList<>();
        for (Long memeberId: memberIds) {
            Member member = memberMap.get(memeberId);                   // ③

            if (member != null)
                sortedMembers.add(member);
        }

        return sortedMembers;
    }
}
```

```java
public interface JpaRepository<T, ID> extends PagingAndSortingRepository<T, ID>, QueryByExampleExecutor<T> {

    ...

    @Override
    List<T> findAllById(Iterable<ID> ids);

    ...
}
```

- Example 1
    - ① : id list를 조회
    - ② : 조회한 멤버 리스트에서 알맞는 id의 회원을 조회
    - ③ : 해당 회원을 리스트에 추가

- Example 2
    - ① : id list를 조회
    - ② : 조회한 멤버 리스트를 Map 형태로 보관
    - ③ : Map에서 아이디를 조회해서 리스트에 추가



## 해결
- 지금의 나는 코드를 작성할 때 `Example1`이나 `Example2`를 생각했을 거고, `Example2`를 활용해서 값을 반환했을 것 같다.
- 하지만 이 방법은 매우 비효율적이다. DB에서 값을 조회하고, 그 다음 다시 반복문을 실행시켜야 한다는 점과 굳이 불필요한 Map 객체를 생성해서 처리해야 한다.
- 따라서 이 부분을 애초에 쿼리를 실행할 때 정렬 순서를 조건으로 설정해서 값을 반환받을 수 있다면 불필요한 반복문 실행이나 객체 생성이 필요없을 수 있다.


### 기존 쿼리
```sql
SELECT * FROM member WHERE id IN (2, 1, 3);
```
- 조회할 아이디 값과 정렬된 리스트가 `2,1,3` 일 때

![image](https://user-images.githubusercontent.com/59307414/162125038-84207543-7236-4850-9a61-9c84cdf7ed9e.png)

### 해결
```sql
SELECT * FROM member WHERE id IN (2, 1, 3) ORDER BY FIELD(id, 2, 1, 3);
```
- `ORDER BY FEILD`를 설정한다.
    - `ORDER BY FELID('컬럼명', '정렬 순서 1', '정렬 순서 2', '정렬 순서 3', ...)`

![image](https://user-images.githubusercontent.com/59307414/162125100-a0b52288-9d19-4e9b-a9c1-ca56e2ba1a9e.png)

```java
public interface MemberRepository extends JpaRepository<Member, Long> {
    @Query(value = "SELECT * FROM member WHERE id IN (?1) ORDER BY FIELD(id, ?1);", nativeQuery = true)
    List<Member> findMembersByIdOrderByFeild(String ids); // 리스트를 문자열로 변환
}