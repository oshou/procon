n = gets.to_i;
ary = gets.split(" ").map(&:to_i);
sum = 0;
ary.each{|x| sum+=x};
print("#{ary.min} #{ary.max} #{sum}\n");
