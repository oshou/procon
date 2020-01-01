def solution(arr)
  max = -63
  (0..3).each do |i|
    (0..3).each do |j|
      sum = arr[i][j] + arr[i][j + 1] + arr[i][j + 2] + arr[i + 1][j + 1] + arr[i + 2][j] + arr[i + 2][j + 1] + arr[i + 2][j + 2]
      if max < sum
        max = sum
      end
    end
  end
  return max
end

require "minitest/autorun"

class SolutionTest < Minitest::Test
  def test_solution
    arr = [
      [1, 1, 1, 0, 0, 0],
      [0, 1, 0, 0, 0, 0],
      [1, 1, 1, 0, 0, 0],
      [0, 0, 2, 4, 4, 0],
      [0, 0, 0, 2, 0, 0],
      [0, 0, 1, 2, 4, 0],
    ]
    assert_equal(19, solution(arr))

    arr = [
      [1, 1, 1, 0, 0, 0],
      [0, 1, 0, 0, 0, 0],
      [1, 1, 1, 0, 0, 0],
      [0, 9, 2, -4, -4, 0],
      [0, 0, 0, -2, 0, 0],
      [0, 0, -1, -2, -4, 0],
    ]
    assert_equal(13, solution(arr))

    arr = [
      [-1, -1, 0, -9, -2, -2],
      [-2, -1, -6, -8, -2, -5],
      [-1, -1, -1, -2, -3, -4],
      [-1, -9, -2, -4, -4, -5],
      [-7, -3, -3, -2, -9, -9],
      [-1, -3, -1, -2, -4, -5],
    ]
    assert_equal(-6, solution(arr))
  end
end
