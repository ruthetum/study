package ch1.echo;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import java.nio.charset.Charset;

public class EchoClientHandler extends ChannelInboundHandlerAdapter {
    
    // 소켓 채널이 최초로 활성화되었을 때 실행
    @Override
    public void channelActive(ChannelHandlerContext ctx) {
        String sendMessage = "Hello netty";

        ByteBuf messageBuffer = Unpooled.buffer();
        messageBuffer.writeBytes(sendMessage.getBytes());

        StringBuilder builder = new StringBuilder();
        builder.append("전송한 문자열 [");
        builder.append(sendMessage);
        builder.append("]");

        System.out.println(builder.toString());
        
        // writeAndFlush() : 내부적으로 데이터 기록과 전송의 두 가지 메서드를 호출
        // write : 채널에 데이터 기록, flush : 채널에 기록된 데이털르 서버로 전송
        ctx.writeAndFlush(messageBuffer);
    }

    // 서버로부터 수신된 데이터가 있을 때 호출되는 이벤트 메서드
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        String readMessage = ((ByteBuf) msg).toString(Charset.defaultCharset());

        StringBuilder builder = new StringBuilder();
        builder.append("수신한 문자열 [");
        builder.append(readMessage);
        builder.append("]");

        System.out.println(builder.toString());
    }

    // 수신된 데이터를 모두 읽었을 때 호출되는 이벤트 메서드
    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) {
        ctx.close();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
}