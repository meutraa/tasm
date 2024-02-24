mov r0, 10
mov r1, 20
# should be 30 at this point
start:
call addthirty
jmplte start, r2, 100
# loop until it hits 100
sub r2, r2 10
not r0, r0
push r1
pop r3
jmp end

addthirty:
add r2, r2, 20
call addten
ret

addten:
add r2, r2, 10
ret

end:
