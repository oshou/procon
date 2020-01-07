n, k = gets.chomp.split.map(&:to_i)
bills = gets.chomp.split.map(&:to_i)
paid = gets.to_i

refund = paid - ((bills.sum - bills[k]) / 2)
puts (refund == 0) ? "Bon Appetit" : refund
