package com.architecture.clean.account.domain;

import lombok.Getter;

import java.time.LocalDateTime;
import java.util.*;

@Getter
public class ActivityWindow {
    private List<Activity> activities = new ArrayList<>();

    public LocalDateTime getStartTimestamp() {
        return activities.stream()
                .min(Comparator.comparing(Activity::getTimestamp))
                .orElseThrow(IllegalStateException::new)
                .getTimestamp();
    }

    public LocalDateTime getEndTimestamp() {
        return activities.stream()
                .max(Comparator.comparing(Activity::getTimestamp))
                .orElseThrow(IllegalStateException::new)
                .getTimestamp();
    }

    public Money calculateBalance(Long accountId) {
        Money depositBalance = activities.stream()
                .filter(a -> a.getId().equals(accountId))
                .map(Activity::getMoney)
                .reduce(Money.ZERO, Money::add);

        Money withdrawalBalance = activities.stream()
                .filter(a -> a.getSourceAccountId().equals(accountId))
                .map(Activity::getMoney)
                .reduce(Money.ZERO, Money::add);

        return Money.add(depositBalance, withdrawalBalance.negate());
    }

    public ActivityWindow(List<Activity> activities) {
        this.activities = activities;
    }

    public ActivityWindow(Activity... activities) {
        this.activities = new ArrayList<>(Arrays.asList(activities));
    }

    public void addActivity(Activity activity) {
        this.activities.add(activity);
    }
}
