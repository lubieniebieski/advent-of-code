require 'pry'

class D04
  def initialize(data)
    @data = data.map(&:strip)
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    numbers = @data.first.split(',').map(&:to_i)
    boards = find_boards(@data[2..])
    numbers.each do |n|
      boards.each do |board|
        board.mark!(n)
        return board.unmarked_sum * n if board.ready?
      end
    end
  end

  def part2
    numbers = @data.first.split(',').map(&:to_i)
    boards = find_boards(@data[2..])
    numbers.each do |n|
      boards.each do |board|
        board.mark!(n)

        return board.unmarked_sum * n if boards.count(&:ready?) == boards.count
      end
    end
  end

  private

  def find_boards(data)
    boards = data.delete_if(&:empty?).each_slice(5).to_a
    boards.map { |b| Board.new(b) }
  end

  class Board
    def initialize(data)
      @rows = data.map { |row| row.split(/[^\d]+/).map { |e| Element.new(e) } }
      @columns = @rows.transpose
    end

    def to_s
      @rows.map { |row| row.map(&:to_s).join("\t") }.join("\n")
    end

    def unmarked_sum
      @rows.flatten.select(&:unmarked?).map(&:value).sum
    end

    def mark!(value)
      @rows.each do |row|
        row.each do |element|
          element.mark! if element.value == value
        end
      end
    end

    def ready?
      @rows.any? { |row| row.all?(&:marked?) } || @columns.any? { |column| column.all?(&:marked?) }
    end
  end

  class Element
    attr_reader :value, :marked

    def initialize(value)
      @value = value.to_i
      @marked = false
    end

    def to_s
      value.to_s
    end

    def marked?
      @marked
    end

    def unmarked?
      !marked?
    end

    def mark!
      @marked = true
    end
  end
end
