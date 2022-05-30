package simple;

import com.linecorp.armeria.common.HttpResponse;
import com.linecorp.armeria.server.Server;

public class Main {
    public static void main(String[] args) {
        Server server = Server.builder()
                .http(8080)
                // .https(8443) // HTTPS support
                // .tlsSelfSigned()
                .service("/", (ctx, req) -> HttpResponse.of("Hello, World!"))
                .build();
        server.start();
    }
}
