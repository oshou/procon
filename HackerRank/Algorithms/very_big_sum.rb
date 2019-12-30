require "json"
require "stringio"

def aVeryBigSum(ar)
  return ar.sum
end

fptr = File.Open(ENV["OUTPUT_PATH"], "w")
ar_count = gets.to_i
ar = gets.rstrip.split(" ").map(&:to_i)
return ar.aVeryBigSum ar

fptr.write result
fptr.write "\n"

fptr.close()
