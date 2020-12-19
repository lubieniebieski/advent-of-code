require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @lines = IO.readlines('./D03-input')
    @trees = 0
  end

  def run
    x = 1
    @lines.map do |line|
      add_tree if tree?(line.strip, x)
      x += 3
    end
    @trees
  end

  private

  def add_tree
    @trees += 1
  end

  def tree?(line, position)
    chars = line.chars
    if chars.count <= position
      tree?((chars * 2).join, position)
    else
      chars[position - 1] == '#'
    end
  end
end

Measure.new('D03-01', Calculation.new).run
