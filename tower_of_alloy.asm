mov r0, in ; stack size
mov r1, in ; r1 is the starting location 0 - 2
mov r2, in ; r2 is the r3ination 0 - 2
mov r3, in ; r3 is the r2 place
mov out, r1

; n = goto spot n, 5 turn magnet on off

;r2 should be r2
;r1 r1
;r3 r3

push r2
push r1
push r2
push r3
push r1
push r3
push r2
push r1
push r3
push r2
push r3
push r1
push r2
push r1
push r2
push r3
push r1
push r3
push r1
push r2
push r3
push r2
push r1
push r3
push r2
push r1
push r2
push r3
push r1
push r3
push r2
push r1
push r3
push r2
push r3
push r1
push r2
push r1
push r3
push r2
push r1
push r3
push r1
push r2
push r3
push r2
push r3
push r1
push r2
push r1
push r2
push r3
push r1
push r3
push r2
push r1
push r3 ; spare
push r2 ; dest
push r3 ; spare
push r1 ; start
push r2 ; dest
;push r1 ; start

;test
start:
mov out, 5
pop out
jmpe start, 0, 0

