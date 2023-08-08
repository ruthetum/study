package com.architecture.clean.account.domain;

import lombok.AllArgsConstructor;
import lombok.Getter;

import java.time.LocalDateTime;

@Getter
@AllArgsConstructor(access = lombok.AccessLevel.PROTECTED)
public class Activity {

    private Long id;

    private Long sourceAccountId;

    private Long targetAccountId;

    private LocalDateTime timestamp;

    private Money money;

    public static Activity withoutId(Long sourceAccountId, Long targetAccountId, LocalDateTime timestamp, Money money) {
        return new Activity(null, sourceAccountId, targetAccountId, timestamp, money);
    }

    public static Activity withId(Long id, Long sourceAccountId, Long targetAccountId, LocalDateTime timestamp, Money money) {
        return new Activity(id, sourceAccountId, targetAccountId, timestamp, money);
    }
}
