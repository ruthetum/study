package labs.wilump.encryption.twoway

import java.security.MessageDigest
import java.util.*
import javax.crypto.Cipher
import javax.crypto.spec.GCMParameterSpec
import javax.crypto.spec.SecretKeySpec

object Aes256Utils {
    enum class Algorithm(val transformation: String) {
        GCM("AES/GCM/NoPadding"),
    }

    private const val KEY_SPEC_ALGORITHM = "AES"
    private const val IV_HASH_ALGORITHM = "SHA-1"
    private const val IV_LENGTH_BYTE = 12
    private const val KEY_LENGTH_BIT = 256
    private const val GCM_SPEC_TAG_LEN = 128

    private fun secretKeySpec(key: String): SecretKeySpec {
        /**
         * AES-256 알고리즘을 사용하기 위해 key를 32바이트로 설정한다
         */
        require(
            key.isNotEmpty() && key.length == KEY_LENGTH_BIT / 8
        ) { "Key must be ${KEY_LENGTH_BIT / 8} characters long, but was ${key.length} characters long." }
        val keyBytes = key.toByteArray()
        return SecretKeySpec(keyBytes, 0, KEY_LENGTH_BIT / 8, KEY_SPEC_ALGORITHM)

    }

    private fun iv(plainText: String): ByteArray {
        val ivHash = MessageDigest.getInstance(IV_HASH_ALGORITHM).digest(plainText.toByteArray())
        return Arrays.copyOf(ivHash, IV_LENGTH_BYTE)
    }

    fun encrypt(plainText: String, key: String, algorithm: Algorithm = Algorithm.GCM): String {
        /**
         * 대칭키를 생성한다
         */
        val secretKey = secretKeySpec(key)

        /**
         * 블록 암호화에서 사용되는 IV(Initialization Vector)를 생성한다
         */
        val iv = iv(plainText)

        /**
         * 암호화를 위한 파라미터를 생성한다
         * @param: tLen: 암/복호화에 사용될 인증 태그 길이 (128, 120, 112, 104, 96)
         * @param: iv: 초기화 벡터
         */
        val parameterSpec = GCMParameterSpec(GCM_SPEC_TAG_LEN, iv)

        /**
         * 암호화를 위한 Cipher 객체를 생성한다
         */
        val cipher = Cipher.getInstance(algorithm.transformation)
        cipher.init(Cipher.ENCRYPT_MODE, secretKey, parameterSpec)

        val cipherText = cipher.doFinal(plainText.toByteArray())

        return Base64.getEncoder().encodeToString(iv) +
                Base64.getEncoder().encodeToString(cipherText)
    }

    fun decrypt(input: String, password: String, algorithm: Algorithm = Algorithm.GCM): String {
        /**
         * 대칭키를 생성한다
         */
        val secretKey = secretKeySpec(password)

        /**
         * 암호화된 문자열을 디코딩한다
         */
        val ivAndCipherText = Base64.getDecoder().decode(input)

        /**
         * IV와 암호화된 문자열을 분리한다
         */
        val iv = ivAndCipherText.sliceArray(0 until IV_LENGTH_BYTE)
        val cipherText = ivAndCipherText.sliceArray(IV_LENGTH_BYTE until ivAndCipherText.size)

        /**
         * 복호화를 위한 파라미터를 생성한다
         * @param: tLen: 암/복호화에 사용될 인증 태그 길이 (128, 120, 112, 104, 96)
         * @param: iv: 초기화 벡터
         */
        val parameterSpec = GCMParameterSpec(GCM_SPEC_TAG_LEN, iv)

        /**
         * 복호화를 위한 Cipher 객체를 생성한다
         */
        val cipher = Cipher.getInstance(algorithm.transformation)
        cipher.init(Cipher.DECRYPT_MODE, secretKey, parameterSpec)

        val plainText = cipher.doFinal(cipherText)

        return String(plainText)
    }
}