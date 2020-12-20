require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @groups = IO.read('./D06-input').split("\n\n").map { |g| Group.new(g) }
  end

  def run
    @groups.sum(&:sum_of_yeses)
  end
end

class Group
  def initialize(data)
    @raw_data = data
  end

  def sum_of_yeses
    @raw_data.split("\n").join.chars.uniq.count
  end
end

Measure.new('D06-01', Calculation.new).run
