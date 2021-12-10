require 'pry'

class D10
  class Line
    attr_reader :score

    def initialize(text)
      @text = text
      @score = 0
    end

    def corrupted?
      complete? && chunk.is_a?(String)
    end

    def chunk
      @chunk ||= find_chunk(@text.dup)
    end

    def incomplete?
      !complete?
    end

    def complete?
      closing_chars = '} ) ] >'.chars
      return true if chunk.empty?

      index = closing_chars.map do |c|
        chunk.chars.index(c)
      end.compact.min
      calculate_score!(chunk[index]) unless index.nil?
      !index.nil?
    end

    def calculate_score!(char)
      @score = case char
               when ')'
                 3
               when ']'
                 57
               when '}'
                 1197
               when '>'
                 25_137
               end
    end

    def find_chunk(text)
      return [] if text.empty?

      matches = text.scan(/(\(\))|(\{\})|(<>)|(\[\])/).map(&:compact).flatten
      return text if matches.empty?

      matches.each do |chars|
        text.gsub!(chars, '')
        return text unless find_chunk(text).empty?
      end
      []
    end
  end

  def initialize(data)
    @lines = data.map(&:strip).map { |l| Line.new(l) }
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    @lines.select(&:corrupted?).sum(&:score)
  end

  def part2
    all_scores = @lines.select(&:incomplete?).map do |line|
      scores = line.chunk.chars.reverse.map do |char|
        {
          '(' => 1,
          '<' => 4,
          '{' => 3,
          '[' => 2
        }[char]
      end
      scores.inject { |sum, n| (sum * 5) + n }
    end
    all_scores.sort[all_scores.size / 2]
  end
end
