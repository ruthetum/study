package com.example.mvcorder;

import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@Service
public class OrderService {

    private static final String ITEM_SERVICE_URL = "http://localhost:8081";
    private static final HttpEntity<?> entity = new HttpEntity<>(null, null);

    public int order(int quantity) {
        RestTemplate restTemplate = new RestTemplate();
        ResponseEntity<Integer> responseEntity =
            restTemplate.exchange(ITEM_SERVICE_URL + "/item?quantity=" + quantity, HttpMethod.GET, entity, Integer.class);
        return responseEntity.getBody();
    }

    public int remain() {
        RestTemplate restTemplate = new RestTemplate();
        ResponseEntity<Integer> responseEntity =
                restTemplate.exchange(ITEM_SERVICE_URL + "/remain", HttpMethod.GET, entity, Integer.class);
        return responseEntity.getBody();
    }
}
