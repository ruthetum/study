package chap02.password;

public class PasswordStrengthMeter {

    public PasswordStrength meter(String s) {
        if (s == null || s.isEmpty()) return PasswordStrength.INVALID;

        int count = getMerCriteriaCounts(s);

        if (count <= 1) return PasswordStrength.WEAK;
        if (count == 2) return PasswordStrength.NORMAL;
        return PasswordStrength.STRONG;
    }

    private int getMerCriteriaCounts(String s) {
        int count = 0;
        if (s.length() >= 8) count++;
        if (meetsContainingNumberCriteria(s)) count++;
        if (meetContainingUppercaseCriteria(s)) count++;
        return count;
    }

    private boolean meetsContainingNumberCriteria(String s) {
        for (char ch: s.toCharArray()) {
            if (ch >= '0' && ch <= '9') { return true; }
        }
        return false;
    }

    private boolean meetContainingUppercaseCriteria(String s) {
        for (char ch: s.toCharArray()) {
            if (Character.isUpperCase(ch)) { return true; }
        }
        return false;
    }
}
