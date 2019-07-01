import socket
import sys
import json
import numpy as np

#kinectから受け取ったデータを保存
# [x,y,z,time]
fr = [0.0,0.0,0.0,0.0] #追加
    
def main():
    # 待ち受け用のソケットオブジェクト
    server = socket.socket()
    #tello用のソケットオブジジェクト
    host = ''
    port = 9000
    locaddr = (host,port)
    tello = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)   
    tello_address = ('192.168.10.1', 8889) 

    move = ['speed', 'cw','ccw','forward 21','up 21','down 21','takeoff','land']

    try:
        # 待ち受けポートに割り当て
        tello.bind(locaddr)
        server.bind(('0.0.0.0', 1323))
        
        while True:
            # 待ち受け開始
            server.listen(5)

            # 要求がいたら受け付け
            client, addr = server.accept()
            
            #データ構造は辞書 ex. {"Time":1560232730357,"X":0,"Y":2,"Z":0,"Shape":0}
            
            try:
                recieve_message_json = client.recv(4096).decode()
                recieve_message = json.loads(recieve_message_json)
                #recieve_message = json.loads(recieve_message_json[94:])
                #print(type(recieve_message))

                print('{}'.format(recieve_message))  # 返答
            
                #離陸するならtakeffland = 6, 着陸するならtakeoffland = 7
                takeoffland = -1

                if recieve_message['command'] == 1:
                    takeoffland = 6
                elif recieve_message['command'] == 2:
                    takeoffland = 7
            
                #toは手の動きの相対的な変化を記録する
                to = [-1,-1,-1,-1]                       
                to[0] =  recieve_message['x']
                to[1] =  recieve_message['y']
                to[2] =  recieve_message['z']
                to[3] =  recieve_message['time']

                try:
                    sent = self.sock.sendto('command'.encode(encoding="utf-8"), self.tello)
                    print("OK")
                except:
                    print("No")
                    pass

            
                #手がグーなら移動
                if recieve_message['shape'] == 0: 
                    #time,x,y,z各々相対的な変化を記録
                    time = to[3]-fr[3]
                    x = max(min(to[0]-fr[0],500),-500)
                    y = max(min(to[1]-fr[1],500),-500)
                    z = max(min(to[2]-fr[2],500),-500)
                    #xy平面内を直線で移動する距離
                    forward = int(np.sqrt(x*x+y*y))
                    #z方向に移動する距離
                    updown = abs(z)
                    #移動速度
                    volb = int((forward+updown)/(time+2))
                    vol = max(min(volb,100),10)
                
                    #回転角度
                    theta = ''
                    #時計回りなら1,反時計回りなら2
                    rotate = -1

                    #theta,rotateを決定
                    if x >= 0: 
                        if y==0:
                            theta = 90
                        else:
                            theta = int(np.arctan(x/y)*180/np.pi)
                        rotate = 1
                    else:
                        if y==0:
                            theta = 90
                        else:
                            theta = int(np.arctan(x/y)*180/np.pi)
                        rotate = 2

                
                    #最小で20cmまでしか移動できないので、前方向、上下方向の移動する回数を求めている
                    forwardnum = int(forward+30/21)
                    updownnum = int(updown+30/21)

                    #上昇するなら4,下降するなら5
                    up_down = -1

                    #updownを求める
                    if z >= 0:
                        up_down = 4
                    else:
                        up_down = 5

                    #離着陸
                    if takeoffland > 0:
                        sent = tello.sendto(move[takeoffland].encode("utf-8"),tello_address)
                        print(move[takeoffland]);

                    #向きを変える 
                    sent = tello.sendto('speed 100'.encode("utf-8"),tello_address)
                    sent = tello.sendto((move[rotate]+' '+str(theta)).encode("utf-8"),tello_address)
                    print(move[rotate],end = ' ')
                    print(theta)
                
                    #スピードを設定
                    sent = tello.sendto(('speed '+str(vol)).encode("utf-8"),tello_address)
                    print(move[0],end = ' ')
                    print(vol)

                    #移動
                    print(move[3],end = ' ')
                    print(move[up_down])
                    for i in range(0,forwardnum):
                        recieve_message_json = client.recv(4096).decode()
                        if (updownnum):
                            sent = tello.sendto(move[up_down].encode("utf-8"),tello_address)
                            updownnum -= 1
                    print('-----')
                
                
            
                #frの値を更新

                fr[0] = to[0]
                fr[1] = to[1]
                fr[2] = to[2]
                fr[3] = to[3] #追加↑
           
                # client.send(b"I am socket server...\n")
                client.close()
            except Exception as e:
                print(e)
    except KeyboardInterrupt:
        server.close()
    finally:
        server.close()

if __name__ == "__main__":
    main()