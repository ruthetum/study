package com.example.dbtest.service;

import com.example.dbtest.domain.Message;
import com.example.dbtest.dto.MessageDto;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public interface MessageService {
    void save(MessageDto messageDto);
    Message findById(Long id);
}
