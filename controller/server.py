import socket
import sys
import numpy as np

#問題点: go to x y z speedは

#kinectから受け取ったデータ
# [x,y,z,time]
fr = [0.0,0.0,0.0,0.0] #追加

def main():
    # 待ち受け用のソケットオブジェクト
    server = socket.socket()
    #tello用のソケットオブジジェクト
    host = ''
    port = 9000
    locaddr = (host,port)
    tello = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)   #追加
    tello_address = ('192.168.10.1', 8889)  #追加
    recieve_message = {"time":-1,"x":0,"y":2,"z":0,"shape":0}
    try:
        # 待ち受けポートに割り当て
        tello.bind(locaddr)
        server.bind(('0.0.0.0', 1323))
        
        while True:
            # 待ち受け開始
            server.listen(5)

            # 要求がいたら受け付け
            client, addr = server.accept()

            #受け取ったメッセージを出力
            #データ構造は辞書 ex. {"Time":1560232730357,"X":0,"Y":2,"Z":0,"Shape":0}
            recieve_messege = client.recv(4096).decode()
            if recieve_massag['time'] < 0:
                continue
            #toは手の動きの相対的な変化を記録する
            to = []                       #追加↓
            to[0] =  recieve_message['x']
            to[1] =  recieve_message['y']
            to[2] =  recieve_message['z']
            to[3] =  recieve_message['time']
            #手がグーなら移動
            if recieve_message['shape']: 
                #time,x,y,z各々相対的な変化を記録
                time = to[3]-fr[3]
                x = max(min(to[0]-fr[0],500),-500)
                y = max(min(to[1]-fr[1],500),-500)
                z = max(min(to[2]-fr[2],500),-500)
                #xy平面内を直線で移動する距離
                forward = int(np.sqrt(x^2+y^2))
                #z方向に移動する距離
                updown = abs(z)
                volb = int((forward+updown)/time)
                #移動速度
                vol = max(min(speedb,100),10)
                #移動速度をセット
                sent = tello.sendto('speed vol'.encode(encoding="utf-8"),tello_addres)

                #xy平面上でドローンの向きを適切に変更
                if x >= 0: 
                    theta = int(np.arctan(x/y)*180/np.pi)
                    sent = tello.sendto('cw theta'.encode(encoding="utf-8"),tello_addres)
                else:
                    theta = int(np.arctan(x/y)*180/np.pi)
                    sent = tello.sendto('ccw theta'.encode(encoding="utf-8"),tello_addres)
                
                #最小で20cmまでしか移動できないので、前方向、上下方向の移動する回数を求めている
                forwardnum = forward/21
                updownnum = updown/21

                #前方向→上下方向とできるだけ滑らかに移動
                if z >= 0:
                    for i in raneg(0,forwardnum):
                        sent = tello.sendto('forward 21'.encode(encoding="utf-8"),tello_addres)
                        if (updown):
                            sent = tello.sendto('up 21'.encode(encoding="utf-8"),tello_addres)
                            updown -= 1
                else:
                    for i in raneg(0,forwardnum):
                        sent = ello.sendto('forward 21'.encode(encoding="utf-8"),tello_addres)
                        if (updown):
                            sent = tello.sendto('down 21'.encode(encoding="utf-8"),tello_addres)
                            updown -= 1
            
            #frの値を更新
            fr[0] = to[0]
            fr[1] = to[1]
            fr[2] = to[2]
            fr[3] = to[3] #追加↑
            print('{}'.format(recieve_messege))  # 返答
            # client.send(b"I am socket server...\n")
            client.close()
    except KeyboardInterrupt:
        server.close()
    finally:
        server.close()

if __name__ == "__main__":
    main()
