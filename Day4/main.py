import argparse
import re

class Card():
  def __init__(self, id, winner_numbers, card_numbers) -> None:
    self.id = id
    self.winner_numbers = winner_numbers
    self.card_numbers = card_numbers

  def get_points(self):
    points = self.get_matching_numbers()
    if points == 0:
      return 0
    return 2**(points-1)
  
  def get_matching_numbers(self):
    matching_numbers = 0
    for number in self.card_numbers:
      if number in self.winner_numbers:
        matching_numbers += 1
    return matching_numbers

class CardDeck():
  def __init__(self, cards) -> None:
    self.cards = {}
    for card in cards:
      self.cards[card.id] = card
  
  def get_card(self, id):
    return self.cards.get(id)
  
  def get_number_of_cards(self):
    return sum([len(self.visit_card(id)) for id in self.cards.keys()])
  
  def visit_card(self, card_id):
    card = self.get_card(card_id)
    cards_visited = [card]
    card_matching_numbers = card.get_matching_numbers()
    # Here we should store the path of a specific number so in future 
    # executions of the same number we already know what it means in terms
    # of cards.
    for i in range(1, card_matching_numbers + 1):
      cards_visited.extend(self.visit_card(card_id + i))
    return cards_visited

def parse_line(line):
  clean_line = re.sub(" +", " ", line.strip())
  card, numbers = clean_line.split(":")
  card_id = int(card.split(" ")[-1])
  winner_numbers, card_numbers = numbers.split("|")
  return Card(card_id, winner_numbers.strip().split(" "), card_numbers.strip().split(" "))

def solve_puzzle(file_name):
  # Read the input file
  with open(file_name, "r") as f:
    lines = f.readlines()

  # Solve the puzzle
  # Part One
  cards = [parse_line(line) for line in lines]
  solution = sum([card.get_points() for card in cards])

  # Part Two
  card_deck = CardDeck(cards)
  solution_2 = card_deck.get_number_of_cards()
  print(solution_2)

  # Return the solution
  return solution

def main():

  # Parse any command line arguments
  parser = argparse.ArgumentParser()
  parser.add_argument("-f", "--file-name", help="File name")
  args = parser.parse_args()

  # Solve the puzzle
  solution = solve_puzzle(args.file_name)
  print(f"Solution: {solution}")

# Call the main method
if __name__ == "__main__":
  main()
