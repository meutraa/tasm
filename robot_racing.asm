push 49
push 100
push 164
push 91
push 100
push 49
push 241
push 59
push 144
push 197
push 196
push 59
push 111
push 58
push 59
push 196

start:
pop r0
mov r1, 192 ; r1 mask
mov r2, 6 ; shift count
shift:
and r3, r0, r1 ; 0b10xxxxxx
shr r3, r3, r2 ; 0bxxxxxx10
mov out, r3
sub r2, r2, 2
shr r1, r1, 2
jmpe start, r1, 0
jmpe shift, 0, 0
