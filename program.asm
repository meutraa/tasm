mov r0, 4
mull r1, 5, r0

start:
dec r1
jmpne start, r1, 0
