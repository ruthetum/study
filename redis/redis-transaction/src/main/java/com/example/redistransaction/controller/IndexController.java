package com.example.redistransaction.controller;

import com.example.redistransaction.service.IndexService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Slf4j
@Controller
@RequiredArgsConstructor
@RequestMapping("/redis")
public class IndexController {

    private final IndexService indexService;

    @GetMapping("/v1")
    public void func1() {
        log.info("Use RedisTemplate");
        indexService.useOperations();
    }

    @GetMapping("/v2")
    public void func2() {
        log.info("Use Transactional Annotation");
        indexService.useTransactionalAnnotation();
    }
}
