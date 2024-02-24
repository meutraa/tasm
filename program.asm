mov r0, in
mov r1, in

# calculate r0/r1, 7/3

# quotient
mov r2, 0
start:
jmplt end, r0, r1
# n = 7 - 3
sub r0, r0, r1
# n = 4
inc r2
# go to start if n >= 3
jmpgte start, r0, r1
# else n < 3, n = remainder
# r2 = quotient
end:
mov out, r2
mov out, r0
