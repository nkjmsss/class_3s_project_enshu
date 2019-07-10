void setup() {
  // put your setup code here, to run once:
  pinMode(14,OUTPUT);
  pinMode(12,OUTPUT);
  pinMode(13,OUTPUT);
  pinMode(15,OUTPUT);
  pinMode(2,OUTPUT);
  pinMode(0,OUTPUT);

  digitalWrite(14,LOW);
  digitalWrite(12,LOW);
  digitalWrite(13,LOW);
  digitalWrite(15,LOW);
  digitalWrite(2,LOW);
  digitalWrite(0,LOW);
}

void loop() {
  // put your main code here, to run repeatedly:
  right();
  delay(200);
}

void right(){
  digitalWrite(14,HIGH);
  delay(200);
  digitalWrite(12,HIGH);
  delay(200);
  digitalWrite(14,LOW);
  delay(150);
  digitalWrite(12,LOW);  
}

void top(){
  digitalWrite(15,HIGH);
  delay(200);
  digitalWrite(13,HIGH);
  delay(200);
  digitalWrite(15,LOW);
  delay(150);
  digitalWrite(13,LOW);  
}
