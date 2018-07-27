/*
    Main.ino
    システムメイン

    Created 2018/06/28
    By Nakajim Yam
*/
/* Include Files */
#include "ArduinoLibrary.h"
#include "ServerCommunication.h"
#include "function.h"
#include "constants.h"
#include "System.h"

Servo servo;          //Servo Lib
WiFiClient client;     //Wifi Instance
WIFICONFIG internet;

/* --- プロトタイプ宣言 --- */
void deviceInit( void );

void setup() 
{
  deviceInit();
  
  if(internet.device_id==NULL){
    registerDevice();
  }
}


SINT servo_angle=145;
void loop()
{
  if(digitalRead(BTN_PIN) == HIGH){
    //サーボを回転
    servo.write(servo_angle);
    //POST & PUT
    getAngle(&servo_angle);
    servo_angle = 180 - servo_angle;
    servo.write(servo_angle);
    delay(500);
    //ブザーを鳴らす
    digitalWrite(LED_PIN,LOW);
    servo_angle-=35;
 }
  delay(100);
}

void deviceInit()
{
  /* シリアル通信 */
  Serial.begin(115200);
  Serial.println();
  /* ピン設定 */
  pinMode(APSWT, INPUT );     /* push Button */
  pinMode(LED_PIN, OUTPUT);   /* LED         */
  pinMode(BTN_PIN, OUTPUT);   /* Button      */
  digitalWrite(LED_PIN, HIGH);/* LED ON      */
  /* サーボモータ設定*/
  servo.attach(2);
  servo.write(0);
  /* ファイルシステム */
  SPIFFS.begin();
  /* マイコンのモード設定
     設定ボタンが押されたか３秒間待つ*/
  delay(3000); //ボタン待ち
  digitalWrite(LED_PIN, LOW); /* LED OFF     */
  delay(100);
  if(digitalRead(APSWT)==LOW){
    setupWifi();
  } 
  else if(digitalRead(APSWT)==HIGH) {
    connectRouter();
  }
}

