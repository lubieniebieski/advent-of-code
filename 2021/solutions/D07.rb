require 'pry'

class D07
  def initialize(data)
    @positions = data.split(',').map(&:to_i)
    @solutions = {}
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    (@positions.min..@positions.max).map do |pos|
      calculate_cost(pos, @positions)
    end.min
  end

  def part2
    (@positions.min..@positions.max).map do |pos|
      calculate_cost_v2(pos, @positions)
    end.min
  end

  private

  def calculate_cost(current_position, positions)
    positions.sum { |p| (current_position - p).abs }
  end

  def calculate_cost_v2(current_position, positions)
    positions.sum do |p|
      @solutions[(current_position - p).abs] ||=
        (0..(current_position - p).abs).inject(0) { |sum, x| sum + x }
      @solutions[(current_position - p).abs]
    end
  end
end
