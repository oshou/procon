n = gets.chomp.to_i
scores = gets.chomp.split.map(&:to_i)
uscores = scores.uniq
m = gets.chomp.to_i
ascores = gets.chomp.split.map(&:to_i)

ascores.each do |ascore|
  uscores.each.with_index(1) do |uscore, i|
    if uscore <= ascore
      puts i
      break
    end

    if uscore > ascore
      if i == uscores.length
        puts i + 1
        break
      end
      next
    end
  end
end
