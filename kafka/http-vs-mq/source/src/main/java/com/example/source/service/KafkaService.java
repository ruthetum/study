package com.example.source.service;

import lombok.RequiredArgsConstructor;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class KafkaService {

    private final KafkaTemplate<String, String> kafkaTemplate;
    private final String TOPIC = "topic";

    public void send(String data) {
        kafkaTemplate.send(TOPIC, data);
    }
}
