require 'pry'

class D05
  class Line
    attr_reader :x1, :y1, :x2, :y2

    def initialize(line)
      matches = line.match(/^(?<x1>\d+),(?<y1>\d+).*\D(?<x2>\d+),(?<y2>\d+)$/)
      @x1 = matches[:x1].to_i
      @y1 = matches[:y1].to_i
      @x2 = matches[:x2].to_i
      @y2 = matches[:y2].to_i
    end

    def marked_coordinates
      if vertical?
        if y2 > y1
          (y1..y2).to_a.map { |y| [x1, y] }
        else
          (y2..y1).to_a.map { |y| [x1, y] }
        end
      elsif horizontal?
        if x2 > x1
          (x1..x2).to_a.map { |x| [x, y1] }
        else
          (x2..x1).to_a.map { |x| [x, y1] }
        end
      elsif x2 > x1
        y = y1
        if y2 > y1
          (x1..x2).to_a.map do |x|
            y += 1
            [x, y - 1]
          end
        else
          (x1..x2).to_a.map do |x|
            y -= 1
            [x, y + 1]
          end
        end
      elsif x2 < x1
        y = y2
        if y2 > y1
          (x2..x1).to_a.map do |x|
            y -= 1
            [x, y + 1]
          end
        else
          (x2..x1).to_a.map do |x|
            y += 1
            [x, y - 1]
          end
        end
      end
    end

    def vertical?
      x1 == x2
    end

    def horizontal?
      y1 == y2
    end
  end

  def initialize(data)
    @lines = data.map { |line| Line.new(line) }
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    results = {}
    @lines.select { |line| line.vertical? || line.horizontal? }.each do |line|
      line.marked_coordinates.each do |coord|
        results[coord] = results.fetch(coord, 0) + 1
      end
    end
    results.select { |_, v| v > 1 }.size
  end

  def part2
    results = {}
    @lines.each do |line|
      line.marked_coordinates.each do |coord|
        results[coord] = results.fetch(coord, 0) + 1
      end
    end
    results.select { |_, v| v > 1 }.size
  end
end
