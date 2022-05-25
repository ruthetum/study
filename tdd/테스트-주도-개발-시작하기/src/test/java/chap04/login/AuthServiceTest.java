package chap04.login;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class AuthServiceTest {

    @Test
    void 정상_로그인() {
        AuthService authService = new AuthService();
        authService.login("id1", "pw1");
    }

    @Test
    void 비밀번호_불일치() {
        Assertions.assertThrows(PwNotMatchException.class, () -> {
            AuthService authService = new AuthService();
            authService.login("id1", "pw2");
        });
    }
}
