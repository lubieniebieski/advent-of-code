require_relative '../lib/utils.rb'

class Calculation
  def initialize
    @numbers = IO.read('./D09-input').split("\n").map(&:to_i)
  end

  def run
    window_size = 25
    i = window_size
    loop do
      unless @numbers[i - window_size..i - 1].permutation(2).any? { |p| p.sum == @numbers[i] }
        return @numbers[i]
      end
      i += 1
    end
  end
end

Measure.new('D09-01', Calculation.new).run
