; temp1 = seed  xor (seed  shr 1)
; temp2 = temp1 xor (temp1 shl 1)
; nseed = temp2 xor (temp2 shr 2)  

; assume r0 is our seed
mov r0, in

next:
shr r1, r0, 1
xor r1, r0, r1
shl r2, r1, 1
xor r2, r2, r1
shr r3, r2, 2
xor r0, r3, r2
; r0 is now our next seed
and out, r0, 3
jmp next
