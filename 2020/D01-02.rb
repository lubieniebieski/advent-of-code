require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @numbers = IO.readlines('./D01-input').map(&:to_i).sort
  end

  def run
    result = @numbers.permutation(3).find { |p| p.sum == 2020 }
    result.inject(:*)
  end
end

Measure.new('D01-02', Calculation.new).run
