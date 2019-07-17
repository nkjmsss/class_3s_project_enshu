# class_3s_project_enshu

- Kinect -> KinectForDroneControl -> Middleware -> Controller -> Drone
- Drone -> Hardware -> feedback

## KinectForDroneControl

- C++
- Kinectから人間の座標を取る

## Middleware

- Go
- Docker
- 複雑な処理等を書く

## Controller

- Python
- Docker
- TelloとのUDP通信が主なタスク

## Hardware

- Python
- Arduino
- Telloに搭載したセンサーの情報を受取り、フィードバックを返す
