n = gets.to_i

class Stack
  attr_accessor :stack

  def initialize(stack,n)
    @stack = stack
  end

stack = []
n.times {
  order = gets.chomp.split.map(&:to_i)
  case order[0]
  when 1
    stack << order[1]
  when 2
    stack.delete_at(-1)
  when 3
    puts stack.max
  end
}
