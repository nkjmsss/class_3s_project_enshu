#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>
#include <VL53L1X.h>

VL53L1X sensor1;
VL53L1X sensor2;
VL53L1X sensor3;
VL53L1X sensor4;
VL53L1X sensor5;
VL53L1X sensor6;

const char SSID[] = "aterm-602fae-g";
const char PASSWORD[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
String body;
int value1;
int value2;
int value3;
int value4;
int value5;
int value6;

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

  sensor1.setTimeout(500);
  if (!sensor1.init())
  {
    Serial.println("Failed to detect and initialize sensor1!");
    while (1);
  }
  sensor1.setDistanceMode(VL53L1X::Long);
  sensor1.setMeasurementTimingBudget(50000);
  sensor1.startContinuous(50);

  sensor2.setTimeout(500);
  if (!sensor2.init())
  {
    Serial.println("Failed to detect and initialize sensor2!");
    while (1);
  }
  sensor2.setDistanceMode(VL53L1X::Long);
  sensor2.setMeasurementTimingBudget(50000);
  sensor2.startContinuous(50);

  sensor3.setTimeout(500);
  if (!sensor3.init())
  {
    Serial.println("Failed to detect and initialize sensor3!");
    while (1);
  }
  sensor3.setDistanceMode(VL53L1X::Long);
  sensor3.setMeasurementTimingBudget(50000);
  sensor3.startContinuous(50);

  sensor4.setTimeout(500);
  if (!sensor4.init())
  {
    Serial.println("Failed to detect and initialize sensor4!");
    while (1);
  }
  sensor4.setDistanceMode(VL53L1X::Long);
  sensor4.setMeasurementTimingBudget(50000);
  sensor4.startContinuous(50);

  sensor5.setTimeout(500);
  if (!sensor5.init())
  {
    Serial.println("Failed to detect and initialize sensor5!");
    while (1);
  }
  sensor5.setDistanceMode(VL53L1X::Long);
  sensor5.setMeasurementTimingBudget(50000);
  sensor5.startContinuous(50);

  sensor6.setTimeout(500);
  if (!sensor6.init())
  {
    Serial.println("Failed to detect and initialize sensor6!");
    while (1);
  }
  sensor6.setDistanceMode(VL53L1X::Long);
  sensor6.setMeasurementTimingBudget(50000);
  sensor6.startContinuous(50);

  pinMode(15,OUTPUT);
}

void loop() {
  HTTPClient http;
  http.begin(URL);
  value1 = sensor1.read();
  value2 = sensor2.read();
  value3 = sensor3.read();
  value4 = sensor4.read();
  value5 = sensor5.read();
  value6 = sensor6.read();
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
