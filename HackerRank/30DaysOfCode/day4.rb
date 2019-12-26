=begin
t=gets.to_i
ages=gets.chomp.split.map{|i|i.to_i}
ages.each do |age|
	if age < 0
		puts "Age is not valid, setting age to 0.."
	elsif age < 13
		puts "You are young.."
	elsif age < 18
		puts "You are a teenager.."
	else
		puts "You are old.."
	end
end
=end

class Person
	attr_accessor :age

	def initialize(initialAge)
		@age=initialAge
		if age < 0
			@age=0
			puts "Age is not valid, setting age to 0."
		end
	end

	def amIOld()
		if age < 13
			puts "You are young."
		elsif age < 18
			puts "You are a teenager."
		else
			puts "You are old."
		end
	end

	def yearPasses()
		@age += 1
	end
end

t=gets.to_i
for i in (1..t) do
	age = gets.to_i
	p = Person.new(age)
	p.amIOld()
	for j in (1..3) do
		p.yearPasses()
	end
	p.amIOld()
	puts ""
end

=begin
class Person
	attr_accessor :age

	def initialize(initialAge)
		if initialAge < 0
			print("Age is not valid, setting age to 0.\n")
			@age=0
		else
			@age=initialAge
		end
	end

	def yearPasses()
		@age=age+1
	end

	def amIOld()
		if @age>=18
			print("You are old.\n")
		elsif @age>=13
			print("You are a teenager.\n")
		else
			print("You are young.\n")
		end
	end
end

T=gets.to_i
for i in (1..T) do
	age=gets.to_i
	p=Person.new(age)
	p.amIOld()
	for j in (1..3) do
		p.yearPasses()
	end
	p.amIOld()
end
=end
