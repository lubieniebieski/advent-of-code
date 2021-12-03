require 'pry'

class D03
  def initialize(data)
    @data = data.map(&:strip)
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    gg = ''
    ee = ''
    digits = @data.first.length
    digits.times do |d|
      g = 0
      e = 0
      @data.each do |n|
        g += 1 if n[d] == '1'
        e += 1 if n[d] == '0'
      end

      gg += g > e ? '1' : '0'
      ee += g > e ? '0' : '1'
    end
    gg.to_i(2) * ee.to_i(2)
  end

  def part2
    calculate_oxygen(@data.dup) * calculate_co2(@data.dup)
  end

  private

  def calculate_oxygen(array)
    iterations = array.first.length
    selected_oxygen = array
    iterations.times do |i|
      transposed = array.map(&:chars).transpose
      zeros = transposed[i].count('0')
      ones = transposed[i].count('1')
      if zeros > ones
        selected_oxygen.reject! { |a| a[i] == '1' }
      else
        selected_oxygen.reject! { |a| a[i] == '0' }
      end
      return selected_oxygen.first.to_i(2) if selected_oxygen.size == 1
    end
  end

  def calculate_co2(array)
    iterations = array.first.length
    selected_co2 = array
    iterations.times do |i|
      transposed = array.map(&:chars).transpose
      zeros = transposed[i].count('0')
      ones = transposed[i].count('1')
      if zeros > ones
        selected_co2.reject! { |a| a[i] == '0' }
      else
        selected_co2.reject! { |a| a[i] == '1' }
      end

      return selected_co2.first.to_i(2) if selected_co2.size == 1
    end
  end
end
