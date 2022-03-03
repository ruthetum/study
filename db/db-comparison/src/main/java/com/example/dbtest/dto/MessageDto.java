package com.example.dbtest.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class MessageDto {
    private int type;
    private Long id;
    private String content;
}
