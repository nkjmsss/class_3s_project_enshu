#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>
#include <VL53L1X.h>

VL53L1X sensor;

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

  sensor.setTimeout(500);
  if (!sensor.init())
  {
    Serial.println("Failed to detect and initialize sensor!");
    while (1);
  }
  sensor.setDistanceMode(VL53L1X::Long);
  sensor.setMeasurementTimingBudget(50000);
  sensor.startContinuous(50);

  pinMode(15,OUTPUT);
}

void loop() {
  HTTPClient http;
  http.begin(URL);
  String requestBody = (String)sensor.read();
  int httpCode = http.POST(requestBody);

  Serial.printf("Response: %d", httpCode);
  Serial.println();
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
