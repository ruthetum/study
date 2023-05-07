package com.example.demo.future.repository;

import com.example.demo.common.repository.ArticleEntity;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.function.Supplier;
import java.util.stream.Collectors;

@Slf4j
public class ArticleFutureRepository {
    private static List<ArticleEntity> articleEntities;

    public ArticleFutureRepository() {
        articleEntities = List.of(
                new ArticleEntity("1", "제목1", "내용1", "1234"),
                new ArticleEntity("2", "제목2", "내용2", "1234"),
                new ArticleEntity("3", "제목3", "내용3", "12345")
        );
    }

    @SneakyThrows
    public CompletableFuture<List<ArticleEntity>> findAllByUserId(String userId) {
        log.info("ArticleRepository.findAllByUserId: {}", userId);
        return CompletableFuture.supplyAsync((Supplier<List<ArticleEntity>>) () -> {
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
            return articleEntities.stream()
                    .filter(articleEntity -> articleEntity.getUserId().equals(userId))
                    .collect(Collectors.toList());
        });
    }
}
