package com.example.source.controller;

import com.example.source.service.HttpService;
import com.example.source.service.KafkaService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;

@Slf4j
@RestController
@RequiredArgsConstructor
public class SourceController {

    private final HttpService httpService;
    private final KafkaService kafkaService;

    @PostMapping("/http")
    public void sendByHttp() {
        log.info("/http");
        String data = LocalDateTime.now().toString();
        httpService.send(data);
    }

    @PostMapping("/kafka")
    public void sendByKafka() {
        log.info("/kafka");
        String data = LocalDateTime.now().toString();
        kafkaService.send(data);
    }
}
