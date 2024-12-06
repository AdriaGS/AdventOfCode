import argparse
from collections import Counter

def parse_file(filename):
  first_column = []
  second_column = []
  try:
    with open(filename, 'r') as file:
      for line in file:
        # Split each line and convert to integers
        num1, num2 = map(int, line.strip().split())
        first_column.append(num1)
        second_column.append(num2)

      return first_column, second_column
  except FileNotFoundError:
    raise FileNotFoundError(f"Error: File '{filename}' not found")
  except ValueError:
    raise ValueError("Error: Invalid data format in file")

def solve_puzzle_part_1(file_name):
  locations_1, locations_2 = parse_file(file_name)
  # sort both lists
  sorted_locations_1 = sorted(locations_1)
  sorted_locations_2 = sorted(locations_2)
  return sum([abs(a - b) for a, b in zip(sorted_locations_1, sorted_locations_2)])

def solve_puzzle_part_2(file_name):
  locations_1, locations_2 = parse_file(file_name)
  
  counter_locations_2 = Counter(locations_2)

  return sum(
    location * counter_locations_2[location] 
    for location in locations_1 
    if location in counter_locations_2
  )

def main():
  # Parse any command line arguments
  parser = argparse.ArgumentParser()
  parser.add_argument("-f", "--file-name", help="File name")
  args = parser.parse_args()

  # Solve the puzzle
  solution_part_1 = solve_puzzle_part_1(args.file_name)
  print(f"Solution Part 1: {solution_part_1}")

  solution_part_2 = solve_puzzle_part_2(args.file_name)
  print(f"Solution Part 2: {solution_part_2}")

# Call the main method
if __name__ == "__main__":
  main()
