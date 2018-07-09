// pin attached to measurement device
const int measurePin = A0;

void setup() {
  Serial.begin(9600);
  pinMode(measurePin, INPUT);
}


void do_sleep() {
  // do 10 measurements every second
  delay(1000/10);
}

// volts_per_a is the ACS712 resolution, in this case 185mV/A
void do_measurement(float volts_per_a) {
  // ACS712 returns a value on analog input with the reading from 
  // the hall effect sensor
  //
  // this puts the 0-1024 value in the 0-5000 mV range
  int volts = (analogRead(measurePin)/1024.0)*5000;
  // 2500 is the 0 value (no current passing)
  // use this to calculate the mAmps value for the 
  // last measurement
  int mAmps = (volts - 2500)/volts_per_a;

  if (mAmps < 0) {
    // return a positive value regardless of current direction
    mAmps *= -1;
  }

  Serial.print(mAmps);
  Serial.println();
}

void loop() {
  do_measurement(0.185);
  do_sleep();
}
