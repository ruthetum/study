package com.example.webfluxitem;

import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.ReactiveRedisTemplate;
import org.springframework.data.redis.core.ReactiveValueOperations;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import reactor.core.publisher.Mono;

@Service
@RequiredArgsConstructor
public class ItemServiceImplReactive{

    private final String ITEM_KEY = "item";
    private final ReactiveRedisTemplate reactiveRedisTemplate;

    @Transactional
    public Mono<Integer> order(int quantity) {
        ReactiveValueOperations<String, String> valueOperations = reactiveRedisTemplate.opsForValue();
        Mono<String> remain = valueOperations.get(ITEM_KEY);
        return remain.flatMap(v -> {
            int remainNum = Integer.parseInt(v) - quantity;
            valueOperations.set(ITEM_KEY, String.valueOf(remainNum));
            return Mono.just(remainNum);
        });
    }

    public Mono<Integer> remain() {
        ReactiveValueOperations<String, String> valueOperations = reactiveRedisTemplate.opsForValue();
        Mono<String> remain = valueOperations.get(ITEM_KEY);
        return remain.flatMap(v -> Mono.just(Integer.parseInt(v)));
    }
}
