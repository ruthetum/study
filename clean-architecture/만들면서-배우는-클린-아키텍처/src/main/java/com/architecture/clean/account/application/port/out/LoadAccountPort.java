package com.architecture.clean.account.application.port.out;

import com.architecture.clean.account.domain.Account;

import java.time.LocalDateTime;

public interface LoadAccountPort {
    Account loadAccount(Long accountId, LocalDateTime baselineDate);
}
