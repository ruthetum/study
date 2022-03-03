package com.example.webfluxorder;

import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;

@Service
public class OrderService {

    private static final String ITEM_SERVICE_URL = "http://localhost:8081";

    public Mono<Integer>order(int quantity) {
        WebClient client = WebClient.builder()
                .baseUrl(ITEM_SERVICE_URL)
                .build();

        Mono<Integer> response = client.get()
                .uri("/item?quantity=" + quantity)
                .retrieve().bodyToMono(Integer.class);

        return response;
    }

    public Mono<Integer> remain() {
        WebClient client = WebClient.builder()
                .baseUrl(ITEM_SERVICE_URL)
                .build();

        Mono<Integer> response = client.get()
                .uri("/remain")
                .retrieve().bodyToMono(Integer.class);

        return response;
    }
}
