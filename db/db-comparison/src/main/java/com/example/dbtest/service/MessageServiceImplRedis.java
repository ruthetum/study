package com.example.dbtest.service;

import com.example.dbtest.domain.Message;
import com.example.dbtest.dto.MessageDto;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.ValueOperations;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
@RequiredArgsConstructor
public class MessageServiceImplRedis implements MessageService {

    private final RedisTemplate redisTemplate;

    @Override
    public void save(MessageDto messageDto) {
        String id = String.valueOf(LocalDateTime.now().getSecond()) + messageDto.getId();
        Message message = new Message();
        message.setId(Long.parseLong(id));
        message.setMid(messageDto.getId());
        message.setContent(messageDto.getContent());
        ValueOperations<String, String> valueOperations = redisTemplate.opsForValue();
        valueOperations.set(messageDto.getId().toString(), message.toString());
    }

    @Override
    public Message findById(Long id) {
        ValueOperations<String, String> valueOperations = redisTemplate.opsForValue();
        valueOperations.get(id.toString());
        return null;
    }
}
