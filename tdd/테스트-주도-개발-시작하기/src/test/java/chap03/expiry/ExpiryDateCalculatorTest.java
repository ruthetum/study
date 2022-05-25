package chap03.expiry;

import org.junit.jupiter.api.Test;

import java.time.LocalDate;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class ExpiryDateCalculatorTest {

    @Test
    void 만원_납부하면_한달_뒤가_만료일이_됨() {
        PayData payData = PayData.builder()
                .billingDate(LocalDate.of(2022, 5, 25))
                .payAmount(10000)
                .build();
        assertExpiryDate(payData, LocalDate.of(2022, 6, 25));
    }

    @Test
    void 납부일과_한달_뒤_일자가_같지_않음() {
        PayData payData = PayData.builder()
                .billingDate(LocalDate.of(2022, 5, 31))
                .payAmount(10000)
                .build();
        assertExpiryDate(payData, LocalDate.of(2022, 6, 30));
    }

    @Test
    void 첫_납부일과_만료일_일자가_다를때_만원_납부() {
        PayData payData = PayData.builder()
                .firstBillingDate(LocalDate.of(2022, 5, 31))
                .billingDate(LocalDate.of(2022, 6, 30))
                .payAmount(10000)
                .build();
        assertExpiryDate(payData, LocalDate.of(2022, 7, 31));
    }

    @Test
    void 이만원_이상_납부하면_비례해서_만료일_계산() {
        PayData payData = PayData.builder()
                .billingDate(LocalDate.of(2022,3,1))
                .payAmount(20000)
                .build();
        assertExpiryDate(payData, LocalDate.of(2022, 5, 1));
    }

    @Test
    void 첫_납부일과_만료일_일자가_다를때_이만원_이상_납부() {
        PayData payData = PayData.builder()
                .firstBillingDate(LocalDate.of(2022, 1, 31))
                .billingDate(LocalDate.of(2022, 2, 28))
                .payAmount(20000)
                .build();
        assertExpiryDate(payData, LocalDate.of(2022, 4, 30));
    }

    @Test
    void 십만원을_납부하면_1년_제공() {
        PayData payData = PayData.builder()
                .billingDate(LocalDate.of(2022, 2, 28))
                .payAmount(100000)
                .build();
        assertExpiryDate(payData, LocalDate.of(2023, 2, 28));
    }

    private void assertExpiryDate(PayData payData, LocalDate expectedExpiryDate) {
        ExpiryDateCalculator cal = new ExpiryDateCalculator();
        LocalDate expiryDate = cal.calculateExpiryDate(payData);
        assertEquals(expectedExpiryDate, expiryDate);
    }
}
