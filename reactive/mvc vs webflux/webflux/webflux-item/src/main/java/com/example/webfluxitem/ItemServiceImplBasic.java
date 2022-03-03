package com.example.webfluxitem;

import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.ValueOperations;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Objects;

@Service
@RequiredArgsConstructor
public class ItemServiceImplBasic {

    private final String ITEM_KEY = "item";
    private final String DEFAULT_QUANTITY_STR = "1000000";
    private final RedisTemplate redisTemplate;

    @Transactional
    public int order(int quantity) {
        ValueOperations<String, String> valueOperations = redisTemplate.opsForValue();
        String remain = valueOperations.get(ITEM_KEY);
        if (Objects.isNull(remain))
            remain = DEFAULT_QUANTITY_STR;
        int remainNum = Integer.parseInt(remain) - quantity;
        valueOperations.set(ITEM_KEY, String.valueOf(remainNum));
        return remainNum;
    }

    public int remain() {
        ValueOperations<String, String> valueOperations = redisTemplate.opsForValue();
        String remain = valueOperations.get(ITEM_KEY);
        return Integer.parseInt(remain);
    }
}
