package chap04.login;

import java.util.HashMap;
import java.util.Map;

public class AuthService {

    private final Map<String, User> memory = new HashMap<>() {{
        put("id1", new User("id1", "user1", "pw1"));
        put("id2", new User("id2", "user2", "pw2"));
    }};

    public void login(String id, String pw) {
        User user = memory.get(id);

        if (!user.getPassword().equals(pw)) {
            throw new PwNotMatchException();
        }
    }
}
