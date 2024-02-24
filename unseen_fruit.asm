# left 0, forward 1, right 2, rest 3, use 4, shoot 5

mov out, 2
mov out, 1
mov out, 2
mov out, 1
mov out, 1
mov out, 1
mov out, 1
mov out, 2
mov out, 1
mov out, 0
mov out, 1
# we are at the belt

# if we see 92 (belt), rest
start:
mov out, 3
mov r0, in
jmpe start, r0, 92

# here we have a fruit
# r0 is a fruit
# store the raddr to restore 
jmpe store, ra, 0

mov r1, ra 
mov ra, 0

check:

jmpe found, rm, r0
inc ra
jmpne check, ra, r1

store:
mov rm, r0
inc ra
jmp start

found:
mov out, 2
mov out, 4
