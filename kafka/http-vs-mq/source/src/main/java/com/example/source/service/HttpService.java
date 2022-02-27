package com.example.source.service;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
@RequiredArgsConstructor
public class HttpService {

    private static final String TARGET_SERVICE_1 = "http://localhost:8081?data=";
    private static final String TARGET_SERVICE_2 = "http://localhost:8082?data=";
    private static final HttpEntity<?> entity = new HttpEntity<>(null, null);

    public void send(String data) {
        RestTemplate restTemplate = new RestTemplate();
        restTemplate.exchange(TARGET_SERVICE_1+data, HttpMethod.GET, entity, String.class);
        restTemplate.exchange(TARGET_SERVICE_2+data, HttpMethod.GET, entity, String.class);
    }
}
