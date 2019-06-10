import socket
import sys

def main():

    # 待ち受け用のソケットオブジェクト
    server = socket.socket()
    try:
        # 待ち受けポートに割り当て
        server.bind(('0.0.0.0', 1323))
        while True:
            # 待ち受け開始
            server.listen(5)

            # 要求がいたら受け付け
            client, addr = server.accept()

            # 受け取ったメッセージを出力
            recieve_messege = client.recv(4096).decode()
            print('{}'.format(recieve_messege))

            # 返答
            # client.send(b"I am socket server...\n")
            client.close()
    except KeyboardInterrupt:
        server.close()
    finally:
        server.close()

if __name__ == "__main__":
    main()
