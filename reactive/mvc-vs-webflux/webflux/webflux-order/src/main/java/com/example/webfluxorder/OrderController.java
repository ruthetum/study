package com.example.webfluxorder;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@Slf4j
@RestController
@RequiredArgsConstructor
public class OrderController {

    private final OrderService orderService;

    /**
     * 주문을 받으면 item service에 주문받은 갯수만큼 재고 수량 감소 요청
     */
    @GetMapping("/order")
    public Mono<Integer> order(
            @RequestParam("quantity") int quantity
    ) {
        log.info("[webflux-order] quantity : {}", quantity);
        Mono<Integer> reamin = orderService.order(quantity);
        return reamin;
    }

    /**
     * 재고 수량 확인
     */
    @GetMapping("/remain")
    public Mono<Integer> remain() {
        log.info("[webflux-order] remain");
        Mono<Integer> remain = orderService.remain();
        return remain;
    }
}
