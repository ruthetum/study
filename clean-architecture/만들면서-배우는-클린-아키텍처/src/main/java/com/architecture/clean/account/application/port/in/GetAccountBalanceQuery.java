package com.architecture.clean.account.application.port.in;

import com.architecture.clean.account.domain.Money;

public interface GetAccountBalanceQuery {
    Money getAccountBalance(Long accountId);
}
