=begin
 * ITP_1_6_B
 *
 * input:
 * In the first line, the number of cards n (n ≤ 52) is given.
 * In the following n lines, data of the n cards are given.
 * Each card is given by a pair of a character and an integer which represent its suit and rank respectively.
 * A suit is represented by 'S', 'H', 'C' and 'D' for spades, hearts, clubs and diamonds respectively.
 * A rank is represented by an integer from 1 to 13.
 *
 * output:
 * Print the missing cards. The same as the input format,
 * each card should be printed with a character and an integer separated by a space character in a line.
 * Arrange the missing cards in the following priorities:
 * Print cards of spades, hearts, clubs and diamonds in this order.
 * If the ranks are equal, print cards with lower ranks first.
=end
n = gets.to_i
h = {}
n.times do
  s = gets
  h[s.chomp] = 1
end

for m in ["S", "H", "C", "D"] do
  for i in 1..13 do
    puts "#{m} #{i}" if !h.has_key?(m+" "+i.to_s)
  end
end
