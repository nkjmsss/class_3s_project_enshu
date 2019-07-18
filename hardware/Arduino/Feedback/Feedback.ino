#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

char ssid[] = "aterm-602fae-g";
char password[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
//int httpCode;
String requestBody;
int value;
int threshold=500;  
const int delay_time = 200;
HTTPClient http;
String body;

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

  pinMode(12,OUTPUT);
  pinMode(14,OUTPUT);
  http.begin(URL);
}

void loop() {
  String cmds[2] = {"\0"};
  int httpCode = http.GET();
  if (httpCode == HTTP_CODE_OK) {
    
    body = http.getString();
    split(body,',',cmds);
    Serial.println(body);
    Serial.println(cmds[1]);
    if(cmds[1].toInt()<500){
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
    }
  }
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
