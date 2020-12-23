require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @numbers = IO.read('./D09-input').split("\n").map(&:to_i)
  end

  def run
    not_so_luck_number(lucky_number)
  end

  private

  def not_so_luck_number(number, window_size = 2)
    i = window_size
    loop do
      while i < @numbers.count - 1
        window = @numbers[i - window_size..i - 1]
        return window.min + window.max if window.sum == number
        i += 1
      end
      window_size += 1
      i = window_size
    end
  end

  def lucky_number
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

Measure.new('D09-02', Calculation.new).run
