mealCost = gets.to_f
tipPercent = gets.to_i
taxPercent = gets.to_i
puts (mealCost * (1 + (tipPercent.to_f + taxPercent) /100)).round
