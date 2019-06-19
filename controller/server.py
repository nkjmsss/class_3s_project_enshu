import socket
import sys
import json
import numpy as np


#kinectから受け取ったデータを保存
# [x,y,z,time]
fr = [0.0,0.0,0.0,0.0] #追加


#離着陸を判断するオートマトン
class takeofflandAutomata():
    def __init__(self,states,alphabets,transitions,current_state,final_state):
        self.states = states
        self.alphabets = alphabets
        self.transitions = transitions
        self.current_state = current_state
        self.final_state = final_state
    def run(self,shape):
        self.current_state = self.transitions[self.current_state][shape]
    



def main():
    # 待ち受け用のソケットオブジェクト
    server = socket.socket()
    #tello用のソケットオブジジェクト
    host = ''
    port = 9000
    locaddr = (host,port)
    tello = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)   
    tello_address = ('192.168.10.1', 8889) 

    #離着陸オートマトンの引数
    states = {"a","b","c","d","e","f","g"}
    alphabets = {0,1,2,3,4,5,6,7,8,9}
    transitions = {
        "a":{0:"b",1:"a",2:"a",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
        "b":{0:"b",1:"c",2:"a",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
        "c":{0:"a",1:"c",2:"d",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
        "d":{0:"e",1:"a",2:"d",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
        "e":{0:"e",1:"f",2:"a",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
        "f":{0:"a",1:"e",2:"g",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
        "g":{0:"a",1:"a",2:"a",3:"a",4:"a",5:"a",6:"a",7:"a",8:"a",9:"a"},
    }
    current_state = "a"
    final_state = "g"

    #takeofflandAutomata classの変数
    takeoff = takeofflandAutomata(states,alphabets,transitions,current_state,final_state)
    land = takeofflandAutomata(states,alphabets,transitions,current_state,final_state)

    #空中にいるならsky = 1,地上にいるならsky=0
    sky = 0

    #ドローンを操作するコマンド
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
            recieve_message_json = client.recv(4096).decode()
            #print('----')
            #print(recieve_message_json[94:])
            #print('----')
            recieve_message = json.loads(recieve_message_json[94:])
            #print(type(recieve_message))
            
            #離陸するならtakeffland = 6, 着陸するならtakeoffland = 7
            takeoffland = -1

            #地上にいるなら
            if sky == 0:
                #手の形を読み取ってオートマトン遷移
                takeoff.run(recieve_message['shape'])
                #離陸
                if (takeoff.current_state == takeoff.final_state):
                    takeoffland = 6
                    sky = 1
            #空中にいるなら
            else:
                #手の形を読み取ってオートマトン遷移
                land.run(recieve_message['shape'])
                #着陸
                if (land.current_state == land.final_state):
                    takeoffland = 7
                    sky = 0
            
            #toは手の動きの相対的な変化を記録する
            to = [-1,-1,-1,-1]                       
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
                forward = int(np.sqrt(x*x+y*y))
                #z方向に移動する距離
                updown = abs(z)
                #移動速度
                volb = int((forward+updown)/time)
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
                forwardnum = int(forward/21)
                updownnum = int(updown/21)

                #上昇するなら4,下降するなら5
                up_down = -1

                #updownを求める
                if z >= 0:
                    up_down = 4
                else:
                    up_down = 5
                
                
                #離着陸
                if takeoffland > 0:
                    sent = tello.sendto('move[takeoffland]'.encode(encoding="utf-8"),tello_address)
                    print(move[takeoffland]);

                #向きを変える 
                sent = tello.sendto('move[0] 100'.encode(encoding="utf-8"),tello_address)
                sent = tello.sendto('move[rotate] theta'.encode(encoding="utf-8"),tello_address)
                print(move[rotate],end = ' ')
                print(theta)
                
                #スピードを設定
                sent = tello.sendto('move[0] vol'.encode(encoding="utf-8"),tello_address)
                print(move[0],end = ' ')
                print(vol)

                #移動
                print(move[3],end = ' ')
                print(move[up_down])
                for i in range(0,forwardnum):
                    sent = tello.sendto('move[3]'.encode(encoding="utf-8"),tello_address)
                    if (updownnum):
                        sent = tello.sendto('move[updown]'.encode(encoding="utf-8"),tello_address)
                        updownnum -= 1
                print('-----')
                
                
            
            #frの値を更新
            fr[0] = to[0]
            fr[1] = to[1]
            fr[2] = to[2]
            fr[3] = to[3] #追加↑
            #print('{}'.format(recieve_message))  # 返答
            # client.send(b"I am socket server...\n")
            client.close()
    except KeyboardInterrupt:
        server.close()
    finally:
        server.close()

if __name__ == "__main__":
    main()