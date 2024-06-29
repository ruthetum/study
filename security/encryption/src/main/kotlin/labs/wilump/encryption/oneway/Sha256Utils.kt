package labs.wilump.encryption.oneway

import java.security.DigestException
import java.security.MessageDigest

object Sha256Utils {
    private const val algorithm = "SHA-256"
    private const val digits = "0123456789ABCDEF"

    fun encrypt(plainText: String): String {
        val hash: ByteArray
        try {
            val md = MessageDigest.getInstance(algorithm)
            md.update(plainText.toByteArray())
            hash = md.digest()
        } catch (e: CloneNotSupportedException) {
            throw DigestException("couldn't make digest of partial content");
        }

        return bytesToHex(hash)
    }

    private fun bytesToHex(byteArray: ByteArray): String {
        val hexChars = CharArray(byteArray.size * 2)
        for (i in byteArray.indices) {
            val v = byteArray[i].toInt() and 0xff
            hexChars[i * 2] = digits[v shr 4]
            hexChars[i * 2 + 1] = digits[v and 0xf]
        }
        return String(hexChars)
    }
}