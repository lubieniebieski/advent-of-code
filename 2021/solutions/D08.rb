require 'pry'

class D08
  class Entry
    def initialize(line)
      @decoder = decoder(line.split('|')[0].strip)
      encode_numbers!
      @digits = decode(line.split('|')[1].strip)
    end

    def number_of_simple_digits
      @digits.count { |digit| [1, 4, 7, 8].include?(digit) }
    end

    def output
      @digits.join.to_i
    end

    def decoder(pattern)
      output = {}

      eight = pattern.split.max_by(&:size)
      one = pattern.split.min_by(&:size)
      four = pattern.split.find { |digit| digit.size == 4 }
      seven = pattern.split.find { |digit| digit.size == 3 }
      six = pattern.split.select do |a|
              a.size == 6
            end.reject { |a| a.match(one[0]) && a.match(one[1]) }.first
      output[0] = (seven.chars - one.chars).first
      output[5] = (six.chars & one.chars).first
      output[2] = (one.chars - six.chars).first
      nine = (pattern.split.select do |a|
                a.size == 6
              end - [six]).select { |a| (a.chars & four.chars).size == 4 }.first
      output[4] = (eight.chars - nine.chars).first
      zero = (pattern.split.select { |a| a.size == 6 } - [six, nine]).first
      output[3] = (eight.chars - zero.chars).first
      output[1] = ((zero.chars & six.chars & (nine.chars - one.chars)) & four.chars).first
      output[6] = (eight.chars - output.values).first

      output
    end

    def encode_numbers!
      @decoder_map ||= {}
      @decoder_map[0] = encode_digit [0, 1, 2, 4, 5, 6]
      @decoder_map[1] = encode_digit [2, 5]
      @decoder_map[2] = encode_digit [0, 2, 3, 4, 6]
      @decoder_map[3] = encode_digit [0, 2, 3, 5, 6]
      @decoder_map[4] = encode_digit [1, 2, 3, 5]
      @decoder_map[5] = encode_digit [0, 1, 3, 5, 6]
      @decoder_map[6] = encode_digit [0, 1, 3, 4, 5, 6]
      @decoder_map[7] = encode_digit [0, 2, 5]
      @decoder_map[8] = encode_digit [0, 1, 2, 3, 4, 5, 6]
      @decoder_map[9] = encode_digit [0, 1, 2, 3, 5, 6]
    end

    def encode_digit(array)
      array.map { |i| @decoder[i] }.sort.join
    end

    def decode_digit(digit)
      @decoder_map.key(digit.chars.sort.join)
    end

    def decode(pattern)
      pattern.split.map do |digit|
        decode_digit(digit)
      end
    end
  end

  def initialize(data)
    @data = data.map { |line| Entry.new(line) }
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    @data.sum(&:number_of_simple_digits)
  end

  def part2
    @data.sum(&:output)
  end
end
