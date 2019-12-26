class Person
  def initialize(firstName, lastName, id)
    @firstName = firstName
    @lastName = lastName
    @id = id
  end

  def printPerson()
    print("Name: ", @lastName, ", " + @firstName, "\nID: ", @id)
  end
end

class Student < Person
  def initialize(firstName, lastName, id, scores)
    @scores = scores
  end

  def calculate()
    avg = (@scores.sum.to_f / @scores.length)
    case avg
    when 90...100
      return "O"
    when 80..90
      return "E"
    when 70..80
      return "A"
    when 55..70
      return "P"
    when 40..55
      return "D"
    else
      return "T"
    end
  end
end

nput = gets.split()
firstName = input[0]
lastName = input[1]
id = input[2].to_i
numScores = gets.to_i
scores = gets.strip().split().map!(&:to_i)
s = Student.new(firstName, lastName, id, scores)
s.printPerson
print("\nGrade: " + s.calculate)
