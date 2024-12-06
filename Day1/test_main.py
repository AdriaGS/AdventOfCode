import unittest
from unittest.mock import patch
from main import solve_puzzle_part_1, solve_puzzle_part_2

class TestPuzzleSolver(unittest.TestCase):
  def test_main_solve_puzzle_part_1(self):
    actual = solve_puzzle_part_1('test_input.txt')
    assert actual == 11, f"Expected 11, but got {actual}"

  def test_main_solve_puzzle_part_2(self):
    actual = solve_puzzle_part_2('test_input.txt')
    assert actual == 31, f"Expected 31, but got {actual}"

if __name__ == '__main__':
  unittest.main()
