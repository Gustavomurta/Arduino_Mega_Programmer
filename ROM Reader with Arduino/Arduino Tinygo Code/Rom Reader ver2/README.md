Version 2

You may read EPROMs up to 512 KBytes. 

Change these lines of programs to adjust the size of your eprom.

for block64K = 0; block64K < 2; block64K++ { // count blocks of 64K => 128K  (TinyGo program) 

But the size must be the same in both TinyGo and Go programs. 

for i := 0; i < 131072; i++ { // read 512K bytes  => 128K   (Go program) 


Compiling the program:

C:\Users\jgust\tinygo\programas\programador_eprom\EPROM_512K>go mod init main.go

go: creating new go.mod: module main.go

go: to add module requirements and sums:

        go mod tidy
        
C:\Users\jgust\tinygo\programas\programador_eprom\EPROM_512K>go mod tidy

tinygo flash -target=arduino-mega2560 main.go  (flashing FW on Arduino) 

