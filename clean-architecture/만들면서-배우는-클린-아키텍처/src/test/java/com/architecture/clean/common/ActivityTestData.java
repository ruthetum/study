package com.architecture.clean.common;

import com.architecture.clean.account.domain.Activity;
import com.architecture.clean.account.domain.Money;
import lombok.Builder;

import java.time.LocalDateTime;

public class ActivityTestData {

    public static ActivityBuilder defaultActivity(){
        return new ActivityBuilder()
                .withOwnerAccount(42L)
                .withSourceAccount(42L)
                .withTargetAccount(41L)
                .withTimestamp(LocalDateTime.now())
                .withMoney(Money.of(999L));
    }

    public static class ActivityBuilder {
        private Long id;
        private Long ownerAccountId;
        private Long sourceAccountId;
        private Long targetAccountId;
        private LocalDateTime timestamp;
        private Money money;

        public ActivityBuilder withId(Long id) {
            this.id = id;
            return this;
        }

        public ActivityBuilder withOwnerAccount(Long accountId) {
            this.ownerAccountId = accountId;
            return this;
        }

        public ActivityBuilder withSourceAccount(Long accountId) {
            this.sourceAccountId = accountId;
            return this;
        }

        public ActivityBuilder withTargetAccount(Long accountId) {
            this.targetAccountId = accountId;
            return this;
        }

        public ActivityBuilder withTimestamp(LocalDateTime timestamp) {
            this.timestamp = timestamp;
            return this;
        }

        public ActivityBuilder withMoney(Money money) {
            this.money = money;
            return this;
        }

        public Activity build() {
            return Activity.withId(
                    this.id,
                    this.ownerAccountId,
                    this.sourceAccountId,
                    this.targetAccountId,
                    this.timestamp,
                    this.money);
        }
    }
}


