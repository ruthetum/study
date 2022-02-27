package com.example.target.service;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.clients.consumer.Consumer;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@Service
@RequiredArgsConstructor
public class ConsumerService {

    private final String TOPIC = "topic";
    private final String GROUP_ID = "test";

    @KafkaListener(topics = TOPIC, groupId = GROUP_ID)
    @Transactional
    public void consumeMsg(String data, Consumer<String, String> consumer)  {
        log.info(data);
        consumer.commitAsync();
    }
}
