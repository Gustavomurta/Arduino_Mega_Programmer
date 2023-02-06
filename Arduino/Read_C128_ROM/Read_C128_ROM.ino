
/* Programador EPROM - Arduino MEGA - Commodore C128 ROM
   Arduino IDE 2.0.3
   Gustavo Murta   2022/12/28
   atualizado em 2023/02/05

   PORT A = Memory Adress A0 to A7
   PORT C = Memory Adress A8 to A15
   PORT L = Data Memory D0 to D7
   PORT B = A16 to A18 and VPP
   PG0 = -OE  /  PG1 = -CE

*/

int address0F = 0;      // endereço 00 a 0F
int addressF0 = 0;      // endereço 00 a F0
word addressFF = 0;     // endreço 0000 a FF00
int atraso = 5;         // atraso de 1 us

void setup() {

  DDRA = 0xFF;  // configure PORTA pins as output - Memory Adress A0 to A7
  DDRC = 0xFF;  // configure PORTC pins as output - Memory Adress A8 to A15
  DDRB = 0x0F;  // configure PORT B0 to B3 pins as output
  DDRG = 0x0F;  // configure PORT G0 to G3 pins as output
  DDRL = 0x00;  // configure PORTL pins as input - Memory Data D0 to D7

  PORTA = 0x00;  // SET Memory address A0 to A7 = 0
  PORTC = 0xC0;  // SET Memory address A8 to A13 = 0 / A15 e A14 = 1
  PORTB = 0x08;  // SET A16 to A18 = 0 and VPP = 1
  PORTG = 0x0F;  // SET -CE(PG1) and -OE(PG0) = 1

  Serial.begin(115200);  // Monitor serial 115200 bps
  // Serial.println("Lendo C128 ROM");  // print

  printEPROM();
}

void printEPROM() {

  // Serial.write(0);   //  caracter nulo

  for (addressFF = 0; addressFF < 16384; addressFF = addressFF + 256) {  // incrementa as páginas em 256 Bytes

    for (addressF0 = 0; addressF0 < 256; addressF0 = addressF0 + 16) {   // incrementa as páginas em 16 Bytes

      for (address0F = 0; address0F < 16; address0F++)  // endereços 0 a 15
      {
        PORTC = (addressFF >> 8) | 0xC0;    // shift right 8 bits - Memory address A8 to A13 / bitwise OR A14 e A15 = 1
        PORTA = addressF0 + address0F;      // set Memory Adress A0 to A7
        delayMicroseconds(atraso);          // atraso de X us
        PORTG = 0x00;                       // SET -CE(PG1) = 0 and -OE(PG0) = 0
        delayMicroseconds(atraso);          // atraso de X us
        Serial.write(PINL);                 // read Data Memory - D0 to D7
        PORTG = 0x01;                       // SET -CE(PG1) = 0 and -OE(PG0) = 1
        delayMicroseconds(atraso);          // atraso de X us
        PORTA = 0x00;                       // set Memory Adress A0 to A7 = 0
        PORTC = 0xC0;                       // set Memory Adress A8 to A13 = 0 / A14 e A15 = 1
      }
    }
  }
  PORTA = 0x00;  // set Memory Adress A0 to A7 = 0
  PORTC = 0x00;  // set Memory Adress A8 to A15 = 0
  PORTG = 0x0F;  // SET -OE and -CE = 1
}

void loop() {
}
