package com.example.demo.common.domain;

import lombok.Data;

@Data
public class Article {
    private final String id;
    private final String title;
    private final String content;
    private final String userId;
}
