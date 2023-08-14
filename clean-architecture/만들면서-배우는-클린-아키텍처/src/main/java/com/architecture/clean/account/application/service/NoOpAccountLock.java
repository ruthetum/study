package com.architecture.clean.account.application.service;

import com.architecture.clean.account.application.port.out.AccountLock;
import org.springframework.stereotype.Component;

@Component
class NoOpAccountLock implements AccountLock {

    @Override
    public void lockAccount(Long accountId) {
        // do nothing
    }

    @Override
    public void releaseAccount(Long accountId) {
        // do nothing
    }

}