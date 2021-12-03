require_relative '../lib/utils.rb'

class D09
  def initialize(data)
    @numbers = data.map(&:to_i)
  end

  def self.run(data)
    puts new(data).part1, new(data).part2
  end

  def part1(window_size = 25)
    lucky_number(window_size)
  end

  def part2(window_size = 25)
    not_so_lucky_number(lucky_number(window_size))
  end

  private

  def not_so_lucky_number(number)
    i = 2
    window_size = 2
    loop do
      while i < @numbers.count - 1
        window = @numbers[i - window_size..i - 1]
        return(window.min + window.max) if window.sum == number
        i += 1
      end
      window_size += 1
      i = 1
    end
  end

  def lucky_number(window_size)
    i = window_size
    loop do
      unless @numbers[i - window_size..i - 1].permutation(2).any? { |p| p.sum == @numbers[i] }
        return @numbers[i]
      end
      i += 1
    end
  end
end
