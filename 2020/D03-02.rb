require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @lines = IO.readlines('./D03-input')
    @trees = 0
    @slopes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]
  end

  def run
    @slopes.map do |slope|
      x = 1
      y = 1
      trees = 0
      while y <= @lines.count
        if is_tree?(@lines[y - 1].strip, x)
          trees +=1
        end
        x += slope[0]
        y += slope[1]
      end
      trees
    end.inject(:*)

  end

  private


  def is_tree?(line, position)
    chars = line.chars
    if (chars.count) <= position
      is_tree?((chars * 2).join, position)
    else
      chars[position - 1] == "#"
    end
  end
end

Measure.new('D03-01', Calculation.new).run
