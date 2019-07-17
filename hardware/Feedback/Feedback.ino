#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

char ssid[] = "aterm-602fae-g";
char password[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
int httpCode;
String requestBody;
String body;
int value;
int threshold=200;

const int right_pin = 14;
const int left_pin = 12;
const int front_pin = 15;
const int back_pin = 13;
const int top_pin = 2;
const int bottom_pin = 0;

const int delay_time = 200;


String cmds[2] = {"\0"}; // 分割された文字列を格納する配列 

void split(String data, char delimiter, String *dst) {
    int index = 0;
    int arraySize = (sizeof(data)/sizeof((data)[0]));  
    int datalength = data.length();
    for (int i = 0; i < datalength; i++) {
        char tmp = data.charAt(i);
        if ( tmp == delimiter ) {
            index++;
        }
        else dst[index] += tmp;
    }
}

void setup() {
  Serial.begin(115200);
  Serial.println("");

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    delay(100);
  }
  Serial.println("connected!");

  pinMode(2,OUTPUT);
}

void loop() {
  HTTPClient http;
  http.begin(URL);
  httpCode = http.GET();
  if (httpCode == HTTP_CODE_OK) {
    body = http.getString();
    split(body,',',cmds);
    if(cmds[1].toInt()<100){
      if(cmds[0]=="1"){
        right();
        delay(200);
        right();
        }
      else if(cmds[0]=="2"){
        left();
        delay(200);
        left();
        }
      else if(cmds[0]=="3"){
        front();
        delay(200);
        front();
        }
      else if(cmds[0]=="4"){
        back();
        delay(200);
        back();
        }
      else if(cmds[0]=="5"){
        top();
        delay(200);
        top();
        }
      else if(cmds[0]=="6"){
        bottom();
        delay(200);
        bottom();
        }
    }else{
    }
  }
  delay(100); 
}

void right() {
  digitalWrite(14,HIGH);
  delay(delay_time);
  digitalWrite(12,HIGH);
  delay(delay_time);
  digitalWrite(14,LOW);
  delay(delay_time);
  digitalWrite(12,LOW);  
}

void left() {
  digitalWrite(12,HIGH);
  delay(delay_time);
  digitalWrite(14,HIGH);
  delay(delay_time);
  digitalWrite(12,LOW);
  delay(delay_time);
  digitalWrite(14,LOW);  
}

void front() {
  digitalWrite(15,HIGH);
  delay(delay_time);
  digitalWrite(13,HIGH);
  delay(delay_time);
  digitalWrite(15,LOW);
  delay(delay_time);
  digitalWrite(13,LOW);  
}

void back() {
  digitalWrite(13,HIGH);
  delay(delay_time);
  digitalWrite(15,HIGH);
  delay(delay_time);
  digitalWrite(13,LOW);
  delay(delay_time);
  digitalWrite(15,LOW);  
}

void top() {
  digitalWrite(2,HIGH);
  delay(delay_time);
  digitalWrite(0,HIGH);
  delay(delay_time);
  digitalWrite(2,LOW);
  delay(delay_time);
  digitalWrite(0,LOW);  
}

void bottom() {
  digitalWrite(0,HIGH);
  delay(delay_time);
  digitalWrite(2,HIGH);
  delay(delay_time);
  digitalWrite(0,LOW);
  delay(delay_time);
  digitalWrite(2,LOW);  
}
