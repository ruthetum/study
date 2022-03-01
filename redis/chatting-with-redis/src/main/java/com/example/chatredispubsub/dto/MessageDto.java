package com.example.chatredispubsub.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class MessageDto {
    private String roomId;
    private String writer;
    private String message;
}
