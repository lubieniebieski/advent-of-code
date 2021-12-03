require 'pry'

class D01
  def initialize(data)
    @depths = data.map(&:to_i)
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    previous = 0
    increased = 0
    @depths.each do |depth|
      increased += 1 if depth > previous
      previous = depth
    end
    increased - 1
  end

  def part2
    previous = 0
    increased = 0
    @depths.each_with_index do |depth, index|
      next if index < 2

      sum = depth + @depths[index - 2] + @depths[index - 1]

      increased += 1 if sum > previous
      previous = sum
    end
    increased - 1
  end
end
