package com.architecture.clean.account.adapter.out.persistence;

import com.architecture.clean.account.domain.Account;
import com.architecture.clean.account.domain.Activity;
import com.architecture.clean.account.domain.ActivityWindow;
import com.architecture.clean.account.domain.Money;
import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;

@Component
class AccountMapper {

    Account mapToDomainEntity(
            AccountJpaEntity account,
            List<ActivityJpaEntity> activities,
            Long withdrawalBalance,
            Long depositBalance) {

        Money baselineBalance = Money.subtract(
                Money.of(depositBalance),
                Money.of(withdrawalBalance));

        return Account.withId(
                account.getId(),
                baselineBalance,
                mapToActivityWindow(activities));

    }

    ActivityWindow mapToActivityWindow(List<ActivityJpaEntity> activities) {
        List<Activity> mappedActivities = new ArrayList<>();

        for (ActivityJpaEntity activity : activities) {
            mappedActivities.add(Activity.withId(
                    activity.getId(),
                    activity.getOwnerAccountId(),
                    activity.getSourceAccountId(),
                    activity.getTargetAccountId(),
                    activity.getTimestamp(),
                    Money.of(activity.getAmount())));
        }

        return new ActivityWindow(mappedActivities);
    }

    ActivityJpaEntity mapToJpaEntity(Activity activity) {
        return new ActivityJpaEntity(
                activity.getId() == null ? null : activity.getId(),
                activity.getTimestamp(),
                activity.getOwnerAccountId(),
                activity.getSourceAccountId(),
                activity.getTargetAccountId(),
                activity.getMoney().getAmount().longValue());
    }

}
