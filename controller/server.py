import socket
import sys
import json
import numpy as np

#kinectから受け取ったデータを保存
# [x,y,z,time]
fr = [0.0, 0.0, 0.0, 0.0, 0]  #追加


def main():
    # 待ち受け用のソケットオブジェクト
    server = socket.socket()
    #tello用のソケットオブジジェクト
    host = ''
    port = 9000
    locaddr = (host, port)
    tello = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    tello_address = ('192.168.10.1', 8889)

    move = [
        'speed', 'cw', 'ccw', 'forward', 'back', 'left', 'right', 'up', 'down',
        'takeoff', 'land'
    ]

    try:
        # 待ち受けポートに割り当て
        tello.bind(locaddr)
        server.bind(('0.0.0.0', 1324))
        # 待ち受け開始
        server.listen(5)
        print('Server is listening on controller:1324')
        ready = 0
        land = 0
        cur = [0,0,0]

        while True:

            # 要求がいたら受け付け
            client, addr = server.accept()

            #データ構造は辞書 ex. {"Time":1560232730357,"X":0,"Y":2,"Z":0,"Shape":0}

            try:
                # send command to tello
                # also, send to client
                def sendTello(message):
                    tello.sendto('command'.encode(encoding="utf-8"), tello_address)
                    tello.sendto(message.encode(encoding="utf-8"), tello_address)
                    print(message)
                    client.sendall((message + '\n').encode('utf-8'))

                recieve_message_json = client.recv(4096).decode()
                recieve_message = json.loads(recieve_message_json)
                #print(type(recieve_message))

                #print('{}'.format(recieve_message))  # 返答

                takeoffland = -1

                if recieve_message['command'] == 1:
                    takeoffland = 9
                elif recieve_message['command'] == 2:
                    takeoffland = 10
                    land = 1

                if takeoffland > 0:
                    sendTello(move[takeoffland])


                #toは手の動きの相対的な変化を記録する
                to = [-1, -1, -1, -1,0]
                to[0] = recieve_message['x']
                to[1] = recieve_message['y']
                to[2] = recieve_message['z']
                to[3] = recieve_message['time']
                to[4] = recieve_message['shape']
                #手がグーなら移動
                if recieve_message['shape'] != 3:
                    ready = 0
                if recieve_message['shape'] == 3:
                    #time,x,y,z各々相対的な変化を記録
                    cmd = -1
                    dis = -1
                    if ready:
                        dx = int((to[0] - fr[0])*0.01)
                        dy = int((to[1] - fr[1])*0.005)
                        dz = int((to[2] - fr[2])*0.01)
                        dis = np.sqrt(dx**2+dy**2+dz**2)
                        dt = to[3] - fr[3]
                        vol = int(dis*0.5/dt)
                        vol = min(100,vol)
                        vol = max(30,vol)
                        x = max(min(cur[0]+dx,500),-500)
                        y = max(min(cur[1]-dz,500),-500)
                        z = max(min(cur[2]+dy,500),-500)
                        if land != 1 and fr[4] == 3: 
                            sendTello('go'+' '+str(x)+' '+str(y)+' '+str(z)+' '+str(vol))
                            cur[0] = x
                            cur[1] = y
                            cur[2] = z
                    else:
                        ready = 1
                #frの値を更新
                fr[0] = to[0]
                fr[1] = to[1]
                fr[2] = to[2]
                fr[3] = to[3] 
                fr[4] = to[4]

                # client.send(b"I am socket server...\n")
                client.close()
            except Exception as e:
                print(e)
    except KeyboardInterrupt:
        print('keyboard interrupt')
        server.close()
    finally:
        print('finish')
        server.close()


if __name__ == "__main__":
    main()
