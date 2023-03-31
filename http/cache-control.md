# Cache-Control
## Prologue
- [RIDI 개인정보 유출 사고](https://help.ridibooks.com/hc/ko/articles/15257092388883--%EC%82%AC%EA%B3%BC%EB%AC%B8-%EA%B0%9C%EC%9D%B8%EC%A0%95%EB%B3%B4-%EC%9C%A0%EC%B6%9C-%EC%82%AC%EA%B3%BC) 관련해서 CDN 서버 캐시 설정 오류 문제 발생
- CDN 설정의 경우 여러 서비스 및 부서에서 활용하다보니 의도치 않게 오설정될 수 있는 가능성 존재
- 결국 내부적으로 이런 오설정 문제를 방지할 필요가 있고, 이에 대한으로 API 서버에 `Cache-Control: private` 헤더를 추가하여 문제를 방지할 수 있음

## Cache-Control Header
> The `Cache-Control` HTTP header field holds **directives** (instructions) — in both **requests and responses** — that control **caching in browsers and shared caches** (e.g. Proxies, CDNs).
- Cache-Control은 브라우저 및 공유 캐시에 캐싱 동작을 지시하는 HTTP 헤더

#### Browser Cache
- 브라우저 캐싱은 웹 브라우저에서 웹 사이트 리소스를 저장하여 서버에서 다시 가져오는 것을 방지
- 예를 들어 웹 사이트의 배경 이미지를 캐시에 로컬로 저장하면 사용자가 해당 페이지를 두 번째로 방문할 때 이미지가 사용자의 로컬 파일에서 로드되므로 페이지가 훨씬 빠르게 로드

#### Shared cache
- 서버와 클라이언트 사이에 존재하는 캐시 (e.g. Proxy, CDN)
- 단일 응답을 저장한 후 여러 사용자에게 재사용 가능 (It stores a single response and reuses it with multiple users)
  - so developers should avoid storing personalized contents to be cached in the shared cache.

## Cache directives
|     Request     |        Response        |
|:---------------:|:----------------------:|
|     max-age     |        max-age         |
|    max-stale    |           -            |
|    min-fresh    |           	-           |
|        -        |        s-maxage        |
|    **no-cache**     |       	**no-cache**        |
|    **no-store**	    |        **no-store**        |
|  no-transform   |     	no-transform      |
| only-if-cached  |           	-           |
|        -        |    must-revalidate     |
|        -        |    proxy-revalidate    |
|        -        |    must-understand     |
|        -        |        **private**         |
|        -        |         public         |
|        -        |       immutable        |
|        -        | stale-while-revalidate |
| stale-if-error	 |     stale-if-error     |

## no-store / no-cache / private
### no-store
> The `no-store` response directive indicates that any caches of any kind (private or shared) should not store this response.
```
Cache-Control: no-store
```
- 아무것도 캐싱하지 않음

### no-cache
> The `no-cache` response directive indicates that the response can be stored in caches, but the response must be validated with the origin server before each reuse, even when the cache is disconnected from the origin server.
```
Cache-Control: no-cache
```
- 응답을 캐시에 저장할 수는 있지만, 오리전 서버와 연결이 끊어졌다면 재사용 전에 오리진 서버를 통해 확인(유효성 검사)를 진행해야 함

#### must-revalidate
> The `must-revalidate` response directive indicates that the response can be stored in caches and can be reused while fresh. If the response becomes stale, it must be validated with the origin server before reuse.
- 만료된 캐시만 서버에 확인하고, 만료되지 않은 캐시는 캐시에서 바로 사용 (no-cache 로직이 다름)

### private
> The `private` response directive indicates that the response can be stored only in a private cache (e.g. local caches in browsers).
```
Cache-Control: private
```
- 브라우저와 같은 특정 사용자 환경에 한해 저장 가능

#### public
> The `public` response directive indicates that the response can be stored in a shared cache. Responses for requests with Authorization header fields must not be stored in a shared cache; however, the public directive will cause such responses to be stored in a shared cache.
- 공유 캐시(Proxy or CDN)에 저장 가능

## cf-cache-status
- cf-cache-status는 Cloudflare에서 제공하는 헤더로, 캐시 상태를 확인할 수 있음
- `cf-cache-status: DYNAMIC`: 기본적으로 캐시를 사용하지 않음

## Reference
- [MDN - Cache-Control](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control)
- [Cloudflare - What is cache-control](https://www.cloudflare.com/ko-kr/learning/cdn/glossary/what-is-cache-control)
- [Toss - 웹 서비스 캐시 똑똑하게 다루기](https://toss.tech/article/smart-web-service-cache)
- [Effective HTTP Caching Part III: Public, Private and No-Store](https://software-factotum.medium.com/effective-http-caching-part-iii-public-private-and-no-store-b64f0452325)
- [알아둬야 할 HTTP 쿠키 & 캐시 헤더](https://www.zerocho.com/category/HTTP/post/5b594dd3c06fa2001b89feb9)
- [Stackoverflow - What is Cache-Control: private](https://stackoverflow.com/questions/12908766/what-is-cache-control-private)
- [w3 - Cache-Control](https://www.w3.org/Protocols/HTTP/Issues/cache-private.html)
- [imperva - Cache-Control](https://www.imperva.com/learn/performance/cache-control/)
- [Cloudflare - cf-cache-status](https://developers.cloudflare.com/cache/about/default-cache-behavior/)