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
  delay(5000);
}

void right(){
  analogWrite(14,255);
  delay(30);
  for(int i=200;i>100;i--){
    analogWrite(14,i);
    analogWrite(12,300-i);
    delay(1);
  }
  analogWrite(12,255);
  delay(50);
}
