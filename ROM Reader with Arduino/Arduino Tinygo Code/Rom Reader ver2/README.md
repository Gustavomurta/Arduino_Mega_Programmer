Version 2

You may read EPROMs up to 512 KBytes. 

Change these lines of programs to adjust the size of your eprom.

for block64K = 0; block64K < 2; block64K++ { // count blocks of 64K => 128K  (TinyGo program) 

But the size must be the same in both TinyGo and Go programs. 

for i := 0; i < 131072; i++ { // read 512K bytes  => 128K   (Go program) 

