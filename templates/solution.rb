require 'pry'

class D00
  def initialize(data)
    @data = data.map(&:to_i)
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1; end

  def part2; end
end
