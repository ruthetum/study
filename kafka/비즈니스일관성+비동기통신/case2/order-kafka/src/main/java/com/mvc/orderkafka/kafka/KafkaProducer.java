package com.mvc.orderkafka.kafka;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Slf4j
@Service
@RequiredArgsConstructor
public class KafkaProducer {

    @Value("${spring.kafka.topic.order}")
    private String orderTopic;

    @Value("${spring.kafka.topic.payment}")
    private String paymentTopic;

    private final KafkaTemplate<String, String> kafkaTemplate;
}
