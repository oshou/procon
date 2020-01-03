class Person
  attr_accessor :age

  def initialize(initialAge)
  end

  def amIOld()
    return @age
  end

  def yearPasses()
    @age += 1
  end
end

require "minitest/autorun"

class PersonTest < Minitest::Test
  def setUp()
  end

  def test_amIOld()
    assert_equal(-1, amIOld())
    assert_equal(-1, amIOld())
  end

  def test_yearPasses()
  end
end

#class Person
#	attr_accessor :age
#
#	def initialize(initialAge)
#		if initialAge < 0
#			print("Age is not valid, setting age to 0.\n")
#			@age=0
#		else
#			@age=initialAge
#		end
#	end
#
#	def yearPasses()
#		@age=age+1
#	end
#
#	def amIOld()
#		if @age>=18
#			print("You are old.\n")
#		elsif @age>=13
#			print("You are a teenager.\n")
#		else
#			print("You are young.\n")
#		end
#	end
#end
#
#T=gets.to_i
#for i in (1..T) do
#	age=gets.to_i
#	p=Person.new(age)
#	p.amIOld()
#	for j in (1..3) do
#		p.yearPasses()
#	end
#	p.amIOld()
#end
#
