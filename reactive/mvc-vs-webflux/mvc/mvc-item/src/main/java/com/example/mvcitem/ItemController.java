package com.example.mvcitem;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
@RequiredArgsConstructor
public class ItemController {

    private final ItemService itemService;

    /**
     * 주문받은 order의 수에 따라 재고 수량 감소
     */
    @GetMapping("/item")
    public ResponseEntity<Integer> order(
            @RequestParam("quantity") int quantity
    ) {
        log.info("[mvc-item] quantity : {}", quantity);
        int remain = itemService.order(quantity);
        return ResponseEntity.ok().body(remain);
    }

    /**
     * 재고 수량 조회
     */
    @GetMapping("/remain")
    public ResponseEntity<Integer> remain() {
        log.info("[mvc-item] remain");
        int remain = itemService.remain();
        return ResponseEntity.ok().body(remain);
    }
}
