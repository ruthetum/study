package com.mvc.order.controller;

import com.mvc.order.dto.OrderRequest;
import com.mvc.order.dto.OrderResponse;
import com.mvc.order.service.OrderService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
@RequiredArgsConstructor
public class OrderController {

    private final OrderService orderService;

    @PostMapping("/order")
    public OrderResponse order(@RequestBody OrderRequest request) {
        return orderService.order(request);
    }
}
