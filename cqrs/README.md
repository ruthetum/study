# CQRS Pattern

## 개요

> CQRS stands for Command and Query Responsibility Segregation, a pattern that **separates read and update operations** for a data store.
> 
> Implementing CQRS in your application can maximize its performance, scalability, and security.
> 
> The flexibility created by migrating to CQRS allows a system to better evolve over time and prevents update commands from causing merge conflicts at the domain level.

데이터를 읽는데 사용하는 모델과 데이터를 업데이트할 수 있는 모델을 분리하여 사용하는 패턴

모델(command, query)을 분리한다는 것은 서로 다른 객체 모델을 의미하기 때문에 서로 다른 논리적 프로세스에서 구동될 수 있다.

#### Single model
![crud](https://martinfowler.com/bliki/images/cqrs/single-model.png)

#### CQRS
![cqrs](https://martinfowler.com/bliki/images/cqrs/cqrs.png)

## 사용 시기
### [CQRS - Martin Fowler](https://martinfowler.com/bliki/CQRS.html)
> For some situations, this separation can be valuable, but beware that for most systems CQRS adds risky complexity.

몇몇 상황에서는 이러한 분리가 유용할 수 있지만, 대부분의 시스템에서 CQRS는 위험한 복잡성을 추가

> In particular CQRS should only be used on specific portions of a system (a BoundedContext in DDD lingo) and not the system as a whole.

CQRS는 시스템 전체가 아닌 시스템의 특정 부분에서만 사용해야 함

> So far I see benefits in two directions. Firstly is that a few complex domains may be easier to tackle by using CQRS. I must stress, however, that such suitability for CQRS is very much the minority case. Usually there's enough overlap between the command and query sides that sharing a model is easier. Using CQRS on a domain that doesn't match it will add complexity, thus reducing productivity and increasing risk.

CQRS 사용 시 두 가지 이점을 생각해볼 수 있는데, 첫 번째는 CQRS를 사용하여 몇 가지 복잡한 도메인을 다루기가 더 쉬울 수 있음 (하지만 그런 경우는 매우 소수)

일반적으로 single model(sharing a model)이 조회(query)와 업데이트(command) 사이에 겹치는 부분(공유할 수 있는 부분)이 많음

그래서 오히려 맞지 않는 상황에 CQRS를 사용하면 복잡성이 추가되어 생산성이 감소하고 위험성이 증가함

> The other main benefit is in handling high performance applications. CQRS allows you to separate the load from reads and writes allowing you to scale each independently. If your application sees a big disparity between reads and writes this is very handy. Even without that, you can apply different optimization strategies to the two sides. An example of this is using different database access techniques for read and update.

고성능 애플리케이션을 만들 때 이점이 있음

읽기와 쓰기를 분리하기 때문에 각각의 로드를 분리하여 확장 가능함

> Despite these benefits, you should be very cautious about using CQRS. Many information systems fit well with the notion of an information base that is updated in the same way that it's read, adding CQRS to such a system can add significant complexity.

그래도 CQRS 사용은 매우 신중하게 선택해라. 생각보다 CQRS 추가로 인해 위험성 증가 및 생산성 하락이 된 경우를 많이 봤다.

## Reference
- https://martinfowler.com/bliki/CQRS.html
- https://learn.microsoft.com/en-us/azure/architecture/patterns/cqrs