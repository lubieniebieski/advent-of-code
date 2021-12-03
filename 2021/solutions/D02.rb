require 'pry'

class D02
  def initialize(data)
    @data = data
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    horizontal = 0
    depth = 0
    @data.each do |line|
      direction, value = line.split
      case direction
      when 'forward'
        horizontal += value.to_i
      when 'down'
        depth += value.to_i
      else
        depth += -value.to_i
      end
    end
    horizontal * depth
  end

  def part2
    horizontal = 0
    depth = 0
    aim = 0
    @data.each do |line|
      direction, value = line.split
      case direction

      when 'up'
        aim -= value.to_i
      when 'down'
        aim += value.to_i
      else
        horizontal += value.to_i
        depth += aim * value.to_i
      end
    end

    horizontal * depth
  end
end
