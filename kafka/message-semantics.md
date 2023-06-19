# Message Semantics
> ref. https://github.com/ruthetum/study/wiki/Comparison-message-queue#%EB%A9%94%EC%84%B8%EC%A7%80-%EC%A0%84%EB%8B%AC

![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_webp,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F933ba5a8-cf94-4da2-86f2-53e2dc57d5cc_1999x1215.png)

- At most once(최대 한번): 최대 한 번만 전송, 메시지를 한번만 전송하고 상대가 받았는지 받지 못했는지는 확인하지 않음
- At least once(최소 한번): 메시지를 전송하고 최소한 상대방이 하나의 메시지는 받았는지 확인
- Exactly once(정확히 한번): 메시지를 정확히 한번만 전송

## In Kafka
| Project | QoS / Guarantees |
|:---:|:---:|
| **Apache Kafka** | **At least once, Exactly once** |
| RabbitMQ | At most once, At least once | 
| NATS (with JetStream) | At most once, At least once, Exactly once |

### Exactly-one Support in Apache Kafka

_TBD_

## Reference
- https://www.confluent.io/blog/exactly-once-semantics-are-possible-heres-how-apache-kafka-does-it/
- https://medium.com/@jaykreps/exactly-once-support-in-apache-kafka-55e1fdd0a35f
  - ko: https://barunmo.blogspot.com/2017/07/apache-kafka-exactly-once.html
- https://medium.com/naver-cloud-platform/%EC%9D%B4%EB%A0%87%EA%B2%8C-%EA%B0%9C%EB%B0%9C%ED%96%88%EC%8A%B5%EB%8B%88%EB%8B%A4-simple-easy-notification-service-2-%EB%A9%94%EC%8B%9C%EC%A7%80-%EB%B0%9C%EC%86%A1-%EC%B2%98%EB%A6%AC%EB%A5%BC-%EC%9C%84%ED%95%9C-kafka-%EC%82%AC%EC%9A%A9%EA%B8%B0-60beda9d5773
- https://huisam.tistory.com/entry/kafka-message-semantics
  - https://huisam.tistory.com/entry/kafka-producer?category=849126%20
- https://velog.io/@xogml951/Kafka%EC%99%80-Exactly-Once
- https://blog.voidmainvoid.net/438
- https://godekdls.github.io/Apache%20Kafka/producer-configuration/
- https://medium.com/@sdjemails/kafka-producer-delivery-semantics-be863c727d3f
- https://dhkdn9192.github.io/apache-kafka/kakfa-exactly-once-delivery/
- https://4betterme.tistory.com/177
- https://bistros.tistory.com/entry/Kafka-idempotent-producer-%EB%A9%B1%EB%93%B1%EC%84%B1%EC%97%90-%EA%B4%80%ED%95%B4%EC%84%9C