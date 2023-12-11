import argparse
import re

NUMBER_MAP = {
  "zero": 0,
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9,
}

second_pattern = r"(?=(zero)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(\d))"

def get_number_from_match(match):
  return NUMBER_MAP[match] if match in NUMBER_MAP else int(match)

def parse_line(line):
  numbers = []
  for matches in re.finditer(second_pattern, line):
    numbers.append(
      get_number_from_match([match for match in matches.groups() if match != None][0])
    )
  return numbers

def solve_puzzle(file_name):
  # Read the input file
  with open(file_name, "r") as f:
    lines = f.readlines()

  # Solve the puzzle
  solution = 0
  for line in lines:
    digits = parse_line(line)
    if len(digits) < 1:
      print(f"Invalid line: {line}")
      break
    print(f"Line: {line} -> {digits}")
    solution += int(str(digits[0]) + str(digits[-1]))

  # Return the solution
  # 53539
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
