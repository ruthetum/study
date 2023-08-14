package com.architecture.clean.account.application.port.out;

public interface AccountLock {

    void lockAccount(Long accountId);

    void releaseAccount(Long accountId);
}
