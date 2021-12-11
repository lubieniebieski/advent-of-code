require 'pry'

class D11
  def initialize(data)
    @data = clean_input(data)
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    step!(100)
  end

  def flashes_in_step(step)
    step!(step)
    @flashes[step]
  end

  def compare(data, step)
    step!(step)
    input = clean_input(data)
    output = Array.new(@data.size) { Array.new(@data[0].size) }
    input.each_with_index do |row, y|
      row.each_with_index do |cell, x|
        output[y][x] = if @data[y][x] == cell
                         '.'
                       else
                         @data[y][x]
                       end
      end
    end
    @data.map(&:inspect).join("\n") == input.map(&:inspect).join("\n")
  end

  def part2
    1000.times do |i|
      step!(1)
      return i if @flashes[i] == @data.flatten.size
    end
  end

  private

  def clean_input(data)
    data.map(&:strip).map(&:chars).map { |a| a.map(&:to_i) }
  end

  def step!(number_of_steps)
    @flashes ||= { 0 => 0 }
    number_of_steps.times do |_i|
      increase_all!
      flash_flash_flash!
      @flashes[@flashes.keys.max + 1] = @data.flatten.count(0)
    end
    @flashes.values.sum
  end

  def increase_all!
    @data.each do |row|
      row.each_with_index do |_cell, i|
        row[i] += 1
      end
    end
  end

  def flash_flash_flash!
    while @data.flatten.count { |a| a > 9 } > 0
      to_flash = []
      @data.each_with_index do |row, i|
        row.each_with_index do |cell, j|
          to_flash << [i, j] if cell > 9
        end
      end
      to_flash.each do |y, x|
        @data[y][x] = 0
        enlight_neighbours!(y, x)
      end
    end
  end

  def enlight_neighbours!(i, j)
    enlight_cell(i - 1, j) if i > 0
    enlight_cell(i + 1, j) if i < @data.length - 1
    enlight_cell(i, j - 1) if j > 0
    enlight_cell(i, j + 1) if j < @data[i].length - 1
    enlight_cell(i - 1, j - 1) if i > 0 && j > 0
    enlight_cell(i - 1, j + 1) if i > 0 && j < @data[i].length - 1
    enlight_cell(i + 1, j - 1) if i < @data.length - 1 && j > 0
    enlight_cell(i + 1, j + 1) if i < @data.length - 1 && j < @data[i].length - 1
  end

  def enlight_cell(y, x)
    return @data[y][x] if @data[y][x].zero?

    @data[y][x] += 1
  end
end
