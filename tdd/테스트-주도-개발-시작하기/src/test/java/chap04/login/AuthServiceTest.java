package chap04.login;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

// assertJ
// import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class AuthServiceTest {

    @Test
    void 정상_로그인() {
        AuthService authService = new AuthService();
        authService.login("id1", "pw1");
    }

    @Test
    void 비밀번호_불일치() {
        AuthService authService = new AuthService();
        Assertions.assertThrows(PwNotMatchException.class, () -> {
            authService.login("id1", "pw2");
        });
    }

    // @Test
    // void 비밀번호_불일치_With_AssertJ() {
    // AuthService authService = new AuthService();
    // assertThatThrownBy(() -> authService.login("id1", "pw2"))
    //      .isInstanceOf(PwNotMatchException.class);
    // }
}
