
/* Programador EPROM - Arduino MEGA - 27C020
   Arduino IDE 2.0.3
   Gustavo Murta   2022/12/28
   atualizado em 2023/01/12

   PORT A = Memory Adress A0 to A7
   PORT C = Memory Adress A8 to A15
   PORT L = Data Memory D0 to D7
   PORT B = A16 to A18 and VPP
   PG0 = /OE  /  PG1 = /CE

*/

int address0F = 0;   // endereço 00 a 0F
int addressF0 = 0;   // endereço 00 a F0
long addressFF = 0;  // endreço 0000 a FF00
int atraso = 1;      // atraso de 1 us

void setup() {

  Serial.begin(115200);  // Monitor serial 115200 bps
  Serial.println();
  Serial.println("Lendo a EPROM 27C020");  // print
  Serial.println();

  DDRA = 0xFF;  // configure PORTA pins as output - Memory Adress A0 to A7
  DDRC = 0xFF;  // configure PORTC pins as output - Memory Adress A8 to A15
  DDRB = 0x0F;  // configure PORT B0 to B3 pins as output
  DDRG = 0x0F;  // configure PORT G0 to G3 pins as output
  DDRL = 0x00;  // configure PORTL pins as input - Memory Data D0 to D7

  PORTG = 0x0F;  // SET /OE and /CE = 1
  PORTA = 0x00;  // SET Memory address A0 to A7 = 0
  PORTC = 0x00;  // SET Memory address A8 to A15 = 0
  PORTB = 0x08;  // SET A16 to A18 = 0 and VPP = 1

  printEPROM();
}

void printEPROM() {

  Serial.println("      00 01 02 03 04 05 06 07   08 09 0A 0B 0C 0D 0E 0F");  // print
  PORTG = 0x01;                                                          // SET /CE = 0 and OE = 1
  for (addressFF = 0; addressFF < 65536; addressFF = addressFF + 256) {  // incrementa as páginas em 256 Bytes

    for (addressF0 = 0; addressF0 < 256; addressF0 = addressF0 + 16)  // incrementa as páginas em 16 Bytes
    {
      if (addressFF < 0x1000) {
        Serial.print("0");
      }
      if (addressFF < 0x100) {
        Serial.print("0");
      }
      dataByte(addressFF + addressF0);  // imprime o endereço da linha de memoria com 16 bytes
      Serial.print(")");

      for (address0F = 0; address0F < 8; address0F++)  // endereços 0 a 7
      {
        PORTC = addressFF >> 8;         // shift right 8 bits - Memory address A8 to A15
        PORTA = addressF0 + address0F;  // set Memory Adress A0 to A7
        PORTG = 0x00;                   // SET /CE = 0 and OE = 0
        Serial.print(" ");
        dataByte(readEPROM(PORTA));  // read Data Memory - D0 to D7
        // Serial.print(char(readEPROM(PORTA)));  // print ASCII Data
        // dataByte(PORTA);         // teste
        PORTA = 0x00;               // set Memory Adress A0 to A7 = 0
        PORTC = 0x00;               // set Memory Adress A8 to A15 = 0
        delayMicroseconds(atraso);  // atraso de X us
        PORTG = 0x01;               // SET /CE = 0 and OE = 1
      }

      Serial.print(" -");

      for (address0F = 8; address0F < 16; address0F++)  // endereços 8 a 15
      {
        PORTC = addressFF >> 8;         // shift right 8 bits - Memory address A8 to A15
        PORTA = addressF0 + address0F;  // set Memory Adress A0 to A7
        PORTG = 0x00;                   // SET /CE = 0 and OE = 0
        Serial.print(" ");
        dataByte(readEPROM(PORTA));  // read Data Memory - D0 to D7
        // Serial.print(char(readEPROM(PORTA)));  // print ASCII Data
        // dataByte(PORTA);         // teste
        PORTA = 0x00;               // set Memory Adress A0 to A7 = 0
        PORTC = 0x00;               // set Memory Adress A8 to A15 = 0
        delayMicroseconds(atraso);  // atraso de X us
        PORTG = 0x01;               // SET /CE = 0 and OE = 1
      }
      Serial.println();
    }
  }
  PORTA = 0x00;  // set Memory Adress A0 to A7 = 0
  PORTC = 0x00;  // set Memory Adress A8 to A15 = 0
  PORTG = 0x0F;  // SET /OE and /CE = 1
}


void dataByte(unsigned int X) {
  if (X < 16) Serial.print("0");  // para imprimir Zero a esquerda do Hexadecimal
  Serial.print(X, HEX);           // imprime dados em hexadecimais
}

byte readEPROM(unsigned int Address16bit) {
  byte epromData = 0x00;
  epromData = PINL;  // read data port L - Memory data D0 to D7
  return epromData;
}


void loop() {
}
