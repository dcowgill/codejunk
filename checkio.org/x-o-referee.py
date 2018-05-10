# https://checkio.org/

def checkio(game_result):
    rows = list(zip(*game_result))  # columns
    rows.append(r[i] for i, r in enumerate(game_result))  # first diagonal
    rows.append(r[-i-1] for i, r in enumerate(game_result))  # second diagonal
    rows = list(map("".join, rows))
    rows.extend(game_result)
    x_wins = "XXX" in rows
    o_wins = "OOO" in rows
    return "D" if x_wins == o_wins else "X" if x_wins else "O"


#These "asserts" using only for self-checking and not necessary for auto-testing
if __name__ == '__main__':
    assert checkio([
        "X.O",
        "XX.",
        "XOO"]) == "X", "Xs wins"
    assert checkio([
        "OO.",
        "XOX",
        "XOX"]) == "O", "Os wins"
    assert checkio([
        "OOX",
        "XXO",
        "OXX"]) == "D", "Draw"
    assert checkio([
        "O.X",
        "XX.",
        "XOO"]) == "X", "Xs wins again"
