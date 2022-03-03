package com.example.dbtest.service;

import com.example.dbtest.domain.Message;
import com.example.dbtest.dto.MessageDto;
import com.example.dbtest.repository.MessageRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
@RequiredArgsConstructor
public class MessageServiceImplMysql implements MessageService {

    private final MessageRepository messageRepository;

    @Override
    public void save(MessageDto messageDto) {
        String now = LocalDateTime.now().toString();
        Message message = new Message();
        message.setMid(messageDto.getId());
        message.setContent(now);
        messageRepository.save(message);
    }

    @Override
    public Message findById(Long id) {
        return messageRepository.findById(id).get();
    }
}
