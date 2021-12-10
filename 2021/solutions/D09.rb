require 'pry'

class D09
  class Basin
    def initialize(locations)
      @locations = locations
    end

    def size
      @locations.size
    end
  end

  class Window
    attr_reader :x, :y

    def initialize(main_number, adjacent_locations, x, y)
      @main_number = main_number
      @adjacent_locations = adjacent_locations
      @x = x
      @y = y
    end

    def risk
      @main_number + 1
    end

    def low_point?
      @adjacent_locations.all? { |value| value > @main_number }
    end
  end

  def initialize(data)
    @rows = data.map(&:strip).map { |a| a.chars.map(&:to_i) }
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    data = add_some_numbers(@rows.dup)
    windows = get_windows(data)
    windows.select(&:low_point?).sum(&:risk)
  end

  def part2
    data = add_some_numbers(@rows.dup)
    windows = get_windows(data)
    data.each do |row|
      row.pop
      row.shift
    end
    data.pop
    data.shift
    low_point_windows = windows.select(&:low_point?)
    basins = low_point_windows.map do |window|
      Basin.new(explore(@rows, window.x - 1, window.y - 1))
    end
    basins.sort_by(&:size).last(3).map(&:size).inject(:*)
  end

  private

  def neighbours(pos)
    y, x = pos
    n = []
    v = @rows[y][x]
    n.push([y, x + 1]) if (x + 1) < @rows[y].size && @rows[y][x + 1] != 9 && @rows[y][x + 1] > v
    n.push([y + 1, x]) if (y + 1) < @rows.size && @rows[y + 1][x] != 9 && @rows[y + 1][x] > v
    n.push([y - 1, x]) if (y - 1) >= 0 && @rows[y - 1][x] != 9 && @rows[y - 1][x] > v
    n.push([y, x - 1]) if (x - 1) >= 0 && @rows[y][x - 1] != 9 && @rows[y][x - 1] > v
    n
  end

  def explore(data, start_x, start_y)
    locations = [[start_x, start_y]]
    locations.each do |loc|
      x, y = loc
      right = [x + 1, y] if (x + 1) < data[y].size && data[y][x + 1] != 9
      left = [x - 1, y] if (x - 1) >= 0 && data[y][x - 1] != 9
      up = [x, y - 1] if (y - 1) >= 0 && data[y - 1][x] != 9
      down = [x, y + 1] if (y + 1) < data.size && data[y + 1][x] != 9
      [right, left, up, down].compact.each do |coords|
        locations << coords if data[coords[1]][coords[0]] > data[y][x]
      end
    end
    locations.uniq
  end

  def add_some_numbers(data)
    data.each do |row|
      row.push(-1)
      row.unshift(-1)
    end
    data.push(Array.new(data.first.size, -1))
    data.unshift(Array.new(data.first.size, -1))
  end

  def get_windows(data)
    windows = []
    data.each_with_index do |row, y|
      next if [0, (data.count - 1)].include?(y)

      row.each_with_index do |element, x|
        next if [0, (row.count - 1)].include?(x)

        neighbours = [data[y - 1][x], data[y + 1][x], row[x - 1], row[x + 1]].reject(&:negative?)
        windows << Window.new(element, neighbours, x, y)
      end
    end
    windows
  end
end
