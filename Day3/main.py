import argparse

class Engine():
  def  __init__(self, numeric_parts, symbol_parts) -> None:
    self.numeric_parts = numeric_parts
    self.symbol_parts = symbol_parts

    self.symbol_map = {}
    for symbol_part in self.symbol_parts:
      self.symbol_map[symbol_part.coordinates] = symbol_part.value

    self.numeric_map = {}
    for numeric_part in self.numeric_parts:
      start_x, start_y = numeric_part.start_coordinate
      _, end_y = numeric_part.end_coordinate
      for y in range(start_y, end_y+1):
        self.numeric_map[(start_x, y)] = numeric_part.value
    
    self.gear_parts = filter(lambda symbol_part: symbol_part.value == "*", symbol_parts)

  def sum_of_engine_parts(self):
    sum = 0
    for numeric_part in self.numeric_parts:
      if self.is_adjacent(numeric_part.start_coordinate, numeric_part.end_coordinate):
        sum += numeric_part.value
    return sum

  def get_gear_ratio(self):
    gear_ratio = 0
    for gear_part in self.gear_parts:
      # print(f"Gear part with coorinates {gear_part.coordinates}")
      adjacent_numbers = self.find_adjacent_numbers(gear_part.coordinates)
      # print(f"Adjacent numbers: {adjacent_numbers}")
      if len(adjacent_numbers) == 2:
        gear_ratio += (adjacent_numbers[0] * adjacent_numbers[1])
    return gear_ratio

  def is_adjacent(self, start_coordinate, end_coordinate):
    start_x, start_y = start_coordinate
    end_x, end_y = end_coordinate
    for x in range(start_x-1, end_x+2):
        for y in range(start_y-1, end_y+2):
          if (x, y) in self.symbol_map:
            return True
    
  def find_adjacent_numbers(self, gear_coordinates):
    x, y = gear_coordinates
    adjacent_numbers = []
    for i in range(x-1, x+2):
      for j in range(y-1, y+2):
        if (i, j) in self.numeric_map:
          # Check if we have already appended it, this is a hack since it will not work in case there are two numbers that have different coordinates
          # and both are adjacent to the gear but it worked lol
          if self.numeric_map[(i, j)] in adjacent_numbers:
            continue
          adjacent_numbers.append(self.numeric_map[(i, j)])
    return adjacent_numbers

class NumberPart():
  def __init__(self, value, start, end) -> None:
    self.value = value
    if len(start) != 2 or len(end) != 2:
      raise Exception("Engine parts coordinates must be of length 2")
    self.start_coordinate = start
    self.end_coordinate = end

class SymbolPart():
  def __init__(self, value, coordinates) -> None:
    self.value = value
    if len(coordinates) != 2:
      raise Exception("Symbol parts coordinates must be of length 2")
    self.coordinates = coordinates

class Gear(SymbolPart):
  def __init__(self, coordinates) -> None:
    super().__init__("*", coordinates)

def is_number(char):
  return len(char) != 1 or (ord(char) >= 48 and ord(char) < 58)

def is_symbol(char):
  return len(char) == 1 and char != '.' and not is_number(char)

def is_adjacent(engine_map, line_index, char_index):
  for i in range(max(line_index-1, 0), min(line_index+2, len(engine_map))):
    for j in range(max(char_index-1, 0), min(char_index+2, len(engine_map[i]))):
      if is_symbol(engine_map[i][j]):
        # print(f"Found symbol {engine_map[i][j]} at ({i}, {j})")
        return True
  return False

def parse_line_to_engine_parts(line_index, line):
  numeric_parts = []
  symbol_parts = []
  numeric_value_start = None
  for char_index, char in enumerate(line):
    if is_number(char):
      if numeric_value_start == None:
        numeric_value_start = char_index
    else:
      if is_symbol(char):
        symbol_parts.append(SymbolPart(char, (line_index, char_index)))
      if numeric_value_start != None:
        numeric_parts.append(NumberPart(int("".join(line[numeric_value_start:char_index])), (line_index, numeric_value_start), (line_index, char_index-1)))
      numeric_value_start = None
  if numeric_value_start != None:
    numeric_parts.append(NumberPart(int("".join(line[numeric_value_start:])), (line_index, numeric_value_start), (line_index, char_index)))
  return numeric_parts, symbol_parts

def parse_line(line):
  start = None
  parsed_line = []
  strip_line = line.strip()
  for i, char in enumerate(strip_line):
    if is_number(char):
      if start == None:
        start = i
    else:
      if start != None:
        parsed_line.extend(["".join(line[start:i])]*(i-start))
        start = None
      parsed_line.append(char)
  if start != None:
    parsed_line.extend(["".join(strip_line[start:])]*(len(line)-start-1))
  return parsed_line

def evaluate_engine(engine_map):
  engine_value = 0
  for i, line in enumerate(engine_map):
    current_number = None
    adjacent_to_symbol = False
    for j, char in enumerate(line):
      if is_number(char):
        current_number = char
        if adjacent_to_symbol:
          pass
        else:
          if is_adjacent(engine_map, i, j):
            adjacent_to_symbol = True
            engine_value += int(current_number)
      else:
        current_number = None
        adjacent_to_symbol = False
  return engine_value

def solve_puzzle(file_name):
  # Read the input file
  with open(file_name, "r") as f:
    lines = f.readlines()

  # Solve the puzzle

  # Part One
  # Example: 4361
  # Input: 540025
  engine_map = [parse_line(line) for line in lines]
  solution = evaluate_engine(engine_map)
  print(f"Solution with approach 1: {solution}")
  
  numeric_parts = []
  symbol_parts = []
  for i, line in enumerate(lines):
    engine_parts = parse_line_to_engine_parts(i, line.strip())
    numeric_parts.extend(engine_parts[0])
    symbol_parts.extend(engine_parts[1])
  engine = Engine(numeric_parts, symbol_parts)
  solution_2 = engine.sum_of_engine_parts()
  print(f"Solution with approach 2: {solution_2}")

  # Part Two
  # Example: 451490
  solution_3 = engine.get_gear_ratio()
  print(f"Solution part two: {solution_3}")

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
