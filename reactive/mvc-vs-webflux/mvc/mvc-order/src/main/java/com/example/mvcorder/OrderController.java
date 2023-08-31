package com.example.mvcorder;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
@RequiredArgsConstructor
public class OrderController {

    private final OrderService orderService;

    /**
     * 주문을 받으면 item service에 주문받은 갯수만큼 재고 수량 감소 요청
     */
    @GetMapping("/order")
    public ResponseEntity<Integer> order(
            @RequestParam("quantity") int quantity
    ) {
        log.info("[mvc-order] quantity : {}", quantity);
        int remain = orderService.order(quantity);
        return ResponseEntity.ok().body(remain);
    }

    /**
     * 재고 수량 확인
     */
    @GetMapping("/remain")
    public ResponseEntity<Integer> remain() {
        log.info("[mvc-order] remain");
        int remain = orderService.remain();
        return ResponseEntity.ok().body(remain);
    }
}
