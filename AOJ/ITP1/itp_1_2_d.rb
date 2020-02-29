W, H, x, y, r = gets.chomp.split(" ").map(&:to_i);
puts (r <= x && r <= y && r + x < W && r + y < H) ? "Yes" : "No";
