package chap07.autodebit;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;

import static chap07.autodebit.CardValidity.INVALID;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class AutoDebitRegister_Fake_Test {

    private AutoDebitRegister register;
    private StubCardNumberValidator stubValidator;
    private MemoryAutoDebitInfoRepository repository;

    @BeforeEach
    void setUp() {
        stubValidator = new StubCardNumberValidator();
        repository = new MemoryAutoDebitInfoRepository();
        register = new AutoDebitRegister(stubValidator, repository);
    }

    @Test
    void alreadyRegistered_InfoUpdated() {
        repository.save(new AutoDebitInfo("user1", "111122223333", LocalDateTime.now()));

        AutoDebitReq req = new AutoDebitReq("user1", "11232321321321");
        RegisterResult result = this.register.register(req);

        AutoDebitInfo saved = repository.findOne("user1");
        assertEquals("11232321321321", saved.getCardNumber());
    }
}
