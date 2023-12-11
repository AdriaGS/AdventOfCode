import argparse


NUM_RED_SQUARES = 12
NUM_GREEN_SQUARES = 13
NUM_BLUE_SQUARES = 14


def insert_sorted(arr, val):
  i = 0
  while i < len(arr) and arr[i] < val:
    i += 1
  arr.insert(i, val)
  return arr


def parse_line(line):
  game, squares_list = line.replace("\n", "").split(":")

  game_id = int(game.split(" ")[1])

  squares_by_color = {"red": [], "blue": [], "green": []}
  for square_set in squares_list.split(";"):
    for square in square_set.split(","):
      square_number, square_color = square.strip().split(" ")
      insert_sorted(squares_by_color[square_color], int(square_number))

  return game_id, squares_by_color

def is_possible_game(squares_by_color):
  return max(squares_by_color["red"]) <= NUM_RED_SQUARES and \
    max(squares_by_color["green"]) <= NUM_GREEN_SQUARES and \
    max(squares_by_color["blue"]) <= NUM_BLUE_SQUARES

def game_power(squares_by_color):
  return squares_by_color["red"][-1] * squares_by_color["green"][-1] * squares_by_color["blue"][-1]

def solve_puzzle(file_name):
  # Read the input file
  with open(file_name, "r") as f:
    lines = f.readlines()

  # Solve the puzzle
  solution = 0
  for line in lines:
    game_id, squares_by_color = parse_line(line)
    # if is_possible_game(squares_by_color):
    #   solution += game_id
    solution += game_power(squares_by_color)

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
