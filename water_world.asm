
start:
mov rm, in
add ra, ra, 1
jmpne start, ra, 16

tryagain:
mov ra, 0
wander:
mov r0, rm
add ra, ra, 1
mov r1, rm
add ra, ra, 1
mov r2, rm
; now r0 = L, r1 = C, r2 = R
jmpgte next, r1, r0
jmpgte next, r1, r2

; at this point we now C is lower than both L and R
jmplte lislower, r0, r2

; r is lower here
sub r3, r2, r1 ; this is the difference
jmpe set, 0, 0

lislower:
sub r3, r0, r1 ; this is the difference

set:
add r6, r6, r3 ; add difference to total
sub ra, ra, 1 ; go back to the middle in ram
add rm, r3, r1 ; center is now the same as r
jmpe final, 0, 0

next:
sub ra, ra, 1

final:
jmpe tryagain, ra, 15
jmpe wander, 0, 0
