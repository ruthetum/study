package com.example.dbtest.service;

import com.example.dbtest.domain.Message;
import com.example.dbtest.dto.MessageDto;
import lombok.RequiredArgsConstructor;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
@RequiredArgsConstructor
public class MessageServiceImplMongo implements MessageService {

    private final MongoTemplate mongoTemplate;

    @Override
    public void save(MessageDto messageDto) {
        String id = String.valueOf(LocalDateTime.now().getSecond()) + messageDto.getId();
        Message message = new Message();
        message.setId(Long.parseLong(id));
        message.setMid(messageDto.getId());
        message.setContent(messageDto.getContent());
        mongoTemplate.save(message);
    }

    @Override
    public Message findById(Long id) {
        return mongoTemplate.findById(id, Message.class);
    }
}
