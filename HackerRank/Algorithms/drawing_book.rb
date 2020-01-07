n = gets.to_i
p = gets.to_i

last = (n % 2 == 0) ? n : n - 1

from_first = ((p - 1).to_f / 2).ceil
from_last = ((last - p).to_f / 2).ceil
puts [from_first, from_last].min
