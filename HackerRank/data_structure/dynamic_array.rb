
def solution(n, queries)
  ans = ""
  last_answer = 0
  arr = (0..n - 1).map { [] }
  queries.each do |query|
    case query[0]
    when 1
      seq = (query[1] + last_answer) % n
      arr[seq] << query[2]
    when 2
      seq = (query[1] + last_answer) % n
      last_answer = arr[seq][-1]
      puts last_answer
    end
  end
end

require "minitest/autorun"

class SolutionTest < Minitest::Test
  def test_solution
    n = 2
    queries = [
      [1, 0, 5],
      [1, 1, 7],
      [1, 0, 3],
      [2, 1, 0],
      [2, 1, 1],
    ]
    assert_equal(solution(n, queries), 7)
  end
end
