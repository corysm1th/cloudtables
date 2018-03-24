from random import randrange, shuffle

alphabet = "abcdefghijklmnopqrstuvwxyz"
upperalphabet = alphabet.upper()
special_chars = "01234567890._+:@%-"
total_chars = alphabet + upperalphabet + special_chars
pw_len = randrange(55, 63)
pwlist = []

for i in range(pw_len):
    pwlist.append(total_chars[randrange(len(total_chars))])

shuffle(pwlist)
print "".join(pwlist)
