# https://checkio.org/

FIRST_TEN = [
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
]
SECOND_TEN = [
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fifteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen",
]
OTHER_TENS = [
    "twenty",
    "thirty",
    "forty",
    "fifty",
    "sixty",
    "seventy",
    "eighty",
    "ninety",
]
HUNDRED = "hundred"

def checkio(number):
    words = []
    # Hundreds: 100-999
    d, number = divmod(number, 100)
    if 1 <= d <= 9:
        words.append(FIRST_TEN[d - 1])
        words.append(HUNDRED)
    # Tens: 10-99
    d, number = divmod(number, 10)
    if d == 1:
        words.append(SECOND_TEN[number])
        number = 0
    elif d > 1:
        words.append(OTHER_TENS[d - 2])
    # Ones: 0-9
    if number > 0:
        words.append(FIRST_TEN[number - 1])
    return " ".join(words)

# These "asserts" using only for self-checking and not necessary for auto-testing
if __name__ == '__main__':
    assert checkio(4) == 'four', "1st example"
    assert checkio(133) == 'one hundred thirty three', "2nd example"
    assert checkio(12) == 'twelve', "3rd example"
    assert checkio(101) == 'one hundred one', "4th example"
    assert checkio(212) == 'two hundred twelve', "5th example"
    assert checkio(40) == 'forty', "6th example"
    assert not checkio(212).endswith(' '), "Don't forget to remove whitespace from end of string"
