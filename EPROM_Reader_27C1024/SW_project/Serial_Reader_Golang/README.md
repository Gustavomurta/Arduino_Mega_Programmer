This program captures bytes read from EPROM, dumps bytes to the console, and saves a file in binary format.

You must change the Arduino COM port of your PC in the go program.

Original 87F4794 EPROM BIN file (IBM PS/2 BIOS)- https://www.ardent-tool.com/firmware/system.html

go version go1.21.3 windows/amd64  - https://go.dev/doc/install

To compile the go program (Windows 10): 

* go mod init main.go
* go mod tidy
* go build main.go

To run this program : execute main.exe


```
C:\Users\jgust\go\programas\bin_format\EPROM serial reader 27C1024>dir
 O volume na unidade C não tem nome.
 O Número de Série do Volume é 709A-B58F

 Pasta de C:\Users\jgust\go\programas\bin_format\EPROM serial reader 27C1024

15/02/2024  18:51    <DIR>          .
15/02/2024  18:51    <DIR>          ..
29/09/2020  23:08           131.072 87F4794 original EPROM.BIN
29/09/2020  23:08           131.072 87F4794 PS1 2123.BIN
14/02/2024  23:21           131.072 EPROM_87F4794.bin
14/02/2024  20:56               139 go.mod
14/02/2024  20:56               372 go.sum
14/02/2024  22:30         2.129.920 main.exe
14/02/2024  23:24             1.972 main.go
               7 arquivo(s)      2.656.691 bytes
               2 pasta(s)   16.411.713.536 bytes disponíveis

C:\Users\jgust\go\programas\bin_format\EPROM serial reader 27C1024>
```

Running the Serial Reader program: 

```

C:\Users\jgust\go\programas\bin_format\EPROM serial reader 27C1024>main
2024/02/15 19:05:01  Read EPROM start
2024/02/15 19:05:14  EPROM Reading ends
00000000  00 00 83 ec 0e 55 8b ec  89 5e 04 8c 5e 02 2e 8e  |.....U...^..^...|
00000010  1e 00 00 8e 1e 0e 04 8a  1e ec 00 80 fa 80 74 0a  |..............t.|
00000020  3a d3 74 04 32 db eb 18  b3 80 86 d3 80 fc 15 74  |:.t.2..........t|
00000030  f3 80 fc 08 74 ee 3d 03  1f 74 e9 3d 02 1f 74 e4  |....t.=..t.=..t.|
00000040  88 5e 14 eb 12 55 8b ec  9c 80 7e 06 00 74 03 8a  |.^...U....~..t..|
00000050  56 06 9d 5d ca 02 00 c5  1e e8 00 c7 46 0a 45 00  |V..]........F.E.|
00000060  8c 4e 0c 89 5e 06 8c 5e  08 5d 1f 5b cb 02 30 31  |.N..^..^.].[..01|
00000070  2f 30 39 2f 39 32 39 39  36 36 46 46 37 37 34 34  |/09/929966FF7744|
00000080  39 39 32 33 20 20 28 28  43 43 29 29 20 20 43 43  |9923  ((CC))  CC|
00000090  4f 4f 50 50 59 59 52 52  49 49 47 47 48 48 54 54  |OOPPYYRRIIGGHHTT|
000000a0  20 20 49 49 42 42 4d 4d  20 20 43 43 4f 4f 52 52  |  IIBBMM  CCOORR|
000000b0  50 50 4f 4f 52 52 41 41  54 54 49 49 4f 4f 4e 4e  |PPOORRAATTIIOONN|
000000c0  20 20 31 31 39 39 38 38  31 31 2c 2c 20 20 31 31  |  11998811,,  11|
000000d0  39 39 39 39 32 32 20 20  41 41 4c 4c 4c 4c 20 20  |999922  AALLLL  |
000000e0  52 52 49 49 47 47 48 48  54 54 53 53 20 20 52 52  |RRIIGGHHTTSS  RR|
000000f0  45 45 53 53 45 45 52 52  56 56 45 45 44 44 20 20  |EESSEERRVVEEDD  |
00000100  20 38 37 46 34 37 39 34  20 28 43 29 20 43 4f 50  | 87F4794 (C) COP|
00000110  59 52 49 47 48 54 20 49  42 4d 20 43 4f 52 50 4f  |YRIGHT IBM CORPO|
00000120  52 41 54 49 4f 4e 20 31  39 38 31 2c 20 31 39 39  |RATION 1981, 199|
00000130  32 20 41 4c 4c 20 52 49  47 48 54 53 20 52 45 53  |2 ALL RIGHTS RES|
00000140  45 52 56 45 44 20 9c d0  c0 f9 d0 d8 fa e6 70 e6  |ERVED ........p.|
00000150  4f e4 71 50 b0 1a d0 d8  e6 70 e6 4f e4 71 58 9d  |O.qP.....p.O.qX.|
00000160  c3 9c 50 d0 c0 f9 d0 d8  fa e6 70 e6 4f 8a c4 e6  |..P.......p.O...|
00000170  71 b0 1a d0 d8 e6 70 e6  4f e4 71 58 9d c3 b0 0e  |q.....p.O.qX....|
00000180  e8 c3 ff a8 80 74 04 b8  ff 80 c3 a8 40 b8 00 00  |.....t......@...|

...

0001ff60  cb cb e8 2a ea cb ea 7a  2a 00 e0 00 00 00 00 00  |...*...z*.......|
0001ff70  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ff80  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ff90  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ffa0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ffb0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ffc0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ffd0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001ffe0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
0001fff0  ea 5b e0 00 f0 30 31 2f  30 39 2f 39 32 00 fc 6d  |.[...01/09/92..m|
2024/02/15 19:05:16  Save file OK!

C:\Users\jgust\go\programas\bin_format\EPROM serial reader 27C1024>
```
