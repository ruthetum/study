package com.architecture.clean.account.application.service;

import com.architecture.clean.account.application.port.in.SendMoneyCommand;
import com.architecture.clean.account.application.port.out.AccountLock;
import com.architecture.clean.account.application.port.out.LoadAccountPort;
import com.architecture.clean.account.application.port.out.UpdateAccountStatePort;
import com.architecture.clean.account.domain.Account;
import com.architecture.clean.account.domain.Money;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mockito;

import java.time.LocalDateTime;
import java.util.List;
import java.util.stream.Collectors;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.BDDMockito.given;
import static org.mockito.BDDMockito.then;
import static org.mockito.Mockito.times;

public class SendMoneyServiceTest {

    private final LoadAccountPort loadAccountPort =
            Mockito.mock(LoadAccountPort.class);

    private final AccountLock accountLock =
            Mockito.mock(AccountLock.class);

    private final UpdateAccountStatePort updateAccountStatePort =
            Mockito.mock(UpdateAccountStatePort.class);

    private final SendMoneyService sendMoneyService =
            new SendMoneyService(loadAccountPort, accountLock, updateAccountStatePort, moneyTransferProperties());

    @Test
    void givenWithdrawalFails_thenOnlySourceAccountIsLockedAndReleased() {

        Long sourceAccountId = 41L;
        Account sourceAccount = givenAnAccountWithId(sourceAccountId);

        Long targetAccountId = 42L;
        Account targetAccount = givenAnAccountWithId(targetAccountId);

        givenWithdrawalWillFail(sourceAccount);
        givenDepositWillSucceed(targetAccount);

        SendMoneyCommand command = new SendMoneyCommand(
                sourceAccountId,
                targetAccountId,
                Money.of(300L));

        boolean success = sendMoneyService.sendMoney(command);

        assertThat(success).isFalse();

        then(accountLock).should().lockAccount(eq(sourceAccountId));
        then(accountLock).should().releaseAccount(eq(sourceAccountId));
        then(accountLock).should(times(0)).lockAccount(eq(targetAccountId));
    }

    @Test
    void transactionSucceeds() {

        Account sourceAccount = givenSourceAccount();
        Account targetAccount = givenTargetAccount();

        givenWithdrawalWillSucceed(sourceAccount);
        givenDepositWillSucceed(targetAccount);

        Money money = Money.of(500L);

        SendMoneyCommand command = new SendMoneyCommand(
                sourceAccount.getId(),
                targetAccount.getId(),
                money);

        boolean success = sendMoneyService.sendMoney(command);

        assertThat(success).isTrue();

        Long sourceAccountId = sourceAccount.getId();
        Long targetAccountId = targetAccount.getId();

        then(accountLock).should().lockAccount(eq(sourceAccountId));
        then(sourceAccount).should().withdraw(eq(money), eq(targetAccountId));
        then(accountLock).should().releaseAccount(eq(sourceAccountId));

        then(accountLock).should().lockAccount(eq(targetAccountId));
        then(targetAccount).should().deposit(eq(money), eq(sourceAccountId));
        then(accountLock).should().releaseAccount(eq(targetAccountId));

        thenAccountsHaveBeenUpdated(sourceAccountId, targetAccountId);
    }

    private void thenAccountsHaveBeenUpdated(Long... accountIds){
        ArgumentCaptor<Account> accountCaptor = ArgumentCaptor.forClass(Account.class);
        then(updateAccountStatePort).should(times(accountIds.length))
                .updateActivities(accountCaptor.capture());

        List<Long> updatedAccountIds = accountCaptor.getAllValues()
                .stream()
                .map(Account::getId)
                .collect(Collectors.toList());

        for(Long accountId : accountIds){
            assertThat(updatedAccountIds).contains(accountId);
        }
    }

    private void givenDepositWillSucceed(Account account) {
        given(account.deposit(any(Money.class), any(Long.class)))
                .willReturn(true);
    }

    private void givenWithdrawalWillFail(Account account) {
        given(account.withdraw(any(Money.class), any(Long.class)))
                .willReturn(false);
    }

    private void givenWithdrawalWillSucceed(Account account) {
        given(account.withdraw(any(Money.class), any(Long.class)))
                .willReturn(true);
    }

    private Account givenTargetAccount(){
        return givenAnAccountWithId(42L);
    }

    private Account givenSourceAccount(){
        return givenAnAccountWithId(41L);
    }

    private Account givenAnAccountWithId(Long id) {
        Account account = Mockito.mock(Account.class);
        given(account.getId())
                .willReturn(id);
        given(loadAccountPort.loadAccount(eq(account.getId()), any(LocalDateTime.class)))
                .willReturn(account);
        return account;
    }

    private MoneyTransferProperties moneyTransferProperties(){
        return new MoneyTransferProperties(Money.of(Long.MAX_VALUE));
    }
}
