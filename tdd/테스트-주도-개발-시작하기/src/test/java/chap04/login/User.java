package chap04.login;

public class User {
    private String id;
    private String name;
    private String password;

    public String getId() { return id; }
    public String getName() { return name; }
    public String getPassword() { return password; }

    public User(String id, String name, String password) {
        this.id = id;
        this.name = name;
        this.password = password;
    }
}
