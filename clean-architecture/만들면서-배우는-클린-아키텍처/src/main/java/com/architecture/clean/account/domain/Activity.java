package com.architecture.clean.account.domain;

import lombok.AllArgsConstructor;
import lombok.Getter;

import java.time.LocalDateTime;

@Getter
@AllArgsConstructor(access = lombok.AccessLevel.PROTECTED)
public class Activity {

    private Long id;

    private Long ownerAccountId;

    private Long sourceAccountId;

    private Long targetAccountId;

    private LocalDateTime timestamp;

    private Money money;

    public static Activity withoutId(Long ownerAccountId, Long sourceAccountId, Long targetAccountId, LocalDateTime timestamp, Money money) {
        return new Activity(null, ownerAccountId, sourceAccountId, targetAccountId, timestamp, money);
    }

    public static Activity withId(Long id, Long ownerAccountId, Long sourceAccountId, Long targetAccountId, LocalDateTime timestamp, Money money) {
        return new Activity(id, ownerAccountId, sourceAccountId, targetAccountId, timestamp, money);
    }
}
