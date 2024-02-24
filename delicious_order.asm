mov r0, 15
read:
mov rm, in
inc ra
dec r0
jmpne read, r0, 0

mov r2, 1
verystart:
# r2 is the number of swaps
jmpe end, r2, 0
mov r2, 0
mov ra, 0

start:
jmpgte verystart, ra, 14
mov r0, rm
inc ra
jmpgt swap, r0, rm
jmp start

swap:
inc r2
mov r1, rm
mov rm, r0
dec ra
mov rm, r1
inc ra
jmp start

end:
mov ra, 0

write:
mov out, rm
inc ra
jmp write
