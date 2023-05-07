package com.example.demo.blocking.repository;

import com.example.demo.common.repository.ArticleEntity;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;

import java.util.List;
import java.util.stream.Collectors;

@Slf4j
public class ArticleRepository {
    private static List<ArticleEntity> articleEntities;

    public ArticleRepository() {
        articleEntities = List.of(
                new ArticleEntity("1", "제목1", "내용1", "1234"),
                new ArticleEntity("2", "제목2", "내용2", "1234"),
                new ArticleEntity("3", "제목3", "내용3", "12345")
        );
    }

    @SneakyThrows
    public List<ArticleEntity> findAllByUserId(String userId) {
        log.info("ArticleRepository.findAllByUserId: {}", userId);
        Thread.sleep(1000);
        return articleEntities.stream()
                .filter(articleEntity -> articleEntity.getUserId().equals(userId))
                .collect(Collectors.toList());
    }
}
