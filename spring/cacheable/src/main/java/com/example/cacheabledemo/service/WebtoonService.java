package com.example.cacheabledemo.service;

import com.example.cacheabledemo.config.CacheKey;
import com.example.cacheabledemo.domain.Webtoon;
import com.example.cacheabledemo.repository.WebtoonRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.DayOfWeek;
import java.time.LocalDate;
import java.util.List;

@Service
@RequiredArgsConstructor
@Transactional(readOnly = true)
public class WebtoonService {

    private final WebtoonRepository webtoonRepository;

    public List<Webtoon> getTodayWebtoon() {
        DayOfWeek todayDayOfWeek = LocalDate.now().getDayOfWeek();
        return webtoonRepository.findByDayOfWeek(todayDayOfWeek);
    }

    @Cacheable(value = CacheKey.TODAY_WEBTOON, key = "0")
    public List<Webtoon> getTodayWebtoonWithCache() {
        DayOfWeek todayDayOfWeek = LocalDate.now().getDayOfWeek();
        return webtoonRepository.findByDayOfWeek(todayDayOfWeek);
    }
}
