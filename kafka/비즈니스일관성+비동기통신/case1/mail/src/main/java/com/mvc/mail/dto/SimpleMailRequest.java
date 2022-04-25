package com.mvc.mail.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class SimpleMailRequest {
    private String to;
    private String title;
    private String content;
}
