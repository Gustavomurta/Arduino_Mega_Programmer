Serial Arduino reader and save bin file - ver 2 

Adjust the size of the eprom here:

for i := 0; i < 131072; i++ { // read 512K bytes  => 128K for example

But the size must be the same in both TinyGo and Go programs.






