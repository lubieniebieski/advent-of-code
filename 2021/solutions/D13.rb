require 'pry'

class D13
  def initialize(data)
    data = data.map(&:strip)
    divider_index = data.index('')
    @dots = data[...divider_index].map { |a| a.split(',').map(&:to_i) }
    @folds = data[(divider_index + 1)..-1]
    @matrix = dots_to_matrix(@dots)
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    fold!(1)
    number_of_dots
  end

  def number_of_dots
    @matrix.flatten.count(true)
  end

  def fold!(fold_number)
    fold_number.times do
      match = @folds.shift.match(/.*(?<axis>\S)=(?<index>\d+)/)
      axis = match[:axis]
      index = match[:index].to_i
      fold_one!(axis, index)
    end
  end

  def part2
    fold!(@folds.count)
    number_of_dots
  end

  def print_code
    puts @matrix.map { |b| b.map { |a| a == true ? '#' : ' ' }.join }
  end

  private

  def dots_to_matrix(dots)
    max_x = dots.map(&:first).max + 1
    max_y = dots.map(&:last).max + 1
    matrix = Array.new(max_y) { Array.new(max_x, false) }
    dots.each do |dot|
      x, y = dot
      matrix[y][x] = true
    end
    matrix
  end

  def fold_one!(axis, index)
    if axis == 'y'
      @matrix[0...index].each_with_index do |row, y|
        row.each_with_index do |_cell, x|
          @matrix[y][x] = @matrix[y][x] || @matrix[@matrix.length - 1 - y][x]
        end
      end

      @matrix = @matrix[0...index]
    else
      @matrix.each_with_index do |row, y|
        row[0...index].each_with_index do |_cell, x|
          @matrix[y][x] = @matrix[y][x] || @matrix[y][row.length - 1 - x]
        end
      end
      @matrix = @matrix.transpose[0...index].transpose
    end
  end
end
