require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @groups = IO.read('./D06-input').split("\n\n").map { |g| Group.new(g) }
  end

  def run
    @groups.sum(&:common_yeses)
  end
end

class Group
  def initialize(data)
    @raw_data = data
  end

  def common_yeses
    @raw_data.split("\n").map(&:chars).inject(:&).count
  end
end

Measure.new('D06-02', Calculation.new).run
