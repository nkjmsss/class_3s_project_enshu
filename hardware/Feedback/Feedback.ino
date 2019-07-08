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
    if(body.toInt()<100){
      digitalWrite(2,HIGH);
    }else{
      digitalWrite(2,LOW);
    }
  }
  delay(500);
}
