package com.architecture.clean.account.domain;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;

import java.time.LocalDateTime;

@Getter
@AllArgsConstructor(access = AccessLevel.PROTECTED)
public class Account {

    private Long id;
    private Money baselineBalance;
    private ActivityWindow activityWindow;

    public static Account withoutId(
            Money baselineBalance,
            ActivityWindow activityWindow) {
        return new Account(null, baselineBalance, activityWindow);
    }

    public static Account withId(
            Long accountId,
            Money baselineBalance,
            ActivityWindow activityWindow) {
        return new Account(accountId, baselineBalance, activityWindow);
    }

    public Money calculateBalance() {
        return Money.add(this.baselineBalance, activityWindow.calculateBalance(this.id));
    }

    public boolean withdraw(Money money, Long targetAccountId) {
        if (!mayWithdraw(money)) {
            return false;
        }

        Activity withdrawal = Activity.withoutId(this.id, this.id, targetAccountId, LocalDateTime.now(), money);
        this.activityWindow.addActivity(withdrawal);
        return true;
    }

    private boolean mayWithdraw(Money money) {
        return Money.add(this.calculateBalance(), money.negate()).isPositive();
    }

    public boolean deposit(Money money, Long sourceAccountId) {
        Activity deposit = Activity.withoutId(this.id, sourceAccountId, this.id, LocalDateTime.now(), money);
        this.activityWindow.addActivity(deposit);
        return true;
    }
}
