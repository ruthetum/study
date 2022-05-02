package com.example.cacheabledemo.repository;

import com.example.cacheabledemo.domain.Webtoon;
import org.springframework.data.jpa.repository.JpaRepository;

import java.time.DayOfWeek;
import java.util.List;

public interface WebtoonRepository extends JpaRepository<Webtoon, Long> {
    List<Webtoon> findByDayOfWeek(DayOfWeek dayOfWeek);
}
