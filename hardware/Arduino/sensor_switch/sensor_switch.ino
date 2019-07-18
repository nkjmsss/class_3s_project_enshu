#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>
#include <VL53L1X.h>

VL53L1X sensor;

const char SSID[] = "aterm-602fae-g";
const char PASSWORD[] = "48e04d7a669be";
const char URL[] = "http://192.168.10.2:8000/data";
String body;
int value;
int threshold=200;
int httpCode;
String requestBody;
int i=0;
bool state=true;

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
  pinMode(17,OUTPUT);
}

void loop() {
  digitalWrite(17,state);
  i++;
  if(i>10){
      state = !state;
      i=0;
  }
  HTTPClient http;
  http.begin(URL);
  value = sensor.read();
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
