package chap02.password;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

/**
 * 암호 강도 검사기
 *
 * 규칙
 * - 길이가 8글자 이상
 * - 0부터 9 사이의 숫자를 포함
 * - 대문자 포함
 *
 * 강도
 * - 강함 : 모든 규칙 충족
 * - 보통 : 2개의 규칙 충족
 * - 약함 : 그 외의 경우
 */

public class PasswordStrengthMeterTest {

    private PasswordStrengthMeter meter = new PasswordStrengthMeter();

    private void assertStrength(String password, PasswordStrength expStr) {
        PasswordStrength result = meter.meter(password);
        assertEquals(expStr, result);
    }

    @DisplayName("모든 규칙을 충족하는 경우")
    @Test
    void meetsAllCriteria_Then_Strong() {
        assertStrength("ab12!@AB", PasswordStrength.STRONG);
    }

    @DisplayName("길이만 8글자 미만이고 나머지 조건은 충족하는 경우")
    @Test
    void meetsOtherCriteria_except_for_Length_Then_Normal() {
        assertStrength("ab12!@A", PasswordStrength.NORMAL);
        assertStrength("Ab12!c", PasswordStrength.NORMAL);
        PasswordStrengthMeter meter = new PasswordStrengthMeter();
    }

    @DisplayName("숫자를 포함하지 않고 나머지 조건은 충족하는 경우")
    @Test
    void meetsOtherCriteria_except_for_number_Then_Normal() {
        assertStrength("ab!@ABqwer", PasswordStrength.NORMAL);
    }

    @DisplayName("값이 없는 경우")
    @Test
    void nullInput_Then_Invalid() {
        assertStrength(null, PasswordStrength.INVALID);
    }

    @DisplayName("값이 비어있는 경우")
    @Test
    void emptyInput_Then_Invalid() {
        assertStrength("", PasswordStrength.INVALID);
    }

    @DisplayName("대문자를 포함하지 않고 나머지 조건을 충족하는 경우")
    @Test
    void meetsOtherCriteria_except_for_Uppercase_Then_Normal() {
        assertStrength("ab12!@df", PasswordStrength.NORMAL);
    }

    @DisplayName("길이가 8글자 이상인 조건만 충족하는 경우")
    @Test
    void meetsOnlyLengthCriteria_Then_Weak() {
        assertStrength("abdefghi", PasswordStrength.WEAK);
    }

    @DisplayName("숫자 포함 조건만 충족하는 경우")
    @Test
    void meetsOnlyNumCriteria_Then_Weak() {
        assertStrength("12345", PasswordStrength.WEAK);
    }

    @DisplayName("대문자 포함 조건만 충족하는 경우")
    @Test
    void meetsOnlyUppercaseCriteria_Then_Weak() {
        assertStrength("ABAZ", PasswordStrength.WEAK);
    }

    @DisplayName("아무 조건도 충족하지 않는 경우")
    @Test
    void meetsNoCriteria_Then_Weak() {
        assertStrength("abc", PasswordStrength.WEAK);
    }
}
