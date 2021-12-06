require 'pry'

class D06
  def initialize(data)
    input = data.split(',').map(&:to_i)
    new_data = {}
    (0..8).each do |number|
      new_data[number] = input.count { |a| a == number }
    end
    @data = new_data
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    80.times do
      pass_the_day!
    end
    @data.values.sum
  end

  def part2
    256.times do
      pass_the_day!
    end
    @data.values.sum
  end

  private

  def pass_the_day!
    initial_zeros = @data[0]
    (1..8).each do |timer|
      @data[timer - 1] = @data[timer]
    end
    @data[6] += initial_zeros
    @data[8] = initial_zeros
  end
end
