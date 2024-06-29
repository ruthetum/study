package labs.wilump.encrpytion.oneway

import labs.wilump.encryption.oneway.Sha256Utils
import org.junit.jupiter.api.Test

class Sha256UtilsTest {
    @Test
    fun `SHA-256 알고리즘을 이용하여 암호화한다`() {
        // given
        val plainText = "12345678"

        // when
        val encryptedText = Sha256Utils.encrypt(plainText)

        // then
        assert(plainText != encryptedText)
        assert(encryptedText == Sha256Utils.encrypt(plainText))
    }
}