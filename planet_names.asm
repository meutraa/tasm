wordstart:
# start of a word
# upper case = -32
sub out, in, 32

word:
mov r0, in
mov out, r0
jmpne word, r0, 32
jmp wordstart
