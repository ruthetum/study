package com.architecture.clean.account.application.port.out;

import com.architecture.clean.account.domain.Account;

public interface UpdateAccountStatePort {
    void updateActivities(Account account);
}
