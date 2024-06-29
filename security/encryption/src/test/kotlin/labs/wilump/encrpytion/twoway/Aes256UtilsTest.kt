package labs.wilump.encrpytion.twoway

import labs.wilump.encryption.twoway.Aes256Utils
import org.junit.jupiter.api.Test

class Aes256UtilsTest {
    @Test
    fun `AES-256 GCM 알고리즘을 사용하여 암호화 및 복호화를 진행한다`() {
        // given
        val key = "12345678901234567890123456789012"
        val plainText = "12345678"

        // when
        val encryptedText = Aes256Utils.encrypt(plainText, key)
        val decryptedText = Aes256Utils.decrypt(encryptedText, key)

        // then
        assert(plainText != encryptedText)
        assert(plainText == decryptedText)
    }
}