EPROM Reader 27C1024 - Arduino MEGA 2560

Tinygo - version: 0.30.0  - https://tinygo.org/getting-started/install/

To compile and flash the TinyGo program:
* go mod init main.go
* go mod tidy
* tinygo flash -target=arduino-mega2560 main.go


```
  
C:\Users\jgust\tinygo\programas\arduino\Eprom_reader_27C1024\versão2>tinygo flash -target=arduino-mega2560 main.go

avrdude.exe: Version 7.0
             Copyright (c) Brian Dean, http://www.bdmicro.com/
             Copyright (c) Joerg Wunsch

             System wide configuration file is C:/Users/jgust/scoop/apps/avr-gcc/current/bin/avrdude.conf

             Using Port                    : COM15
             Using Programmer              : wiring
             Overriding Baud Rate          : 115200
             AVR Part                      : ATmega2560
             Chip Erase delay              : 9000 us
             PAGEL                         : PD7
             BS2                           : PA0
             RESET disposition             : dedicated
             RETRY pulse                   : SCK
             Serial program mode           : yes
             Parallel program mode         : yes
             Timeout                       : 200
             StabDelay                     : 100
             CmdexeDelay                   : 25
             SyncLoops                     : 32
             PollIndex                     : 3
             PollValue                     : 0x53
             Memory Detail                 :

                                               Block Poll               Page                       Polled
               Memory Type Alias    Mode Delay Size  Indx Paged  Size   Size #Pages MinW  MaxW   ReadBack
               ----------- -------- ---- ----- ----- ---- ------ ------ ---- ------ ----- ----- ---------
               eeprom                 65    10     8    0 no       4096    8      0  9000  9000 0x00 0x00
               flash                  65    10   256    0 yes    262144  256   1024  4500  4500 0x00 0x00
               lfuse                   0     0     0    0 no          1    1      0  9000  9000 0x00 0x00
               hfuse                   0     0     0    0 no          1    1      0  9000  9000 0x00 0x00
               efuse                   0     0     0    0 no          1    1      0  9000  9000 0x00 0x00
               lock                    0     0     0    0 no          1    1      0  9000  9000 0x00 0x00
               calibration             0     0     0    0 no          1    1      0     0     0 0x00 0x00
               signature               0     0     0    0 no          3    1      0     0     0 0x00 0x00

             Programmer Type : Wiring
             Description     : Wiring
             Programmer Model: AVRISP
             Hardware Version: 15
             Firmware Version Master : 2.10
             Vtarget         : 0.0 V
             SCK period      : 0.1 us

avrdude.exe: AVR device initialized and ready to accept instructions

Reading | ################################################## | 100% 0.01s

avrdude.exe: Device signature = 0x1e9801 (probably m2560)
avrdude.exe: reading input file C:\Users\jgust\AppData\Local\Temp\tinygo3562371359\main.hex
avrdude.exe: writing flash (3220 bytes):

Writing | ################################################## | 100% 0.53s

avrdude.exe: 3220 bytes of flash written
avrdude.exe: verifying flash memory against C:\Users\jgust\AppData\Local\Temp\tinygo3562371359\main.hex:

Reading | ################################################## | 100% 0.43s

avrdude.exe: 3220 bytes of flash verified

avrdude.exe done.  Thank you.


C:\Users\jgust\tinygo\programas\arduino\Eprom_reader_27C1024\versão2>
```

