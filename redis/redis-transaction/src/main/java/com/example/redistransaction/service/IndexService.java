package com.example.redistransaction.service;

import lombok.RequiredArgsConstructor;
import org.springframework.dao.DataAccessException;
import org.springframework.data.redis.core.RedisOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.SessionCallback;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@RequiredArgsConstructor
@Transactional(readOnly = true)
public class IndexService {

    private final RedisTemplate redisTemplate;

    public void useOperations() {
        try {
            redisTemplate.execute(new SessionCallback<List<Object>>() {
                public List<Object> execute(RedisOperations operations) {
                    // redis transaction 시작
                    operations.multi();

                    // operations.watch("apple"); // watch

                    operations.opsForValue().set("apple", "iphone1");
                    operations.opsForValue().set("samsung", "galaxy1");

                    // operations.discard(); // discard

                    // redis transaction 종료
                    return operations.exec();
                }
            });
        } catch (DataAccessException e) {
            e.printStackTrace();
        }
    }

    @Transactional
    public void useTransactionalAnnotation() {
        redisTemplate.opsForValue().set("apple", "iphone2");
        redisTemplate.opsForValue().set("samsung", "galaxy2");
    }
}
