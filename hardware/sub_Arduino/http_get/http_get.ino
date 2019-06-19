#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>

const char SSID[] = "aterm-602fae-g";
const char PASSWORD[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
String body;

void setup() {
  Serial.begin(115200);
  while (!Serial);

  WiFi.begin(SSID, PASSWORD);
  Serial.print("WiFi connecting");

  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    delay(100);
  }

  Serial.println(" connected");
  
  Wire.begin();
  Wire.setClock(400000); // use 400 kHz I2C

  pinMode(15,OUTPUT);
}

void loop() {
  HTTPClient http;
  http.begin(URL);
  int httpCode = http.GET();

  if (httpCode == HTTP_CODE_OK) {
    body = http.getString();
    Serial.print("Response Body: ");
    Serial.println(body);
  }
  
  if(body.toInt()>200){
    digitalWrite(15,HIGH);
    }else{
    digitalWrite(15,LOW);
    }
  delay(500);
}
