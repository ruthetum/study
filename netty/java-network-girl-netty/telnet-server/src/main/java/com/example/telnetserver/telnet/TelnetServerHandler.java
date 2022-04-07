package com.example.telnetserver.telnet;

import io.netty.channel.*;

import java.net.InetAddress;
import java.util.Date;

/**
 * 텔넷 서버 로직 처리
 */

@ChannelHandler.Sharable
public class TelnetServerHandler extends SimpleChannelInboundHandler<String> {

    private static final String DELIMITER = "\r\n";
    private static final String CLOSE = "bye";

    // 클라이언트 접속 완료
    @Override
    public void channelActive(ChannelHandlerContext ctx) throws Exception {
        ctx.write("환영합니다. "
                + InetAddress.getLocalHost().getHostName() + "에 접속하셨습니다!" + DELIMITER);
        ctx.write("현재 시간은 " + new Date() + " 입니다." + DELIMITER);
        ctx.flush();
    }

    // 데이터 수신
    @Override
    public void channelRead0(ChannelHandlerContext ctx, String request) throws Exception {
        String response;
        boolean close = false;

        if (request.isEmpty()) {
            response = "명령을 입력해주세요." + DELIMITER;
        } else if (CLOSE.equals(request.toLowerCase())) {
            response = "좋은 하루되세요!" + DELIMITER;
            close = true;
        } else {
            response = "입력하신 명령이 '" + request + "' 입니까?" + DELIMITER;
        }

        ChannelFuture future = ctx.write(response);

        if (close) {
            future.addListener(ChannelFutureListener.CLOSE);
        }
    }

    // 데이터 수신 완료
    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) {
        ctx.flush();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
}
