package com.architecture.clean.account.application.service;

import com.architecture.clean.account.application.port.in.GetAccountBalanceQuery;
import com.architecture.clean.account.application.port.out.LoadAccountPort;
import com.architecture.clean.account.domain.Money;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
@RequiredArgsConstructor
class GetAccountBalanceService implements GetAccountBalanceQuery {

    private final LoadAccountPort loadAccountPort;

    @Override
    public Money getAccountBalance(Long accountId) {
        return loadAccountPort.loadAccount(accountId, LocalDateTime.now())
                .calculateBalance();
    }
}