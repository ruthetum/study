package com.architecture.clean.account.application.port.in;

import com.architecture.clean.account.domain.Money;
import com.architecture.clean.common.SelfValidating;
import jakarta.validation.constraints.NotNull;
import lombok.EqualsAndHashCode;
import lombok.Value;

@Value
@EqualsAndHashCode(callSuper = false)
public class SendMoneyCommand extends SelfValidating<SendMoneyCommand> {

    @NotNull
    private final Long sourceAccountId;

    @NotNull
    private final Long targetAccountId;

    @NotNull
    private final Money money;

    public SendMoneyCommand(
            Long sourceAccountId,
            Long targetAccountId,
            Money money) {
        this.sourceAccountId = sourceAccountId;
        this.targetAccountId = targetAccountId;
        this.money = money;
        this.validateSelf();
    }
}
