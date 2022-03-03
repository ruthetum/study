package com.example.webfluxitem;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@Slf4j
@RestController
@RequiredArgsConstructor
public class ItemController {

    // private final ItemServiceImplBasic itemService;
    private final ItemServiceImplReactive itemService;

    /**
     * 주문받은 order의 수에 따라 재고 수량 감소
     */
    @GetMapping("/item")
    public Mono<Integer> order(
            @RequestParam(value = "quantity", defaultValue = "1") int quantity
    ) {
        log.info("[webflux-item] quantity : {}", quantity);

        // basic
//        int remain = itemService.order(quantity);
//        return Mono.just(remain);

        // reactive
        return itemService.order(quantity);
    }

    /**
     * 재고 수량 조회
     */
    @GetMapping("/remain")
    public Mono<Integer> remain() {
        log.info("[webflux-item] remain");

        // basic
//        int remain = itemService.remain();
//        return Mono.just(remain);

        // reactive
        return itemService.remain();
    }
}
