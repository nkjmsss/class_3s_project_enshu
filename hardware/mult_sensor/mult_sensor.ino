#include <HTTPClient.h>

#include <WiFi.h>
#include <Wire.h>
#include <VL53L0X.h>

VL53L0X sensor;

const char SSID[] = "aterm-602fae-g";
const char PASSWORD[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";

const int sensor1=26;
const int sensor2=32;
const int sensor3=34;
const int sensor4=19;
const int sensor5=25;
const int sensor6=23;
String body;
int value;

int threshold=200;
int httpCode;
String requestBody;

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

  sensor.init();
  sensor.setTimeout(500);
  if (!sensor.init())
  {
    Serial.println("Failed to detect and initialize sensor!");
    while (1);
  }
  sensor.setMeasurementTimingBudget(50000);
  sensor.startContinuous();
  
  pinMode(sensor1,OUTPUT);
  pinMode(sensor2,OUTPUT);
  pinMode(sensor3,OUTPUT);
  pinMode(sensor4,OUTPUT);
  pinMode(sensor5,OUTPUT);
  pinMode(sensor6,OUTPUT);
}

void loop() {
  HTTPClient http;
  http.begin(URL);

  //1
  digitalWrite(sensor1,HIGH);
  digitalWrite(sensor2,LOW);
  digitalWrite(sensor3,LOW);
  digitalWrite(sensor4,LOW);
  digitalWrite(sensor5,LOW);
  digitalWrite(sensor6,LOW);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = "1"+String(value);
  http.POST(requestBody);

  //2
  digitalWrite(sensor1,LOW);
  digitalWrite(sensor2,HIGH);
  digitalWrite(sensor3,LOW);
  digitalWrite(sensor4,LOW);
  digitalWrite(sensor5,LOW);
  digitalWrite(sensor6,LOW);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = "2"+String(value);
  http.POST(requestBody);

  //3
  digitalWrite(sensor1,LOW);
  digitalWrite(sensor2,LOW);
  digitalWrite(sensor3,HIGH);
  digitalWrite(sensor4,LOW);
  digitalWrite(sensor5,LOW);
  digitalWrite(sensor6,LOW);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = "3"+String(value);
  http.POST(requestBody);
  
  //4
  digitalWrite(sensor1,LOW);
  digitalWrite(sensor2,LOW);
  digitalWrite(sensor3,LOW);
  digitalWrite(sensor4,HIGH);
  digitalWrite(sensor5,LOW);
  digitalWrite(sensor6,LOW);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = "4"+String(value);
  http.POST(requestBody);

  //5
  digitalWrite(sensor1,LOW);
  digitalWrite(sensor2,LOW);
  digitalWrite(sensor3,LOW);
  digitalWrite(sensor4,LOW);
  digitalWrite(sensor5,HIGH);
  digitalWrite(sensor6,LOW);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = "5"+String(value);
  http.POST(requestBody);


  //6
  digitalWrite(sensor1,LOW);
  digitalWrite(sensor2,LOW);
  digitalWrite(sensor3,LOW);
  digitalWrite(sensor4,LOW);
  digitalWrite(sensor5,LOW);
  digitalWrite(sensor6,HIGH);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = "6"+String(value);
  http.POST(requestBody);
  
  delay(500);
}
