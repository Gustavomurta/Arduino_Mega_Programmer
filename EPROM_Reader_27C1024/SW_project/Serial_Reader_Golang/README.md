This program capture bytes readed from EPROM, dump bytes at console and save a binary format file. 

You must change the Arduino COM port of your PC in the go program. 

To compile the go program: 

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
