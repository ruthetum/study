package com.example.cacheabledemo.controller;

import com.example.cacheabledemo.domain.Webtoon;
import com.example.cacheabledemo.service.WebtoonService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequiredArgsConstructor
@RequestMapping("/webtoon")
public class WebtoonController {

    private final WebtoonService webtoonService;

    // 캐시 미적용
    @GetMapping
    public List<Webtoon> getTodayWebtoon() {
        return webtoonService.getTodayWebtoon();
    }

    // 캐시 적용
    @GetMapping("/cache")
    public List<Webtoon> getTodayWebtoonWithCache() {
        return webtoonService.getTodayWebtoonWithCache();
    }
}
