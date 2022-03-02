package ch1.echo;


import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import java.nio.charset.Charset;

/**
 * - SimpleChannelInboundHandler와 ChannelInboundHandlerAdapter 모두 네티 기본 제공 클래스
 * - 클라이언트로부터 수신한 데이터를 처리하는 이벤트를 제공
 * - SimpleChannelInboundHandler는 ChannelInboundHandlerAdapter를 상속
 * - SimpleChannelInboundHandler는 데이터가 수신되었을 때 호출되는 channelRead 이벤트에 대한 처리가 이미 구현
 */
public class EchoServerHandler extends ChannelInboundHandlerAdapter {

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        String readMessage = ((ByteBuf) msg).toString(Charset.defaultCharset());

        StringBuilder builder = new StringBuilder();
        builder.append("수신한 문자열 [");
        builder.append(readMessage);
        builder.append("]");
        System.out.println(builder.toString());

        // 채널 파이프라인에 대한 이벤트 처리
        ctx.write(msg);
    }

    // channelRead 이벤트의 처리가 완료된 후 자동으로 수행되는 이벤트 메서드
    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) {
        ctx.flush();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        // Close the connection when an exception is raised.
        cause.printStackTrace();
        ctx.close();
    }

}
