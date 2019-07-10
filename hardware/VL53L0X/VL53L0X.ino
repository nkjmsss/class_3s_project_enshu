#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>
#include <VL53L0X.h>

VL53L0X sensor;

const char SSID[] = "aterm-602fae-g";
const char PASSWORD[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
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

  pinMode(26,OUTPUT);
  pinMode(32,OUTPUT);
  
  pinMode(34,OUTPUT);
  pinMode(19,OUTPUT);

  pinMode(25,OUTPUT);
  pinMode(23,OUTPUT);
}

void loop() {
  digitalWrite(26,LOW);
  digitalWrite(32,HIGH);
  digitalWrite(34,LOW);
  digitalWrite(19,LOW);
  digitalWrite(25,LOW);
  digitalWrite(23,LOW);
  
  HTTPClient http;
  http.begin(URL);
  value = sensor.readRangeContinuousMillimeters();
  requestBody = (String)value;
  httpCode = http.POST(requestBody);
  
  Serial.printf("Response: %d", httpCode);
  Serial.println();
  if (httpCode == HTTP_CODE_OK) {
    body = http.getString();
    Serial.print("Response Body: ");
    Serial.println(body);
  }
  delay(500);
}
