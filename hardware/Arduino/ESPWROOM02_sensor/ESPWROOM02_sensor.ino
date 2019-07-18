#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <Wire.h>
#include <VL53L0X.h>

VL53L0X sensor1;
VL53L0X sensor2;

const char SSID[] = "aterm-602fae-g";
const char PASSWORD[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
const int sensor1_PIN=12;
const int sensor2_PIN=14;
String body;
int value1;
int value2;
#define ADDRESS_DEFALT 0b0101001
int threshold=200;
int httpCode;
String requestBody;

void setup() {
  pinMode(sensor1_PIN,OUTPUT);
  pinMode(sensor2_PIN,OUTPUT);
  digitalWrite(sensor1_PIN,LOW);
  digitalWrite(sensor2_PIN,LOW);
  
  Serial.begin(115200);
  while (!Serial);
  WiFi.mode(WIFI_STA);
  WiFi.begin(SSID, PASSWORD);
  Serial.print("WiFi connecting");

  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    delay(100);
  }
  Serial.println(" connected");
  
  Wire.begin();
  Wire.setClock(400000); // use 400 kHz I2C

  digitalWrite(sensor1_PIN,HIGH);
  sensor1.init();
  sensor1.setTimeout(500);
  if (!sensor1.init())
  {
    Serial.println("Failed to detect and initialize sensor!");
    while (1);
  }
  sensor1.setMeasurementTimingBudget(50000);
  sensor1.startContinuous();
  sensor1.setAddress(ADDRESS_DEFALT + 100002);

  delay(100);
  digitalWrite(sensor2_PIN,HIGH);
  sensor2.init();
  sensor2.setTimeout(500);
  if (!sensor2.init())
  {
    Serial.println("Failed to detect and initialize sensor!");
    while (1);
  }
  sensor2.setMeasurementTimingBudget(50000);
  sensor2.startContinuous();
  sensor2.setAddress(ADDRESS_DEFALT + 100008);
  
}

void loop() {
  HTTPClient http1;
  http1.begin(URL);

  //1
  value1 = sensor1.readRangeSingleMillimeters();
  requestBody = "1,"+String(value1);
  http1.POST(requestBody);
  http1.end();

  delay(200);
  
  HTTPClient http2;
  http2.begin(URL);
  //2
  value2 = sensor2.readRangeSingleMillimeters();
  requestBody = "2,"+String(value2);
  http2.POST(requestBody);
  http2.end();
  
  delay(200);
}
