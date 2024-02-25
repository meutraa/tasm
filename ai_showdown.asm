mov out, 3      
mov r1, 9


; x o o o o they lose, I want to get it to 5
mov r0, in     
sub r1, r1, r0 ; they picked some up, 6-8
; r1 is how many they picked up
; if 1, I want 3
; if 3, I want 1
sub r2, 4, r1
; r2 is now 1-3
mov out, r2
; now 5 cards
mov r0, in
; they picked some up, so 2-4
; I want it to 1
; if 2, I want 1
; if 4, I want 3
dec r0
; r2 is how many they picked up
mov out, r0
