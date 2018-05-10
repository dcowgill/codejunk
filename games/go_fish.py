#!/usr/bin/env python3

from __future__ import division, print_function

import collections
import itertools
import random
import time


RANKS = {
    1: "A", 2: "2", 3: "3", 4: "4", 5: "5",
    6: "6", 7: "7", 8: "8", 9: "9", 10: "10",
    11: "J", 12: "Q", 13: "K",
}
RANK_STR = dict(map(reversed, RANKS.items()))
SUITS = ["H", "S", "D", "C"]
STARTING_HAND_SIZE = 7
BOOK_SIZE = len(SUITS)

COMPUTER_NAMES = [
    "1040-EZ", "976-KILL", "AK-47", "BHS-79", "BIM-XT", "BOOJI-1",
    "CRC-16", "DORN-3", "DOS-1.0", "HAL-2001", "JOR-15", "KORB-7B",
    "ME-262", "NECRO-99", "SR-71", "XR4-TI",
]


Card = collections.namedtuple('Card', ['rank', 'suit'])


def create_ocean(num_decks=1):
    """Return a shuffled list of cards."""
    cards = [Card(rank, suit) for rank, suit in itertools.product(RANKS, SUITS)]
    cards *= num_decks
    random.shuffle(cards)
    return cards


def ranks_in_hand(hand):
    """Return the unique ranks in a hand."""
    return set(c.rank for c in hand)


def count_by_rank(hand, rank):
    """Return the number of cards in hand of the given rank."""
    return sum(1 for c in hand if c.rank == rank)


def remove_by_rank(hand, rank):
    """
    Return two lists: the first containing cards in hand of the given
    rank, the second containing the remaining cards.
    """
    return ([card for card in hand if card.rank == rank],
            [card for card in hand if card.rank != rank])


def make_books(hand):
    """
    Return two lists: the first is a list of books that can be made from
    the hand, and the second is a new hand with the books removed.
    """
    # Group the cards by rank
    byrank = collections.defaultdict(list)
    for card in hand:
        byrank[card.rank].append(card)

    # Create the books, removing them from "byrank"
    books = []
    for rank, cards in byrank.items():
        if len(cards) >= BOOK_SIZE:
            while len(cards) >= BOOK_SIZE:
                book, cards = cards[:BOOK_SIZE], cards[BOOK_SIZE:]
                books.append(book)
            byrank[rank] = cards  # this is safe

    # Build a new hand from the remaining cards
    newhand = []
    for cards in byrank.values():
        newhand.extend(cards)

    return books, newhand


def plural(n):
    """Return "s" if n is not equal to 1, else the empty string."""
    if not isinstance(n, int):
        n = len(n)
    return "" if n==1 else "s"


def ncards_to_string(cards):
    """
    Return a human-readable string resembling "N card(s)", indicating
    the length of the given list of cards.
    """
    return "%d card%s" % (len(cards), plural(cards))


def card_to_string(card):
    """Convert a card to human-friendly text."""
    return RANKS[card.rank] + card.suit


def cards_to_string(cards):
    """Convert a list of cards to human-friendly text."""
    return ", ".join(map(card_to_string, cards))


__message_counter__ = 0

def message(s):
    """Print an output message."""
    global __message_counter__
    print("[{}] {}".format(__message_counter__, s))
    __message_counter__ += 1
    time.sleep(0.1)


__observers__ = []

def register_observer(obs):
    __observers__.append(obs)

def notify_observers(event):
    for obs in __observers__:
        obs.notify(event)


AskEvent = collections.namedtuple('AskEvent', ['player', 'opponent', 'rank', 'cards'])
FishEvent = collections.namedtuple('FishEvent', ['player', 'card'])
BookEvent = collections.namedtuple('BookEvent', ['player', 'cards'])


class BasePlayer(object):
    def __init__(self, name, is_human):
        self.name = name
        self.is_human = is_human
        self.hand = []
        self.score = 0

    def notify(self, event):
        pass

    def __str__(self):
        return self.name


class HumanPlayer(BasePlayer):
    def __init__(self, name):
        super(HumanPlayer, self).__init__(name=name, is_human=True)

    def choose_rank(self):
        valid_ranks = ranks_in_hand(self.hand)
        if not valid_ranks:
            return None
        rank = input("Enter a rank: ")
        while True:
            while rank.upper() not in RANK_STR:
                rank = input("Enter a valid rank (2-10, J, Q, K, A): ")
            ranknum = RANK_STR[rank.upper()]
            if ranknum in ranks_in_hand(self.hand):
                return ranknum
            message("You can't ask for a rank you don't have!")
            rank = input("Enter a rank: ")


class ComputerPlayer(BasePlayer):
    def __init__(self, name):
        super(ComputerPlayer, self).__init__(name=name, is_human=False)
        self._enemy_ranks = collections.defaultdict(bool)
        self._last_asked = collections.defaultdict(int)
        self._ask_counter = 0
        self._last_fished = None

    def choose_rank(self):
        time.sleep(0.5)
        valid_ranks = ranks_in_hand(self.hand)

        if not valid_ranks:
            return None

        def _ask(rank):
            self._ask_counter += 1
            self._last_asked[rank] = self._ask_counter
            return rank

        can_steal = valid_ranks.intersection(
            rank for rank, has_some in self._enemy_ranks.items() if has_some)
        if can_steal:
            return _ask(random.choice(list(can_steal)))

        if self._last_fished is not None:
            rank = self._last_fished.rank
            if count_by_rank(self.hand, rank) == 1:
                self._last_fished = None
                return _ask(rank)

        def sort_key(r):
            return self._last_asked[r], random.random()
        least_recently_fished = list(sorted(valid_ranks, key=sort_key))
        return _ask(least_recently_fished[0])

    def notify(self, event):
        if isinstance(event, AskEvent):
            if event.player is not self:
                self._enemy_ranks[event.rank] = True
            else:
                self._enemy_ranks[event.rank] = False
        elif isinstance(event, FishEvent):
            if event.player is not self:
                if event.card is not None:
                    self._enemy_ranks[event.card.rank] = True
            else:
                self._last_fished = event.card
        elif isinstance(event, BookEvent):
            if event.player is not self:
                self._enemy_ranks[event.cards[0].rank] = False


def draw_if_empty(ocean, player):
    """
    If the player's hand is empty and there are cards left in the ocean,
    draw a card and return true. Otherwise, return false.
    """
    if not ocean or player.hand:
        return False
    card = ocean.pop()
    message("{player}'s hand is empty, drawing a card. Got {card}".format(
        player=player, card=card_to_string(card)))
    player.hand.append(card)
    return True


def make_and_score_books(player):
    """
    If the player can make one or more books, remove them from the hand,
    add them to the player's score, and return true. Else return false.
    """
    books, player.hand = make_books(player.hand)
    if not books:
        return False
    for book in books:
        message("{player} made a book! {book}".format(player=player, book=cards_to_string(book)))
        notify_observers(BookEvent(player=player, cards=book))
    player.score += len(books)
    return True


def process_turn(ocean, player, opponent):
    def quiesce():
        while draw_if_empty(ocean, player) or make_and_score_books(player):
            pass

    turn_finished = False
    while not turn_finished:
        quiesce()

        if player.is_human:
            message("{player} has {player_ncards}: {player_hand}".format(
                player=player, player_ncards=ncards_to_string(player.hand),
                player_hand=cards_to_string(player.hand)))
            message("{opponent} has {opponent_ncards}".format(
                opponent=opponent, opponent_ncards=ncards_to_string(opponent.hand)))
            message("Current score: {player} {player.score}, {opponent} {opponent.score}".format(
                player=player, opponent=opponent))

        rank = player.choose_rank()
        if rank is None:
            return

        message("{player} asks {opponent} if they have any {rank}s".format(
            player=player, opponent=opponent, rank=RANKS[rank]))
        cards, opponent.hand = remove_by_rank(opponent.hand, rank)
        notify_observers(AskEvent(player=player, opponent=opponent, rank=rank, cards=cards))

        if cards:
            message("{player} takes {ncards} from {opponent}!".format(
                player=player, opponent=opponent, ncards=ncards_to_string(cards)))
            for card in cards:
                player.hand.append(card)
        else:
            message("{opponent} says: Go fish!".format(opponent=opponent))
            if not ocean:
                message("The ocean is empty. There are no more cards to draw!")
                return
            card = ocean.pop()
            player.hand.append(card)
            if player.is_human:
                message("{player} draws a card: {card}".format(
                    player=player, card=card_to_string(card)))
            else:
                message("{player} draws a card".format(player=player))
            if card.rank == rank:
                message("It's {card}, the rank {player} fished for!".format(
                    player=player, card=card_to_string(card)))
                notify_observers(FishEvent(player=player, card=card))
            else:
                player.notify(FishEvent(player=player, card=card))
                opponent.notify(FishEvent(player=player, card=None))
                quiesce()
                return


def play():
    ocean = create_ocean(num_decks=1)
    max_books = len(ocean) // BOOK_SIZE

    human = HumanPlayer(input("What is your name? "))
    computer = ComputerPlayer(random.choice(COMPUTER_NAMES))
    message("Your A.I. opponent is {computer}!".format(computer=computer))

    players = [human, computer]
    current_player, opponent = players

    register_observer(computer)

    for _ in range(STARTING_HAND_SIZE*2):
        for p in players:
            p.hand.append(ocean.pop())

    while sum(p.score for p in players) < max_books:
        message("********************")
        message("It's {player}'s turn".format(player=current_player))
        message("********************")
        process_turn(ocean, current_player, opponent)
        current_player, opponent = opponent, current_player

    message("Final score: {human} {human.score}, {computer} {computer.score}".format(
        human=human, computer=computer))

    if human.score > computer.score:
        print("You won!")
    elif human.score < computer.score:
        print("You lost :(")
    else:
        print("It's a tie!")


if __name__ == '__main__':
    play()
