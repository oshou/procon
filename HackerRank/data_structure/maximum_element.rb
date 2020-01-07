
class Stack
  attr_accessor :stack

  def initialize(stack = [])
    @stack = stack
  end

  def element_op(order)
    case order[0]
    when 1
      stack << order[1]
    when 2
      stack.delete_at(-1)
    when 3
      puts stack.max
    end
  end
end

require "minitest/autorun"

class StackTest < Minitest::Test
  def test_element_op
    stack = []
    stack = Stack.new([])
    tests = [
      {
        "input" => [1, 25],
        "output" => [25],
      },
      {
        "input" => [1, 96],
        "output" => [25, 96],
      },
      {
        "input" => [2],
        "output" => 96,
      },
      {
        "input" => [3],
        "output" => nil,
      },
    ]
    tests.each do |test|
      assert_equal(test["output"], stack.element_op(test["input"]))
    end
  end
end
